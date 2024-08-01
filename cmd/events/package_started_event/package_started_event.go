package package_started_event

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageStartedEvent struct {
	time        time.Time
	packageName string
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageStartedEvent {
	return PackageStartedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
	}
}

func (evt PackageStartedEvent) PackageName() string {
	return evt.packageName
}
