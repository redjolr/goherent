package unbounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	Error()
	TestingStarted()
	PackageStarted(packageUt ctests_tracker.PackageUnderTest)
	EraseScreen()
	Packages(packages []*ctests_tracker.PackageUnderTest)
	TestingFinishedSummaryLabel()
	TestingFinishedSummary(summary ctests_tracker.TestingSummary)
}
