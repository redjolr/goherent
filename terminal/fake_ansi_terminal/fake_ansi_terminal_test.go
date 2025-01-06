package fake_ansi_terminal_test

import (
	"math"
	"testing"

	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	. "github.com/redjolr/goherent/test"
)

func TestNewFakeAnsiTerminal(t *testing.T) {
	Test("it should return an instance of type FakeAnsiTerminal", func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(10, 10)
		Expect(fakeTerminal).ToBeOfSameTypeAs(fake_ansi_terminal.FakeAnsiTerminal{})
	}, t)
}

func TestPrintBasic(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A "
	Then the terminal should store "A "`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A ")
		Expect(fakeTerminal.Text()).ToEqual("A ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀",
	Then the terminal should store "🚀"	`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀")
		Expect(fakeTerminal.Text()).ToEqual("🚀")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print"\n🚀", if we print
	Then the terminal should store "\n🚀"`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n🚀")
		Expect(fakeTerminal.Text()).ToEqual("\n🚀")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀A"
	Then the terminal should store "🚀A"`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀A")
		Expect(fakeTerminal.Text()).ToEqual("🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀🚀A",
	Then the terminal should store "🚀🚀A"`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀🚀A")
		Expect(fakeTerminal.Text()).ToEqual("🚀🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A🚀🚀A",
	Then the terminal should store "A🚀🚀A"	`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A🚀🚀A")
		Expect(fakeTerminal.Text()).ToEqual("A🚀🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n",
	Then the terminal should store "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n"	`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
		Expect(fakeTerminal.Text()).ToEqual("\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"
	Then the terminal should store "Hello".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		Expect(fakeTerminal.Text()).ToEqual("Hello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello " and then "World"
	Then the terminal should store "Hello World"
		if we print .`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello ")
		fakeTerminal.Print("World")
		Expect(fakeTerminal.Text()).ToEqual("Hello World")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWorld"
	Then the terminal should store "Hello\nWorld".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		Expect(fakeTerminal.Text()).ToEqual("Hello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello" and then "\n" and then "World"
	Then the terminal should store "Hello\nWorld".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print("\n")
		fakeTerminal.Print("World")
		Expect(fakeTerminal.Text()).ToEqual("Hello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AAA\n", then "AAA" then MoveCursorUpNRows(1), then "\n", then "B"
	Then the terminal should store "AAA\nBAA",
	`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AAA\n")
		fakeTerminal.Print("AAA")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("\n")
		fakeTerminal.Print("B")
		Expect(fakeTerminal.Text()).ToEqual("AAA\nBAA")
	}, t)
}

func TestPrintEraseScreen(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print the ERASE_SCREEN ansi code,
	Then the terminal should store an empty string.`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual("")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A", and then ERASE_SCREEN.
	Then the terminal should store " ".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual(" ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print"AB", and then ERASE_SCREEN
	Then the terminal should store "  "`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual("  ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB", and then ERASE_SCREEN,
	Then the terminal should store ""`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual("\n ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print ERASE_SCREEN and then print "A"
	Then the terminal should store "A".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("A")
		Expect(fakeTerminal.Text()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print print "A", ERASE_SCREEN, and then "B"
	Then the terminal should store "B".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("B")
		Expect(fakeTerminal.Text()).ToEqual(" B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB", and then ERASE_SCREEN and then print "C"
	Then the terminal should store "  C"`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("C")
		Expect(fakeTerminal.Text()).ToEqual("  C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN
	Then the terminal should store "   ".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual("\n   ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN and then print "A"
	Then the terminal should store  "   A".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("A")
		Expect(fakeTerminal.Text()).ToEqual("\n   A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN and then print "Line 1\nLine 2"
	Then the terminal should store "   Line 1\nLine 2"`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("Line 1\nLine 2")
		Expect(fakeTerminal.Text()).ToEqual("\n   Line 1\nLine 2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "     OTHER LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OTHER LINE")
		Expect(fakeTerminal.Text()).ToEqual("\n     OTHER LINE")
	}, t)

	Test(`
		Given that there is a terminal with an height=5 and width=infinity
		When we print "Line1\nLine2\nLine3", and then ERASE_SCREEN, and then "OVERWRITE LINE"
		Then the terminal should store "     OVERWRITE LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\n     ".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		Expect(fakeTerminal.Text()).ToEqual("Line1\n\n\n     ")
	}, t)

	Test(`
		Given that there is a terminal with an height=3 and width=infinity
		When we print "Line1\nLine2\nLine3\nLine4", and then ERASE_SCREEN, and then "OVERWRITE LINE"
		Then the terminal should store "Line1\n     OVERWRITE LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\n\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\n\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(1), then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\n\n     OVERWRITE LINE\n")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(2), then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\n     OVERWRITE LINE\n\n")
	}, t)
}

func TestPrintCursorToHome(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "J"
	Then the terminal should store "Jello".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "J")
		Expect(fakeTerminal.Text()).ToEqual("Jello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then"J"
	Then the terminal should store "Jello".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("J")
		Expect(fakeTerminal.Text()).ToEqual("Jello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Cond"
	Then the terminal should store "Condo".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Cond")
		Expect(fakeTerminal.Text()).ToEqual("Condo")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Cond"
	Then the terminal should store "Condo".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Cond")
		Expect(fakeTerminal.Text()).ToEqual("Condo")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Candy"
	Then the terminal should store "Condo".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Candy")
		Expect(fakeTerminal.Text()).ToEqual("Candy")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Candy"
	Then the terminal should store "Candy".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Candy")
		Expect(fakeTerminal.Text()).ToEqual("Candy")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Granny"
	Then the terminal should store "Granny".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Granny")
		Expect(fakeTerminal.Text()).ToEqual("Granny")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Granny"
	Then the terminal should store "Granny".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Granny")
		Expect(fakeTerminal.Text()).ToEqual("Granny")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWor;d", and then CURSOR_TO_HOME+"J"
	Then the terminal should store "Jello\nWorld".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "J")
		Expect(fakeTerminal.Text()).ToEqual("Jello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWorld", and then CURSOR_TO_HOME+"J"
	Then the terminal should store "Candy\nWorld".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Candy")
		Expect(fakeTerminal.Text()).ToEqual("Candy\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "OVERWRITE LINE\nLine2".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("OVERWRITE LINE\nLine2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "OVERWRITE LINE\nLine2\nLine3".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("OVERWRITE LINE\nLine2\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nOVERWRITE LINE\nLine3\nLine4".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nOVERWRITE LINE\nLine3\nLine4")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(1), then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(2), then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)
}

func TestPrintMoveCursorLeft(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "⏳"+MoveCursorLeftNCols(1)+"✅".
	Then the terminal should store "✅".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("⏳" + ansi_escape.MoveCursorLeftNCols(1) + "✅")
		Expect(fakeTerminal.Text()).ToEqual("✅")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "H"+MoveCursorLeftNCols(1)+"A"
	Then the terminal should store "A".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("H" + ansi_escape.MoveCursorLeftNCols(1) + "A")
		Expect(fakeTerminal.Text()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "H"+MoveCursorLeftNCols(0)+"A"
	Then the terminal should store "HA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("H" + ansi_escape.MoveCursorLeftNCols(0) + "A")
		Expect(fakeTerminal.Text()).ToEqual("HA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "H"+MoveCursorLeftNCols(2)+"A"
	Then the terminal should store "H".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("H" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		Expect(fakeTerminal.Text()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(n>>1)+"A" ,
	Then the terminal should store "AR".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		Expect(fakeTerminal.Text()).ToEqual("AR")

		fakeTerminal = fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(10000) + "A")
		Expect(fakeTerminal.Text()).ToEqual("AR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(1)+"A"
	Then the terminal should store "RA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(1) + "A")
		Expect(fakeTerminal.Text()).ToEqual("RA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"A"
	Then the terminal should store "RAR".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		Expect(fakeTerminal.Text()).ToEqual("RAR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"AA"
	Then the terminal should store "RAA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "AA")
		Expect(fakeTerminal.Text()).ToEqual("RAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"AAA"
	Then the terminal should store "RAAA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "AAA")
		Expect(fakeTerminal.Text()).ToEqual("RAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(0)+"A" 
	Then the terminal should store "RRA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(0) + "A")
		Expect(fakeTerminal.Text()).ToEqual("RRA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorLeftNCols(1)+"a"
	Then the terminal should store "Hella".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(1) + "a")
		Expect(fakeTerminal.Text()).ToEqual("Hella")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorLeftNCols(5)+"Juice"
	Then the terminal should store "Juice".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(5) + "Juice")
		Expect(fakeTerminal.Text()).ToEqual("Juice")
	}, t)
}

func TestPrintMoveCursorRight(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorRightNCols(1)
	Then the terminal should store  "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorRightNCols(1))
		Expect(fakeTerminal.Text()).ToEqual("")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorRightNCols(3)
	Then the terminal should store "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorRightNCols(3))
		Expect(fakeTerminal.Text()).ToEqual("")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorRightNCols(1)+"R"
	Then the terminal should store "R R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorRightNCols(1) + "R")
		Expect(fakeTerminal.Text()).ToEqual("R R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorRightNCols(3)+"World"
	Then the terminal should store "Hello   World".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorRightNCols(3) + "World")
		Expect(fakeTerminal.Text()).ToEqual("Hello   World")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB"+MoveCursorRightNCols(0) + "C"
	Then the terminal should store "AC\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorRightNCols(0) + "C")
		Expect(fakeTerminal.Text()).ToEqual("A C")
	}, t)
}

func TestPrintMoveCursorLeftAndRight(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print  MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)
	Then the terminal should store "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1))
		Expect(fakeTerminal.Text()).ToEqual("")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1),
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1))
		Expect(fakeTerminal.Text()).ToEqual("R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(2)
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(2))
		Expect(fakeTerminal.Text()).ToEqual("R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)+"R"
	Then the terminal should store "RR".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1) + "R")
		Expect(fakeTerminal.Text()).ToEqual("RR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(4)
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(4))
		Expect(fakeTerminal.Text()).ToEqual("R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+ MoveCursorLeftNCols(2) + MoveCursorRightNCols(1) + "p"
	Then the terminal should store "Hellp".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(2) + ansi_escape.MoveCursorRightNCols(1) + "p")
		Expect(fakeTerminal.Text()).ToEqual("Hellp")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+ MoveCursorLeftNCols(2) + "ix" + MoveCursorRightNCols(1) + "shaped" 
	Then the terminal should store "Helix shaped".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(2) + "ix" + ansi_escape.MoveCursorRightNCols(1) + "shaped")
		Expect(fakeTerminal.Text()).ToEqual("Helix shaped")
	}, t)
}

func TestPrintMoveCursorUp(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorUpNRows(1)
	Then the terminal should store "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorUpNRows(10)
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorUpNRows(10))
		Expect(fakeTerminal.Text()).ToEqual("R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB"+MoveCursorUpNRows(1)
	Then the terminal should store "A\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB" + ansi_escape.MoveCursorUpNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(1)+"B"
	Then the terminal should store "B\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(2)+"B"
	Then the terminal should store "B\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual("B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(10)+"B"
	Then the terminal should store "B\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(10) + "B")
		Expect(fakeTerminal.Text()).ToEqual("B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB\n"+MoveCursorUpNRows(1)+"CD"
	Then the terminal should store "CD\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB\n" + ansi_escape.MoveCursorUpNRows(1) + "CD")
		Expect(fakeTerminal.Text()).ToEqual("CD\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB\n"+MoveCursorUpNRows(1)+"CDE"
	Then the terminal should store "CDE\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB\n" + ansi_escape.MoveCursorUpNRows(1) + "CDE")
		Expect(fakeTerminal.Text()).ToEqual("CDE\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+MoveCursorUpNRows(1)
	Then the terminal should store "\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+MoveCursorUpNRows(1) + "A"
	Then the terminal should store "A\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1) + "A")
		Expect(fakeTerminal.Text()).ToEqual("A\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n" + MoveCursorUpNRows + "AAA"
	Then the terminal should store "AAA\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1) + "AAA")
		Expect(fakeTerminal.Text()).ToEqual("AAA\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(1) + "B"
	Then the terminal should store " B\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual(" B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(2) + "B"
	Then the terminal should store " B\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual(" B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(5) + "B"
	Then the terminal should store " B\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(5) + "B")
		Expect(fakeTerminal.Text()).ToEqual(" B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA" + MoveCursorUpNRows(2) + "B"
	Then the terminal should store " B\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA" + ansi_escape.MoveCursorUpNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual("  B\nAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA" + MoveCursorUpNRows(2) + "BBB"
	Then the terminal should store "     BBB\nAAAAA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA" + ansi_escape.MoveCursorUpNRows(2) + "BBB")
		Expect(fakeTerminal.Text()).ToEqual("     BBB\nAAAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA\nA" + MoveCursorUpNRows(1) + "B"
	Then the terminal should store "\nAB\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA\nA" + ansi_escape.MoveCursorUpNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("\nAB\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA\nBBBBB" + MoveCursorUpNRows(1) + "C"
	Then the terminal should store "\nAA   C\nBBBBB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA\nBBBBB" + ansi_escape.MoveCursorUpNRows(1) + "C")
		Expect(fakeTerminal.Text()).ToEqual("\nAA   C\nBBBBB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nA D\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\nA D\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\n  D\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\n  D\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAAD\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\nAAD\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\nAAD\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAADA\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\nAADA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		Expect(fakeTerminal.Text()).ToEqual("\nAADAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(3) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(3) + "D")
		Expect(fakeTerminal.Text()).ToEqual("  D\nAAAAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(4) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(4) + "D")
		Expect(fakeTerminal.Text()).ToEqual("  D\nAAAAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB"+MoveCursorUpNRows(0) + "C"
	Then the terminal should store "AC\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB" + ansi_escape.MoveCursorUpNRows(0) + "C")
		Expect(fakeTerminal.Text()).ToEqual("AC\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A"+MoveCursorUpNRows(0) + "C"
	Then the terminal should store "AC".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorUpNRows(0) + "C")
		Expect(fakeTerminal.Text()).ToEqual("AC")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then MoveCursorUpNRows(1), and then "HELLO"
	Then the terminal should store "OVERWRITE LINE\nLine2".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("HELLO")
		Expect(fakeTerminal.Text()).ToEqual("Line1HELLO\nLine2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(1), and then "HELLO"
	Then the terminal should store "Line1\nLine2HELLO\nLine3".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("HELLO")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2HELLO\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(2), and then "HELLO"
	Then the terminal should store "Line1HELLO\nLine2\nLine3".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print("HELLO")
		Expect(fakeTerminal.Text()).ToEqual("Line1HELLO\nLine2\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(3), and then "HELLO"
	Then the terminal should store "Line1HELLO\nLine2\nLine3".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(3))
		fakeTerminal.Print("HELLO")
		Expect(fakeTerminal.Text()).ToEqual("Line1HELLO\nLine2\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nOVERWRITE LINE\nLine3\nLine4".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(3))
		fakeTerminal.Print("HELLO")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2HELLO\nLine3\nLine4")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		Expect(fakeTerminal.Text()).ToEqual("Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)
}

func TestPrintMoveCursorDown(t *testing.T) {
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorDownNRows(1)
	Then the terminal should store "".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorDownNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorDownNRows(1)
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorDownNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("R\n ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorDownNRows(2)
	Then the terminal should store "R".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorDownNRows(2))
		Expect(fakeTerminal.Text()).ToEqual("R\n\n ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(1)+MoveCursorDwonNRows(1)+"B"
	Then the terminal should store "A\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(1) + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"B"
	Then the terminal should store "B\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(2)+"B"
	Then the terminal should store "\nA\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual("\nA\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(2)+"B"
	Then the terminal should store "B\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(5) + "B")
		Expect(fakeTerminal.Text()).ToEqual("\nA\n\n\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAB"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"CD"
	Then the terminal should store "\nCD".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAB" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "CD")
		Expect(fakeTerminal.Text()).ToEqual("\nCD")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAB"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"CDE"
	Then the terminal should store "\nCDE".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAB" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "CDE")
		Expect(fakeTerminal.Text()).ToEqual("\nCDE")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)
	Then the terminal should store "\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1))
		Expect(fakeTerminal.Text()).ToEqual("\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+ CURSOR_TO_HOME + MoveCursorDownNRows(1) + "A"
	Then the terminal should store "\nA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "A")
		Expect(fakeTerminal.Text()).ToEqual("\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "AAA"
	Then the terminal should store "\nAAA".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "AAA")
		Expect(fakeTerminal.Text()).ToEqual("\nAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(2) + "B"
	Then the terminal should store "A\n\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(5) + "B"
	Then the terminal should store "A\n\n\n\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(5) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\n\n\n\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A" + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\n B".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\n B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A" + MoveCursorDownNRows(2) + "B"
	Then the terminal should store "A\n\n B".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorDownNRows(2) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\n\n B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AA" + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "AA\n  B".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AA" + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("AA\n  B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AA" + MoveCursorDownNRows(3) + "B"
	Then the terminal should store "AA\n\n\n  B".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AA" + ansi_escape.MoveCursorDownNRows(3) + "B")
		Expect(fakeTerminal.Text()).ToEqual("AA\n\n\n  B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AAAAA" + MoveCursorDownNRows(2) + "BBB"
	Then the terminal should store "AAAAA\n\n     BBB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AAAAA" + ansi_escape.MoveCursorDownNRows(2) + "BBB")
		Expect(fakeTerminal.Text()).ToEqual("AAAAA\n\n     BBB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nA\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\nB\n".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nA\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		Expect(fakeTerminal.Text()).ToEqual("A\nB\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(1) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA C".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(1) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		Expect(fakeTerminal.Text()).ToEqual("BBBBB\nAA C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(2) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA  C".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(2) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		Expect(fakeTerminal.Text()).ToEqual("BBBBB\nAA  C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(3) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA  C".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(3) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		Expect(fakeTerminal.Text()).ToEqual("BBBBB\nAA   C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(6) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA C".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(6) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		Expect(fakeTerminal.Text()).ToEqual("BBBBB\nAA      C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB"+MoveCursorUpNRows(1) + MoveCursorDownNRows(0) "C"
	Then the terminal should store "AC\nB".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB" + ansi_escape.MoveCursorUpNRows(1) + ansi_escape.MoveCursorDownNRows(0) + "C")
		Expect(fakeTerminal.Text()).ToEqual("A\nBC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A"+MoveCursorDownNRows(0) + "C"
	Then the terminal should store "A\n C".`, func(Expect expect.F) {
		fakeTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorDownNRows(0) + "C")
		Expect(fakeTerminal.Text()).ToEqual("A\n C")
	}, t)
}
