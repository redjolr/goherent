package concurrent_events

import (
	"fmt"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type UnboundedTerminalPresenter struct {
	terminal terminal.Terminal
}

func NewUnboundedTerminalPresenter(term terminal.Terminal) UnboundedTerminalPresenter {
	return UnboundedTerminalPresenter{
		terminal: term,
	}
}

func (tp *UnboundedTerminalPresenter) PackageStarted(packageUt ctests_tracker.PackageUnderTest) {
	tp.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
}

func (tp *UnboundedTerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}

func (tp *UnboundedTerminalPresenter) EraseScreen() {
	tp.terminal.Print(ansi_escape.ERASE_SCREEN)
	tp.terminal.Print(ansi_escape.CURSOR_TO_HOME)
}

func (tp *UnboundedTerminalPresenter) Packages(packages []*ctests_tracker.PackageUnderTest) {
	for _, packageUt := range packages {
		if packageUt.TestsAreRunning() {
			tp.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
		}
		if packageUt.HasPassed() {
			tp.terminal.Print(fmt.Sprintf("\n✅ %s", packageUt.Name()))
		}
		if packageUt.IsSkipped() {
			tp.terminal.Print(fmt.Sprintf("\n⏩ %s", packageUt.Name()))
		}
		if packageUt.HasAtLeastOneFailedTest() {
			tp.terminal.Print(fmt.Sprintf("\n❌ %s", packageUt.Name()))
		}
	}
}
