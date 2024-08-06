package events

import (
	"time"
)

type NoPackageTestsFoundEvent struct {
	Time        time.Time
	PackageName string
}

func NewNoPackageTestsFoundEvent(jsonEvt JsonTestEvent) NoPackageTestsFoundEvent {
	return NoPackageTestsFoundEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
	}
}
