package fake_ansi_terminal

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal/internal"
)

type FakeAnsiTerminal struct {
	width  int
	height int
	lines  [][]string
	coords Coordinates
}

func NewFakeAnsiTerminal(width, height int) FakeAnsiTerminal {
	origin := Origin()
	return FakeAnsiTerminal{
		width:  width,
		height: height,
		lines:  [][]string{{}},
		coords: origin,
	}
}

func (fat *FakeAnsiTerminal) Print(t string) {
	text := internal.NewText(t)
	for !text.IsEmpty() {
		curSequence := text.PopLeft()
		if curSequence.Equals(ansi_escape.CURSOR_TO_HOME) {
			fat.cursorToVisibleUpperLeftCorner()
			continue
		}
		if curSequence.Equals(ansi_escape.ERASE_SCREEN) {
			for i := fat.visibleUpperLine(); i < len(fat.lines); i++ {
				fat.lines[i] = []string{}
			}
			if fat.coords.X > 0 {
				fat.lines[fat.coords.Y] = strings.Split(strings.Repeat(" ", fat.coords.X), "")
			}
			continue
		}
		if curSequence.Equals("\n") {
			if fat.coords.Y == len(fat.lines)-1 {
				fat.lines = append(fat.lines, []string{""})
			}
			fat.coords.OffsetY(1)
			fat.coords.X = 0
			continue
		}

		// Move left
		if curSequence.Matches("\033\\[[0-9]{1,}D") {
			fat.coords.OffsetX(-min(curSequence.MoveLeftCount(), fat.coords.X))
		}

		// Move right
		if curSequence.Equals(ansi_escape.MoveCursorRightNCols(0)) {
			fat.coords.OffsetX(1)
		} else if curSequence.Matches("\033\\[[0-9]{1,}C") {
			fat.coords.OffsetX(curSequence.MoveRightCount())
		}

		// Move up
		if curSequence.Equals(ansi_escape.MoveCursorUpNRows(0)) {
			fat.coords.OffsetY(-min(1, fat.coords.Y))
		} else if curSequence.Matches("\033\\[[0-9]{1,}A") {
			if fat.coords.Y-curSequence.MoveUpCount() < fat.visibleUpperLine() {
				fat.coords.Y = fat.visibleUpperLine()
			} else {
				fat.coords.OffsetY(-min(curSequence.MoveUpCount(), fat.coords.Y))
			}
			if fat.coords.X > len(fat.lines[fat.coords.Y]) {
				newLineStr := strings.Join(fat.lines[fat.coords.Y], "") + strings.Repeat(" ", fat.coords.X-len(fat.lines[fat.coords.Y]))
				fat.lines[fat.coords.Y] = strings.Split(newLineStr, "")
			}
		}

		// Move down
		if curSequence.Equals(ansi_escape.MoveCursorDownNRows(0)) {
			fat.coords.OffsetY(1)
		} else if curSequence.Matches("\033\\[[0-9]{1,}B") {
			fat.coords.OffsetY(curSequence.MoveDownCount())
		}

		// Append empty strings to the right
		if fat.coords.Y >= len(fat.lines) {
			linesToAddCount := fat.coords.Y - len(fat.lines) + 1
			fat.lines = append(fat.lines, make([][]string, linesToAddCount)...)
			if fat.coords.X > len(fat.lines[fat.coords.Y]) {
				newLineStr := strings.Join(fat.lines[fat.coords.Y], "") + strings.Repeat(" ", fat.coords.X-len(fat.lines[fat.coords.Y]))
				fat.lines[fat.coords.Y] = append(fat.lines[fat.coords.Y], strings.Split(newLineStr, "")...)
			}
		}

		if curSequence.IsPrintable() {
			fat.writeGlyph(curSequence.Value())
		}
	}
}

// wideCharContinuation is the placeholder stored in the second cell of a
// double-width glyph (emoji, CJK). A real terminal advances the cursor by two
// columns for such a glyph, so the fake reserves two slice cells for it: the
// glyph itself, followed by this sentinel. Text() renders the sentinel as
// nothing.
const wideCharContinuation = "\x00"

// writeGlyph places a single printable glyph at the cursor, modelling the cell
// width of a real terminal: a double-width glyph occupies two columns, and
// overwriting one half of an existing wide glyph blanks its orphaned partner.
func (fat *FakeAnsiTerminal) writeGlyph(glyph string) {
	width := displayWidth(glyph)
	line := fat.coords.Y
	col := fat.coords.X
	fat.padLineTo(line, col+width)

	// Landing on the right half of a wide glyph strands its start cell.
	if col-1 >= 0 && fat.lines[line][col] == wideCharContinuation {
		fat.lines[line][col-1] = " "
	}
	// Writing over the start of a wide glyph strands its continuation cell,
	// unless we are about to overwrite that cell ourselves.
	for j := col; j <= col+width-1; j++ {
		partner := j + 1
		if displayWidth(fat.lines[line][j]) == 2 &&
			partner < len(fat.lines[line]) &&
			fat.lines[line][partner] == wideCharContinuation &&
			partner > col+width-1 {
			fat.lines[line][partner] = " "
		}
	}

	fat.lines[line][col] = glyph
	if width == 2 {
		fat.lines[line][col+1] = wideCharContinuation
	}
	fat.coords.OffsetX(width)
}

// padLineTo right-pads a line with blank cells until it has at least length
// cells, so that columns up to length-1 can be written.
func (fat *FakeAnsiTerminal) padLineTo(line, length int) {
	for len(fat.lines[line]) < length {
		fat.lines[line] = append(fat.lines[line], " ")
	}
}

func (fat *FakeAnsiTerminal) Printf(text string, args ...any) {
	print := fmt.Sprintf(text, args...)
	fat.Print(print)
}

func (fat *FakeAnsiTerminal) Text() string {
	text := ""
	for lineIndex, line := range fat.lines {
		for _, char := range line {
			if char == wideCharContinuation {
				continue
			}
			text += char
		}
		if lineIndex < len(fat.lines)-1 {
			text += "\n"
		}
	}
	return text
}

func (fat *FakeAnsiTerminal) GoToOrigin() {
	fat.coords.SetToOrigin()
}

func (fat *FakeAnsiTerminal) MoveUp(n int) {
	fat.Print(ansi_escape.MoveCursorUpNRows(n))
}

func (fat *FakeAnsiTerminal) MoveDown(n int) {
	fat.Print(ansi_escape.MoveCursorDownNRows(n))
}

func (fat *FakeAnsiTerminal) MoveRight(n int) {
	fat.Print(ansi_escape.MoveCursorRightNCols(n))
}

func (fat *FakeAnsiTerminal) MoveLeft(n int) {
	fat.Print(ansi_escape.MoveCursorLeftNCols(n))
}

func (fat *FakeAnsiTerminal) cursorToVisibleUpperLeftCorner() {
	fat.coords.X = 0
	fat.coords.Y = fat.visibleUpperLine()
}

func (fat *FakeAnsiTerminal) visibleUpperLine() int {
	var visibleUpperLine int
	if len(fat.lines) <= fat.height {
		visibleUpperLine = 0
	} else {
		visibleUpperLine = len(fat.lines) - fat.height
	}
	return visibleUpperLine
}

func (fat *FakeAnsiTerminal) Height() int {
	return fat.height
}
