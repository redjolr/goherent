package events

import "time"

type TestingEvent interface {
	Pictogram() string
	Timestamp() time.Time
	HasDuration() bool
	Duration() float64
	Equals(evt TestingEvent) bool
}

type PackageEvent interface {
	TestingEvent
	PackageName() string
}

type CtestEvent interface {
	TestingEvent
	CtestName() string
}
