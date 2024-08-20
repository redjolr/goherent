package bounded_terminal_handler

import (
	"strings"

	"github.com/redjolr/goherent/cmd/events"
)

type Router struct {
	eventsMapper events.EventsMapper
	interactor   *Interactor
}

func NewRouter(interactor *Interactor) Router {

	return Router{
		interactor:   interactor,
		eventsMapper: events.NewEventsMapper(),
	}
}

func (router *Router) Route(jsonEvt events.JsonEvent) {
	if jsonEvt.Test == nil && jsonEvt.Action == "pass" {
		packagePassedEvt := router.eventsMapper.JsonTestEvt2PackagePassedEvt(jsonEvt)
		router.interactor.HandlePackagePassed(packagePassedEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "fail" {
		packageFailedEvt := router.eventsMapper.JsonTestEvt2PackageFailedEvt(jsonEvt)
		router.interactor.HandlePackageFailed(packageFailedEvt)
	}
	if jsonEvt.Test == nil && jsonEvt.Action == "start" {
		packageStartedEvt := router.eventsMapper.JsonTestEvt2PackageStartedEvt(jsonEvt)
		router.interactor.HandlePackageStartedEvent(packageStartedEvt)
	}

	// if jsonEvt.Test == nil && jsonEvt.Action == "skip" {
	// 	noPackageTestsFoundEvt := router.eventsMapper.JsonTestEvt2NoPackTestsFoundEvent(jsonEvt)
	// 	router.interactor.HandleNoPackageTestsFoundEvent(noPackageTestsFoundEvt)
	// }

	if jsonEvt.Test != nil && jsonEvt.Action == "pass" && strings.Contains(*jsonEvt.Test, "/") {
		ctestPassedEvt := router.eventsMapper.JsonTestEvt2CtestPassedEvt(jsonEvt)
		router.interactor.HandleCtestPassedEvent(ctestPassedEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "fail" && strings.Contains(*jsonEvt.Test, "/") {
		ctestFailedEvt := router.eventsMapper.JsonTestEvt2CtestFailedEvt(jsonEvt)
		router.interactor.HandleCtestFailedEvent(ctestFailedEvt)
	}
	if jsonEvt.Test != nil && jsonEvt.Action == "skip" && strings.Contains(*jsonEvt.Test, "/") {
		ctestSkippedEvt := router.eventsMapper.JsonTestEvt2CtestSkippedEvt(jsonEvt)
		router.interactor.HandleCtestSkippedEvent(ctestSkippedEvt)
	}
}
