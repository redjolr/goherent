package bounded_terminal_handler

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
