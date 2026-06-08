package sequential_events

import (
	"strings"
	"testing"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

// When a package fails to build its tests never run, so the pass rate (computed
// only over tests that ran) would misleadingly read "100% passed". It must be
// hidden in that case.
func TestPassRateLabelHiddenWhenAPackageFailedToBuild(t *testing.T) {
	withBuildFailure := ctests_tracker.TestingSummary{
		TestsCount:               10,
		PassedTestsCount:         10,
		BuildFailedPackagesCount: 1,
	}
	if got := passRateLabel(withBuildFailure); got != "" {
		t.Errorf("pass rate should be hidden when a package failed to build, got %q", got)
	}

	allBuilt := ctests_tracker.TestingSummary{TestsCount: 10, PassedTestsCount: 10}
	if got := passRateLabel(allBuilt); !strings.Contains(got, "100% passed") {
		t.Errorf("expected the pass rate to be shown when everything built, got %q", got)
	}
}

func TestBuildFailuresNote(t *testing.T) {
	if got := buildFailuresNote(ctests_tracker.TestingSummary{}); got != "" {
		t.Errorf("expected no note when nothing failed to build, got %q", got)
	}

	one := buildFailuresNote(ctests_tracker.TestingSummary{BuildFailedPackagesCount: 1})
	if !strings.Contains(one, "1 package failed to build") || !strings.Contains(one, "its tests did not run") {
		t.Errorf("unexpected single-package note: %q", one)
	}

	two := buildFailuresNote(ctests_tracker.TestingSummary{BuildFailedPackagesCount: 2})
	if !strings.Contains(two, "2 packages failed to build") || !strings.Contains(two, "their tests did not run") {
		t.Errorf("unexpected plural note: %q", two)
	}
}
