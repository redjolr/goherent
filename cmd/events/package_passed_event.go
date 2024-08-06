package events

import (
	"time"
)

type PackagePassedEvent struct {
	Time        time.Time
	PackageName string
	Elapsed     *float64
}

func NewPackagePassedEvent(jsonEvt JsonTestEvent) PackagePassedEvent {
	return PackagePassedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		Elapsed:     jsonEvt.Elapsed,
	}
}
