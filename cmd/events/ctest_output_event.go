package events

import (
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestOutputEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
	Output      string
}

func NewCtestOutputEvent(jsonEvt JsonTestEvent) CtestOutputEvent {
	return CtestOutputEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		Output:      jsonEvt.Output,
	}
}
