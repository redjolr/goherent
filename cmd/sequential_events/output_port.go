package sequential_events

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	Error()
	TestingStarted()
	PackageTestsStartedRunning(packageName string)
	CtestPassed(ctest *ctests_tracker.Ctest, duration float64)
	CtestFailed(ctest *ctests_tracker.Ctest, duration float64)
	Print(output string)
	CtestSkipped(ctest *ctests_tracker.Ctest)
	CtestStartedRunning(ctest *ctests_tracker.Ctest)
	CtestOutput(ctest *ctests_tracker.Ctest)
	TestingFinishedSummary(summary ctests_tracker.TestingSummary)
}
