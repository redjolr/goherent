package testing_started_event

import "time"

type TestingStartedEvent struct {
	time time.Time
}

func NewTestingStartedEvent(timestamp time.Time) TestingStartedEvent {
	return TestingStartedEvent{
		time: timestamp,
	}
}

func (evt TestingStartedEvent) Timestamp() time.Time {
	return evt.time
}
