package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestPassedEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
	Elapsed     float64
}

func NewCtestPassedEvent(jsonEvt JsonTestEvent) CtestPassedEvent {
	return CtestPassedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		Elapsed:     *jsonEvt.Elapsed,
	}
}
