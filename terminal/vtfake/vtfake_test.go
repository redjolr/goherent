package vtfake_test

import (
	"reflect"
	"testing"

	"github.com/redjolr/goherent/terminal/vtfake"
)

// The expectations below were recorded from a REAL terminal with
// capture_terminal.sh. Each test feeds the same byte sequence to the emulator
// and asserts the visible grid, scrollback, and cursor match what the real
// terminal produced. If the emulator drifts from real behavior, these fail.

func check(t *testing.T, term *vtfake.Terminal, visible, scrollback []string, cx, cy int) {
	t.Helper()
	if got := term.VisibleLines(); !reflect.DeepEqual(got, visible) {
		t.Errorf("visible:\n got  = %#v\n want = %#v", got, visible)
	}
	if got := term.ScrollbackLines(); !reflect.DeepEqual(got, scrollback) {
		t.Errorf("scrollback:\n got  = %#v\n want = %#v", got, scrollback)
	}
	if term.CursorX() != cx || term.CursorY() != cy {
		t.Errorf("cursor: got (x=%d y=%d), want (x=%d y=%d)", term.CursorX(), term.CursorY(), cx, cy)
	}
}

// exp 1: '\n' on the bottom row scrolls; top rows go to scrollback.
func TestScrollOnNewlineAtBottom(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("L1\nL2\nL3\nL4\nL5")
	check(t, term, []string{"L3", "L4", "L5"}, []string{"L1", "L2"}, 2, 2)
}

// exp 2: cursor down past the bottom clamps, no scroll.
func TestCursorDownClampsAtBottom(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("L1\nL2\nL3\033[3BX")
	check(t, term, []string{"L1", "L2", "L3X"}, nil, 3, 2)
}

// exp 3: cursor up past the top clamps at row 0.
func TestCursorUpClampsAtTop(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("AAA\033[9AX")
	check(t, term, []string{"AAAX", "", ""}, nil, 4, 0)
}

// exp 4: \033[0J erases from the cursor (inclusive) to end of screen.
func TestEraseToEndOfDisplay(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("AAA\nBBB\nCCC\033[2;2H\033[0J")
	check(t, term, []string{"AAA", "B", ""}, nil, 1, 1)
}

// exp 5: \r returns to column 0 on the same row.
func TestCarriageReturn(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABCDEF\rXY")
	check(t, term, []string{"XYCDEF", ""}, nil, 2, 0)
}

// exp 6a: DECSC/DECRC save & restore the ABSOLUTE cursor cell — it does not
// follow content across a scroll (Z lands on L4, the content that scrolled into
// the saved row, not on L2).
func TestSaveRestoreIsAbsoluteAcrossScroll(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("L1\nL2\0337\nL3\nL4\nL5\0338Z")
	check(t, term, []string{"L3", "L4Z", "L5"}, []string{"L1", "L2"}, 3, 1)
}

// exp 7: autowrap at the right margin.
func TestAutowrap(t *testing.T) {
	term := vtfake.New(4, 3)
	term.Print("ABCDEFG")
	check(t, term, []string{"ABCD", "EFG", ""}, nil, 3, 1)
}

// exp 7b: autowrap on the bottom row scrolls.
func TestAutowrapScrolls(t *testing.T) {
	term := vtfake.New(4, 2)
	term.Print("ABCDEFGHIJ")
	check(t, term, []string{"EFGH", "IJ"}, []string{"ABCD"}, 2, 1)
}
