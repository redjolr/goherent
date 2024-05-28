package ctests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events"
)

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	name        string
	packageName string
	events      []events.CtestEvent
	isRunning   bool
}

func NewCtest(testName string, packageName string) Ctest {
	return Ctest{
		name:        testName,
		packageName: packageName,
		events:      []events.CtestEvent{},
		isRunning:   false,
	}
}

func (ctest *Ctest) IsRunning() bool {
	return ctest.isRunning
}

func (ctest *Ctest) HasName(name string) bool {
	return ctest.name == name
}

func (ctest *Ctest) HasEvent(evt events.CtestEvent) bool {
	return slices.ContainsFunc(ctest.events, func(otherEvt events.CtestEvent) bool {
		return evt.Equals(otherEvt)
	})
}

func (ctest *Ctest) EventCount() int {
	return len(ctest.events)
}

func (ctest *Ctest) Equals(otherCtest Ctest) bool {
	return ctest.name == otherCtest.name &&
		ctest.packageName == otherCtest.packageName
}
