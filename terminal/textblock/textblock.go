package textblock

import (
	"regexp"
	"strings"

	"github.com/redjolr/goherent/terminal/coordinates"
)

type Textblock struct {
	lines          []string
	cursorPosition coordinates.Coordinates
}

func EmptyTextblock() Textblock {
	return Textblock{
		lines:          []string{""},
		cursorPosition: coordinates.Origin(),
	}
}

func FromString(text string) Textblock {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)

	return Textblock{
		lines:          lines,
		cursorPosition: coordinates.New(len(lines[len(lines)-1]), len(lines)-1),
	}
}

func (tb *Textblock) Write(writeStr string) error {
	x := tb.cursorPosition.X
	y := tb.cursorPosition.Y
	line := tb.lines[y]

	if x == len(line) {
		tb.lines[y] += writeStr
	} else {
		lineChars := strings.Split(tb.lines[y], "")
		writeChars := strings.Split(writeStr, "")
		for i, char := range writeChars {
			if x+i < len(lineChars) {
				lineChars[x+i] = char
			} else {
				lineChars = append(lineChars, char)
			}
		}
		tb.lines[y] = strings.Join(lineChars, "")
	}
	tb.MoveCursorRight(len(writeStr))
	return nil
}

func (tb Textblock) Lines() []string {
	return tb.lines
}

func (tb *Textblock) MoveCursorRight(offset int) {
	tb.MoveCursorTo(tb.cursorPosition.X+offset, tb.cursorPosition.Y)
}

func (tb *Textblock) MoveCursorLeft(offset int) {
	tb.MoveCursorTo(tb.cursorPosition.X-offset, tb.cursorPosition.Y)
}

func (tb *Textblock) MoveCursorToOrigin() {
	if len(tb.lines) == 0 {
		tb.lines = []string{""}
	}

	tb.cursorPosition.X = 0
	tb.cursorPosition.Y = 0
}

func (tb *Textblock) MoveCursorTo(x int, y int) {
	if x < 0 || y < 0 {
		panic("Textblock Coordinates cannot be negative.")
	}
	if y >= len(tb.lines) {
		panic("Textblock cannot move cursor to Y coordinate that is >= len(lines) - 1.")
	}
	if x > len(tb.lines[y]) {
		panic("Textblock cannot move cursor to X coordinate that is > len(lines[y]) - 1.")
	}

	tb.cursorPosition.X = x
	tb.cursorPosition.Y = y
}
