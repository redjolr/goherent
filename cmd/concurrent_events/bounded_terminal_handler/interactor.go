package bounded_terminal_handler

import (
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

func (i Interactor) HandlePackageStartedEvent(evt events.PackageStartedEvent) error {
	existingPackageUt := i.ctestsTracker.FindPackageWithName(evt.PackageName)
	if existingPackageUt != nil {
		return nil
	}

	i.ctestsTracker.InsertPackageUt(evt.PackageName)
	if i.ctestsTracker.HasPackages() {
		i.output.EraseScreen()
	}

	i.output.DisplayPackages(i.ctestsTracker.Packages())
	return nil
}
