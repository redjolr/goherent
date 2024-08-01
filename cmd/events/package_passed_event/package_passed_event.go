package package_passed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackagePassedEvent struct {
	time        time.Time
	packageName string
	elapsed     *float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackagePassedEvent {
	return PackagePassedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackagePassedEvent) Pictogram() string {
	return "📦✅"
}

func (evt PackagePassedEvent) PackageName() string {
	return evt.packageName
}

func (evt PackagePassedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackagePassedEvent) Equals(otherEvt events.PackageEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.PackageName() == otherEvt.PackageName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
