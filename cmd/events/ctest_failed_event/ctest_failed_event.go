package ctest_failed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestFailedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestFailedEvent {
	return CtestFailedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt CtestFailedEvent) Pictogram() string {
	return "❌"
}

func (evt CtestFailedEvent) CtestName() string {
	return evt.testName
}

func (evt CtestFailedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestFailedEvent) HasDuration() bool {
	return true
}

func (evt CtestFailedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt CtestFailedEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
