package testing_finished_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/testing_finished_event"
)

type EventsHandler struct {
	output        OutputPort
	ctestsTracker *ctests_tracker.CtestsTracker
}

func NewEventsHandler(output OutputPort, ctestTracker *ctests_tracker.CtestsTracker) EventsHandler {
	return EventsHandler{
		output:        output,
		ctestsTracker: ctestTracker,
	}
}

func (eh EventsHandler) HandleTestingFinished(evt testing_finished_event.TestingFinishedEvent) {

	eh.ctestsTracker.MarkAllPackagesAsFinished()

	summary := TestingSummary{
		packagesCount:        eh.ctestsTracker.PackagesCount(),
		passedPackagesCount:  eh.ctestsTracker.PassedPackagesCount(),
		failedPackagesCount:  eh.ctestsTracker.FailedPackagesCount(),
		skippedPackagesCount: eh.ctestsTracker.SkippedPackagesCount(),

		testsCount:        eh.ctestsTracker.CtestsCount(),
		passedTestsCount:  eh.ctestsTracker.PassedCtestsCount(),
		failedTestsCount:  eh.ctestsTracker.FailedCtestsCount(),
		skippedTestsCount: eh.ctestsTracker.SkippedCtestsCount(),

		durationS: evt.DurationS(),
	}
	eh.output.TestingFinishedSummary(summary)
}
