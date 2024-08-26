package cmd

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
)

type Router struct {
	sequential *sequential_events.Router
	concurrent *concurrent_events.Router
}

func NewRouter(
	sequential *sequential_events.Router,
	concurrent *concurrent_events.Router,
) Router {
	return Router{
		sequential: sequential,
		concurrent: concurrent,
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

func (r *Router) RouteTestingStartedEvent(concurrently bool) {
	testingStartedEvt := events.NewTestingStartedEvent(time.Now())
	if concurrently {
		r.concurrent.Route(testingStartedEvt)
	} else {
		r.sequential.Route(testingStartedEvt)
	}
}

func (r *Router) RouteTestingFinishedEvent(concurrently bool) {
	testingFinishedEvt := events.NewTestingFinishedEvent(time.Now())
	if concurrently {
		r.concurrent.Route(testingFinishedEvt)
	} else {
		r.sequential.Route(testingFinishedEvt)
	}
}
