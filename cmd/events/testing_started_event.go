package events

import "time"

type TestingStartedEvent struct {
	Timestamp time.Time
}

func NewTestingStartedEvent(timestamp time.Time) TestingStartedEvent {
	return TestingStartedEvent{
		Timestamp: timestamp,
	}
}
