package cmd

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper  EventsMapper
	eventsHandler EventsHandler
}

func NewRouter() Router {
	terminalPresenter := NewTerminalPresenter()
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	return Router{
		eventsMapper:  NewEventsMapper(),
		eventsHandler: NewEventsHandler(terminalPresenter, &ctestsTracker),
	}
}

func (router Router) RouteJsonEvent(jsonEvt events.JsonEvent) {
	if jsonEvt.Test != nil && jsonEvt.Action == "pass" {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "run" {
		ctestRanEvt := router.eventsMapper.JsonTestEvt2CtestRanEvt(jsonEvt)
		router.eventsHandler.HandleCtestRanEvt(ctestRanEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "fail" {
		ctestFailedEvt := router.eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
		router.eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)
	}
}
