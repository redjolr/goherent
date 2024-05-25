package package_ctests_failed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageCTestsFailedEvent struct {
	time        time.Time
	packageName string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageCTestsFailedEvent {
	return PackageCTestsFailedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackageCTestsFailedEvent) Pictogram() string {
	return "📦❌"
}

func (evt PackageCTestsFailedEvent) PackageName() string {
	return evt.packageName
}

func (evt PackageCTestsFailedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackageCTestsFailedEvent) HasDuration() bool {
	return true
}

func (evt PackageCTestsFailedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt PackageCTestsFailedEvent) Equals(otherEvt events.PackageEvent) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.PackageName() == otherEvt.PackageName() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
