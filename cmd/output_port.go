package cmd

import (
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	TestingStarted(timestamp time.Time)
	PackageTestsStartedRunning(packageName string)
	CtestPassed(ctest *ctests_tracker.Ctest, duration float64)
	CtestFailed(ctest *ctests_tracker.Ctest, duration float64)
	CtestStartedRunning(ctest *ctests_tracker.Ctest)
	CtestOutput(ctest *ctests_tracker.Ctest)
	Error()
}
