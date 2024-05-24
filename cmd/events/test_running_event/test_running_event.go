package test_running_event

import "time"

type TestRunningEvent struct {
	time        time.Time
	packageName string
	testName    string
}

func (evt TestRunningEvent) Pictogram() string {
	return "🏃"
}

func (evt TestRunningEvent) Message() string {
	return evt.testName
}

func (evt TestRunningEvent) Timestamp() time.Time {
	return evt.time
}

func (evt TestRunningEvent) HasDuration() bool {
	return false
}

func (evt TestRunningEvent) Duration() float64 {
	return 0
}
