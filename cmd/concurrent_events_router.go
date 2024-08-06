package cmd

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events_handler"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
)

type ConcurrentEventsRouter struct {
	eventsMapper EventsMapper

	eventsHandler          *concurrent_events_handler.EventsHandler
	testingFinishedHandler *testing_finished_handler.EventsHandler
}

func NewConcurrentEventsRouter(
	eventsHandler *concurrent_events_handler.EventsHandler,
	testingFinishedHandler *testing_finished_handler.EventsHandler,
) ConcurrentEventsRouter {

	return ConcurrentEventsRouter{
		eventsHandler:          eventsHandler,
		testingFinishedHandler: testingFinishedHandler,
		eventsMapper:           NewEventsMapper(),
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

func (router ConcurrentEventsRouter) RouteTestingStartedEvent(timestamp time.Time) {
	testingStartedEvt := events.NewTestingStartedEvent(timestamp)
	router.eventsHandler.HandleTestingStarted(testingStartedEvt)
}

func (router ConcurrentEventsRouter) RouteTestingFinishedEvent(duration time.Duration) {
	testingFinishedEvt := events.NewTestingFinishedEvent(duration)
	router.testingFinishedHandler.HandleTestingFinished(testingFinishedEvt)
}
