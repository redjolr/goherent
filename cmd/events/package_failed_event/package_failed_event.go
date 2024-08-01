package package_failed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageFailedEvent struct {
	time        time.Time
	packageName string
	elapsed     *float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageFailedEvent {
	return PackageFailedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackageFailedEvent) Pictogram() string {
	return "📦❌"
}

func (evt PackageFailedEvent) PackageName() string {
	return evt.packageName
}

func (evt PackageFailedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackageFailedEvent) Equals(otherEvt events.PackageEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.PackageName() == otherEvt.PackageName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
