package tests_tracker

import (
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
)

type Cest struct {
	name      string
	events    []events.Event
	isRunning bool
}

func NewCest(name string) Cest {
	return Cest{
		name:   name,
		events: []events.Event{},
	}
}

func (cest *Cest) IsRunning() bool {
	return cest.isRunning
}

func (cest *Cest) NewRanEvent(evt test_ran_event.TestRanEvent) {
	cest.isRunning = true
	cest.events = append(cest.events, evt)
}
