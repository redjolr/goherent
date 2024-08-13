package testing_finished

import (
	"fmt"

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

func (tp *TerminalPresenter) TestingFinishedSummary(summary TestingSummary) {

	packagesSummary := ansi_escape.BOLD + "\nPackages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     %.3fs", summary.durationS)

	if summary.failedPackagesCount > 0 {
		packagesSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.failedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.skippedPackagesCount > 0 {
		packagesSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.skippedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.passedPackagesCount > 0 {
		packagesSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.passedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}

	if summary.failedTestsCount > 0 {
		testsSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.failedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.skippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.skippedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.passedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.passedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
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
