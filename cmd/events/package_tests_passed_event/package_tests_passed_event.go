package package_tests_passed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageTestsPassedEvent struct {
	time        time.Time
	packageName string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageTestsPassedEvent {
	return PackageTestsPassedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackageTestsPassedEvent) Pictogram() string {
	return "📦✅"
}

func (evt PackageTestsPassedEvent) Message() string {
	return evt.packageName
}

func (evt PackageTestsPassedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackageTestsPassedEvent) HasDuration() bool {
	return true
}

func (evt PackageTestsPassedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt PackageTestsPassedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
