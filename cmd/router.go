package cmd

import (
	"strings"
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

	startedHandler  *testing_started.Interactor
	finishedHandler *testing_finished.Interactor
}

func NewRouter(
	sequential *sequential_events.Router,
	concurrent *concurrent_events.Router,
	startedHandler *testing_started.Interactor,
	finishedHandler *testing_finished.Interactor,
) Router {
	return Router{
		sequential:      sequential,
		concurrent:      concurrent,
		startedHandler:  startedHandler,
		finishedHandler: finishedHandler,
	}
}

func (r *Router) Route(jsonEvt events.JsonEvent, concurrently bool) {
	eventsMapper := events.NewEventsMapper()
	var evt any
	if jsonEvt.Test == nil && jsonEvt.Action == "pass" {
		evt = eventsMapper.JsonTestEvt2PackagePassedEvt(jsonEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "fail" {
		evt = eventsMapper.JsonTestEvt2PackageFailedEvt(jsonEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "start" {
		evt = eventsMapper.JsonTestEvt2PackageStartedEvt(jsonEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "skip" {
		evt = eventsMapper.JsonTestEvt2NoPackTestsFoundEvent(jsonEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		evt = eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "run" && strings.Contains(*jsonEvt.Test, "/") {
		evt = eventsMapper.JsonTestEvt2CtestRanEvt(jsonEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "output" && strings.Contains(*jsonEvt.Test, "/") {
		evt = eventsMapper.JsonTestEvt2CtestOutputEvt(jsonEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "fail" && strings.Contains(*jsonEvt.Test, "/") {
		evt = eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "skip" && strings.Contains(*jsonEvt.Test, "/") {
		evt = eventsMapper.JsonTestEvt2CtestSkippedEvt(jsonEvt)
	}

	if concurrently {
		r.concurrent.Route(evt)
	} else {
		r.sequential.Route(evt)
	}
}

func (router *Router) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := events.NewTestingStartedEvent(timestamp)
	router.startedHandler.HandleTestingStarted(testingStartedEvt)
}

func (r *Router) RouteTestingFinishedEvent(duration time.Duration, concurrently bool) {
	testingFinishedEvt := events.NewTestingFinishedEvent(duration)
	if concurrently {
		r.concurrent.Route(testingFinishedEvt)
	} else {
		r.sequential.Route(testingFinishedEvt)
	}
}
