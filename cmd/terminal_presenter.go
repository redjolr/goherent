package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/console"
	"github.com/redjolr/goherent/console/elements"
)

const testsListId string = "testsList"
const startingTestsTextblockId string = "startingTestsTextBlock"

type TerminalPresenter struct {
	console *console.Console
}

func NewTerminalPresenter(console *console.Console) TerminalPresenter {
	return TerminalPresenter{
		console: console,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.console.NewTextBlock(
		startingTestsTextblockId,
		fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")),
	)
	tp.console.Render()
}

func (tp *TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	if !tp.console.HasElementWithId(testsListId) {
		tp.console.NewUnorderedList(testsListId, fmt.Sprintf("📦⏳ %s", packageName))
	}
	tp.console.Render()
}

func (tp *TerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	testsListElement := tp.console.GetElementWithId(testsListId)
	if testsListElement == nil {
		return
	}
	testsList := testsListElement.(*elements.UnorderedList)
	testsList.NewItem(ctest.Id(), fmt.Sprintf("⏳ %s", ctest.Name()))
	tp.console.Render()
}

func (tp *TerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, testDuration float64) {
	var testsList *elements.UnorderedList
	if tp.console.HasElementWithId(testsListId) {
		existingElement := tp.console.GetElementWithId(testsListId)
		testsList = existingElement.(*elements.UnorderedList)
	} else {
		panic("Test list does not exist")
	}
	listItem := testsList.FindItemById(ctest.Id())
	if listItem == nil {
		testsList.NewItem(ctest.Id(), fmt.Sprintf("✅ %s\n  %.2fs", ctest.Name(), testDuration))
	}
	listItem.Edit(fmt.Sprintf("✅ %s\n  %.2fs", ctest.Name(), testDuration))

	tp.console.Render()
}

func (tp *TerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, testDuration float64) {
	var testsList *elements.UnorderedList
	if tp.console.HasElementWithId(testsListId) {
		existingElement := tp.console.GetElementWithId(testsListId)
		testsList = existingElement.(*elements.UnorderedList)
	} else {
		panic("Test list does not exist")
	}
	testsList.NewItem(ctest.Id(), fmt.Sprintf("❌ %s\n  %.2fs", ctest.Name(), testDuration))
	tp.console.Render()
}

func (tp *TerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	testsListElement := tp.console.GetElementWithId(testsListId)
	if testsListElement == nil {
		return
	}
	testsList := testsListElement.(*elements.UnorderedList)
	testItem := testsList.FindItemById(ctest.Id())

	if testItem == nil {
		return
	}

	testItem.Edit(
		testItem.Text() + fmt.Sprintf("\n%s", ctest.Output()),
	)
	tp.console.Render()
}
