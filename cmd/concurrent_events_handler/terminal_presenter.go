package concurrent_events_handler

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

const testsListId string = "testsList"
const startingTestsTextblockId string = "startingTestsTextBlock"

type TerminalPresenter struct {
	terminal terminal.Terminal
}

func NewTerminalPresenter(term terminal.Terminal) TerminalPresenter {
	return TerminalPresenter{
		terminal: term,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.terminal.Print(fmt.Sprintf("\n🚀 Starting... %s", timestamp.Format("2006-01-02 15:04:05.000")))
}

func (tp *TerminalPresenter) PackageStarted(packageUt ctests_tracker.PackageUnderTest) {
	tp.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
}

func (tp *TerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}

func (tp *TerminalPresenter) EraseScreen() {
	tp.terminal.Print(ansi_escape.ERASE_SCREEN)
	tp.terminal.Print(ansi_escape.CURSOR_TO_HOME)
}

func (tp *TerminalPresenter) Packages(packages []*ctests_tracker.PackageUnderTest) {
	for _, packageUt := range packages {
		if packageUt.TestsAreRunning() {
			tp.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
		}
		if packageUt.HasPassed() {
			tp.terminal.Print(fmt.Sprintf("\n✅ %s", packageUt.Name()))
		}
		if packageUt.HasAtLeastOneFailedTest() {
			tp.terminal.Print(fmt.Sprintf("\n❌ %s", packageUt.Name()))
		}
	}
}
