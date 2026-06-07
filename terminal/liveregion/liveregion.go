// Package liveregion renders a terminal UI as a permanent "committed" area that
// scrolls naturally, plus a "live" block at the bottom that is rewritten in
// place on every update (a status footer, a spinner, the currently-running
// item, …).
//
// It redraws the live block using only relative cursor moves, a carriage
// return, and erase-to-end-of-screen — never absolute cursor save/restore —
// because absolute save/restore does not survive scrolling (a real terminal
// restores to the same screen cell, which now shows scrolled content). Relative
// moves scroll together with the content, so the live block stays correct even
// as the committed area scrolls off the top.
package liveregion

import (
	"strings"

	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type LiveRegion struct {
	terminal terminal.Terminal
	live     string // content of the live block currently on screen
	hasLive  bool
}

func New(t terminal.Terminal) *LiveRegion {
	return &LiveRegion{terminal: t}
}

// Render commits `committed` (printed once, permanently, above the live block —
// it scrolls with the rest of the output and is never touched again), then
// (re)draws the live block as `live`. Pass committed == "" to update only the
// live block in place.
func (r *LiveRegion) Render(committed, live string) {
	if r.hasLive {
		// Move to the top-left of the live block and clear it (and anything
		// below). Only the live block's own lines are moved over, so committed
		// content above is untouched.
		if nl := strings.Count(r.live, "\n"); nl > 0 {
			r.terminal.MoveUp(nl)
		}
		r.terminal.Print("\r")
		r.terminal.Print(ansi_escape.ERASE_SCREEN) // \033[0J: cursor -> end of screen
	}
	if committed != "" {
		r.terminal.Print(committed + "\n")
	}
	r.terminal.Print(live)
	r.live = live
	r.hasLive = true
}

// Commit is shorthand for committing a line while leaving the live block as-is.
func (r *LiveRegion) Commit(committed string) {
	r.Render(committed, r.live)
}

// SetLive replaces only the live block.
func (r *LiveRegion) SetLive(live string) {
	r.Render("", live)
}
