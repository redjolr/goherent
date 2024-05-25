package ctest_passed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestPassedEvent struct {
	time        time.Time
	packageName string
	testName    string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestPassedEvent {
	return CtestPassedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt CtestPassedEvent) Pictogram() string {
	return "✅"
}

func (evt CtestPassedEvent) CtestName() string {
	return evt.testName
}

func (evt CtestPassedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestPassedEvent) HasDuration() bool {
	return true
}

func (evt CtestPassedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt CtestPassedEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
