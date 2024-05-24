package ctest_continued_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestContinuedEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestContinuedEvent {
	return CtestContinuedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}

func (evt CtestContinuedEvent) Pictogram() string {
	return "🔁"
}

func (evt CtestContinuedEvent) Message() string {
	return evt.testName
}

func (evt CtestContinuedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestContinuedEvent) HasDuration() bool {
	return true
}

func (evt CtestContinuedEvent) Duration() float64 {
	return 0
}

func (evt CtestContinuedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
