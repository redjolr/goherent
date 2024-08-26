package events

import "time"

type TestingFinishedEvent struct {
	Timestamp time.Time
}

func NewTestingFinishedEvent(timestamp time.Time) TestingFinishedEvent {
	return TestingFinishedEvent{
		Timestamp: timestamp,
	}
}
