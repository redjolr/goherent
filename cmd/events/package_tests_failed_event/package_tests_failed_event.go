package package_tests_failed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageTestsFailedEvent struct {
	time        time.Time
	packageName string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageTestsFailedEvent {
	return PackageTestsFailedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackageTestsFailedEvent) Pictogram() string {
	return "📦❌"
}

func (evt PackageTestsFailedEvent) Message() string {
	return evt.packageName
}

func (evt PackageTestsFailedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackageTestsFailedEvent) HasDuration() bool {
	return true
}

func (evt PackageTestsFailedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt PackageTestsFailedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
