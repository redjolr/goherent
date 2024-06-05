package cmd

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/console"
)

type TerminalPresenter struct {
	terminal *console.Terminal
}

func NewTerminalPresenter(terminal *console.Terminal) TerminalPresenter {
	return TerminalPresenter{
		terminal: terminal,
	}
}

func (presenter TerminalPresenter) TestingStarted(timestamp time.Time) {
	presenter.terminal.NewTextBlock(
		fmt.Sprintf("\n🚀 Starting... %s\n\n", timestamp.Format("2006-01-02 15:04:05.000")),
	)
	presenter.terminal.NewUnorderedList("All tests:")
	presenter.terminal.Render()
}

func (presenter TerminalPresenter) PackageTestsStartedRunning(packageName string) {
	fmt.Printf("📦⏳ %s\n", packageName)
}

func (pressenter TerminalPresenter) CtestStartedRunning(testName string) {
	fmt.Printf("\t⏳ %s\n\n", testName)
}

func (presenter TerminalPresenter) CtestPassed(testName string, testDuration float64) {
	fmt.Printf("\t✅ %s\n \t%f\n\n", testName, testDuration)
}

func (presenter TerminalPresenter) CtestFailed(testName string, testDuration float64) {
	fmt.Printf("\t❌ %s\n \t%f\n\n", testName, testDuration)
}

func (presenter TerminalPresenter) CtestOutput(testName string, packageName string, output string) {
	fmt.Printf("\t %s\n\n %s\n", testName, output)
}
