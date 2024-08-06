package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestSkippedEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
	Elapsed     *float64
}

func NewCtestSkippedEvent(jsonEvt JsonTestEvent) CtestSkippedEvent {
	return CtestSkippedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		Elapsed:     jsonEvt.Elapsed,
	}
}
