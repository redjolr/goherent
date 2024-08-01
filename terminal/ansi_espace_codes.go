package terminal

import "fmt"

const ANSI_CURSOR_TO_HOME = "\033[H"
const ANSI_RED string = "\033[31m"
const ANSI_GREEN string = "\033[32m"
const ANSI_YELLOW string = "\u001B[33m"
const ANSI_COLOR_RESET string = "\033[0m"

const ANSI_BOLD string = "\033[1m"
const ANSI_RESET_BOLD string = "\033[22m"

const ANSI_YELLOW_CIRCLE string = ANSI_YELLOW + "⚬" + ANSI_COLOR_RESET

func MoveCursorUpNRowsAnsiSequence(n int) string {
	return fmt.Sprintf("\033[%dA", n)
}

func MoveCursorDownNRowsAnsiSequence(n int) string {
	return fmt.Sprintf("\033[%dB", n)
}

func MoveCursorRightNColsAnsiSequence(n int) string {
	return fmt.Sprintf("\033[%dC", n)
}

func MoveCursorLeftNColsAnsiSequence(n int) string {
	return fmt.Sprintf("\033[%dD", n)
}
