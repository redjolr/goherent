package test_failed_event

import (
	"reflect"
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

func (evt TestFailedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
