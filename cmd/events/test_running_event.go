package events

import "time"

// {"Time":"2024-05-17T20:33:05.7278673+02:00","Action":"run","Package":"github.com/redjolr/go-iam/src/core/tests","Test":"TestUuid"}
type TestRunningEvent struct {
	Time    time.Time
	Package string `json:"Package"`
	Test    string `json:"Test"`
}

func (evt TestRunningEvent) Pictogram() string {
	return "🏃"
}

func (evt TestRunningEvent) Message() string {
	return evt.Test
}

func (evt TestRunningEvent) Timestamp() time.Time {
	return evt.Time
}

func (evt TestRunningEvent) HasDuration() bool {
	return false
}

func (evt TestRunningEvent) Duration() float64 {
	return 0
}
