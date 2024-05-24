package tests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_continued_event"
	"github.com/redjolr/goherent/cmd/events/test_failed_event"
	"github.com/redjolr/goherent/cmd/events/test_passed_event"
	"github.com/redjolr/goherent/cmd/events/test_paused_event"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
	"github.com/redjolr/goherent/cmd/events/test_skipped_event"
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

func (ctest *Ctest) NewRanEvent(evt test_ran_event.TestRanEvent) {
	ctest.isRunning = true
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewPassedEvent(evt test_passed_event.TestPassedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewFailedEvent(evt test_failed_event.TestFailedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewPausedEvent(evt test_paused_event.TestPausedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewSkippedEvent(evt test_skipped_event.TestSkippedEvent) {
	ctest.isRunning = false
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) NewContinuedEvent(evt test_continued_event.TestContinuedEvent) {
	ctest.isRunning = true
	ctest.events = append(ctest.events, evt)
}

func (ctest *Ctest) EventCount() int {
	return len(ctest.events)
}
