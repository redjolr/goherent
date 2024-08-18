package testing_finished

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
)

type Interactor struct {
	output        OutputPort
	ctestsTracker *ctests_tracker.CtestsTracker
}

func NewInteractor(output OutputPort, ctestTracker *ctests_tracker.CtestsTracker) Interactor {
	return Interactor{
		output:        output,
		ctestsTracker: ctestTracker,
	}
}

func (i Interactor) HandleTestingFinished(evt events.TestingFinishedEvent) {

	i.ctestsTracker.MarkAllPackagesAsFinished()

	summary := TestingSummary{
		packagesCount:        i.ctestsTracker.PackagesCount(),
		passedPackagesCount:  i.ctestsTracker.PassedPackagesCount(),
		failedPackagesCount:  i.ctestsTracker.FailedPackagesCount(),
		skippedPackagesCount: i.ctestsTracker.SkippedPackagesCount(),

		testsCount:        i.ctestsTracker.CtestsCount(),
		passedTestsCount:  i.ctestsTracker.PassedCtestsCount(),
		failedTestsCount:  i.ctestsTracker.FailedCtestsCount(),
		skippedTestsCount: i.ctestsTracker.SkippedCtestsCount(),

		durationS: evt.DurationS(),
	}
	i.output.TestingFinishedSummary(summary)
}
