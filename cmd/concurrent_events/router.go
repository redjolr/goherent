package concurrent_events

import (
	"strings"

	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper  events.EventsMapper
	eventsHandler *Handler
}

func NewRouter(
	eventsHandler *Handler,
) Router {

	return Router{
		eventsHandler: eventsHandler,
		eventsMapper:  events.NewEventsMapper(),
	}
}

func (router Router) Route(jsonEvt events.JsonEvent) {
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

	if jsonEvt.Test == nil && jsonEvt.Action == "skip" {
		noPackageTestsFoundEvt := router.eventsMapper.JsonTestEvt2NoPackTestsFoundEvent(jsonEvt)
		router.eventsHandler.HandleNoPackageTestsFoundEvent(noPackageTestsFoundEvt)
	}

	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.eventsHandler.HandleCtestPassedEvent(ctestPassedEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "fail" && strings.Contains(*jsonEvt.Test, "/") {
		ctestFailedEvt := router.eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
		router.eventsHandler.HandleCtestFailedEvent(ctestFailedEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "skip" && strings.Contains(*jsonEvt.Test, "/") {
		ctestSkippedEvt := router.eventsMapper.JsonTestEvt2CtestSkippedEvt(jsonEvt)
		router.eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)
	}
}