package sequential_events

import (
	"fmt"

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

// passRateLabel returns a dimmed " (NN% passed)" suffix for the tests summary
// line, or an empty string when no tests ran.
func passRateLabel(summary ctests_tracker.TestingSummary) string {
	if summary.TestsCount == 0 {
		return ""
	}
	passRate := summary.PassedTestsCount * 100 / summary.TestsCount
	return ansi_escape.DIM + fmt.Sprintf(" (%d%% passed)", passRate) + ansi_escape.COLOR_RESET
}
