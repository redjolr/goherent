package test_skipped_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestSkippedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestSkippedEvent {
	return TestSkippedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt TestSkippedEvent) Pictogram() string {
	return "⏭"
}

func (evt TestSkippedEvent) Message() string {
	return evt.testName
}

func (evt TestSkippedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestSkippedEvent) HasDuration() bool {
	return true
}

func (evt TestSkippedEvent) Duration() float64 {
	return evt.elapsed
}
