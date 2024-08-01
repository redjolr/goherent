package terminal

import (
	"fmt"

	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type AnsiTerminal struct {
}

func NewAnsiTerminal() AnsiTerminal {
	return AnsiTerminal{}
}

func (at *AnsiTerminal) Print(text string) {
	fmt.Print(text)
}

func (at *AnsiTerminal) MoveDown(n int) {
	fmt.Print(ansi_escape.MoveCursorDownNRows(n))
}

func (at *AnsiTerminal) MoveUp(n int) {
	fmt.Print(ansi_escape.MoveCursorUpNRows(n))
}

func (at *AnsiTerminal) MoveLeft(n int) {
	fmt.Print(ansi_escape.MoveCursorLeftNCols(n))
}
