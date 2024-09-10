package unbounded_terminal_handler

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

func (i *Interactor) HandlePackageStartedEvent(evt events.PackageStartedEvent) error {
	existingPackageUt := i.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt != nil {
		return nil
	}

	packageUt := i.ctestsTracker.InsertPackageUt(evt.PackageName)
	i.output.PackageStarted(packageUt)
	return nil
}

func (i *Interactor) HandlePackagePassed(evt events.PackagePassedEvent) error {
	existingPackageUt := i.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		i.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOnePassedTest() && !existingPackageUt.HasAtLeastOneSkippedTest() {
		i.output.Error()
		return errors.New("No passing test found for the package that received a PackagePassedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	i.output.EraseScreen()
	i.output.Packages(i.ctestsTracker.Packages())
	return nil
}

func (i *Interactor) HandleCtestOutputEvent(evt events.CtestOutputEvent) {
	i.ctestsTracker.HandleCtestOutputEvent(evt)
}

func (i *Interactor) HandlePackageFailed(evt events.PackageFailedEvent) error {
	existingPackageUt := i.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		i.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOneFailedTest() {
		i.output.Error()
		return errors.New("No failing test found for the package that received a PackageFailedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	i.output.EraseScreen()
	i.output.Packages(i.ctestsTracker.Packages())
	return nil
}

func (i *Interactor) HandleCtestFailedEvent(evt events.CtestFailedEvent) {
	i.ctestsTracker.HandleCtestFailedEvent(evt)
}

func (i Interactor) HandleCtestPassedEvent(evt events.CtestPassedEvent) {
	i.ctestsTracker.HandleCtestPassedEvent(evt)
}

func (i *Interactor) HandleCtestSkippedEvent(evt events.CtestSkippedEvent) {
	i.ctestsTracker.HandleCtestSkippedEvent(evt)
}

func (i *Interactor) HandleNoPackageTestsFoundEvent(evt events.NoPackageTestsFoundEvent) error {
	packageUt := i.ctestsTracker.PackageUnderTest(evt.PackageName)
	if packageUt == nil {
		i.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	if packageUt.HasAtLeastOneTest() {
		i.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	i.ctestsTracker.DeletePackage(packageUt)
	i.output.EraseScreen()
	i.output.Packages(i.ctestsTracker.Packages())
	return nil
}

func (i Interactor) HandleTestingStarted(evt events.TestingStartedEvent) {
	i.ctestsTracker.TestingStarted(evt)
	i.output.TestingStarted()
}

func (i Interactor) HandleTestingFinished(evt events.TestingFinishedEvent) {
	i.ctestsTracker.TestingFinished(evt)
	i.output.TestingFinishedSummary(i.ctestsTracker.TestingSummary())
}
