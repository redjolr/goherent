package sequential_events

import (
	"errors"

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

func (h Handler) HandleCtestRanEvt(evt events.CtestRanEvent) error {
	existingCtest := h.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest != nil {
		return nil
	}
	if h.ctestsTracker.RunningCtestsCount() > 0 {
		h.output.Error()
		return errors.New("More than one running test detected.")
	}
	ctest := ctests_tracker.NewRunningCtest(evt)
	h.ctestsTracker.InsertCtest(ctest)

	if h.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		h.output.PackageTestsStartedRunning(evt.PackageName)
		h.output.CtestStartedRunning(&ctest)
		return nil
	}
	h.output.CtestStartedRunning(&ctest)

	return nil
}

func (h Handler) HandleCtestPassedEvt(evt events.CtestPassedEvent) error {
	existingCtest := h.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)

	if existingCtest == nil {
		h.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if existingCtest.HasPassed() {
		return nil
	}
	if existingCtest.IsRunning() {
		existingCtest.MarkAsPassed(evt)
		h.output.CtestPassed(existingCtest, evt.Elapsed)
		return nil
	}
	return nil
}

func (h Handler) HandleCtestFailedEvt(evt events.CtestFailedEvent) error {
	existingCtest := h.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)

	if existingCtest == nil {
		h.output.Error()
		return errors.New("There is no existing test.")
	}
	if existingCtest.HasFailed() {
		return nil
	}
	if !existingCtest.IsRunning() {
		h.output.Error()
		return errors.New("No running test found for test pass event.")
	}

	existingCtest.MarkAsFailed(evt)
	h.output.CtestFailed(existingCtest, evt.Elapsed)

	if existingCtest.ContainsOutput() {
		h.output.CtestOutput(existingCtest)
	}
	return nil
}

func (h Handler) HandleCtestSkippedEvt(evt events.CtestSkippedEvent) error {
	existingCtest := h.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest == nil {
		h.output.Error()
		return errors.New("There is no existing test.")
	}

	if existingCtest.IsSkipped() {
		return nil
	}

	if !existingCtest.IsRunning() {
		h.output.Error()
		return errors.New("No running test found for test pass event.")
	}
	existingCtest.MarkAsSkipped(evt)
	h.output.CtestSkipped(existingCtest)
	return nil
}

func (h Handler) HandleCtestOutputEvent(evt events.CtestOutputEvent) {
	existingCtest := h.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest != nil {
		existingCtest.RecordOutputEvt(evt)
		return
	}

	ctest := ctests_tracker.NewCtest(evt.TestName, evt.PackageName)
	ctest.RecordOutputEvt(evt)
	h.ctestsTracker.InsertCtest(ctest)
}

func (h Handler) HandleTestingStarted(evt events.TestingStartedEvent) {
	h.output.TestingStarted()
}
