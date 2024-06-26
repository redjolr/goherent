package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/console"
)

type TerminalPresenter struct {
	console   *console.Console
	header    *console.Textblock
	testsList *console.UnorderedList
}

func NewTerminalPresenter(console *console.Console) TerminalPresenter {
	return TerminalPresenter{
		console:   console,
		header:    nil,
		testsList: nil,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.header = tp.console.NewTextBlock(
		fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")),
	)
	tp.console.Render()
}

func (tp *TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	if tp.testsList == nil {
		testsList := console.NewUnorderedList(fmt.Sprintf("📦⏳ %s\n", packageName))
		tp.testsList = &testsList
	}
	tp.console.Render()
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
