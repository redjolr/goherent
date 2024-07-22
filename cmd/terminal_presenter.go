package cmd

import (
	"fmt"
	"time"

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

func (pressenter *TerminalPresenter) CtestStartedRunning(testName string) {
	fmt.Printf("\t⏳ %s\n\n", testName)
}

func (tp *TerminalPresenter) CtestPassed(testName string, testDuration float64) {
	var testsList *elements.UnorderedList
	if tp.console.HasElementWithId(testsListId) {
		existingElement := tp.console.GetElementWithId(testsListId)
		testsList = existingElement.(*elements.UnorderedList)
	} else {
		panic("Test list does not exist")
	}
	testsList.NewItem(fmt.Sprintf("✅ %s", testName))
	testsList.NewItem(fmt.Sprintf("  %.2fs", testDuration))
	tp.console.Render()
}

func (tp *TerminalPresenter) CtestFailed(testName string, testDuration float64) {
	var testsList *elements.UnorderedList
	if tp.console.HasElementWithId(testsListId) {
		existingElement := tp.console.GetElementWithId(testsListId)
		testsList = existingElement.(*elements.UnorderedList)
	} else {
		panic("Test list does not exist")
	}
	testsList.NewItem(fmt.Sprintf("❌ %s", testName))
	testsList.NewItem(fmt.Sprintf("  %.2fs", testDuration))
	tp.console.Render()
}

func (presenter *TerminalPresenter) CtestOutput(testName string, packageName string, output string) {
	fmt.Printf("\t %s\n\n %s\n", testName, output)
}
