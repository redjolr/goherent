package test_failed_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestFailedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestFailedEvent {
	return TestFailedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt TestFailedEvent) Pictogram() string {
	return "❌"
}

func (evt TestFailedEvent) Message() string {
	return evt.testName
}

func (evt TestFailedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestFailedEvent) HasDuration() bool {
	return true
}

func (evt TestFailedEvent) Duration() float64 {
	return evt.elapsed
}
