package cmd

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
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

func (eh EventsHandler) HandleCtestPassedEvt(evt ctest_passed_event.CtestPassedEvent) {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())

	if existingCtest != nil && existingCtest.HasPassed() {
		return
	}

	if existingCtest != nil && existingCtest.IsRunning() {
		existingCtest.MarkAsPassed(evt)
		eh.output.CtestPassed(evt.CtestName(), evt.TestDuration())
		return
	}

	ctest := ctests_tracker.NewPassedCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)

	if eh.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		eh.output.FirstCtestOfPackagePassed(evt.CtestName(), evt.PackageName(), evt.TestDuration())
		return
	}

	eh.output.CtestPassed(evt.CtestName(), evt.TestDuration())
}

func (eh EventsHandler) HandleCtestRanEvt(evt ctest_ran_event.CtestRanEvent) {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())
	if existingCtest != nil {
		return
	}
	ctest := ctests_tracker.NewRunningCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)
	if eh.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		eh.output.FirstCtestOfPackageStartedRunning(evt.CtestName(), evt.PackageName())
		return
	}

	eh.output.CtestStartedRunning(evt.CtestName())
}

func (eh EventsHandler) HandleCtestFailedEvt(evt ctest_failed_event.CtestFailedEvent) {
	eh.output.CtestPassed(evt.CtestName(), evt.TestDuration())
}
