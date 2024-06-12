package terminal

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/redjolr/goherent/console/coordinates"
)

type FakeAnsiTerminal struct {
	lines  []string
	cursor coordinates.Coordinates
}

func NewFakeAnsiTerminal() FakeAnsiTerminal {
	return FakeAnsiTerminal{
		lines:  []string{""},
		cursor: coordinates.Origin(),
	}
}

func (fat *FakeAnsiTerminal) Print(text string) {
	for len(text) > 0 {
		if strings.HasPrefix(text, CursorToHomePosEscapeCode) {
			text, _ = strings.CutPrefix(text, CursorToHomePosEscapeCode)
			fat.cursor.GoToOrigin()
			continue
		}
		if strings.HasPrefix(text, "\n") {
			text, _ = strings.CutPrefix(text, "\n")
			fat.lines = append(fat.lines, "")
			fat.cursor.MoveDown(1)
			fat.cursor.X = 0
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

			fat.cursor.MoveLeft(min(moveLeftCount, fat.cursor.X))
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
			fat.cursor.MoveRight(moveRightCount)
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
				panic("Cannot determine the number steps to move left.")
			}
			fat.cursor.MoveUp(min(moveUpCount, fat.cursor.Y))
			if fat.cursor.X > len(fat.lines[fat.cursor.Y]) {
				fat.lines[fat.cursor.Y] = fat.lines[fat.cursor.Y] + strings.Repeat(" ", fat.cursor.X-len(fat.lines[fat.cursor.Y]))
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
			fat.cursor.MoveDown(moveDownCount)
			continue
		}

		// Append empty strings to the right
		if fat.cursor.Y >= len(fat.lines) {
			linesToAddCount := fat.cursor.Y - len(fat.lines) + 1
			fat.lines = append(fat.lines, make([]string, linesToAddCount)...)
			if fat.cursor.X > len(fat.lines[fat.cursor.Y]) {
				fat.lines[fat.cursor.Y] = fat.lines[fat.cursor.Y] + strings.Repeat(" ", fat.cursor.X-len(fat.lines[fat.cursor.Y]))
			}
		}

		firstChar := strings.Split(text, "")[0]
		remainingChars := strings.Split(text, "")[1:]
		if fat.cursor.X >= len(fat.lines[fat.cursor.Y]) {
			emptySpacesToAdd := fat.cursor.X - len(fat.lines[fat.cursor.Y])
			fat.lines[fat.cursor.Y] += strings.Repeat(" ", emptySpacesToAdd)
			fat.lines[fat.cursor.Y] += firstChar
			fat.cursor.MoveRight(1)
			text = strings.Join(remainingChars, "")
		} else {
			lineChars := strings.Split(fat.lines[fat.cursor.Y], "")
			text = strings.Join(remainingChars, "")
			lineChars[fat.cursor.X] = firstChar
			fat.lines[fat.cursor.Y] = strings.Join(lineChars, "")
			fat.cursor.MoveRight(1)
		}
	}
}

func (fat *FakeAnsiTerminal) Text() string {
	return strings.Join(fat.lines, "\n")
}
