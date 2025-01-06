package sequential_events

import (
	"errors"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
)

type Interactor struct {
	output        OutputPort
	ctestsTracker *ctests_tracker.CtestsTracker
}

func NewInteractor(output OutputPort, ctestTracker *ctests_tracker.CtestsTracker) Interactor {
	return Interactor{
		output:        output,
		ctestsTracker: ctestTracker,
	}
}

func (i *Interactor) HandleCtestRanEvt(evt events.CtestRanEvent) error {
	existingCtest := i.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest != nil {
		return nil
	}
	if i.ctestsTracker.RunningCtestsCount() > 0 {
		i.output.Error()
		return errors.New("More than one running test detected.")
	}
	ctest := ctests_tracker.NewRunningCtest(evt)
	i.ctestsTracker.InsertCtest(ctest)

	if i.ctestsTracker.IsCtestFirstOfItsPackage(ctest) {
		i.output.PackageTestsStartedRunning(evt.PackageName)
		i.output.CtestStartedRunning(&ctest)
		return nil
	}
	i.output.CtestStartedRunning(&ctest)

	return nil
}

func (i *Interactor) HandleCtestPassedEvt(evt events.CtestPassedEvent) error {
	existingCtest := i.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)

	if existingCtest == nil {
		i.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if existingCtest.HasPassed() {
		return nil
	}
	if existingCtest.IsRunning() {
		existingCtest.MarkAsPassed(evt)
		i.output.CtestPassed(existingCtest, evt.Elapsed)
		return nil
	}
	return nil
}

func (i *Interactor) HandleCtestFailedEvt(evt events.CtestFailedEvent) error {
	existingCtest := i.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if evt.IsEventOfAParentTest() {
		return nil
	}
	if existingCtest == nil {
		i.output.Error()
		return errors.New("There is no existing test.")
	}
	if existingCtest.HasFailed() {
		return nil
	}
	if !existingCtest.IsRunning() {
		i.output.Error()
		return errors.New("No running test found for test pass event.")
	}

	existingCtest.MarkAsFailed(evt)
	i.output.CtestFailed(existingCtest, evt.Elapsed)

	if existingCtest.ContainsOutput() {
		i.output.CtestOutput(existingCtest)
	}
	return nil
}

func (i *Interactor) HandleCtestSkippedEvt(evt events.CtestSkippedEvent) error {
	existingCtest := i.ctestsTracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	if existingCtest == nil {
		i.output.Error()
		return errors.New("There is no existing test.")
	}

	if existingCtest.IsSkipped() {
		return nil
	}

	if !existingCtest.IsRunning() {
		i.output.Error()
		return errors.New("No running test found for test pass event.")
	}
	existingCtest.MarkAsSkipped(evt)
	i.output.CtestSkipped(existingCtest)
	return nil
}

func (i *Interactor) HandleCtestOutputEvent(evt events.CtestOutputEvent) {
	i.ctestsTracker.HandleCtestOutputEvent(evt)
}

func (i *Interactor) HandlePackageFailedEvt(evt events.PackageFailedEvent) {
	packUt := i.ctestsTracker.PackageUnderTest(evt.PackageName)
	if packUt == nil {
		return
	}
	i.output.Print("\n\n" + packUt.ParentTestsOutput())
}

func (i *Interactor) HandleTestingStarted(evt events.TestingStartedEvent) {
	i.ctestsTracker.TestingStarted(evt)
	i.output.TestingStarted()
}

func (i Interactor) HandleTestingFinished(evt events.TestingFinishedEvent) {
	i.ctestsTracker.TestingFinished(evt)
	i.output.TestingFinishedSummary(i.ctestsTracker.TestingSummary())
}
