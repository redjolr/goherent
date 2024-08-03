package no_package_tests_found_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type NoPackageTestsFoundEvent struct {
	time        time.Time
	packageName string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) NoPackageTestsFoundEvent {
	return NoPackageTestsFoundEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
	}
}

func (evt NoPackageTestsFoundEvent) PackageName() string {
	return evt.packageName
}
