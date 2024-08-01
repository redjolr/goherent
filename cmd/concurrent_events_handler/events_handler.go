package concurrent_events_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/package_started_event"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
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

func (eh EventsHandler) HandleTestingStarted(evt testing_started_event.TestingStartedEvent) {
	eh.output.TestingStarted(evt.Timestamp())
}

func (eh EventsHandler) HandlePackageStartedEvent(evt package_started_event.PackageStartedEvent) error {
	existingPackageUt := eh.ctestsTracker.FindPackageWithName(evt.PackageName())
	if existingPackageUt != nil {
		return nil
	}

	packageUt := eh.ctestsTracker.InsertPackageUt(evt.PackageName())
	eh.output.PackageStarted(packageUt)

	return nil
}
