package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestContinuedEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
}

func NewCtestContinuedEvent(jsonEvt JsonTestEvent) CtestContinuedEvent {
	return CtestContinuedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}
