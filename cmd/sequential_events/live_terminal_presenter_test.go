package sequential_events_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/terminal/vtfake"
)

// These tests drive the real Interactor through vtfake. vtfake models SGR color
// codes as zero-width and does not render them, so the assertions are on the
// plain geometry (what occupies the screen). The presenter still emits the
// colors; their bytes just aren't part of vtfake's rendered text.

func setupLive(width, height int) (*sequential_events.Interactor, *vtfake.Terminal) {
	term := vtfake.New(width, height)
	presenter := sequential_events.NewLiveTerminalPresenter(term)
	tracker := ctests_tracker.NewCtestsTracker()
	interactor := sequential_events.NewInteractor(presenter, &tracker)
	return &interactor, term
}

func ranEvt(test, pkg string) events.CtestRanEvent {
	return events.NewCtestRanEvent(events.JsonTestEvent{
		Time: time.Now(), Action: "run", Test: test, Package: pkg,
	})
}

func passedEvt(test, pkg string, elapsed float64) events.CtestPassedEvent {
	return events.NewCtestPassedEvent(events.JsonTestEvent{
		Time: time.Now(), Action: "pass", Test: test, Package: pkg, Elapsed: &elapsed,
	})
}

func failedEvt(test, pkg string, elapsed float64) events.CtestFailedEvent {
	return events.NewCtestFailedEvent(events.JsonTestEvent{
		Time: time.Now(), Action: "fail", Test: test, Package: pkg, Elapsed: &elapsed,
	})
}

func skippedEvt(test, pkg string) events.CtestSkippedEvent {
	return events.NewCtestSkippedEvent(events.JsonTestEvent{
		Time: time.Now(), Action: "skip", Test: test, Package: pkg,
	})
}

func wantText(t *testing.T, term *vtfake.Terminal, want string) {
	t.Helper()
	if got := term.Text(); got != want {
		t.Errorf("Text():\n got  = %q\n want = %q", got, want)
	}
}

// While a test runs, the live block shows ⏳ + the running test (footer is empty
// until something finishes).
func TestLiveRunningTestIsShown(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))

	wantText(t, term, "🚀 Starting...\n\n📦 somePackage\n⏳ ParentTest/testName")
}

// A passed test is committed with its duration on the first line, and the footer
// shows the running tally.
func TestLivePassedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/testName", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"✅ ParentTest/testName (10ms)\n"+
			"1 passed")
}

// Two passed tests: both committed, footer counts up in place (no stacking).
func TestLiveTwoPassedTests(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/test1", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/test1", "somePackage", 0.01))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/test2", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/test2", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"✅ ParentTest/test1 (10ms)\n"+
			"✅ ParentTest/test2 (10ms)\n"+
			"2 passed")
}

// A failed test is committed; the footer shows passed and failed counts.
func TestLiveFailedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestFailedEvt(failedEvt("ParentTest/testName", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"❌ ParentTest/testName (10ms)\n"+
			"0 passed · 1 failed")
}

// A skipped test (no duration) is committed; the footer shows skipped count.
func TestLiveSkippedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestSkippedEvt(skippedEvt("ParentTest/testName", "somePackage"))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"⏩ ParentTest/testName\n"+
			"0 passed · 1 skipped")
}

// A multi-line BDD name keeps the duration on the first line.
func TestLiveMultiLineName(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/Func\nGiven X\nThen Y", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/Func\nGiven X\nThen Y", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"✅ ParentTest/Func (10ms)\nGiven X\nThen Y\n"+
			"1 passed")
}

// At TestingFinished the live footer is dropped and the final summary is
// committed as permanent output.
func TestLiveFinalSummaryReplacesFooter(t *testing.T) {
	interactor, term := setupLive(80, 30)
	t1 := time.Now()
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(t1))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/testName", "somePackage", 0.01))
	interactor.HandleTestingFinished(events.NewTestingFinishedEvent(t1.Add(1200 * time.Millisecond)))

	wantText(t, term,
		"🚀 Starting...\n\n📦 somePackage\n"+
			"✅ ParentTest/testName (10ms)\n"+
			"✓ All tests passed\n"+
			"Packages: 1 passed, 1 total\n"+
			"Tests:    1 passed, 1 total (100% passed)\n"+
			"Time:     1.200s\n"+
			"Ran all tests.")
}
