package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestPausedEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
}

func NewCtestPausedEvent(jsonEvt JsonTestEvent) CtestPausedEvent {
	return CtestPausedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}
