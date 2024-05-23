package textblock

import (
	"errors"
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

	if line == "" {
		tb.lines[y] = writeStr
	} else if x == len(line) {
		tb.lines[y] += writeStr
	} else {
		lineChars := strings.Split(tb.lines[y], "")
		writeChars := strings.Split(writeStr, "")
		for i, char := range writeChars {
			lineChars[x+i] = char
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
	tb.cursorPosition = tb.cursorPosition.OffsetX(offset)
}

func (tb *Textblock) MoveCursorLeft(offset int) {
	tb.cursorPosition = tb.cursorPosition.OffsetX(offset)
}

func (tb *Textblock) MoveCursorToOrigin() {
	if len(tb.lines) == 0 {
		tb.lines = []string{""}
	}

	tb.cursorPosition.X = 0
	tb.cursorPosition.Y = 0
}

func (tb *Textblock) MoveCursorTo(x int, y int) error {
	if x < 0 || y < 0 {
		return errors.New("Coordinates cannot be negative.")
	}
	if y >= len(tb.lines) {
		return errors.New("Cannot move cursor to Y coordinate that is larger than len(lines) - 1.")
	}

	tb.cursorPosition.X = x
	tb.cursorPosition.Y = y
	return nil
}
