package sequential_events

import (
	"fmt"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type UnboundedTerminalPresenter struct {
	terminal terminal.Terminal
}

func NewUnboundedTerminalPresenter(term terminal.Terminal) UnboundedTerminalPresenter {
	return UnboundedTerminalPresenter{
		terminal: term,
	}
}

func (tp UnboundedTerminalPresenter) TestingStarted() {
	tp.terminal.Print("\n🚀 Starting...")
}

func (tp UnboundedTerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.terminal.Print(fmt.Sprintf("\n\n📦 %s\n", packageName))
}

func (tp UnboundedTerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print(fmt.Sprintf("\n   • %s    ⏳", ctest.Name()))
}

func (tp UnboundedTerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	tp.terminal.MoveLeft(utils.DisplayWidth("⏳"))
	tp.terminal.Print("✅" + formatDurationLabel(duration) + "\n")
}

func (tp UnboundedTerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	tp.terminal.MoveLeft(utils.DisplayWidth("⏳"))
	tp.terminal.Print("❌" + formatDurationLabel(duration) + "\n")
}

func (tp UnboundedTerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	tp.terminal.MoveLeft(utils.DisplayWidth("⏳"))
	tp.terminal.Print("⏩\n")
}

func (tp UnboundedTerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print("\n" + ctest.Output())
}

func (tp UnboundedTerminalPresenter) FailedTestsList(failedPackages []*ctests_tracker.PackageUnderTest) {
	tp.terminal.Print("\n\nFailed tests:")
	for _, packageUt := range failedPackages {
		tp.terminal.Print("\n\n❌ " + ansi_escape.BOLD + ansi_escape.RED + packageUt.Name() + ansi_escape.COLOR_RESET)
		if packageUt.HasBuildFailure() {
			tp.terminal.Print("  " + ansi_escape.RED + "[build failed]" + ansi_escape.COLOR_RESET)
			if packageUt.BuildOutput() != "" {
				tp.terminal.Print("\n\n  " + packageUt.BuildOutput())
			}
		}
		for _, ctest := range packageUt.FailedCtests() {
			tp.terminal.Print("\n\n  " + ansi_escape.RED + "● " + ctest.Name() + ansi_escape.COLOR_RESET)
			if ctest.ContainsOutput() {
				tp.terminal.Print("\n\n  " + ctest.Output())
			}
		}
		if packageUt.HasOutputOfParentTests() {
			tp.terminal.Print("\n\n" + packageUt.ParentTestsOutput())
		}
	}
	tp.terminal.Print("\n")
}

func (tp UnboundedTerminalPresenter) SlowestTests(tests []*ctests_tracker.Ctest) {
	tp.terminal.Print("\n\n" + buildSlowestTestsReport(tests) + "\n\n")
}

// Tick is a no-op: piped/non-TTY output has no animated live region.
func (tp UnboundedTerminalPresenter) Tick() {}

func (tp UnboundedTerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}

func (tp UnboundedTerminalPresenter) Print(output string) {
	tp.terminal.Print(output)
}

func (tp UnboundedTerminalPresenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {

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
	testsSummary += fmt.Sprintf("%d total", summary.TestsCount) + passRateLabel(summary)

	tp.terminal.Print("\n" + testingVerdictHeadline(summary) + "\n")
	tp.terminal.Print(
		packagesSummary + "\n" +
			testsSummary + "\n" +
			timeSummary + "\n" +
			"Ran all tests.",
	)
}
