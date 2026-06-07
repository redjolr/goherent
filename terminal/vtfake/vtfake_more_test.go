package vtfake_test

import (
	"reflect"
	"testing"

	"github.com/redjolr/goherent/terminal/vtfake"
)

// ---- helpers ----

func vis(t *testing.T, term *vtfake.Terminal, want ...string) {
	t.Helper()
	if got := term.VisibleLines(); !reflect.DeepEqual(got, want) {
		t.Errorf("visible:\n got  = %#v\n want = %#v", got, want)
	}
}

func curs(t *testing.T, term *vtfake.Terminal, x, y int) {
	t.Helper()
	if term.CursorX() != x || term.CursorY() != y {
		t.Errorf("cursor: got (x=%d y=%d), want (x=%d y=%d)", term.CursorX(), term.CursorY(), x, y)
	}
}

func scrollEq(t *testing.T, term *vtfake.Terminal, want ...string) {
	t.Helper()
	if got := term.ScrollbackLines(); !reflect.DeepEqual(got, []string(want)) {
		t.Errorf("scrollback:\n got  = %#v\n want = %#v", got, []string(want))
	}
}

// ---- printing ----

func TestPrintSingleChar(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("A")
	vis(t, term, "A", "")
	curs(t, term, 1, 0)
	scrollEq(t, term)
}

func TestPrintWord(t *testing.T) {
	term := vtfake.New(10, 2)
	term.Print("Hello")
	vis(t, term, "Hello", "")
	curs(t, term, 5, 0)
}

func TestPrintAccumulatesAcrossCalls(t *testing.T) {
	term := vtfake.New(12, 2)
	term.Print("Hello ")
	term.Print("World")
	vis(t, term, "Hello World", "")
	curs(t, term, 11, 0)
}

func TestPrintWideGlyphTakesTwoColumns(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("🚀")
	vis(t, term, "🚀", "")
	curs(t, term, 2, 0)
}

func TestPrintWideThenNarrow(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("🚀A")
	vis(t, term, "🚀A", "")
	curs(t, term, 3, 0)
}

func TestPrintNarrowWideNarrow(t *testing.T) {
	term := vtfake.New(8, 2)
	term.Print("A🚀B")
	vis(t, term, "A🚀B", "")
	curs(t, term, 4, 0)
}

// ---- newlines & scrolling ----

func TestNewlineMovesToNextRowColumnZero(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A\nB")
	vis(t, term, "A", "B", "")
	curs(t, term, 1, 1)
	scrollEq(t, term)
}

func TestNewlineAsSeparateCall(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A")
	term.Print("\n")
	term.Print("B")
	vis(t, term, "A", "B", "")
	curs(t, term, 1, 1)
}

func TestFillExactlyNoScroll(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("L1\nL2\nL3")
	vis(t, term, "L1", "L2", "L3")
	curs(t, term, 2, 2)
	scrollEq(t, term)
}

func TestOverflowByOneScrollsOnce(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("L1\nL2\nL3\nL4")
	vis(t, term, "L2", "L3", "L4")
	scrollEq(t, term, "L1")
	curs(t, term, 2, 2)
}

func TestMultipleScrolls(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("L1\nL2\nL3\nL4")
	vis(t, term, "L3", "L4")
	scrollEq(t, term, "L1", "L2")
	curs(t, term, 2, 1)
}

// ---- carriage return ----

func TestCarriageReturnOverwrites(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABC\rZ")
	vis(t, term, "ZBC", "")
	curs(t, term, 1, 0)
}

func TestCarriageReturnThenNewline(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABC\r\nX")
	vis(t, term, "ABC", "X")
	curs(t, term, 1, 1)
}

// ---- cursor up / down ----

func TestCursorUpOverwritesPreviousRow(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A\nB\nC\033[1AX")
	vis(t, term, "A", "BX", "C")
	curs(t, term, 2, 1)
}

func TestCursorUpClampsAtTopRow(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A\nB\033[5AX")
	vis(t, term, "AX", "B", "")
	curs(t, term, 2, 0)
}

func TestCursorDownClampsAtBottomRow(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A\033[9BX")
	vis(t, term, "A", "", " X")
	curs(t, term, 2, 2)
}

// ---- cursor left / right ----

func TestCursorLeftOverwrites(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("Hello\033[1Da")
	vis(t, term, "Hella", "")
	curs(t, term, 5, 0)
}

func TestCursorLeftClampsAtColumnZero(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("AB\033[9DZ")
	vis(t, term, "ZB", "")
	curs(t, term, 1, 0)
}

func TestCursorRightLeavesGap(t *testing.T) {
	term := vtfake.New(8, 2)
	term.Print("A\033[2CB")
	vis(t, term, "A  B", "")
	curs(t, term, 4, 0)
}

func TestCursorRightClampsAtRightMargin(t *testing.T) {
	term := vtfake.New(4, 2)
	term.Print("A\033[9CB")
	vis(t, term, "A  B", "")
	curs(t, term, 3, 0)
}

// ---- cursor position (CUP) ----

func TestCursorPosition(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("\033[2;3HX")
	vis(t, term, "", "  X", "")
	curs(t, term, 3, 1)
}

func TestCursorPositionHomeDefault(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABC\033[HZ")
	vis(t, term, "ZBC", "")
	curs(t, term, 1, 0)
}

func TestCursorPositionClamps(t *testing.T) {
	term := vtfake.New(4, 3)
	term.Print("\033[9;9HX")
	vis(t, term, "", "", "   X")
	curs(t, term, 3, 2)
}

// ---- erase in display ----

func TestEraseDisplayToEndFromHome(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("AB\nCD\033[H\033[0J")
	vis(t, term, "", "")
	curs(t, term, 0, 0)
}

func TestEraseDisplayAllLeavesCursor(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("AB\nCD\033[2JX")
	vis(t, term, "", "  X")
	curs(t, term, 3, 1)
}

func TestEraseDisplayToStart(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("AAA\nBBB\033[1;2H\033[1J")
	vis(t, term, "  A", "BBB")
	curs(t, term, 1, 0)
}

// ---- erase in line ----

func TestEraseLineToEnd(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABCDEF\033[3D\033[K")
	vis(t, term, "AB", "")
	curs(t, term, 2, 0)
}

func TestEraseLineToStart(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABCDEF\033[3D\033[1K")
	vis(t, term, "   DEF", "")
	curs(t, term, 2, 0)
}

func TestEraseLineAll(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("ABC\033[2K")
	vis(t, term, "", "")
	curs(t, term, 3, 0)
}

// ---- save / restore (DECSC/DECRC) ----

func TestSaveRestoreNoScroll(t *testing.T) {
	term := vtfake.New(8, 2)
	term.Print("AB\0337CD\0338X")
	vis(t, term, "ABXD", "")
	curs(t, term, 3, 0)
}

// ---- double-width glyph overwrite (orphan handling) ----

func TestOverwriteWideGlyphRightHalfBlanksLeft(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("⏳\033[1D✅")
	vis(t, term, " ✅", "")
	curs(t, term, 3, 0)
}

func TestOverwriteWideGlyphLeftHalfBlanksRight(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("A⏳\033[2DX")
	vis(t, term, "AX", "")
	curs(t, term, 2, 0)
}

// ---- SGR colors are zero-width ----

func TestSgrIsZeroWidthAndStripped(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("\033[31mAB\033[0mC")
	vis(t, term, "ABC", "")
	curs(t, term, 3, 0)
}

func TestSgrDoesNotMoveCursor(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("A\033[1;31mB")
	vis(t, term, "AB", "")
	curs(t, term, 2, 0)
}

// ---- autowrap ----

func TestWrapExactFitParksCursor(t *testing.T) {
	term := vtfake.New(4, 3)
	term.Print("ABCD")
	vis(t, term, "ABCD", "", "")
	curs(t, term, 3, 0)
	scrollEq(t, term)
}

func TestWrapToNextRow(t *testing.T) {
	term := vtfake.New(4, 3)
	term.Print("ABCDE")
	vis(t, term, "ABCD", "E", "")
	curs(t, term, 1, 1)
}

func TestWideGlyphWrapsWhenItDoesNotFit(t *testing.T) {
	term := vtfake.New(3, 3)
	term.Print("AB🚀")
	vis(t, term, "AB", "🚀", "")
	curs(t, term, 2, 1)
}

// ---- terminal.Terminal move methods ----

func TestMoveLeftMethod(t *testing.T) {
	term := vtfake.New(8, 2)
	term.Print("Hello")
	term.MoveLeft(2)
	term.Print("X")
	vis(t, term, "HelXo", "")
	curs(t, term, 4, 0)
}

func TestMoveUpMethod(t *testing.T) {
	term := vtfake.New(6, 3)
	term.Print("A\nB\nC")
	term.MoveUp(2)
	term.Print("Z")
	vis(t, term, "AZ", "B", "C")
	curs(t, term, 2, 0)
}

// ---- Text() rendering ----

func TestTextTrimsTrailingBlankLines(t *testing.T) {
	term := vtfake.New(6, 4)
	term.Print("A\nB")
	if got := term.Text(); got != "A\nB" {
		t.Errorf("Text() = %q, want %q", got, "A\nB")
	}
}

func TestTextTrimsTrailingSpaces(t *testing.T) {
	term := vtfake.New(6, 2)
	term.Print("AB")
	if got := term.Text(); got != "AB" {
		t.Errorf("Text() = %q, want %q", got, "AB")
	}
}
