package concurrent_events_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	TestingStarted()
	TestingFinishedSummary(summary TestingSummary)
	PackageStarted(packageUt ctests_tracker.PackageUnderTest)
	EraseScreen()
	Packages(packages []*ctests_tracker.PackageUnderTest)
	Error()
}
