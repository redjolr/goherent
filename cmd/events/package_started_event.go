package events

import (
	"time"
)

type PackageStartedEvent struct {
	Time        time.Time
	PackageName string
}

func NewPackageStartedEvent(jsonEvt JsonTestEvent) PackageStartedEvent {
	return PackageStartedEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
	}
}
