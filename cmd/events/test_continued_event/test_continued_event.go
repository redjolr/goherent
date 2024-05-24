package test_continued_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
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
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
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

func (evt TestContinuedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
