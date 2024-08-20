package sequential_events

import (
	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper  events.EventsMapper
	eventsHandler *Interactor
}

func NewRouter(
	eventsHandler *Interactor,
) Router {
	return Router{
		eventsHandler: eventsHandler,
		eventsMapper:  events.NewEventsMapper(),
	}
}

func (router Router) Route(unknownEvt any) {

	switch evt := unknownEvt.(type) {
	case events.CtestPassedEvent:
		router.eventsHandler.HandleCtestPassedEvt(evt)
	case events.CtestRanEvent:
		router.eventsHandler.HandleCtestRanEvt(evt)
	case events.CtestOutputEvent:
		router.eventsHandler.HandleCtestOutputEvent(evt)
	case events.CtestFailedEvent:
		router.eventsHandler.HandleCtestFailedEvt(evt)
	case events.CtestSkippedEvent:
		router.eventsHandler.HandleCtestSkippedEvt(evt)
	case events.TestingFinishedEvent:
		router.eventsHandler.HandleTestingFinished(evt)
	}
}
