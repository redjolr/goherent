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

func (evt CtestOutputEvent) CtestName() string {
	return evt.testName
}

func (evt CtestOutputEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestOutputEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
