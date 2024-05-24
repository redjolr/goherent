package ctest_output_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type CtestOutputEvent struct {
	time        time.Time
	packageName string
	testName    string
	output      string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) CtestOutputEvent {
	return CtestOutputEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		output:      jsonEvt.Output,
	}
}

func (evt CtestOutputEvent) Pictogram() string {
	return ""
}

func (evt CtestOutputEvent) Message() string {
	return evt.testName
}

func (evt CtestOutputEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestOutputEvent) HasDuration() bool {
	return true
}

func (evt CtestOutputEvent) Duration() float64 {
	return 0
}

func (evt CtestOutputEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
