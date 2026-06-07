// Package vtfake is a faithful (small) ANSI terminal emulator used as a test
// double. Unlike the older FakeAnsiTerminal, it models a real terminal the way
// xterm/VTE actually behave — a fixed width×height viewport with scrollback,
// scroll-on-newline, autowrap, double-width glyphs (emoji), and the cursor /
// erase / save-restore semantics captured from a real terminal (see
// vtfake_test.go, whose expectations were recorded with capture_terminal.sh).
//
// It implements terminal.Terminal, so presenters can be driven by it in tests,
// and it exposes VisibleLines / ScrollbackLines / CursorX / CursorY for
// assertions.
//
// Colors: SGR sequences ("\033[…m") are treated as genuinely zero-width — they
// are consumed and do not occupy a cell or move the cursor, and they do not
// appear in the rendered text. vtfake is about geometry (where glyphs land,
// scrolling, wrapping); byte-level color assertions stay in the existing
// fake_ansi_terminal tests.
package vtfake

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/redjolr/goherent/internal/utils"
)

const (
	blank = " "
	// continuation marks the right half of a double-width glyph. Rendering skips
	// it; it exists so the column count matches a real terminal.
	continuation = "\x00"
)

// Terminal is an in-memory ANSI terminal of a fixed size.
type Terminal struct {
	width, height  int
	rows           [][]string // exactly height rows, each exactly width cells
	scrollback     [][]string
	cx, cy         int // cursor column/row, 0-based, within the viewport
	savedX, savedY int // DECSC / DECRC (absolute — does not follow scroll)
	wrapPending    bool
}

func New(width, height int) *Terminal {
	t := &Terminal{width: width, height: height}
	t.rows = make([][]string, height)
	for y := range t.rows {
		t.rows[y] = t.blankRow()
	}
	return t
}

func (t *Terminal) blankRow() []string {
	row := make([]string, t.width)
	for x := range row {
		row[x] = blank
	}
	return row
}

// ---- terminal.Terminal interface ----

func (t *Terminal) Height() int { return t.height }

func (t *Terminal) Printf(text string, args ...any) { t.Print(fmt.Sprintf(text, args...)) }

func (t *Terminal) MoveUp(n int)    { t.Print(fmt.Sprintf("\033[%dA", n)) }
func (t *Terminal) MoveDown(n int)  { t.Print(fmt.Sprintf("\033[%dB", n)) }
func (t *Terminal) MoveRight(n int) { t.Print(fmt.Sprintf("\033[%dC", n)) }
func (t *Terminal) MoveLeft(n int)  { t.Print(fmt.Sprintf("\033[%dD", n)) }

func (t *Terminal) Print(s string) {
	for i := 0; i < len(s); {
		switch c := s[i]; {
		case c == 0x1b:
			i += t.handleEscape(s[i:])
		case c == '\n':
			t.lineFeed()
			i++
		case c == '\r':
			t.cx = 0
			t.wrapPending = false
			i++
		default:
			r, size := utf8.DecodeRuneInString(s[i:])
			t.writeGlyph(string(r))
			i += size
		}
	}
}

// ---- escape handling ----

func (t *Terminal) handleEscape(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	switch s[1] {
	case '7': // DECSC: save cursor (absolute position)
		t.savedX, t.savedY = t.cx, t.cy
		return 2
	case '8': // DECRC: restore cursor (to the saved absolute cell)
		t.cx, t.cy = t.savedX, t.savedY
		t.wrapPending = false
		return 2
	case '[': // CSI
		j := 2
		for j < len(s) && (s[j] == ';' || (s[j] >= '0' && s[j] <= '9')) {
			j++
		}
		if j >= len(s) {
			return len(s)
		}
		t.handleCSI(s[j], s[2:j])
		return j + 1
	default:
		return 2
	}
}

func (t *Terminal) handleCSI(final byte, params string) {
	switch final {
	case 'A': // cursor up
		t.cy = clamp(t.cy-moveCount(params), 0, t.height-1)
		t.wrapPending = false
	case 'B': // cursor down
		t.cy = clamp(t.cy+moveCount(params), 0, t.height-1)
		t.wrapPending = false
	case 'C': // cursor forward
		t.cx = clamp(t.cx+moveCount(params), 0, t.width-1)
		t.wrapPending = false
	case 'D': // cursor back
		t.cx = clamp(t.cx-moveCount(params), 0, t.width-1)
		t.wrapPending = false
	case 'H', 'f': // cursor position (1-based row;col)
		row, col := 1, 1
		parts := strings.Split(params, ";")
		if len(parts) > 0 && parts[0] != "" {
			row = atoi(parts[0], 1)
		}
		if len(parts) > 1 && parts[1] != "" {
			col = atoi(parts[1], 1)
		}
		t.cy = clamp(row-1, 0, t.height-1)
		t.cx = clamp(col-1, 0, t.width-1)
		t.wrapPending = false
	case 'J': // erase in display
		t.eraseDisplay(atoi(params, 0))
	case 'K': // erase in line
		t.eraseLine(atoi(params, 0))
	case 'm': // SGR (colors): zero width — consumed, never stored or rendered
	}
}

// ---- grid mutations ----

func (t *Terminal) lineFeed() {
	t.cx = 0
	t.wrapPending = false
	if t.cy == t.height-1 {
		t.scrollUp()
	} else {
		t.cy++
	}
}

func (t *Terminal) scrollUp() {
	t.scrollback = append(t.scrollback, t.rows[0])
	copy(t.rows, t.rows[1:])
	t.rows[t.height-1] = t.blankRow()
}

func (t *Terminal) writeGlyph(g string) {
	w := utils.DisplayWidth(g)
	if w < 1 {
		w = 1
	}
	if t.wrapPending {
		t.lineFeed()
	}
	if t.cx+w > t.width { // a wide glyph that won't fit wraps to the next row first
		t.lineFeed()
	}
	row := t.rows[t.cy]
	col := t.cx

	// Overwriting half of an existing double-width glyph orphans its other half,
	// which a real terminal blanks.
	if col-1 >= 0 && row[col] == continuation {
		row[col-1] = blank // we landed on the right half; blank the orphaned left
	}
	for j := col; j <= col+w-1 && j < t.width; j++ {
		partner := j + 1
		if utils.DisplayWidth(row[j]) == 2 &&
			partner < t.width &&
			row[partner] == continuation &&
			partner > col+w-1 {
			row[partner] = blank // blank a continuation cell we won't overwrite
		}
	}

	row[col] = g
	if w == 2 && col+1 < t.width {
		row[col+1] = continuation
	}
	t.cx += w
	if t.cx >= t.width {
		t.cx = t.width - 1
		t.wrapPending = true
	}
}

func (t *Terminal) eraseDisplay(mode int) {
	switch mode {
	case 0: // cursor (inclusive) -> end of screen
		t.eraseInRow(t.cy, t.cx, t.width-1)
		for y := t.cy + 1; y < t.height; y++ {
			t.rows[y] = t.blankRow()
		}
	case 1: // start of screen -> cursor (inclusive)
		for y := 0; y < t.cy; y++ {
			t.rows[y] = t.blankRow()
		}
		t.eraseInRow(t.cy, 0, t.cx)
	case 2: // whole screen
		for y := 0; y < t.height; y++ {
			t.rows[y] = t.blankRow()
		}
	}
}

func (t *Terminal) eraseLine(mode int) {
	switch mode {
	case 0:
		t.eraseInRow(t.cy, t.cx, t.width-1)
	case 1:
		t.eraseInRow(t.cy, 0, t.cx)
	case 2:
		t.eraseInRow(t.cy, 0, t.width-1)
	}
}

func (t *Terminal) eraseInRow(y, from, to int) {
	for x := from; x <= to && x < t.width; x++ {
		t.rows[y][x] = blank
	}
}

// ---- inspection ----

func (t *Terminal) CursorX() int { return t.cx }
func (t *Terminal) CursorY() int { return t.cy }

// VisibleLines returns the height rows of the viewport, each trailing-trimmed
// (matching how a real terminal capture reports them).
func (t *Terminal) VisibleLines() []string {
	out := make([]string, t.height)
	for y := 0; y < t.height; y++ {
		out[y] = renderRow(t.rows[y])
	}
	return out
}

// ScrollbackLines returns the rows that have scrolled off the top, oldest first,
// or nil if nothing has scrolled.
func (t *Terminal) ScrollbackLines() []string {
	if len(t.scrollback) == 0 {
		return nil
	}
	out := make([]string, len(t.scrollback))
	for i, row := range t.scrollback {
		out[i] = renderRow(row)
	}
	return out
}

// Text is the visible viewport joined by newlines, with trailing blank lines
// removed.
func (t *Terminal) Text() string {
	lines := t.VisibleLines()
	end := len(lines)
	for end > 0 && lines[end-1] == "" {
		end--
	}
	return strings.Join(lines[:end], "\n")
}

func renderRow(row []string) string {
	var b strings.Builder
	for _, c := range row {
		if c == continuation {
			continue
		}
		b.WriteString(c)
	}
	return strings.TrimRight(b.String(), " ")
}

// ---- helpers ----

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func atoi(s string, def int) int {
	if s == "" {
		return def
	}
	if i := strings.IndexByte(s, ';'); i >= 0 {
		s = s[:i]
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}

// moveCount is the count for a cursor-move CSI: default 1, and an explicit 0 is
// treated as 1 (matching xterm).
func moveCount(params string) int {
	n := atoi(params, 1)
	if n < 1 {
		n = 1
	}
	return n
}
