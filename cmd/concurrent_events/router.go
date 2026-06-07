package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper events.EventsMapper
	interactor   *Interactor
}

func NewRouter(interactor *Interactor) Router {

	return Router{
		interactor:   interactor,
		eventsMapper: events.NewEventsMapper(),
	}
}

// Tick drives a periodic redraw of the current state (no new event), so the
// running "Time:" keeps advancing during quiet stretches.
func (router *Router) Tick() {
	router.interactor.HandleTick()
}

func (router *Router) Route(unknwonEvt any) {
	switch evt := unknwonEvt.(type) {
	case events.CtestPassedEvent:
		router.interactor.HandleCtestPassedEvent(evt)
	case events.CtestFailedEvent:
		router.interactor.HandleCtestFailedEvent(evt)
	case events.CtestSkippedEvent:
		router.interactor.HandleCtestSkippedEvent(evt)
	case events.PackagePassedEvent:
		router.interactor.HandlePackagePassed(evt)
	case events.PackageFailedEvent:
		router.interactor.HandlePackageFailed(evt)
	case events.CtestOutputEvent:
		router.interactor.HandleCtestOutputEvent(evt)
	case events.PackageStartedEvent:
		router.interactor.HandlePackageStartedEvent(evt)
	case events.NoPackageTestsFoundEvent:
		router.interactor.HandleNoPackageTestsFoundEvent(evt)
	case events.TestingStartedEvent:
		router.interactor.HandleTestingStarted(evt)
	case events.TestingFinishedEvent:
		router.interactor.HandleTestingFinished(evt)
	}
}
