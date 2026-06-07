package sequential_events

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/liveregion"
	"github.com/redjolr/goherent/terminal/spinner"
)

// LiveTerminalPresenter renders a sequential run as a committed area (the run
// banner, package headers, and finished test lines, which scroll naturally) plus
// a live block at the bottom: the currently-running test and a running tally
// footer that updates in place. It carries no cursor bookkeeping of its own —
// LiveRegion owns that — so it just composes strings.
type LiveTerminalPresenter struct {
	region       *liveregion.LiveRegion
	passed       int
	failed       int
	skipped      int
	runningName  string // raw name of the running test, or "" when none is running
	spinnerFrame int
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
	p.runningName = ctest.Name()
	p.region.SetLive(p.liveBlock())
}

// Tick advances the running test's spinner and redraws the live block in place.
// It is a no-op when no test is running.
func (p *LiveTerminalPresenter) Tick() {
	if p.runningName == "" {
		return
	}
	p.spinnerFrame++
	p.region.SetLive(p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	p.passed++
	p.runningName = ""
	p.region.Render("\n"+testLine("✅", ctest.Name(), formatDurationLabel(duration)), p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	p.failed++
	p.runningName = ""
	p.region.Render("\n"+testLine("❌", ctest.Name(), formatDurationLabel(duration)), p.liveBlock())
}

func (p *LiveTerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	p.skipped++
	p.runningName = ""
	p.region.Render("\n"+testLine("⏩", ctest.Name(), ""), p.liveBlock())
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
	p.region.Render("\n"+buildFailedTestsList(failedPackages), "")
}

func (p *LiveTerminalPresenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {
	// Drop the live footer and commit the final summary as permanent output.
	p.region.Render("\n"+buildFinalSummary(summary), "")
}

// liveBlock is the content of the live region: the running test (if any) and the
// running tally footer. Each part is preceded by a blank line so every entry in
// the run is uniformly separated.
func (p *LiveTerminalPresenter) liveBlock() string {
	f := p.footer()
	if p.runningName == "" {
		if f == "" {
			return ""
		}
		return "\n" + f
	}
	running := testLine(p.spinnerIcon(), p.runningName, "")
	if f == "" {
		return "\n" + running
	}
	return "\n" + running + "\n\n" + f
}

func (p *LiveTerminalPresenter) spinnerIcon() string {
	return spinner.Frame(p.spinnerFrame)
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

// testLine formats one test's line(s): the icon and the test name's first line
// (with the duration appended), then the remaining name lines. Leading and
// trailing blank lines in the test message are stripped so every test block is
// separated by exactly one newline, regardless of stray newlines at the start or
// end of the message.
func testLine(icon, name, durationLabel string) string {
	head, body := cleanNameLines(name)
	out := icon + " " + head + durationLabel
	if len(body) > 0 {
		out += "\n" + strings.Join(body, "\n")
	}
	return out
}

// cleanNameLines splits a (possibly multi-line) test name into its first line and
// the remaining body lines, dropping blank lines at the start of the body and at
// the very end.
func cleanNameLines(name string) (head string, body []string) {
	name = strings.TrimRight(name, " \t\n")
	lines := utils.SplitStringByNewLine(name)
	head = lines[0]
	body = lines[1:]
	for len(body) > 0 && strings.TrimSpace(body[0]) == "" {
		body = body[1:]
	}
	return head, body
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
