package terminal

import (
	"fmt"
	"math"

	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type AnsiTerminal struct {
	height int
	width  int
}

func NewBoundedAnsiTerminal(width, height int) AnsiTerminal {
	return AnsiTerminal{
		height: height,
		width:  width,
	}
}

func NewUnboundedAnsiTerminal() AnsiTerminal {
	return AnsiTerminal{
		height: math.MaxInt,
		width:  math.MaxInt,
	}
}

func (at *AnsiTerminal) IsBounded() bool {
	return at.height < math.MaxInt
}

func (at *AnsiTerminal) Print(text string) {
	fmt.Print(text)
}

func (at *AnsiTerminal) Printf(text string, args ...any) {
	print := fmt.Sprintf(text, args...)
	at.Print(print)
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

func (at *AnsiTerminal) MoveRight(n int) {
	fmt.Print(ansi_escape.MoveCursorRightNCols(n))
}

func (at *AnsiTerminal) Height() int {
	return at.height
}
