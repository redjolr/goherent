package cmd

import (
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/cmd/testing_finished"
	"github.com/redjolr/goherent/cmd/testing_started"
)

type Router struct {
	sequential *sequential_events.Router
	concurrent *concurrent_events.Router

	startedHandler  *testing_started.Handler
	finishedHandler *testing_finished.Handler
}

func NewRouter(
	sequential *sequential_events.Router,
	concurrent *concurrent_events.Router,
	startedHandler *testing_started.Handler,
	finishedHandler *testing_finished.Handler,
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
