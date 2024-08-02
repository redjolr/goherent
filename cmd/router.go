package cmd

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/testing_finished_event"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	"github.com/redjolr/goherent/cmd/sequential_events_handler"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
)

type Router struct {
	eventsMapper EventsMapper

	eventsHandler          *sequential_events_handler.EventsHandler
	testingFinishedHandler *testing_finished_handler.EventsHandler
}

func NewRouter(
	eventsHandler *sequential_events_handler.EventsHandler,
	testingFinishedHandler *testing_finished_handler.EventsHandler,
) Router {
	return Router{
		eventsHandler:          eventsHandler,
		testingFinishedHandler: testingFinishedHandler,
		eventsMapper:           NewEventsMapper(),
	}
}

func (router Router) RouteJsonEvent(jsonEvt events.JsonEvent) {
	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "run" && strings.Contains(*jsonEvt.Test, "/") {
		ctestRanEvt := router.eventsMapper.JsonTestEvt2CtestRanEvt(jsonEvt)
		router.eventsHandler.HandleCtestRanEvt(ctestRanEvt)
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

func (router Router) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := testing_started_event.NewTestingStartedEvent(timestamp)
	router.eventsHandler.HandleTestingStarted(testingStartedEvt)
}

func (router Router) RouteTestingFinishedEvent(duration time.Duration) {
	testingFinishedEvt := testing_finished_event.NewTestingFinishedEvent(duration)
	router.testingFinishedHandler.HandleTestingFinished(testingFinishedEvt)
}
