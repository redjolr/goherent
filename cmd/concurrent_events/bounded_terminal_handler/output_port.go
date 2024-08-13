package bounded_terminal_handler

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
)

type OutputPort interface {
	DisplayCurrentState(runningTests []ctests_tracker.PackageUnderTest)
}
