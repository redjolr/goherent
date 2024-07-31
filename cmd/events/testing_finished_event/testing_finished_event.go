package testing_finished_event

import "time"

type TestingFinishedEvent struct {
	duration time.Duration
}

func NewTestingFinishedEvent(duration time.Duration) TestingFinishedEvent {
	return TestingFinishedEvent{
		duration: duration,
	}
}

func (evt TestingFinishedEvent) DurationS() float32 {
	return float32(evt.duration.Seconds())
}
