package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestFailedEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
	Elapsed     float64
}

func NewCtestFailedEvent(jsonEvt JsonTestEvent) CtestFailedEvent {
	return CtestFailedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		Elapsed:     *jsonEvt.Elapsed,
	}
}
