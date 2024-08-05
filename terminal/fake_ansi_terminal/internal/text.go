package internal

import (
	"regexp"
	"strings"

	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type Text struct {
	value string
}

func NewText(value string) Text {
	return Text{
		value: value,
	}
}

func (t *Text) PopLeft() string {
	if t.value == "" {
		return ""
	}
	if strings.HasPrefix(t.value, ansi_escape.CURSOR_TO_HOME) {
		t.value, _ = strings.CutPrefix(t.value, ansi_escape.CURSOR_TO_HOME)
		return ansi_escape.CURSOR_TO_HOME
	}
	if strings.HasPrefix(t.value, ansi_escape.ERASE_SCREEN) {
		t.value, _ = strings.CutPrefix(t.value, ansi_escape.ERASE_SCREEN)
		return ansi_escape.ERASE_SCREEN
	}
	ansiEscapeSequenceRegex, _ := regexp.Compile("\033\\[[0-9]{1,}[A-Z]{1}")
	ansiEscapeSequenceSeqLoc := ansiEscapeSequenceRegex.FindStringIndex(t.value)
	if ansiEscapeSequenceSeqLoc != nil && ansiEscapeSequenceSeqLoc[0] == 0 {
		ansiEscapeSequenceLeft := t.value[0:ansiEscapeSequenceSeqLoc[1]]
		t.value = t.value[ansiEscapeSequenceSeqLoc[1]:]
		return ansiEscapeSequenceLeft
	}
	chars := strings.Split(t.value, "")
	firstChar := chars[0]
	remainingChars := chars[1:]
	t.value = strings.Join(remainingChars, "")
	return firstChar
}

func (t *Text) Value() string {
	return t.value
}
