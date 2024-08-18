package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	EraseScreen()
	DisplayPackages(runningTests []*ctests_tracker.PackageUnderTest)
}
