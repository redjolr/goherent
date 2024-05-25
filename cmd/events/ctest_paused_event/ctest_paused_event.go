package ctest_paused_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestPausedEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestPausedEvent {
	return CtestPausedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}

func (evt CtestPausedEvent) Pictogram() string {
	return "⏸️"
}

func (evt CtestPausedEvent) CtestName() string {
	return evt.testName
}

func (evt CtestPausedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestPausedEvent) HasDuration() bool {
	return true
}

func (evt CtestPausedEvent) Duration() float64 {
	return 0
}

func (evt CtestPausedEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
