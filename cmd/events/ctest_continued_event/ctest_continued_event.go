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

func (evt CtestContinuedEvent) CtestName() string {
	return evt.testName
}

func (evt CtestContinuedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt CtestContinuedEvent) Equals(otherEvt events.CtestEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.CtestName() == otherEvt.CtestName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
