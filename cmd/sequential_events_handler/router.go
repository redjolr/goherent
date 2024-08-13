package sequential_events_handler

import (
	"strings"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
)

type Router struct {
	eventsMapper events.EventsMapper

	eventsHandler          *EventsHandler
	testingFinishedHandler *testing_finished_handler.EventsHandler
}

func NewSequentialEventsRouter(
	eventsHandler *EventsHandler,
) Router {
	return Router{
		eventsHandler: eventsHandler,
		eventsMapper:  events.NewEventsMapper(),
	}
}

func (router Router) Route(jsonEvt events.JsonEvent) {
	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "run" && strings.Contains(*jsonEvt.Test, "/") {
		ctestRanEvt := router.eventsMapper.JsonTestEvt2CtestRanEvt(jsonEvt)
		router.eventsHandler.HandleCtestRanEvt(ctestRanEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "output" && strings.Contains(*jsonEvt.Test, "/") {
		ctestRanEvt := router.eventsMapper.JsonTestEvt2CtestOutputEvt(jsonEvt)
		router.eventsHandler.HandleCtestOutputEvent(ctestRanEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "fail" && strings.Contains(*jsonEvt.Test, "/") {
		ctestFailedEvt := router.eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
		router.eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "skip" && strings.Contains(*jsonEvt.Test, "/") {
		ctestSkippedEvt := router.eventsMapper.JsonTestEvt2CtestSkippedEvt(jsonEvt)
		router.eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)
	}
}
