package cmd

import (
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events_handler"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events_handler"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
	"github.com/redjolr/goherent/cmd/testing_started_handler"
)

type Router struct {
	sequential *sequential_events_handler.Router
	concurrent *concurrent_events_handler.Router

	startedHandler  *testing_started_handler.EventsHandler
	finishedHandler *testing_finished_handler.EventsHandler
}

func NewRouter(
	sequential *sequential_events_handler.Router,
	concurrent *concurrent_events_handler.Router,
	startedHandler *testing_started_handler.EventsHandler,
	finishedHandler *testing_finished_handler.EventsHandler,
) Router {
	return Router{
		sequential:      sequential,
		concurrent:      concurrent,
		startedHandler:  startedHandler,
		finishedHandler: finishedHandler,
	}
}

func (r Router) Route(jsonEvt events.JsonEvent, concurrently bool) {
	if concurrently {
		r.concurrent.Route(jsonEvt)
	} else {
		r.sequential.Route(jsonEvt)
	}
}

func (router Router) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := events.NewTestingStartedEvent(timestamp)
	router.startedHandler.HandleTestingStarted(testingStartedEvt)
}

func (router Router) RouteTestingFinishedEvent(duration time.Duration) {
	testingFinishedEvt := events.NewTestingFinishedEvent(duration)
	router.finishedHandler.HandleTestingFinished(testingFinishedEvt)
}
