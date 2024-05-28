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
	ctestExists := eh.ctestsTracker.CtestWithNameInPackageExists(evt.CtestName(), evt.PackageName())
	if ctestExists {
		return
	}

	ctest := ctests_tracker.NewCtest(evt.CtestName(), evt.PackageName())
	eh.ctestsTracker.InsertCtest(ctest)

	if eh.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		eh.output.FirstCtestOfPackagePassed(evt.CtestName(), evt.PackageName(), evt.TestDuration())
		return
	}

	eh.output.CtestPassed(evt.CtestName(), evt.TestDuration())
}

func (eh EventsHandler) HandleCtestRanEvt(evt ctest_ran_event.CtestRanEvent) {
	ctestExists := eh.ctestsTracker.CtestWithNameInPackageExists(evt.CtestName(), evt.PackageName())
	if ctestExists {
		return
	}
	ctest := ctests_tracker.NewCtest(evt.CtestName(), evt.PackageName())
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
