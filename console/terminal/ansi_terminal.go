package terminal

import "fmt"

type AnsiTerminal struct {
}

func NewAnsiTerminal() AnsiTerminal {
	return AnsiTerminal{}
}

func (at *AnsiTerminal) Print(text string) {
	fmt.Print(text)
}
