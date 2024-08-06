package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestRanEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
}

func NewCtestRanEvent(jsonEvt JsonTestEvent) CtestRanEvent {
	return CtestRanEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
	}
}
