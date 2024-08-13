package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
)

type BoundedTerminalPresenter struct {
	terminal terminal.Terminal
}

func NewBoundedTerminalPresenter(term terminal.Terminal) BoundedTerminalPresenter {
	return BoundedTerminalPresenter{
		terminal: term,
	}
}

func (tp *BoundedTerminalPresenter) DisplayCurrentState(runningTests []ctests_tracker.PackageUnderTest) {

}
