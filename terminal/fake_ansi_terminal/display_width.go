package fake_ansi_terminal

import "github.com/redjolr/goherent/internal/utils"

// displayWidth reports the terminal-cell width of a single stored cell value.
// It delegates to the shared width table so that the fake terminal and the
// presenters that drive it agree on how many columns each glyph occupies.
func displayWidth(cell string) int {
	return utils.DisplayWidth(cell)
}
