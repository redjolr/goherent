package concurrent_events

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

func (h Handler) HandlePackageStartedEvent(evt events.PackageStartedEvent) error {
	existingPackageUt := h.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt != nil {
		return nil
	}

	packageUt := h.ctestsTracker.InsertPackageUt(evt.PackageName)
	h.output.PackageStarted(packageUt)

	return nil
}

func (h Handler) HandlePackagePassed(evt events.PackagePassedEvent) error {
	existingPackageUt := h.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		h.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOnePassedTest() && !existingPackageUt.HasAtLeastOneSkippedTest() {
		h.output.Error()
		return errors.New("No passing test found for the package that received a PackagePassedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	h.output.EraseScreen()
	h.output.Packages(h.ctestsTracker.Packages())
	return nil
}

func (h Handler) HandlePackageFailed(evt events.PackageFailedEvent) error {
	existingPackageUt := h.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt == nil {
		h.output.Error()
		return errors.New("No existing test found for test pass event.")
	}
	if !existingPackageUt.HasAtLeastOneFailedTest() {
		h.output.Error()
		return errors.New("No failing test found for the package that received a PackageFailedEvent.")
	}
	existingPackageUt.MarkAsFinished()
	h.output.EraseScreen()
	h.output.Packages(h.ctestsTracker.Packages())
	return nil
}

func (h Handler) HandleCtestFailedEvent(evt events.CtestFailedEvent) {
	ctest := ctests_tracker.NewFailedCtest(evt)
	h.ctestsTracker.InsertCtest(ctest)
}

func (h Handler) HandleCtestPassedEvent(evt events.CtestPassedEvent) {
	ctest := ctests_tracker.NewPassedCtest(evt)
	h.ctestsTracker.InsertCtest(ctest)
}

func (h Handler) HandleCtestSkippedEvt(evt events.CtestSkippedEvent) {
	ctest := ctests_tracker.NewSkippedCtest(evt)
	h.ctestsTracker.InsertCtest(ctest)
}

func (h Handler) HandleNoPackageTestsFoundEvent(evt events.NoPackageTestsFoundEvent) error {
	packageUt := h.ctestsTracker.PackageUnderTest(evt.PackageName)
	if packageUt == nil {
		h.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	if packageUt.HasAtLeastOneTest() {
		h.output.Error()
		return errors.New("No existing test found for NoPackageTestsFoundEvent event.")
	}
	h.ctestsTracker.DeletePackage(packageUt)
	h.output.EraseScreen()
	h.output.Packages(h.ctestsTracker.Packages())
	return nil
}
