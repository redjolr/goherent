package sequential_events

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

// slowTestThresholdS is the elapsed time (in seconds) at or above which a test's
// duration is highlighted in yellow so slow tests stand out at a glance.
const slowTestThresholdS = 1.0

// formatDurationLabel returns a colorized, space-prefixed " (12ms)" suffix to
// append after a test result. Slow tests (see slowTestThresholdS) are rendered
// in yellow; everything else is dimmed so it stays unobtrusive. A non-positive
// duration (e.g. a test the runner reported as 0s) yields an empty string so we
// don't clutter the output with meaningless "(<1ms)" labels.
func formatDurationLabel(seconds float64) string {
	if seconds <= 0 {
		return ""
	}
	color := ansi_escape.DIM
	if seconds >= slowTestThresholdS {
		color = ansi_escape.YELLOW
	}
	return " " + color + "(" + utils.FormatDuration(seconds) + ")" + ansi_escape.COLOR_RESET
}

// testingVerdictHeadline returns a single bold, colored line summarizing the run
// at a glance, shown above the detailed packages/tests/time breakdown.
func testingVerdictHeadline(summary ctests_tracker.TestingSummary) string {
	switch {
	case summary.TestsCount == 0:
		return ansi_escape.BOLD + ansi_escape.YELLOW + "⚠ No tests ran" + ansi_escape.COLOR_RESET
	case summary.FailedTestsCount == 0 && summary.FailedPackagesCount == 0:
		return ansi_escape.BOLD + ansi_escape.GREEN + "✓ All tests passed" + ansi_escape.COLOR_RESET
	default:
		return ansi_escape.BOLD + ansi_escape.RED + "✗ Tests failed" + ansi_escape.COLOR_RESET
	}
}

// buildSlowestTestsReport renders the "N slowest tests" block shown at the end
// of a run: a header followed by one line per test with its (colored) duration
// and the first line of its name. Returns "" when there are no timed tests.
func buildSlowestTestsReport(tests []*ctests_tracker.Ctest) string {
	if len(tests) == 0 {
		return ""
	}
	noun := "tests"
	if len(tests) == 1 {
		noun = "test"
	}
	report := fmt.Sprintf("🐢 %d slowest %s:", len(tests), noun)
	for _, ctest := range tests {
		report += "\n  " + slowestDurationLabel(ctest.DurationS()) + " " + firstLine(ctest.Name())
	}
	return report
}

// slowestDurationLabel formats a "(12ms)" label for the slowest-tests report,
// yellow for slow tests (see slowTestThresholdS) and dimmed otherwise.
func slowestDurationLabel(seconds float64) string {
	color := ansi_escape.DIM
	if seconds >= slowTestThresholdS {
		color = ansi_escape.YELLOW
	}
	return color + "(" + utils.FormatDuration(seconds) + ")" + ansi_escape.COLOR_RESET
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return s[:i]
	}
	return s
}

// passRateLabel returns a dimmed " (NN% passed)" suffix for the tests summary
// line, or an empty string when no tests ran.
func passRateLabel(summary ctests_tracker.TestingSummary) string {
	if summary.TestsCount == 0 {
		return ""
	}
	passRate := summary.PassedTestsCount * 100 / summary.TestsCount
	return ansi_escape.DIM + fmt.Sprintf(" (%d%% passed)", passRate) + ansi_escape.COLOR_RESET
}
