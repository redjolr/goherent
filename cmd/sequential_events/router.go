package sequential_events

import (
	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper events.EventsMapper
	interactor   *Interactor
}

func NewRouter(
	interactor *Interactor,
) Router {
	return Router{
		interactor:   interactor,
		eventsMapper: events.NewEventsMapper(),
	}
}

// Tick drives a periodic redraw (animates the running test's spinner).
func (router Router) Tick() {
	router.interactor.HandleTick()
}

// RouteBuildFailure records that a package failed to build, with its captured
// compiler output, so it is reported as failed (not skipped).
func (router Router) RouteBuildFailure(packageName string, buildOutput string) {
	router.interactor.HandleBuildFailure(packageName, buildOutput)
}

func (router Router) Route(unknownEvt any) {

	switch evt := unknownEvt.(type) {
	case events.CtestPassedEvent:
		router.interactor.HandleCtestPassedEvt(evt)
	case events.CtestRanEvent:
		router.interactor.HandleCtestRanEvt(evt)
	case events.CtestOutputEvent:
		router.interactor.HandleCtestOutputEvent(evt)
	case events.CtestFailedEvent:
		router.interactor.HandleCtestFailedEvt(evt)
	case events.CtestSkippedEvent:
		router.interactor.HandleCtestSkippedEvt(evt)
	case events.PackageFailedEvent:
		router.interactor.HandlePackageFailedEvt(evt)
	case events.TestingStartedEvent:
		router.interactor.HandleTestingStarted(evt)
	case events.TestingFinishedEvent:
		router.interactor.HandleTestingFinished(evt)
	}
}
