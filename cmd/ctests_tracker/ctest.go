package ctests_tracker

import (
	"strings"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
	"github.com/redjolr/goherent/internal/uuidgen"
)

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	id          string
	name        string
	packageName string
	events      []events.CtestEvent
	output      []string
	isRunning   bool
	hasPassed   bool
	hasFailed   bool
	isSkipped   bool
}

func NewCtest(testName string, packageName string) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        testName,
		packageName: packageName,
		events:      []events.CtestEvent{},
		output:      []string{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewRunningCtest(ranEvt ctest_ran_event.CtestRanEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        ranEvt.CtestName(),
		packageName: ranEvt.PackageName(),
		events:      []events.CtestEvent{ranEvt},
		output:      []string{},
		isRunning:   true,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewPassedCtest(passedEvt ctest_passed_event.CtestPassedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        passedEvt.CtestName(),
		packageName: passedEvt.PackageName(),
		events:      []events.CtestEvent{passedEvt},
		output:      []string{},
		isRunning:   false,
		hasPassed:   true,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewFailedCtest(passedEvt ctest_failed_event.CtestFailedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        passedEvt.CtestName(),
		packageName: passedEvt.PackageName(),
		events:      []events.CtestEvent{passedEvt},
		output:      []string{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   true,
		isSkipped:   false,
	}
}

func NewSkippedCtest(passedEvt ctest_skipped_event.CtestSkippedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        passedEvt.CtestName(),
		packageName: passedEvt.PackageName(),
		events:      []events.CtestEvent{passedEvt},
		output:      []string{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   true,
	}
}

func (ctest *Ctest) Id() string {
	return ctest.id
}

func (ctest *Ctest) Name() string {
	return ctest.name
}

func (ctest *Ctest) PackageName() string {
	return ctest.packageName
}

func (ctest *Ctest) HasName(name string) bool {
	return ctest.name == name
}

func (ctest *Ctest) IsRunning() bool {
	return ctest.isRunning
}

func (ctest *Ctest) IsSkipped() bool {
	return ctest.isSkipped
}

func (ctest *Ctest) HasPassed() bool {
	return ctest.hasPassed
}

func (ctest *Ctest) HasFailed() bool {
	return ctest.hasFailed
}

func (ctest *Ctest) LogOutput(log string) {
	ctest.output = append(ctest.output, log)
}

func (ctest *Ctest) ContainsOutput() bool {
	return len(ctest.output) > 0
}

func (ctest *Ctest) Output() string {
	return strings.Join(ctest.output, "\n")
}

func (ctest *Ctest) MarkAsPassed(passedEvt ctest_passed_event.CtestPassedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = true
	ctest.hasFailed = false
	ctest.isSkipped = false
	ctest.events = append(ctest.events, passedEvt)
}

func (ctest *Ctest) MarkAsFailed(passedEvt ctest_failed_event.CtestFailedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = false
	ctest.hasFailed = true
	ctest.isSkipped = false
	ctest.events = append(ctest.events, passedEvt)
}

func (ctest *Ctest) MarkAsSkipped(passedEvt ctest_skipped_event.CtestSkippedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = false
	ctest.hasFailed = false
	ctest.isSkipped = true
	ctest.events = append(ctest.events, passedEvt)
}

func (ctest *Ctest) Equals(otherCtest Ctest) bool {
	return ctest.name == otherCtest.name &&
		ctest.packageName == otherCtest.packageName
}
