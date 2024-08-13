package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	PackageStarted(packageUt ctests_tracker.PackageUnderTest)
	EraseScreen()
	Packages(packages []*ctests_tracker.PackageUnderTest)
	Error()
}
