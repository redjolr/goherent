package sequential_events

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/liveregion"
)

// LiveTerminalPresenter renders a sequential run as a committed area (the run
// banner, package headers, and finished test lines, which scroll naturally) plus
// a live block at the bottom: the currently-running test and a running tally
// footer that updates in place. It carries no cursor bookkeeping of its own —
// LiveRegion owns that — so it just composes strings.
type LiveTerminalPresenter struct {
	region  *liveregion.LiveRegion
	passed  int
	failed  int
	skipped int
	running string // the running test's line ("⏳ name"), or "" when none is running
}

func NewLiveTerminalPresenter(term terminal.Terminal) *LiveTerminalPresenter {
	return &LiveTerminalPresenter{region: liveregion.New(term)}
}

func (p *LiveTerminalPresenter) TestingStarted() {
	p.region.Render("🚀 Starting...", p.liveBlock())
}

func (p *LiveTerminalPresenter) PackageTestsStartedRunning(packageName string) {
	p.region.Render("\n📦 "+packageName, p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	p.running = "⏳ " + ctest.Name()
	p.region.SetLive(p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	p.passed++
	p.running = ""
	p.region.Render(committedResult("✅", ctest.Name(), formatDurationLabel(duration)), p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	p.failed++
	p.running = ""
	p.region.Render(committedResult("❌", ctest.Name(), formatDurationLabel(duration)), p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	p.skipped++
	p.running = ""
	p.region.Render(committedResult("⏩", ctest.Name(), ""), p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	p.region.Render("\n"+ctest.Output(), p.liveBlock())
}

func (p *LiveTerminalPresenter) Print(output string) {
	p.region.Render(output, p.liveBlock())
}

func (p *LiveTerminalPresenter) Error() {
	p.region.Render("\n\n❗ Error.", p.liveBlock())
}

func (p *LiveTerminalPresenter) FailedTestsList(failedPackages []*ctests_tracker.PackageUnderTest) {
	// Testing is finishing: drop the live footer and commit the failed-tests list.
	p.region.Render(buildFailedTestsList(failedPackages), "")
}

func (p *LiveTerminalPresenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {
	// Drop the live footer and commit the final summary as permanent output.
	p.region.Render(buildFinalSummary(summary), "")
}

// liveBlock is the content of the live region: the running test (if any) and the
// running tally footer, separated by a blank line.
func (p *LiveTerminalPresenter) liveBlock() string {
	f := p.footer()
	switch {
	case p.running == "":
		return f
	case f == "":
		return p.running
	default:
		return p.running + "\n\n" + f
	}
}

func (p *LiveTerminalPresenter) footer() string {
	if p.passed+p.failed+p.skipped == 0 {
		return ""
	}
	parts := []string{ansi_escape.GREEN + fmt.Sprintf("%d passed", p.passed) + ansi_escape.COLOR_RESET}
	if p.failed > 0 {
		parts = append(parts, ansi_escape.RED+fmt.Sprintf("%d failed", p.failed)+ansi_escape.COLOR_RESET)
	}
	if p.skipped > 0 {
		parts = append(parts, ansi_escape.YELLOW+fmt.Sprintf("%d skipped", p.skipped)+ansi_escape.COLOR_RESET)
	}
	return strings.Join(parts, " · ")
}

// committedResult formats a finished test line, placing the duration on the
// first line of a (possibly multi-line) test name.
func committedResult(icon, name, durationLabel string) string {
	lines := utils.SplitStringByNewLine(name)
	out := icon + " " + lines[0] + durationLabel
	if len(lines) > 1 {
		out += "\n" + strings.Join(lines[1:], "\n")
	}
	return out
}

func buildFailedTestsList(failedPackages []*ctests_tracker.PackageUnderTest) string {
	out := "Failed tests:"
	for _, packageUt := range failedPackages {
		out += "\n\n❌ " + ansi_escape.BOLD + ansi_escape.RED + packageUt.Name() + ansi_escape.COLOR_RESET
		for _, ctest := range packageUt.FailedCtests() {
			out += "\n\n  " + ansi_escape.RED + "● " + ctest.Name() + ansi_escape.COLOR_RESET
			if ctest.ContainsOutput() {
				out += "\n\n  " + ctest.Output()
			}
		}
		if packageUt.HasOutputOfParentTests() {
			out += "\n\n" + packageUt.ParentTestsOutput()
		}
	}
	return out
}

func buildFinalSummary(summary ctests_tracker.TestingSummary) string {
	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     %.3fs", summary.DurationS)

	if summary.FailedPackagesCount > 0 {
		packagesSummary += ansi_escape.RED + fmt.Sprintf("%d failed", summary.FailedPackagesCount) + ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedPackagesCount > 0 {
		packagesSummary += ansi_escape.YELLOW + fmt.Sprintf("%d skipped", summary.SkippedPackagesCount) + ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedPackagesCount > 0 {
		packagesSummary += ansi_escape.GREEN + fmt.Sprintf("%d passed", summary.PassedPackagesCount) + ansi_escape.COLOR_RESET + ", "
	}
	if summary.FailedTestsCount > 0 {
		testsSummary += ansi_escape.RED + fmt.Sprintf("%d failed", summary.FailedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW + fmt.Sprintf("%d skipped", summary.SkippedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN + fmt.Sprintf("%d passed", summary.PassedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	packagesSummary += fmt.Sprintf("%d total", summary.PackagesCount)
	testsSummary += fmt.Sprintf("%d total", summary.TestsCount) + passRateLabel(summary)

	return testingVerdictHeadline(summary) + "\n" +
		packagesSummary + "\n" +
		testsSummary + "\n" +
		timeSummary + "\n" +
		"Ran all tests."
}
