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
