package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/console"
)

type TerminalPresenter struct {
	terminal  *console.Terminal
	header    *console.Textblock
	testsList *console.UnorderedList
}

func NewTerminalPresenter(terminal *console.Terminal) TerminalPresenter {
	return TerminalPresenter{
		terminal:  terminal,
		header:    nil,
		testsList: nil,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.header = tp.terminal.NewTextBlock(
		fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")),
	)
	tp.testsList = tp.terminal.NewUnorderedList("All tests:")
	tp.terminal.Render()
}

func (tp *TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.testsList.NewItem(
		fmt.Sprintf("📦⏳ %s\n", packageName),
	)

	tp.terminal.Render()
}

func (pressenter *TerminalPresenter) CtestStartedRunning(testName string) {
	fmt.Printf("\t⏳ %s\n\n", testName)
}

func (presenter *TerminalPresenter) CtestPassed(testName string, testDuration float64) {
	fmt.Printf("\t✅ %s\n \t%f\n\n", testName, testDuration)
}

func (presenter *TerminalPresenter) CtestFailed(testName string, testDuration float64) {
	fmt.Printf("\t❌ %s\n \t%f\n\n", testName, testDuration)
}

func (presenter *TerminalPresenter) CtestOutput(testName string, packageName string, output string) {
	fmt.Printf("\t %s\n\n %s\n", testName, output)
}
