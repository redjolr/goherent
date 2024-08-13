package sequential_events

import (
	"fmt"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
)

type UnboundedTerminalPresenter struct {
	terminal terminal.Terminal
}

func NewUnboundedTerminalPresenter(term terminal.Terminal) UnboundedTerminalPresenter {
	return UnboundedTerminalPresenter{
		terminal: term,
	}
}

func (tp UnboundedTerminalPresenter) TestingStarted() {
	tp.terminal.Print("\n🚀 Starting...")
}

func (tp UnboundedTerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.terminal.Print(fmt.Sprintf("\n\n📦 %s\n", packageName))
}

func (tp UnboundedTerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print(fmt.Sprintf("\n   • %s    ⏳", ctest.Name()))
}

func (tp UnboundedTerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	tp.terminal.MoveLeft(1)
	tp.terminal.Print("✅\n")
}

func (tp UnboundedTerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	tp.terminal.MoveLeft(1)
	tp.terminal.Print("❌\n")
}

func (tp UnboundedTerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	tp.terminal.MoveLeft(1)
	tp.terminal.Print("⏩\n")
}

func (tp UnboundedTerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print("\n" + ctest.Output())
}

func (tp UnboundedTerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}
