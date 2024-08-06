package events

import (
	"time"
)

type PackageFailedEvent struct {
	Time        time.Time
	PackageName string
	Elapsed     *float64
}

func NewPackageFailedEvent(jsonEvt JsonTestEvent) PackageFailedEvent {
	return PackageFailedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		Elapsed:     jsonEvt.Elapsed,
	}
}
