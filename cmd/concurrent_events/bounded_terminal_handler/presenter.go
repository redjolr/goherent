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
	passedPackages []*ctests_tracker.PackageUnderTest,
	failedPackages []*ctests_tracker.PackageUnderTest,
) {
	if p.terminal.Height() <= 5 {
		p.displayPackagesInSmallTerminal(runningPackages, passedPackages, failedPackages)
	} else {
		p.displayPackagesInLargeTerminal(runningPackages, passedPackages)
	}
}

func (p *Presenter) displayPackagesInLargeTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	passedPackages []*ctests_tracker.PackageUnderTest,
) {
	packagesThatFitInTerminalCount := p.terminal.Height() - SummaryLineCount
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), packagesThatFitInTerminalCount)]

	if len(runningPackages) < packagesThatFitInTerminalCount && len(passedPackages) > 0 {
		showPassedPackagesCount := min(len(passedPackages), packagesThatFitInTerminalCount-len(runningPackages))
		latestPassedPackages := passedPackages[len(passedPackages)-showPassedPackagesCount:]
		for i, packageUt := range latestPassedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			p.terminal.Print("✅ " + packageUt.Name())

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

	if p.terminal.Height() > 5 {
		packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
		testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
		timeSummary := ansi_escape.BOLD + "Time:" + ansi_escape.RESET_BOLD + "     0.000s"
		runningPackagesCount := len(runningPackages)
		passedPackagesCount := len(passedPackages)

		packagesSummary += fmt.Sprintf("%d running", runningPackagesCount)
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
}

func (p *Presenter) displayPackagesInSmallTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	passedPackages []*ctests_tracker.PackageUnderTest,
	failedPackages []*ctests_tracker.PackageUnderTest,
) {
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), p.terminal.Height())]

	if len(runningPackages) < p.terminal.Height() && len(passedPackages) > 0 {
		showPassedPackagesCount := min(len(passedPackages), p.terminal.Height()-len(runningPackages))
		latestPassedPackages := passedPackages[len(passedPackages)-showPassedPackagesCount:]
		for i, packageut := range latestPassedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			p.terminal.Print("✅ " + packageut.Name())
		}
		if len(runningPackages) > 0 {
			p.terminal.Print("\n")
		}
	}

	if len(runningPackages) < p.terminal.Height() && len(failedPackages) > 0 {
		showFailedPackagesCount := min(len(failedPackages), p.terminal.Height()-len(runningPackages))
		latestFailedPackages := failedPackages[len(failedPackages)-showFailedPackagesCount:]
		for i, packageut := range latestFailedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			p.terminal.Print("❌ " + packageut.Name())
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
