package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/console/terminal"
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
	tp.terminal.Print(fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")))
}

func (tp *TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.terminal.Print(fmt.Sprintf("📦 %s\n", packageName))
}

func (tp *TerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print(fmt.Sprintf("\n   • %s    ⏳", ctest.Name()))
}

func (tp *TerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	tp.terminal.MoveLeft(1)
	tp.terminal.Print("✅\n")
}

func (tp *TerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	fmt.Printf("\033[%dD", 1)
	fmt.Printf("❌\n")
}

func (tp *TerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	// testsListElement := tp.console.GetElementWithId(testsListId)
	// if testsListElement == nil {
	// 	return
	// }
	// testsList := testsListElement.(*elements.UnorderedList)
	// testItem := testsList.FindItemById(ctest.Id())

	// if testItem == nil {
	// 	return
	// }

	// testItem.Edit(
	// 	testItem.Text() + fmt.Sprintf("\n%s", ctest.Output()),
	// )
	// tp.console.Render()
}

func (tp *TerminalPresenter) GenericError() {
	tp.terminal.Print("\n\n❗ Error.")
}
