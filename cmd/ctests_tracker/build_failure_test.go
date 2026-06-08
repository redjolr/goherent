package ctests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/expect"

	. "github.com/redjolr/goherent/test"
)

func TestBuildFailedPackageIsReportedAsFailed(t *testing.T) {
	Test(`
	Given a package that failed to build (no tests ran)
	When the run finishes
	Then the package is reported as failed, not skipped, with its build output attached.
	`, func(Expect expect.F) {
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.MarkPackageAsBuildFailed(
			"somePackage",
			"./conversation_test.go:95:8: undefined: conversationNew\n",
		)
		tracker.TestingFinished(events.NewTestingFinishedEvent(time.Now()))

		pack := tracker.FindPackageWithName("somePackage")
		Expect(pack.HasBuildFailure()).ToEqual(true)
		Expect(pack.IsSkipped()).ToEqual(false)
		Expect(pack.BuildOutput()).ToEqual("./conversation_test.go:95:8: undefined: conversationNew\n")
		Expect(tracker.FailedPackagesCount()).ToEqual(1)
		Expect(tracker.SkippedPackagesCount()).ToEqual(0)
	}, t)
}
