package events

import "time"

// {"Time":"2024-05-17T20:33:05.750866+02:00","Action":"pass","Package":"github.com/redjolr/go-iam/src/core/tests","Elapsed":0.389}
type PackageTestsPassedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

func (evt PackageTestsPassedEvent) Pictogram() string {
	return "📦✅"
}

func (evt PackageTestsPassedEvent) Message() string {
	return evt.Package
}

func (evt PackageTestsPassedEvent) Timestamp() time.Time {
	return evt.Time
}

func (evt PackageTestsPassedEvent) HasDuration() bool {
	return true
}

func (evt PackageTestsPassedEvent) Duration() float64 {
	return evt.Elapsed
}
