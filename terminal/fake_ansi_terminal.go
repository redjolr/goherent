package terminal

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/redjolr/goherent/terminal/ansi_escape"
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

func (fat *FakeAnsiTerminal) cursorToVisibleUpperLeftCorner() {
	fat.coords.X = 0
	fat.cursorToVisibleUpperLine()
}

func (fat *FakeAnsiTerminal) cursorToVisibleUpperLine() {
	var visibleUpperRow int
	if len(fat.lines) <= fat.height {
		visibleUpperRow = 0
	} else {
		visibleUpperRow = len(fat.lines) - fat.height
	}
	fat.coords.Y = visibleUpperRow
}

func (fat *FakeAnsiTerminal) Print(text string) {
	for len(strings.Split(text, "")) > 0 {
		if strings.HasPrefix(text, ansi_escape.CURSOR_TO_HOME) {
			text, _ = strings.CutPrefix(text, ansi_escape.CURSOR_TO_HOME)
			fat.cursorToVisibleUpperLeftCorner()
			continue
		}
		if strings.HasPrefix(text, ansi_escape.ERASE_SCREEN) {
			text, _ = strings.CutPrefix(text, ansi_escape.ERASE_SCREEN)
			var visibleUpperRow int
			if len(fat.lines) <= fat.height {
				visibleUpperRow = 0
			} else {
				visibleUpperRow = len(fat.lines) - fat.height
			}
			for i := visibleUpperRow; i < len(fat.lines); i++ {
				fat.lines[i] = []string{}
			}
			if fat.coords.X > 0 {
				fat.lines[fat.coords.Y] = strings.Split(strings.Repeat(" ", fat.coords.X), "")
			}
			continue
		}
		if strings.HasPrefix(text, "\n") {
			text, _ = strings.CutPrefix(text, "\n")
			if fat.coords.Y == len(fat.lines)-1 {
				fat.lines = append(fat.lines, []string{""})
			}
			fat.coords.OffsetY(1)
			fat.coords.X = 0
			continue
		}

		// Move left
		moveCursorLeftRegex, _ := regexp.Compile("\033\\[[0-9]{1,}D")
		moveCursorLeftSeqLoc := moveCursorLeftRegex.FindStringIndex(text)
		if moveCursorLeftSeqLoc != nil && moveCursorLeftSeqLoc[0] == 0 {
			moveCursorLeftSeq := text[0:moveCursorLeftSeqLoc[1]]
			text = text[moveCursorLeftSeqLoc[1]:]

			moveLeftCountAsStr, _ := strings.CutPrefix(moveCursorLeftSeq, "\033[")
			moveLeftCountAsStr, _ = strings.CutSuffix(moveLeftCountAsStr, "D")
			moveLeftCount, err := strconv.Atoi(moveLeftCountAsStr)
			if err != nil {
				panic("Cannot determine the number steps to move left.")
			}
			fat.coords.OffsetX(-min(moveLeftCount, fat.coords.X))
			continue
		}

		// Move right
		moveCursorRightRegex, _ := regexp.Compile("\033\\[[0-9]{1,}C")
		moveCursorRightSeqLoc := moveCursorRightRegex.FindStringIndex(text)
		if moveCursorRightSeqLoc != nil && moveCursorRightSeqLoc[0] == 0 {
			moveCursorRightSeq := text[0:moveCursorRightSeqLoc[1]]
			text = text[moveCursorRightSeqLoc[1]:]

			moveRightCountAsStr, _ := strings.CutPrefix(moveCursorRightSeq, "\033[")
			moveRightCountAsStr, _ = strings.CutSuffix(moveRightCountAsStr, "C")
			moveRightCount, err := strconv.Atoi(moveRightCountAsStr)
			if err != nil {
				panic("Cannot determine the number steps to move right.")
			}
			fat.coords.OffsetX(moveRightCount)
			continue
		}

		// Move up
		moveCursorUpRegex, _ := regexp.Compile("\033\\[[0-9]{1,}A")
		moveCursorUpSeqLoc := moveCursorUpRegex.FindStringIndex(text)
		if moveCursorUpSeqLoc != nil && moveCursorUpSeqLoc[0] == 0 {
			moveCursorUpSeq := text[0:moveCursorUpSeqLoc[1]]
			text = text[moveCursorUpSeqLoc[1]:]

			moveUpCountAsStr, _ := strings.CutPrefix(moveCursorUpSeq, "\033[")
			moveUpCountAsStr, _ = strings.CutSuffix(moveUpCountAsStr, "A")
			moveUpCount, err := strconv.Atoi(moveUpCountAsStr)
			if err != nil {
				panic("Cannot determine the number steps to move up.")
			}
			fat.coords.OffsetY(-min(moveUpCount, fat.coords.Y))
			if fat.coords.X > len(fat.lines[fat.coords.Y]) {
				newLineStr := strings.Join(fat.lines[fat.coords.Y], "") + strings.Repeat(" ", fat.coords.X-len(fat.lines[fat.coords.Y]))
				fat.lines[fat.coords.Y] = strings.Split(newLineStr, "")
			}
			continue
		}

		// Move down
		moveCursorDownRegex, _ := regexp.Compile("\033\\[[0-9]{1,}B")
		moveCursorDownSeqLoc := moveCursorDownRegex.FindStringIndex(text)
		if moveCursorDownSeqLoc != nil && moveCursorDownSeqLoc[0] == 0 {
			moveCursorDownSeq := text[0:moveCursorDownSeqLoc[1]]
			text = text[moveCursorDownSeqLoc[1]:]

			moveUpCountAsStr, _ := strings.CutPrefix(moveCursorDownSeq, "\033[")
			moveUpCountAsStr, _ = strings.CutSuffix(moveUpCountAsStr, "B")
			moveDownCount, err := strconv.Atoi(moveUpCountAsStr)
			if err != nil {
				panic("Cannot determine the number steps to move down.")
			}
			fat.coords.OffsetY(moveDownCount)
			continue
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

		firstChar := strings.Split(text, "")[0]
		remainingChars := strings.Split(text, "")[1:]
		if fat.coords.X >= len(fat.lines[fat.coords.Y]) {
			emptySpacesToAdd := fat.coords.X - len(fat.lines[fat.coords.Y])
			for i := 0; i < emptySpacesToAdd; i++ {
				fat.lines[fat.coords.Y] = append(fat.lines[fat.coords.Y], " ")
			}
			fat.lines[fat.coords.Y] = append(fat.lines[fat.coords.Y], firstChar)
			text = strings.Join(remainingChars, "")
			fat.coords.OffsetX(1)
		} else {
			lineChars := fat.lines[fat.coords.Y]
			text = strings.Join(remainingChars, "")
			lineChars[fat.coords.X] = firstChar
			fat.lines[fat.coords.Y] = lineChars
			fat.coords.OffsetX(1)
		}
	}
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
