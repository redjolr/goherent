package test_passed_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestPassedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestPassedEvent {
	return TestPassedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt TestPassedEvent) Pictogram() string {
	return "✅"
}

func (evt TestPassedEvent) Message() string {
	return evt.testName
}

func (evt TestPassedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestPassedEvent) HasDuration() bool {
	return true
}

func (evt TestPassedEvent) Duration() float64 {
	return evt.elapsed
}
