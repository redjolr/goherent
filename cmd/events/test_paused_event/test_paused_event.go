package test_paused_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestPausedEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestPausedEvent {
	return TestPausedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}

func (evt TestPausedEvent) Pictogram() string {
	return "⏸️"
}

func (evt TestPausedEvent) Message() string {
	return evt.testName
}

func (evt TestPausedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestPausedEvent) HasDuration() bool {
	return true
}

func (evt TestPausedEvent) Duration() float64 {
	return 0
}
