package events

import "time"

type TestingEvent interface {
	Pictogram() string
	Timestamp() time.Time
}

type PackageEvent interface {
	TestingEvent
	PackageName() string
	Equals(evt PackageEvent) bool
}

type CtestEvent interface {
	TestingEvent
	CtestName() string
	Equals(evt CtestEvent) bool
}
