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

type Cest struct {
	name      string
	events    []events.Event
	isRunning bool
}

func NewCest(name string) Cest {
	return Cest{
		name:      name,
		events:    []events.Event{},
		isRunning: false,
	}
}

func (cest *Cest) IsRunning() bool {
	return cest.isRunning
}

func (cest *Cest) HasName(name string) bool {
	return cest.name == name
}

func (cest *Cest) HasEvent(evt events.Event) bool {
	return slices.ContainsFunc(cest.events, func(otherEvt events.Event) bool {
		return evt.Equals(otherEvt)
	})
}

func (cest *Cest) NewRanEvent(evt test_ran_event.TestRanEvent) {
	cest.isRunning = true
	cest.events = append(cest.events, evt)
}

func (cest *Cest) NewPassedEvent(evt test_passed_event.TestPassedEvent) {
	cest.isRunning = false
	cest.events = append(cest.events, evt)
}

func (cest *Cest) NewFailedEvent(evt test_failed_event.TestFailedEvent) {
	cest.isRunning = false
	cest.events = append(cest.events, evt)
}

func (cest *Cest) NewPausedEvent(evt test_paused_event.TestPausedEvent) {
	cest.isRunning = false
	cest.events = append(cest.events, evt)
}

func (cest *Cest) NewSkippedEvent(evt test_skipped_event.TestSkippedEvent) {
	cest.isRunning = false
	cest.events = append(cest.events, evt)
}

func (cest *Cest) NewContinuedEvent(evt test_continued_event.TestContinuedEvent) {
	cest.isRunning = true
	cest.events = append(cest.events, evt)
}

func (cest *Cest) EventCount() int {
	return len(cest.events)
}
