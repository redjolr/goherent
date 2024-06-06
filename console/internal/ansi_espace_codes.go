package internal

import "fmt"

const CursorToHomePosEscapeCode = "\033[H"

func MoveCursorUpNRows(n int) string {
	return fmt.Sprintf("\033[%dA", n)
}

func MoveCursorDownNRows(n int) string {
	return fmt.Sprintf("\033[%dB", n)
}

func MoveCursorRightNCols(n int) string {
	return fmt.Sprintf("\033[%dC", n)
}

func MoveCursorLeftNCols(n int) string {
	return fmt.Sprintf("\033[%dD", n)
}
