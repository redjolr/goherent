package terminal

import (
	"fmt"
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
	fmt.Print(MoveCursorDownNRowsAnsiSequence(n))
}

func (at *AnsiTerminal) MoveUp(n int) {
	fmt.Print(MoveCursorUpNRowsAnsiSequence(n))
}

func (at *AnsiTerminal) MoveLeft(n int) {
	fmt.Print(MoveCursorLeftNColsAnsiSequence(n))
}
