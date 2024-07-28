package cmd

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	"github.com/redjolr/goherent/console/terminal"
)

type Router struct {
	eventsMapper  EventsMapper
	eventsHandler EventsHandler
}

func NewRouter() Router {
	ansiTerminal := terminal.NewAnsiTerminal()
	terminalPresenter := NewTerminalPresenter(&ansiTerminal)
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	return Router{
		eventsMapper:  NewEventsMapper(),
		eventsHandler: NewEventsHandler(&terminalPresenter, &ctestsTracker),
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
}

func (router Router) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := testing_started_event.NewTestingStartedEvent(timestamp)
	router.eventsHandler.HandleTestingStarted(testingStartedEvt)
}
