package cmd

import (
	"errors"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_output_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
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

func (eh EventsHandler) HandleCtestRanEvt(evt ctest_ran_event.CtestRanEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())
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
		eh.output.PackageTestsStartedRunning(evt.PackageName())
		eh.output.CtestStartedRunning(&ctest)
		return nil
	}
	eh.output.CtestStartedRunning(&ctest)

	return nil
}

func (eh EventsHandler) HandleCtestPassedEvt(evt ctest_passed_event.CtestPassedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())

	if existingCtest == nil {
		eh.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if existingCtest.HasPassed() {
		return nil
	}
	if existingCtest.IsRunning() {
		existingCtest.MarkAsPassed(evt)
		eh.output.CtestPassed(existingCtest, evt.TestDuration())
		return nil
	}
	return nil
}

func (eh EventsHandler) HandleCtestFailedEvt(evt ctest_failed_event.CtestFailedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())

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
	eh.output.CtestFailed(existingCtest, evt.TestDuration())

	if existingCtest.ContainsOutput() {
		eh.output.CtestOutput(existingCtest)
	}
	return nil
}

func (eh EventsHandler) HandleCtestSkippedEvt(evt ctest_skipped_event.CtestSkippedEvent) error {
	existingCtest := eh.ctestsTracker.FindCtestWithNameInPackage(evt.CtestName(), evt.PackageName())
	if existingCtest == nil {
		eh.output.Error()
		return errors.New("There is no existing test.")
	}

	if existingCtest.IsSkipped() {
		return nil
	}
	existingCtest.MarkAsSkipped(evt)
	eh.output.CtestSkipped(existingCtest)
	return nil
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
