package internal

import (
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
			fat.cursor.OffsetY(1)
			fat.cursor.X = 0
			continue
		}
		x := fat.cursor.X
		y := fat.cursor.Y
		curLine := fat.text[y]
		firstChar := strings.Split(text, "")[0]
		remainingChars := strings.Split(text, "")[1:]
		if fat.cursor.X == len(curLine) {
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
