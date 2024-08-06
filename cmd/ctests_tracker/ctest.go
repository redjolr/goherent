package ctests_tracker

import (
	"strings"

	"github.com/redjolr/goherent/cmd/events"

	"github.com/redjolr/goherent/internal/uuidgen"
)

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	id          string
	name        string
	packageName string
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
		output:      []string{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewRunningCtest(ranEvt events.CtestRanEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        ranEvt.TestName,
		packageName: ranEvt.PackageName,
		output:      []string{},
		isRunning:   true,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewPassedCtest(passedEvt events.CtestPassedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        passedEvt.TestName,
		packageName: passedEvt.PackageName,
		output:      []string{},
		isRunning:   false,
		hasPassed:   true,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewFailedCtest(failedEvt events.CtestFailedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        failedEvt.TestName,
		packageName: failedEvt.PackageName,
		output:      []string{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   true,
		isSkipped:   false,
	}
}

func NewSkippedCtest(skippedEvt events.CtestSkippedEvent) Ctest {
	return Ctest{
		id:          uuidgen.NewString(),
		name:        skippedEvt.TestName,
		packageName: skippedEvt.PackageName,
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

func (ctest *Ctest) MarkAsPassed(passedEvt events.CtestPassedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = true
	ctest.hasFailed = false
	ctest.isSkipped = false
}

func (ctest *Ctest) MarkAsFailed(failedEvt events.CtestFailedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = false
	ctest.hasFailed = true
	ctest.isSkipped = false
}

func (ctest *Ctest) MarkAsSkipped(skippedEvt events.CtestSkippedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = false
	ctest.hasFailed = false
	ctest.isSkipped = true
}

func (ctest *Ctest) Equals(otherCtest Ctest) bool {
	return ctest.name == otherCtest.name &&
		ctest.packageName == otherCtest.packageName
}
