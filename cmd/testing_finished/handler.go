package testing_finished

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
)

type Handler struct {
	output        OutputPort
	ctestsTracker *ctests_tracker.CtestsTracker
}

func NewHandler(output OutputPort, ctestTracker *ctests_tracker.CtestsTracker) Handler {
	return Handler{
		output:        output,
		ctestsTracker: ctestTracker,
	}
}

func (h Handler) HandleTestingFinished(evt events.TestingFinishedEvent) {

	h.ctestsTracker.MarkAllPackagesAsFinished()

	summary := TestingSummary{
		packagesCount:        h.ctestsTracker.PackagesCount(),
		passedPackagesCount:  h.ctestsTracker.PassedPackagesCount(),
		failedPackagesCount:  h.ctestsTracker.FailedPackagesCount(),
		skippedPackagesCount: h.ctestsTracker.SkippedPackagesCount(),

		testsCount:        h.ctestsTracker.CtestsCount(),
		passedTestsCount:  h.ctestsTracker.PassedCtestsCount(),
		failedTestsCount:  h.ctestsTracker.FailedCtestsCount(),
		skippedTestsCount: h.ctestsTracker.SkippedCtestsCount(),

		durationS: evt.DurationS(),
	}
	h.output.TestingFinishedSummary(summary)
}
