package bounded_terminal_handler

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type Presenter struct {
	terminal terminal.Terminal
}

const SummaryLineCount int = 4

func NewPresenter(term terminal.Terminal) Presenter {
	return Presenter{
		terminal: term,
	}
}

func (p *Presenter) IsViewPortLarge() bool {
	return p.terminal.Height() > 5
}

func (p *Presenter) TestingStarted() {
	p.terminal.Print("\n🚀 Starting...")
}

func (p *Presenter) DisplayPackages(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	if p.terminal.Height() <= 5 {
		p.displayPackagesInSmallTerminal(runningPackages, finishedPackages)
	} else {
		p.displayPackagesInLargeTerminal(runningPackages, finishedPackages)
	}
}

func (p *Presenter) displayPackagesInLargeTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	packagesThatFitInTerminalCount := p.terminal.Height() - SummaryLineCount
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), packagesThatFitInTerminalCount)]

	if len(runningPackages) < packagesThatFitInTerminalCount && len(finishedPackages) > 0 {
		showFinishedPackagesCount := min(len(finishedPackages), packagesThatFitInTerminalCount-len(runningPackages))
		latestFinishedPackages := finishedPackages[len(finishedPackages)-showFinishedPackagesCount:]
		for i, packageUt := range latestFinishedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			if packageUt.HasPassed() {
				p.terminal.Print("✅ " + packageUt.Name())
			} else if packageUt.HasAtLeastOneFailedTest() {
				p.terminal.Print("❌ " + packageUt.Name())
			} else if packageUt.IsSkipped() {
				p.terminal.Print("⏩ " + packageUt.Name())
			}
		}
		if len(runningPackages) > 0 {
			p.terminal.Print("\n")
		}
	}

	for i, packageUt := range runningPackagesThatFitInTerminal {
		if i != 0 {
			p.terminal.Print("\n")
		}
		p.terminal.Print("⏳ " + packageUt.Name())
	}
}

func (p *Presenter) displayPackagesInSmallTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), p.terminal.Height())]

	if len(runningPackages) < p.terminal.Height() && len(finishedPackages) > 0 {
		showFinishedPackagesCount := min(len(finishedPackages), p.terminal.Height()-len(runningPackages))
		latestFinishedPackages := finishedPackages[len(finishedPackages)-showFinishedPackagesCount:]
		for i, packageUt := range latestFinishedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			if packageUt.HasPassed() {
				p.terminal.Print("✅ " + packageUt.Name())
			} else if packageUt.HasAtLeastOneFailedTest() {
				p.terminal.Print("❌ " + packageUt.Name())
			} else if packageUt.IsSkipped() {
				p.terminal.Print("⏩ " + packageUt.Name())
			}
		}
		if len(runningPackages) > 0 {
			p.terminal.Print("\n")
		}
	}

	for i, packageut := range runningPackagesThatFitInTerminal {
		if i != 0 {
			p.terminal.Print("\n")
		}
		p.terminal.Print("⏳ " + packageut.Name())
	}
}

func (p *Presenter) RunningTestsSummary(testingSummary ctests_tracker.TestingSummary) {
	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := ansi_escape.BOLD + "Time:" + ansi_escape.RESET_BOLD + "     0.000s"
	runningPackagesCount := testingSummary.RunningPackagesCount
	passedPackagesCount := testingSummary.PassedPackagesCount
	failedPackagesCount := testingSummary.FailedPackagesCount
	skippedPackagesCount := testingSummary.SkippedPackagesCount
	passedTestsCount := testingSummary.PassedTestsCount
	failedTestsCount := testingSummary.FailedTestsCount
	skippedTestsCount := testingSummary.SkippedTestsCount

	packagesSummary += fmt.Sprintf("%d running", runningPackagesCount)

	if failedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.RED + fmt.Sprintf("%d failed", failedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if skippedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.YELLOW + fmt.Sprintf("%d skipped", skippedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if passedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.GREEN + fmt.Sprintf("%d passed", passedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if failedTestsCount > 0 {
		testsSummary += ansi_escape.RED + fmt.Sprintf("%d failed", failedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if skippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW + fmt.Sprintf("%d skipped", skippedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if passedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN + fmt.Sprintf("%d passed", passedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}

	testsSummary, _ = strings.CutSuffix(testsSummary, ", ")

	summary := "\n\n" +
		packagesSummary + "\n" +
		testsSummary + "\n" +
		timeSummary

	p.terminal.Print(summary)
}

func (p *Presenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {
	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     %.3fs", summary.DurationS)

	if summary.FailedPackagesCount > 0 {
		packagesSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.FailedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedPackagesCount > 0 {
		packagesSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.SkippedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedPackagesCount > 0 {
		packagesSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.PassedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}

	if summary.FailedTestsCount > 0 {
		testsSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.FailedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.SkippedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.PassedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	packagesSummary += fmt.Sprintf("%d total", summary.PackagesCount)
	testsSummary += fmt.Sprintf("%d total", summary.TestsCount)

	p.terminal.Print(
		fmt.Sprintf(
			"\n\n" + packagesSummary + "\n" +
				testsSummary + "\n" +
				timeSummary + "\n" +
				"Ran all tests.",
		),
	)
}

func (p *Presenter) EraseScreen() {
	p.terminal.Print(ansi_escape.CURSOR_TO_HOME)
	p.terminal.Print(ansi_escape.ERASE_SCREEN)
}

func (p *Presenter) Error() {
	p.terminal.Print("\n\n❗ Error.")
}
