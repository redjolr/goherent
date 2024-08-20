package bounded_terminal_handler

import (
	"fmt"

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

func (p *Presenter) DisplayPackages(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
	testingSummary ctests_tracker.TestingSummary,
) {
	if p.terminal.Height() <= 5 {
		p.displayPackagesInSmallTerminal(runningPackages, finishedPackages)
	} else {
		p.displayPackagesInLargeTerminal(runningPackages, finishedPackages, testingSummary)
	}
}

func (p *Presenter) displayPackagesInLargeTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
	testingSummary ctests_tracker.TestingSummary,
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

	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := ansi_escape.BOLD + "Time:" + ansi_escape.RESET_BOLD + "     0.000s"
	runningPackagesCount := len(runningPackages)
	passedPackagesCount := testingSummary.PassedPackagesCount
	failedPackagesCount := testingSummary.FailedPackagesCount
	skippedPackagesCount := testingSummary.SkippedPackagesCount

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
	p.terminal.Print(
		"\n\n" +
			packagesSummary + "\n" +
			testsSummary + "0 running\n" +
			timeSummary,
	)
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

func (p *Presenter) EraseScreen() {
	p.terminal.Print(ansi_escape.ERASE_SCREEN)
	p.terminal.Print(ansi_escape.CURSOR_TO_HOME)
}

func (p *Presenter) Error() {
	p.terminal.Print("\n\n❗ Error.")
}
