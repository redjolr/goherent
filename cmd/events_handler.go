package cmd

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_output_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
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
		eh.output.PackageTestsStartedRunning(evt.PackageName())
		eh.output.CtestPassed(evt.CtestName(), evt.TestDuration())
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
		eh.output.PackageTestsStartedRunning(evt.PackageName())
		eh.output.CtestStartedRunning(evt.CtestName())
		return
	}

	eh.output.CtestStartedRunning(evt.CtestName())
}

func (eh EventsHandler) HandleCtestFailedEvt(evt ctest_failed_event.CtestFailedEvent) {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())

	if existingCtest != nil && existingCtest.HasFailed() {
		return
	}
	if existingCtest != nil {
		existingCtest.MarkAsFailed(evt)

		if eh.ctestsTracker.IsCtestFirstOfItsPackage(*existingCtest) {
			eh.output.PackageTestsStartedRunning(evt.PackageName())
		}
		eh.output.CtestFailed(evt.CtestName(), evt.TestDuration())

		if existingCtest.ContainsOutput() {
			eh.output.CtestOutput(evt.CtestName(), evt.PackageName(), existingCtest.Output())
		}
		return
	}
	ctest := ctests_tracker.NewFailedCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)
	if eh.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		eh.output.PackageTestsStartedRunning(evt.PackageName())
		eh.output.CtestFailed(evt.CtestName(), evt.TestDuration())
		return
	}

	eh.output.CtestFailed(evt.CtestName(), evt.TestDuration())
}

func (eh EventsHandler) HandleCtestOutputEvent(evt ctest_output_event.CtestOutputEvent) {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())
	if existingCtest != nil {
		existingCtest.LogOutput(evt.Output())
		return
	}

	ctest := ctests_tracker.NewCtest(evt.CtestName(), evt.PackageName())
	ctest.LogOutput(evt.Output())
	eh.ctestsTracker.InsertCtest(ctest)
}

func (eh EventsHandler) HandleTestingStarted(evt testing_started_event.TestingStartedEvent) {
	eh.output.TestingStarted(evt.Timestamp())
}
