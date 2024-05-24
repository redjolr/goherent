package ctest_skipped_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestSkippedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestSkippedEvent {
	return CtestSkippedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt CtestSkippedEvent) Pictogram() string {
	return "⏭"
}

func (evt CtestSkippedEvent) Message() string {
	return evt.testName
}

func (evt CtestSkippedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestSkippedEvent) HasDuration() bool {
	return true
}

func (evt CtestSkippedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt CtestSkippedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
