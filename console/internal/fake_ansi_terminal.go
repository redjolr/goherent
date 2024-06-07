package internal

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/redjolr/goherent/console/coordinates"
)

type FakeAnsiTerminal struct {
	text   []string
	cursor coordinates.Coordinates
}

func NewFakeAnsiTerminal() FakeAnsiTerminal {
	return FakeAnsiTerminal{
		text:   []string{""},
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
			fat.text = append(fat.text, "")
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
			for i := 0; i < moveRightCount; i++ {
				if fat.cursor.X == len(fat.text[fat.cursor.Y]) {
					fat.text[fat.cursor.Y] += " "
				}
				fat.cursor.MoveRight(1)
			}
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
			if fat.cursor.X > len(fat.text[fat.cursor.Y]) {
				fat.text[fat.cursor.Y] = fat.text[fat.cursor.Y] + strings.Repeat(" ", fat.cursor.X-len(fat.text[fat.cursor.Y]))
			}
			continue
		}

		x := fat.cursor.X
		y := fat.cursor.Y
		firstChar := strings.Split(text, "")[0]
		remainingChars := strings.Split(text, "")[1:]
		if fat.cursor.X == len(fat.text[y]) {
			fat.text[y] += firstChar
			fat.cursor.MoveRight(1)
			text = strings.Join(remainingChars, "")
		} else {
			lineChars := strings.Split(fat.text[y], "")
			text = strings.Join(remainingChars, "")
			lineChars[x] = firstChar
			fat.text[y] = strings.Join(lineChars, "")
			fat.cursor.MoveRight(1)
		}

	}
}

func (fat *FakeAnsiTerminal) Text() string {
	return strings.Join(fat.text, "\n")
}

// func (fat *FakeAnsiTerminal) padRowsToLongestWidth() {
// 	maxWidth := 0
// 	for _, line := range fat.text {
// 		if len(line) > maxWidth {
// 			maxWidth = len(line)
// 		}
// 	}

// 	for i := 0; i < len(fat.text); i++ {
// 		if len(fat.text[i]) < maxWidth {
// 			fat.text[i] = fat.text[i] + strings.Repeat(" ", maxWidth-len(fat.text[i]))
// 		}
// 	}

// }
