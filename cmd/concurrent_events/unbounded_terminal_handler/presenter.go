package unbounded_terminal_handler

import (
	"fmt"

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

func (p *Presenter) PackageStarted(packageUt ctests_tracker.PackageUnderTest) {
	p.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
}

func (p *Presenter) Error() {
	p.terminal.Print("\n\n❗ Error.")
}

func (p *Presenter) EraseScreen() {
	p.terminal.Print(ansi_escape.ERASE_SCREEN)
	p.terminal.Print(ansi_escape.CURSOR_TO_HOME)
}

func (p *Presenter) Packages(packages []*ctests_tracker.PackageUnderTest) {
	for _, packageUt := range packages {
		if packageUt.TestsAreRunning() {
			p.terminal.Print(fmt.Sprintf("\n⏳ %s", packageUt.Name()))
		}
		if packageUt.HasPassed() {
			p.terminal.Print(fmt.Sprintf("\n✅ %s", packageUt.Name()))
		}
		if packageUt.IsSkipped() {
			p.terminal.Print(fmt.Sprintf("\n⏩ %s", packageUt.Name()))
		}
		if packageUt.HasAtLeastOneFailedTest() {
			p.terminal.Print(fmt.Sprintf("\n❌ %s", packageUt.Name()))
		}
	}
}
