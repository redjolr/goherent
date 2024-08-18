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
	packagesThatFitInTerminal := packagesUt[0:p.terminal.Height()]
	for i, packageut := range packagesThatFitInTerminal {
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
