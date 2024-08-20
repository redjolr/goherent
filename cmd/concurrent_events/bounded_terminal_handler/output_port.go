package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	Error()
	EraseScreen()
	DisplayPackages(
		runningPackages []*ctests_tracker.PackageUnderTest,
		finishedPackages []*ctests_tracker.PackageUnderTest,
		testingSummary ctests_tracker.TestingSummary,
	)
}
