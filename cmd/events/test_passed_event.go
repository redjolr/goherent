package events

import "time"

// {"Time":"2024-05-17T20:33:05.7278673+02:00",
// "Action":"pass","Package":"github.com/redjolr/go-iam/src/core/tests","Test":"TestUuid","Elapsed":0}
type TestPassedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

func (evt TestPassedEvent) Pictogram() string {
	return "✅"
}

func (evt TestPassedEvent) Message() string {
	return evt.Test
}

func (evt TestPassedEvent) Timestamp() time.Time {
	return evt.Time
}

func (evt TestPassedEvent) HasDuration() bool {
	return true
}

func (evt TestPassedEvent) Duration() float64 {
	return evt.Elapsed
}
