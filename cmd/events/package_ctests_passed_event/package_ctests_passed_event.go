package package_ctests_passed_event

import (
	"reflect"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageCtestsPassedEvent struct {
	time        time.Time
	packageName string
	elapsed     float64
}

func NewFromJsonTestEvent(jsonEvt events.JsonTestEvent) PackageCtestsPassedEvent {
	return PackageCtestsPassedEvent{
		time:        jsonEvt.Time,
		packageName: jsonEvt.Package,
		elapsed:     jsonEvt.Elapsed,
	}
}

func (evt PackageCtestsPassedEvent) Pictogram() string {
	return "📦✅"
}

func (evt PackageCtestsPassedEvent) Message() string {
	return evt.packageName
}

func (evt PackageCtestsPassedEvent) Timestamp() time.Time {
	return evt.time
}

func (evt PackageCtestsPassedEvent) HasDuration() bool {
	return true
}

func (evt PackageCtestsPassedEvent) Duration() float64 {
	return evt.elapsed
}

func (evt PackageCtestsPassedEvent) Equals(otherEvt events.Event) bool {
	return evt.Pictogram() == otherEvt.Pictogram() &&
		evt.Message() == otherEvt.Message() &&
		evt.Timestamp() == otherEvt.Timestamp() &&
		evt.HasDuration() == otherEvt.HasDuration() &&
		evt.Duration() == otherEvt.Duration() &&
		reflect.TypeOf(evt) == reflect.TypeOf(otherEvt)
}
