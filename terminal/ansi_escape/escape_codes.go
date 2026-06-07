package ansi_escape

import "fmt"

const CURSOR_TO_HOME = "\033[H"

// ERASE_SCREEN clears from the cursor to the end of the screen ("\033[0J"),
// NOT the whole screen ("\033[2J"). It is always used right after CURSOR_TO_HOME,
// so from the home position it still blanks the entire visible screen — but
// unlike "\033[2J", "\033[0J" does not push the current frame into the terminal's
// scrollback on VTE-based terminals (gnome-terminal, etc.). That scrollback push
// is what made each concurrent redraw stack up instead of refreshing in place.
const ERASE_SCREEN = "\033[0J"
const RED string = "\033[31m"
const GREEN string = "\033[32m"
const YELLOW string = "\u001B[33m"
const COLOR_RESET string = "\033[0m"

const BOLD string = "\033[1m"
const RESET_BOLD string = "\033[22m"

const DIM string = "\033[2m"

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
