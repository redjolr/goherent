package cmd

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	"github.com/redjolr/goherent/terminal"
)

type ConcurrentEventsRouter struct {
	eventsMapper  EventsMapper
	eventsHandler concurrent_events_handler.EventsHandler
}

func NewConcurrentEventsRouter() ConcurrentEventsRouter {
	ansiTerminal := terminal.NewAnsiTerminal()
	terminalPresenter := concurrent_events_handler.NewTerminalPresenter(&ansiTerminal)
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	return ConcurrentEventsRouter{
		eventsMapper:  NewEventsMapper(),
		eventsHandler: concurrent_events_handler.NewEventsHandler(&terminalPresenter, &ctestsTracker),
	}
}

func (router ConcurrentEventsRouter) RouteJsonEvent(jsonEvt events.JsonEvent) {
	if jsonEvt.Test == nil && jsonEvt.Action == "pass" {
		packagePassedEvt := router.eventsMapper.JsonTestEvt2PackagePassedEvt(jsonEvt)
		router.eventsHandler.HandlePackagePassed(packagePassedEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "fail" {
		packageFailedEvt := router.eventsMapper.JsonTestEvt2PackageFailedEvt(jsonEvt)
		router.eventsHandler.HandlePackageFailed(packageFailedEvt)
	}

	if jsonEvt.Test == nil && jsonEvt.Action == "start" {
		packageStartedEvt := router.eventsMapper.JsonTestEvt2PackageStartedEvt(jsonEvt)
		router.eventsHandler.HandlePackageStartedEvent(packageStartedEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.eventsHandler.HandleCtestPassedEvent(ctestPassedEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "fail" && strings.Contains(*jsonEvt.Test, "/") {
		ctestFailedEvt := router.eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
		router.eventsHandler.HandleCtestFailedEvent(ctestFailedEvt)
	}

}

func (router ConcurrentEventsRouter) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := testing_started_event.NewTestingStartedEvent(timestamp)
	router.eventsHandler.HandleTestingStarted(testingStartedEvt)
}

func (router ConcurrentEventsRouter) RouteTestingFinishedEvent(duration time.Duration) {

}
