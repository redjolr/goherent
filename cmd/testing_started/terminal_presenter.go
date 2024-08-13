package testing_started

import (
	"github.com/redjolr/goherent/terminal"
)

type TerminalPresenter struct {
	terminal terminal.Terminal
}

func NewTerminalPresenter(term terminal.Terminal) TerminalPresenter {
	return TerminalPresenter{
		terminal: term,
	}
}

func (tp *TerminalPresenter) TestingStarted() {
	tp.terminal.Print("\n🚀 Starting...")
}
