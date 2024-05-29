package ctests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
)

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	name        string
	packageName string
	events      []events.CtestEvent
	isRunning   bool
	hasPassed   bool
}

func NewCtest(testName string, packageName string) Ctest {
	return Ctest{
		name:        testName,
		packageName: packageName,
		events:      []events.CtestEvent{},
		isRunning:   false,
		hasPassed:   false,
	}
}

func NewRunningCtest(ranEvt ctest_ran_event.CtestRanEvent) Ctest {
	return Ctest{
		name:        ranEvt.CtestName(),
		packageName: ranEvt.PackageName(),
		events:      []events.CtestEvent{ranEvt},
		isRunning:   true,
		hasPassed:   false,
	}
}

func NewPassedCtest(passedEvt ctest_passed_event.CtestPassedEvent) Ctest {
	return Ctest{
		name:        passedEvt.CtestName(),
		packageName: passedEvt.PackageName(),
		events:      []events.CtestEvent{passedEvt},
		isRunning:   false,
		hasPassed:   true,
	}
}

func (ctest *Ctest) HasName(name string) bool {
	return ctest.name == name
}

func (ctest *Ctest) IsRunning() bool {
	return ctest.isRunning
}

func (ctest *Ctest) HasPassed() bool {
	return ctest.hasPassed
}

func (ctest *Ctest) MarkAsPassed(passedEvt ctest_passed_event.CtestPassedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = true
	ctest.events = append(ctest.events, passedEvt)
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
