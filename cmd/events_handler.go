package cmd

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
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

func (handler EventsHandler) HandleCtestPassedEvt(evt ctest_passed_event.CtestPassedEvent) {
	ctest := ctests_tracker.NewCtest(evt.CtestName(), evt.PackageName())

	if handler.ctestsTracker.ContainsCtest(ctest) {
		return
	}
	handler.ctestsTracker.InsertCtest(ctest)

	if handler.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		handler.output.FirstCtestOfPackagePassed(evt.CtestName(), evt.PackageName(), evt.TestDuration())
		return
	}

	handler.output.CtestPassed(evt.CtestName(), evt.TestDuration())
}

func (handler EventsHandler) HandleCtestFailedEvt(evt ctest_failed_event.CtestFailedEvent) {
	handler.output.CtestPassed(evt.CtestName(), evt.TestDuration())
}
