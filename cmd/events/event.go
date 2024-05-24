package events

import "time"

type Event interface {
	Pictogram() string
	Message() string
	Timestamp() time.Time
	HasDuration() bool
	Duration() float64
	Equals(evt Event) bool
}
