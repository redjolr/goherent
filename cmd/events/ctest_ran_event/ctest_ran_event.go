package ctest_ran_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestRanEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestRanEvent {
	return CtestRanEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}

func (evt CtestRanEvent) Pictogram() string {
	return "🏃"
}

func (evt CtestRanEvent) Message() string {
	return evt.testName
}

func (evt CtestRanEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestRanEvent) HasDuration() bool {
	return false
}

func (evt CtestRanEvent) Duration() float64 {
	return 0
}

func (evt CtestRanEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}