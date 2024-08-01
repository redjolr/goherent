package concurrent_events_handler

import (
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	TestingStarted(timestamp time.Time)
	PackageStarted(packageUt ctests_tracker.PackageUnderTest)
	EraseScreen()
	Packages(packages []ctests_tracker.PackageUnderTest)
	Error()
}
