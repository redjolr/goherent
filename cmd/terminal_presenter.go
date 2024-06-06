package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/console"
)

type TerminalPresenter struct {
	container *console.Container
	header    *console.Textblock
	testsList *console.UnorderedList
}

func NewTerminalPresenter(container *console.Container) TerminalPresenter {
	return TerminalPresenter{
		container: container,
		header:    nil,
		testsList: nil,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.header = tp.container.NewTextBlock(
		fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")),
	)
	tp.testsList = tp.container.NewUnorderedList("All tests:")
	tp.container.Render()
}

func (tp *TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.testsList.NewItem(
		fmt.Sprintf("📦⏳ %s\n", packageName),
	)

	tp.container.Render()
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
