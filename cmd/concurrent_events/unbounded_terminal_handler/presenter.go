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

func (p *Presenter) TestingStarted() {
	p.terminal.Print("\n🚀 Starting...")
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

func (p *Presenter) TestingFinishedSummaryLabel() {
	p.terminal.Print("\n\n\n\n\n📋 Tests summary.\n\n")
}

func (p Presenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {

	packagesSummary := ansi_escape.BOLD + "\nPackages:" + ansi_escape.RESET_BOLD + " "
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
			packagesSummary + "\n" +
				testsSummary + "\n" +
				timeSummary + "\n" +
				"Ran all tests.",
		),
	)
}
