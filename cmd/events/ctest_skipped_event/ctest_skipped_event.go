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
	elapsed     *float64
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

func (evt CtestSkippedEvent) CtestName() string {
	return evt.testName
}

func (evt CtestSkippedEvent) PackageName() string {
	return evt.packageName
}

func (evt CtestSkippedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestSkippedEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
