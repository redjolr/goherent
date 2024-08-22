package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	Error()
	TestingStarted()
	EraseScreen()
	DisplayPackages(
		runningPackages []*ctests_tracker.PackageUnderTest,
		finishedPackages []*ctests_tracker.PackageUnderTest,
	)
	RunningTestsSummary(testingSummary ctests_tracker.TestingSummary)
	TestingFinishedSummary(testingSummary ctests_tracker.TestingSummary)
	IsViewPortLarge() bool
}
