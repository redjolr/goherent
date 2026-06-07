package sequential_events

import (
	"errors"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
)

// goReportedElapsedResolutionS is the granularity (in seconds) of the per-test
// Elapsed that Go's test2json reports — it formats test times as "%.2fs", so
// anything under 5ms rounds to 0.00s. Below this we recover a finer-grained
// duration ourselves from the run→pass/fail event timestamps.
const goReportedElapsedResolutionS = 0.01

// ctestDuration returns the test's elapsed time in seconds. It trusts Go's
// reported Elapsed when it carries information (>= 0.01s), and otherwise falls
// back to the wall-clock delta between the run event and this end event, which
// carry full-precision timestamps. The fallback can overestimate t.Parallel()
// tests (they sit paused between those events), but those are not the sub-10ms
// tests this branch handles.
func ctestDuration(ctest *ctests_tracker.Ctest, endTime time.Time, reportedElapsed float64) float64 {
	if reportedElapsed >= goReportedElapsedResolutionS {
		return reportedElapsed
	}
	ranAt := ctest.RanAt()
	if ranAt.IsZero() {
		return reportedElapsed
	}
	if measured := endTime.Sub(ranAt).Seconds(); measured > reportedElapsed {
		return measured
	}
	return reportedElapsed
}

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
		duration := ctestDuration(existingCtest, evt.Time, evt.Elapsed)
		existingCtest.MarkAsPassed(evt)
		i.output.CtestPassed(existingCtest, duration)
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

	duration := ctestDuration(existingCtest, evt.Time, evt.Elapsed)
	existingCtest.MarkAsFailed(evt)
	i.output.CtestFailed(existingCtest, duration)

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

// HandleTick drives a periodic redraw so the running test's spinner animates
// even while a single test runs for a while with no events.
func (i *Interactor) HandleTick() {
	i.output.Tick()
}

func (i Interactor) HandleTestingFinished(evt events.TestingFinishedEvent) {
	i.ctestsTracker.TestingFinished(evt)
	failedPackages := i.ctestsTracker.FinishedFailedPackages()
	if len(failedPackages) > 0 {
		i.output.FailedTestsList(failedPackages)
	}
	i.output.TestingFinishedSummary(i.ctestsTracker.TestingSummary())
}
