package liveregion_test

import (
	"reflect"
	"testing"

	"github.com/redjolr/goherent/terminal/liveregion"
	"github.com/redjolr/goherent/terminal/vtfake"
)

func vis(t *testing.T, term *vtfake.Terminal, want ...string) {
	t.Helper()
	if got := term.VisibleLines(); !reflect.DeepEqual(got, want) {
		t.Errorf("visible:\n got  = %#v\n want = %#v", got, want)
	}
}

func scrollEq(t *testing.T, term *vtfake.Terminal, want ...string) {
	t.Helper()
	if got := term.ScrollbackLines(); !reflect.DeepEqual(got, []string(want)) {
		t.Errorf("scrollback:\n got  = %#v\n want = %#v", got, []string(want))
	}
}

func TestFirstLiveDraw(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("footer1")
	vis(t, term, "footer1", "", "", "", "")
}

func TestLiveUpdatesInPlace(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("footer one")
	r.SetLive("footer two")
	vis(t, term, "footer two", "", "", "", "")
	scrollEq(t, term)
}

func TestLiveShrinksAndClearsLeftovers(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("a long footer")
	r.SetLive("short")
	vis(t, term, "short", "", "", "", "")
}

func TestMultiLineLiveUpdatesInPlace(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("live1\nlive2")
	r.SetLive("only one line")
	vis(t, term, "only one line", "", "", "", "")
}

func TestCommitPushesAboveLive(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("ftr")
	r.Render("commitA", "ftr")
	vis(t, term, "commitA", "ftr", "", "", "")
}

func TestMultipleCommitsAccumulate(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("ftr")
	r.Render("A", "ftr")
	r.Render("B", "ftr")
	r.Render("C", "ftr")
	vis(t, term, "A", "B", "C", "ftr", "")
}

// The running-test -> passed flow: a multi-line live block ("⏳ test" + footer)
// is replaced by committing "✅ test" and shrinking the live block to the footer.
func TestCommitReplacesRunningTestInLive(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("⏳ test\n\nftr")
	r.Render("✅ test", "ftr")
	vis(t, term, "✅ test", "ftr", "", "", "")
}

// The key property: committing past the bottom scrolls the committed area into
// scrollback while the live block stays pinned to the bottom — correct because
// the redraw uses relative moves, not absolute save/restore.
func TestCommitScrollsCommittedAreaKeepingLiveAtBottom(t *testing.T) {
	term := vtfake.New(10, 3)
	r := liveregion.New(term)
	r.SetLive("ftr")
	r.Render("L1", "ftr")
	r.Render("L2", "ftr")
	r.Render("L3", "ftr")
	vis(t, term, "L2", "L3", "ftr")
	scrollEq(t, term, "L1")
}

func TestCommitAndUpdateFooterTogether(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.SetLive("0 passed")
	r.Render("✅ A", "1 passed")
	r.Render("✅ B", "2 passed")
	vis(t, term, "✅ A", "✅ B", "2 passed", "", "")
}

func TestFirstRenderWithCommitted(t *testing.T) {
	term := vtfake.New(20, 5)
	r := liveregion.New(term)
	r.Render("🚀 Starting...", "0 passed")
	vis(t, term, "🚀 Starting...", "0 passed", "", "", "")
}

// A live block taller than the viewport is clamped to its bottom lines so the
// in-place redraw stays correct (it can't move above the top of the screen).
func TestLiveBlockTallerThanViewportIsClamped(t *testing.T) {
	term := vtfake.New(20, 2)
	r := liveregion.New(term)
	r.SetLive("running\n\nfooter") // 3 lines into a height-2 viewport
	vis(t, term, "", "footer")     // bottom 2 lines kept
}

// And it still updates in place once clamped — no stacking from an over-tall
// block whose top scrolled off.
func TestClampedLiveBlockUpdatesInPlace(t *testing.T) {
	term := vtfake.New(20, 2)
	r := liveregion.New(term)
	r.SetLive("running\n\nfooter1")
	r.SetLive("running\n\nfooter2")
	vis(t, term, "", "footer2")
	scrollEq(t, term)
}

// Down to a single-row viewport, only the last line survives.
func TestLiveBlockClampedToSingleRow(t *testing.T) {
	term := vtfake.New(20, 1)
	r := liveregion.New(term)
	r.SetLive("a\nb\nc")
	vis(t, term, "c")
}
