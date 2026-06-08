package sequential_events_test

import (
	"strings"
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
//
// Every entry (package header, each test, the footer) is separated by a blank
// line.

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

// While a test runs, the live block shows the spinner (first frame) + the
// running test (footer is empty until something finishes).
func TestLiveRunningTestIsShown(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))

	wantText(t, term, "🚀 Starting...\n\n╭─ 📦 somePackage\n\n│   🕐 ParentTest/testName")
}

// While a multi-line (BDD) test runs, the live block shows only its first line,
// so the live region stays short enough to redraw in place. The full message is
// committed once the test finishes.
func TestLiveRunningMultiLineNameIsCompact(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/Func\nGiven X\nThen Y", "somePackage"))

	wantText(t, term, "🚀 Starting...\n\n╭─ 📦 somePackage\n\n│   🕐 ParentTest/Func")
}

// Each tick advances the running test's spinner frame, redrawn in place.
func TestLiveSpinnerAdvancesOnTick(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))

	wantText(t, term, "🚀 Starting...\n\n╭─ 📦 somePackage\n\n│   🕐 ParentTest/testName")
	interactor.HandleTick()
	wantText(t, term, "🚀 Starting...\n\n╭─ 📦 somePackage\n\n│   🕑 ParentTest/testName")
	interactor.HandleTick()
	wantText(t, term, "🚀 Starting...\n\n╭─ 📦 somePackage\n\n│   🕒 ParentTest/testName")
}

// A tick does nothing when no test is running.
func TestLiveTickIsNoopWhenNoTestRunning(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/testName", "somePackage", 0.01))

	before := term.Text()
	interactor.HandleTick()
	if after := term.Text(); after != before {
		t.Errorf("tick changed output with no test running:\n before = %q\n after  = %q", before, after)
	}
}

// A passed test is committed with its duration on the first line, and the footer
// shows the running tally.
func TestLivePassedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/testName", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/testName (10ms)\n\n"+
			"1 passed")
}

// Two passed tests: both committed, footer counts up in place (no stacking),
// with a blank line between every entry.
func TestLiveTwoPassedTests(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/test1", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/test1", "somePackage", 0.01))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/test2", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/test2", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/test1 (10ms)\n\n"+
			"│   ✅ ParentTest/test2 (10ms)\n\n"+
			"2 passed")
}

// A failed test is committed; the footer shows passed and failed counts.
func TestLiveFailedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestFailedEvt(failedEvt("ParentTest/testName", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ❌ ParentTest/testName (10ms)\n\n"+
			"0 passed · 1 failed")
}

// A skipped test (no duration) is committed; the footer shows skipped count.
func TestLiveSkippedTest(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/testName", "somePackage"))
	interactor.HandleCtestSkippedEvt(skippedEvt("ParentTest/testName", "somePackage"))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ⏩ ParentTest/testName\n\n"+
			"0 passed · 1 skipped")
}

// A multi-line BDD name keeps the duration on the first line; the body lines stay
// together (no blank line inside the block).
func TestLiveMultiLineName(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/Func\nGiven X\nThen Y", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/Func\nGiven X\nThen Y", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/Func (10ms)\n│   Given X\n│   Then Y\n\n"+
			"1 passed")
}

// A trailing newline in the test message must not add an extra blank line.
func TestLiveTrimsTrailingNewlineInName(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/Func\n  Given X\n", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/Func\n  Given X\n", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/Func (10ms)\n│     Given X\n\n"+
			"1 passed")
}

// Extra leading blank lines in the test message must not add an extra blank line
// after the icon line.
func TestLiveTrimsLeadingBlankLinesInName(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/Func\n\n  Given X", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/Func\n\n  Given X", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/Func (10ms)\n│     Given X\n\n"+
			"1 passed")
}

// Regardless of stray newlines in individual messages, consecutive test blocks
// are separated by exactly one blank line.
func TestLiveUniformSpacingBetweenTests(t *testing.T) {
	interactor, term := setupLive(80, 30)
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(time.Now()))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/T1\n  msg one\n", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/T1\n  msg one\n", "somePackage", 0.01))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/T2\n  msg two", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/T2\n  msg two", "somePackage", 0.01))

	wantText(t, term,
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/T1 (10ms)\n│     msg one\n\n"+
			"│   ✅ ParentTest/T2 (10ms)\n│     msg two\n\n"+
			"2 passed")
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
		"🚀 Starting...\n\n╭─ 📦 somePackage\n\n"+
			"│   ✅ ParentTest/testName (10ms)\n\n"+
			// Closing rule spans the header width: DisplayWidth("╭─ 📦 somePackage") == 17.
			"╰" + strings.Repeat("─", 16) + "\n\n"+
			"✓ All tests passed\n"+
			"Packages: 1 passed, 1 total\n"+
			"Tests:    1 passed, 1 total (100% passed)\n"+
			"Time:     1.200s\n"+
			"Ran all tests.\n\n"+
			"🐢 1 slowest test:\n"+
			"  (10ms) ParentTest/testName")
}

// At the end of a run the slowest tests are listed, slowest first. vtfake strips
// the (zero-width) color codes, so only the plain durations and names show.
func TestLiveSlowestTestsReport(t *testing.T) {
	interactor, term := setupLive(80, 40)
	t1 := time.Now()
	interactor.HandleTestingStarted(events.NewTestingStartedEvent(t1))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/fast", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/fast", "somePackage", 0.02))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/slow", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/slow", "somePackage", 1.5))
	interactor.HandleCtestRanEvt(ranEvt("ParentTest/medium", "somePackage"))
	interactor.HandleCtestPassedEvt(passedEvt("ParentTest/medium", "somePackage", 0.5))
	interactor.HandleTestingFinished(events.NewTestingFinishedEvent(t1.Add(2 * time.Second)))

	got := term.Text()
	want := "🐢 3 slowest tests:\n" +
		"  (1.50s) ParentTest/slow\n" +
		"  (500ms) ParentTest/medium\n" +
		"  (20ms) ParentTest/fast"
	if !strings.Contains(got, want) {
		t.Errorf("slowest report missing or out of order:\n got  = %q\n want substring = %q", got, want)
	}
}
