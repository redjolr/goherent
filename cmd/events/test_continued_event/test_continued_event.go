package test_continued_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type TestContinuedEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestContinuedEvent {
	return TestContinuedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    jsonEvt.Test,
	}
}

func (evt TestContinuedEvent) Pictogram() string {
	return "🔁"
}

func (evt TestContinuedEvent) Message() string {
	return evt.testName
}

func (evt TestContinuedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestContinuedEvent) HasDuration() bool {
	return true
}

func (evt TestContinuedEvent) Duration() float64 {
	return 0
}
