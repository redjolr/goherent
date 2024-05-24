package test_ran_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestRanEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestRanEvent {
	return TestRanEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}

func (evt TestRanEvent) Pictogram() string {
	return "🏃"
}

func (evt TestRanEvent) Message() string {
	return evt.testName
}

func (evt TestRanEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestRanEvent) HasDuration() bool {
	return false
}

func (evt TestRanEvent) Duration() float64 {
	return 0
}

func (evt TestRanEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
