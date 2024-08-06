package events

import "time"

type TestingStartedEvent struct {
	Time time.Time
}

func NewTestingStartedEvent(timestamp time.Time) TestingStartedEvent {
	return TestingStartedEvent{
		Time: timestamp,
	}
}
