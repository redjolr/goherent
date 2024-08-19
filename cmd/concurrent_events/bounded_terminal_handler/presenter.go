package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type Presenter struct {
	terminal terminal.Terminal
}

func NewPresenter(term terminal.Terminal) Presenter {
	return Presenter{
		terminal: term,
	}
}

func (p *Presenter) DisplayPackages(packagesUt []*ctests_tracker.PackageUnderTest) {

	var packagesThatFitInTerminalCount int
	if p.terminal.Height() < 5 {
		packagesThatFitInTerminalCount = p.terminal.Height()
	} else {
		packagesThatFitInTerminalCount = p.terminal.Height() - 4
	}
	packagesThatFitInTerminal := packagesUt[0:min(len(packagesUt), packagesThatFitInTerminalCount)]
	for i, packageut := range packagesThatFitInTerminal {
		if i != 0 {
			p.terminal.Print("\n")
		}
		p.terminal.Print("⏳ " + packageut.Name())
	}

	if p.terminal.Height() >= 5 {
		packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
		testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
		timeSummary := ansi_escape.BOLD + "Time:" + ansi_escape.RESET_BOLD + "     0.000s"
		runningPackagesCount := len(packagesUt)
		p.terminal.Printf(
			"\n"+
				packagesSummary+"%d running\n"+
				testsSummary+"0 running\n"+
				timeSummary,
			runningPackagesCount,
		)
	}

}

func (p *Presenter) EraseScreen() {
	p.terminal.Print(ansi_escape.ERASE_SCREEN)
	p.terminal.Print(ansi_escape.CURSOR_TO_HOME)
}
