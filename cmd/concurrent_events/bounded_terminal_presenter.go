package concurrent_events

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

func (tp *BoundedTerminalPresenter) PackageStarted(packageUt ctests_tracker.PackageUnderTest) {
}

func (tp *BoundedTerminalPresenter) Error() {
}

func (tp *BoundedTerminalPresenter) EraseScreen() {

}

func (tp *BoundedTerminalPresenter) Packages(packages []*ctests_tracker.PackageUnderTest) {

}
