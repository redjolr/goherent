package concurrent_events_handler

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

func (eh EventsHandler) HandleTestingStarted(evt events.TestingStartedEvent) {
	eh.output.TestingStarted()
}

func (eh EventsHandler) HandlePackageStartedEvent(evt events.PackageStartedEvent) error {
	existingPackageUt := eh.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt != nil {
		return nil
	}

	packageUt := eh.ctestsTracker.InsertPackageUt(evt.PackageName)
	eh.output.PackageStarted(packageUt)

	return nil
}

func (eh EventsHandler) HandlePackagePassed(evt events.PackagePassedEvent) error {
	existingPackageUt := eh.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		eh.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOnePassedTest() && !existingPackageUt.HasAtLeastOneSkippedTest() {
		eh.output.Error()
		return errors.New("No passing test found for the package that received a PackagePassedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	eh.output.EraseScreen()
	eh.output.Packages(eh.ctestsTracker.Packages())
	return nil
}

func (eh EventsHandler) HandlePackageFailed(evt events.PackageFailedEvent) error {
	existingPackageUt := eh.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		eh.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOneFailedTest() {
		eh.output.Error()
		return errors.New("No failing test found for the package that received a PackageFailedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	eh.output.EraseScreen()
	eh.output.Packages(eh.ctestsTracker.Packages())
	return nil
}

func (eh EventsHandler) HandleCtestFailedEvent(evt events.CtestFailedEvent) {
	ctest := ctests_tracker.NewFailedCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)
}

func (eh EventsHandler) HandleCtestPassedEvent(evt events.CtestPassedEvent) {
	ctest := ctests_tracker.NewPassedCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)
}

func (eh EventsHandler) HandleCtestSkippedEvt(evt events.CtestSkippedEvent) {
	ctest := ctests_tracker.NewSkippedCtest(evt)
	eh.ctestsTracker.InsertCtest(ctest)
}

func (eh EventsHandler) HandleNoPackageTestsFoundEvent(evt events.NoPackageTestsFoundEvent) error {
	packageUt := eh.ctestsTracker.PackageUnderTest(evt.PackageName)
	if packageUt == nil {
		eh.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	if packageUt.HasAtLeastOneTest() {
		eh.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	eh.ctestsTracker.DeletePackage(packageUt)
	eh.output.EraseScreen()
	eh.output.Packages(eh.ctestsTracker.Packages())
	return nil
}
