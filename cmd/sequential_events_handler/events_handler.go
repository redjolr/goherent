package sequential_events_handler

import (
	"errors"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
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

func (eh EventsHandler) HandleCtestRanEvt(evt events.CtestRanEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest != nil {
		return nil
	}
	if eh.ctestsTracker.RunningCtestsCount() > 0 {
		eh.output.Error()
		return errors.New("More than one running test detected.")
	}
	ctest := ctests_tracker.NewRunningCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)

	if eh.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		eh.output.PackageTestsStartedRunning(evt.PackageName)
		eh.output.CtestStartedRunning(&ctest)
		return nil
	}
	eh.output.CtestStartedRunning(&ctest)

	return nil
}

func (eh EventsHandler) HandleCtestPassedEvt(evt events.CtestPassedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)

	if existingCtest == nil {
		eh.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if existingCtest.HasPassed() {
		return nil
	}
	if existingCtest.IsRunning() {
		existingCtest.MarkAsPassed(evt)
		eh.output.CtestPassed(existingCtest, evt.Elapsed)
		return nil
	}
	return nil
}

func (eh EventsHandler) HandleCtestFailedEvt(evt events.CtestFailedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)

	if existingCtest == nil {
		eh.output.Error()
		return errors.New("There is no existing test.")
	}
	if existingCtest.HasFailed() {
		return nil
	}
	if !existingCtest.IsRunning() {
		eh.output.Error()
		return errors.New("No running test found for test pass event.")
	}

	existingCtest.MarkAsFailed(evt)
	eh.output.CtestFailed(existingCtest, evt.Elapsed)

	if existingCtest.ContainsOutput() {
		eh.output.CtestOutput(existingCtest)
	}
	return nil
}

func (eh EventsHandler) HandleCtestSkippedEvt(evt events.CtestSkippedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest == nil {
		eh.output.Error()
		return errors.New("There is no existing test.")
	}

	if existingCtest.IsSkipped() {
		return nil
	}

	if !existingCtest.IsRunning() {
		eh.output.Error()
		return errors.New("No running test found for test pass event.")
	}
	existingCtest.MarkAsSkipped(evt)
	eh.output.CtestSkipped(existingCtest)
	return nil
}

func (eh EventsHandler) HandleCtestOutputEvent(evt events.CtestOutputEvent) {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest != nil {
		existingCtest.LogOutput(evt.Output)
		return
	}

	ctest := ctests_tracker.NewCtest(evt.TestName, evt.PackageName)
	ctest.LogOutput(evt.Output)
	eh.ctestsTracker.InsertCtest(ctest)
}

func (eh EventsHandler) HandleTestingStarted(evt events.TestingStartedEvent) {
	eh.output.TestingStarted()
}
