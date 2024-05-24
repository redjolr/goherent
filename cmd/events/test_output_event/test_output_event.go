package test_output_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type TestOutputEvent struct {
	time        time.Time
	packageName string
	testName    string
	output      string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) TestOutputEvent {
	return TestOutputEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		testName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		output:      jsonEvt.Output,
	}
}

func (evt TestOutputEvent) Pictogram() string {
	return ""
}

func (evt TestOutputEvent) Message() string {
	return evt.testName
}

func (evt TestOutputEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestOutputEvent) HasDuration() bool {
	return true
}

func (evt TestOutputEvent) Duration() float64 {
	return 0
}
