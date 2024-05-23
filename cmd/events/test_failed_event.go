package events

import "time"

// {"Time":"2024-05-23T17:04:33.6991398+02:00","Action":"fail","Package":"github.com/redjolr/go-iam/src/core/tests",
//
//	"Test":"TestUuid/equals","Elapsed":0}
type TestFailedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

func (evt TestFailedEvent) Pictogram() string {
	return "❌"
}

func (evt TestFailedEvent) Message() string {
	return evt.Test
}

func (evt TestFailedEvent) Timestamp() time.Time {
	return evt.Time
}

func (evt TestFailedEvent) HasDuration() bool {
	return true
}

func (evt TestFailedEvent) Duration() float64 {
	return evt.Elapsed
}
