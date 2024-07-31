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
	tp.terminal.MoveLeft(1)
	tp.terminal.Print("❌\n")
}

func (tp *TerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	tp.terminal.MoveLeft(1)
	tp.terminal.Print(ANSI_YELLOW_CIRCLE + "\n")
}

func (tp *TerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print(ctest.Output())
}

func (tp *TerminalPresenter) TestingFinishedSummary(summary TestingSummary) {

	packagesSummary := ANSI_BOLD + "\nPackages:" + ANSI_RESET_BOLD + " "
	testsSummary := ANSI_BOLD + "Tests:" + ANSI_RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ANSI_BOLD+"Time:"+ANSI_RESET_BOLD+"     %.3fs", summary.durationS)

	if summary.failedPackagesCount > 0 {
		packagesSummary += ANSI_RED + fmt.Sprintf("%d failed", summary.failedPackagesCount) + ANSI_COLOR_RESET + ", "
	}
	if summary.skippedPackagesCount > 0 {
		packagesSummary += ANSI_YELLOW + fmt.Sprintf("%d skipped", summary.skippedPackagesCount) + ANSI_COLOR_RESET + ", "
	}
	if summary.passedPackagesCount > 0 {
		packagesSummary += ANSI_GREEN + fmt.Sprintf("%d passed", summary.passedPackagesCount) + ANSI_COLOR_RESET + ", "
	}

	if summary.failedTestsCount > 0 {
		testsSummary += ANSI_RED + fmt.Sprintf("%d failed", summary.failedTestsCount) + ANSI_COLOR_RESET + ", "
	}
	if summary.skippedTestsCount > 0 {
		testsSummary += ANSI_YELLOW + fmt.Sprintf("%d skipped", summary.skippedTestsCount) + ANSI_COLOR_RESET + ", "
	}
	if summary.passedTestsCount > 0 {
		testsSummary += ANSI_GREEN + fmt.Sprintf("%d passed", summary.passedTestsCount) + ANSI_COLOR_RESET + ", "
	}
	packagesSummary += fmt.Sprintf("%d total", summary.packagesCount)
	testsSummary += fmt.Sprintf("%d total", summary.testsCount)

	tp.terminal.Print(
		fmt.Sprintf(
			packagesSummary + "\n" +
				testsSummary + "\n" +
				timeSummary + "\n" +
				"Ran all tests.",
		),
	)
}

func (tp *TerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}
