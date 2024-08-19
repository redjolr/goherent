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
			if fat.coords.X >= len(fat.lines[fat.coords.Y]) {
				emptySpacesToAddCount := fat.coords.X - len(fat.lines[fat.coords.Y])
				emptySpacesToAdd := strings.Split(strings.Repeat(" ", emptySpacesToAddCount), "")
				fat.lines[fat.coords.Y] = append(fat.lines[fat.coords.Y], emptySpacesToAdd...)
				fat.lines[fat.coords.Y] = append(fat.lines[fat.coords.Y], curSequence.Value())
				fat.coords.OffsetX(1)
			} else {
				fat.lines[fat.coords.Y][fat.coords.X] = curSequence.Value()
				fat.coords.OffsetX(1)
			}
		}
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
