package events

import "time"

type TestingFinishedEvent struct {
	Duration time.Duration
}

func NewTestingFinishedEvent(duration time.Duration) TestingFinishedEvent {
	return TestingFinishedEvent{
		Duration: duration,
	}
}

func (evt TestingFinishedEvent) DurationS() float32 {
	return float32(evt.Duration.Seconds())
}
