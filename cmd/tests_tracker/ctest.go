package tests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_continued_event"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_paused_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
)

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	name      string
	events    []events.Event
	isRunning bool
}

func NewCtest(name string) Ctest {
	return Ctest{
		name:      name,
		events:    []events.Event{},
		isRunning: false,
	}
}

func (ctest *Ctest) IsRunning() bool {
	return ctest.isRunning
}

func (ctest *Ctest) HasName(name string) bool {
	return ctest.name == name
}

func (ctest *Ctest) HasEvent(evt events.Event) bool {
	return slices.ContainsFunc(ctest.events, func(otherEvt events.Event) bool {
		return evt.Equals(otherEvt)
	})
}

func (ctest *Ctest) NewRanEvent(evt ctest_ran_event.CtestRanEvent) {
	ctest.isRunning = true
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewPassedEvent(evt ctest_passed_event.CtestPassedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewFailedEvent(evt ctest_failed_event.CtestFailedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewPausedEvent(evt ctest_paused_event.CtestPausedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewSkippedEvent(evt ctest_skipped_event.CtestSkippedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewContinuedEvent(evt ctest_continued_event.CtestContinuedEvent) {
	ctest.isRunning = true
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) EventCount() int {
	return len(ctest.events)
}
