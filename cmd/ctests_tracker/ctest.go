package ctests_tracker

import (
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

// goReportedElapsedResolutionS is the granularity (in seconds) of the per-test
// Elapsed that Go's test2json reports — it formats test times as "%.2fs", so
// anything under 5ms rounds to 0.00s. Below this we recover a finer-grained
// duration ourselves from the run→pass/fail event timestamps.
const goReportedElapsedResolutionS = 0.01

// Ctest stands for Client Test
// It represents the tests that the client of the Goherent package runs
type Ctest struct {
	name        string
	packageName string
	outputEvts  []events.CtestOutputEvent
	isRunning   bool
	hasPassed   bool
	hasFailed   bool
	isSkipped   bool
	ranAt       time.Time
	durationS   float64
}

// ctestDurationS returns a test's elapsed time in seconds. It trusts Go's
// reported Elapsed when it carries information (>= 0.01s), and otherwise falls
// back to the wall-clock delta between the run event and the end event, which
// carry full-precision timestamps. The fallback can overestimate t.Parallel()
// tests (they sit paused between those events), but those are not the sub-10ms
// tests this branch handles. When there is no run timestamp (e.g. concurrent
// runs, which don't track a per-test run event) the reported Elapsed is used.
func ctestDurationS(ranAt, endTime time.Time, reportedElapsed float64) float64 {
	if reportedElapsed >= goReportedElapsedResolutionS {
		return reportedElapsed
	}
	if ranAt.IsZero() {
		return reportedElapsed
	}
	if measured := endTime.Sub(ranAt).Seconds(); measured > reportedElapsed {
		return measured
	}
	return reportedElapsed
}

func NewCtest(testName string, packageName string) Ctest {
	return Ctest{
		name:        testName,
		packageName: packageName,
		outputEvts:  []events.CtestOutputEvent{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
	}
}

func NewRunningCtest(ranEvt events.CtestRanEvent) Ctest {
	return Ctest{
		name:        ranEvt.TestName,
		packageName: ranEvt.PackageName,
		outputEvts:  []events.CtestOutputEvent{},
		isRunning:   true,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   false,
		ranAt:       ranEvt.Time,
	}
}

func NewPassedCtest(passedEvt events.CtestPassedEvent) Ctest {
	return Ctest{
		name:        passedEvt.TestName,
		packageName: passedEvt.PackageName,
		outputEvts:  []events.CtestOutputEvent{},
		isRunning:   false,
		hasPassed:   true,
		hasFailed:   false,
		isSkipped:   false,
		durationS:   ctestDurationS(time.Time{}, passedEvt.Time, passedEvt.Elapsed),
	}
}

func NewFailedCtest(failedEvt events.CtestFailedEvent) Ctest {
	return Ctest{
		name:        failedEvt.TestName,
		packageName: failedEvt.PackageName,
		outputEvts:  []events.CtestOutputEvent{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   true,
		isSkipped:   false,
		durationS:   ctestDurationS(time.Time{}, failedEvt.Time, failedEvt.Elapsed),
	}
}

func NewSkippedCtest(skippedEvt events.CtestSkippedEvent) Ctest {
	return Ctest{
		name:        skippedEvt.TestName,
		packageName: skippedEvt.PackageName,
		outputEvts:  []events.CtestOutputEvent{},
		isRunning:   false,
		hasPassed:   false,
		hasFailed:   false,
		isSkipped:   true,
	}
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

// RanAt is the timestamp of the CtestRanEvent that started this test. It is the
// zero time for Ctests not created from a run event. Used to recover a
// finer-grained duration than Go's 0.01s-resolution Elapsed for fast tests.
func (ctest *Ctest) RanAt() time.Time {
	return ctest.ranAt
}

// DurationS is the test's elapsed time in seconds, computed when the test was
// marked passed or failed. It is 0 for tests that never finished or were skipped.
func (ctest *Ctest) DurationS() float64 {
	return ctest.durationS
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

func (ctest *Ctest) RecordOutputEvt(evt events.CtestOutputEvent) {
	ctest.outputEvts = append(ctest.outputEvts, evt)
}

func (ctest *Ctest) ContainsOutput() bool {
	return ctest.Output() != ""
}

func (ctest *Ctest) Output() string {
	outputEventsSlice := New_outputEventsSlice(ctest.outputEvts)
	for outputEventsSlice.Contains(ctest.name) {
		first, last := outputEventsSlice.NarrowDownRange(ctest.name, 0, len(outputEventsSlice.outputEvts)-1)
		if last != len(outputEventsSlice.outputEvts) {
			outputEventsSlice.RemoveOrderRange(first, last)
		}
	}
	return outputEventsSlice.Output()
}

func (ctest *Ctest) MarkAsPassed(passedEvt events.CtestPassedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = true
	ctest.hasFailed = false
	ctest.isSkipped = false
	ctest.durationS = ctestDurationS(ctest.ranAt, passedEvt.Time, passedEvt.Elapsed)
}

func (ctest *Ctest) MarkAsFailed(failedEvt events.CtestFailedEvent) {
	ctest.isRunning = false
	ctest.hasPassed = false
	ctest.hasFailed = true
	ctest.isSkipped = false
	ctest.durationS = ctestDurationS(ctest.ranAt, failedEvt.Time, failedEvt.Elapsed)
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
