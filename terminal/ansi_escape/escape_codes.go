package ansi_escape

import "fmt"

const CURSOR_TO_HOME = "\033[H"
const ERASE_SCREEN = "\033[2J"
const RED string = "\033[31m"
const GREEN string = "\033[32m"
const YELLOW string = "\u001B[33m"
const COLOR_RESET string = "\033[0m"

const BOLD string = "\033[1m"
const RESET_BOLD string = "\033[22m"

const YELLOW_CIRCLE string = YELLOW + "⚬" + COLOR_RESET

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
