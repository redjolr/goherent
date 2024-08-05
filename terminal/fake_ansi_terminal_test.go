package terminal_test

import (
	"math"
	"testing"

	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/stretchr/testify/assert"
)

func TestNewFakeAnsiTerminal(t *testing.T) {
	assert := assert.New(t)
	Test("it should return an instance of type FakeAnsiTerminal", func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(10, 10)
		assert.IsType(terminal.FakeAnsiTerminal{}, fakeTerminal)
	}, t)
}

func TestPrintBasic(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A "
	Then the terminal should store "A "`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A ")
		assert.Equal(fakeTerminal.Text(), "A ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀", 
	Then the terminal should store "🚀"	`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀")
		assert.Equal(fakeTerminal.Text(), "🚀")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print"\n🚀", if we print
	Then the terminal should store "\n🚀"`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n🚀")
		assert.Equal(fakeTerminal.Text(), "\n🚀")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀A"
	Then the terminal should store "🚀A"`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀A")
		assert.Equal(fakeTerminal.Text(), "🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "🚀🚀A", 
	Then the terminal should store "🚀🚀A"`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("🚀🚀A")
		assert.Equal(fakeTerminal.Text(), "🚀🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A🚀🚀A",
	Then the terminal should store "A🚀🚀A"	`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A🚀🚀A")
		assert.Equal(fakeTerminal.Text(), "A🚀🚀A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n",
	Then the terminal should store "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n"	`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
		assert.Equal(fakeTerminal.Text(), "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"
	Then the terminal should store "Hello".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello " and then "World"
	Then the terminal should store "Hello World"
		if we print .`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello ")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello World")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWorld"
	Then the terminal should store "Hello\nWorld".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello" and then "\n" and then "World"
	Then the terminal should store "Hello\nWorld".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print("\n")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AAA\n", then "AAA" then MoveCursorUpNRows(1), then "\n", then "B"
	Then the terminal should store "AAA\nBAA", 
	`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AAA\n")
		fakeTerminal.Print("AAA")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("\n")
		fakeTerminal.Print("B")
		assert.Equal(fakeTerminal.Text(), "AAA\nBAA")
	}, t)
}

func TestPrintEraseScreen(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print the ERASE_SCREEN ansi code,
	Then the terminal should store an empty string.`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A", and then ERASE_SCREEN.
	Then the terminal should store " ".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), " ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print"AB", and then ERASE_SCREEN
	Then the terminal should store "  "`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), "  ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB", and then ERASE_SCREEN,
	Then the terminal should store ""`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), "\n ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print ERASE_SCREEN and then print "A"
	Then the terminal should store "A".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print print "A", ERASE_SCREEN, and then "B"
	Then the terminal should store "B".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("B")
		assert.Equal(fakeTerminal.Text(), " B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB", and then ERASE_SCREEN and then print "C"
	Then the terminal should store "  C"`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("C")
		assert.Equal(fakeTerminal.Text(), "  C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN
	Then the terminal should store "   ".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), "\n   ")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN and then print "A"
	Then the terminal should store  "   A".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("A")
		assert.Equal(fakeTerminal.Text(), "\n   A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nBob", and then ERASE_SCREEN and then print "Line 1\nLine 2"
	Then the terminal should store "   Line 1\nLine 2"`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nBob")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("Line 1\nLine 2")
		assert.Equal(fakeTerminal.Text(), "\n   Line 1\nLine 2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "     OTHER LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OTHER LINE")
		assert.Equal(fakeTerminal.Text(), "\n     OTHER LINE")
	}, t)

	Test(`
		Given that there is a terminal with an height=5 and width=infinity
		When we print "Line1\nLine2\nLine3", and then ERASE_SCREEN, and then "OVERWRITE LINE"
		Then the terminal should store "     OVERWRITE LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\n     ".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		assert.Equal(fakeTerminal.Text(), "Line1\n\n\n     ")
	}, t)

	Test(`
		Given that there is a terminal with an height=3 and width=infinity
		When we print "Line1\nLine2\nLine3\nLine4", and then ERASE_SCREEN, and then "OVERWRITE LINE"
		Then the terminal should store "Line1\n     OVERWRITE LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\n\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\n\n\n     OVERWRITE LINE")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(1), then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\n\n     OVERWRITE LINE\n")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(2), then ERASE_SCREEN, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\n     OVERWRITE LINE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print(ansi_escape.ERASE_SCREEN)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\n     OVERWRITE LINE\n\n")
	}, t)
}

func TestPrintCursorToHome(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "J"
	Then the terminal should store "Jello".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then"J"
	Then the terminal should store "Jello".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Cond"
	Then the terminal should store "Condo".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Cond"
	Then the terminal should store "Condo".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Candy"
	Then the terminal should store "Condo".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Candy"
	Then the terminal should store "Candy".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME + "Granny"
	Then the terminal should store "Granny".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello", and then CURSOR_TO_HOME, and then "Granny"
	Then the terminal should store "Granny".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWor;d", and then CURSOR_TO_HOME+"J"
	Then the terminal should store "Jello\nWorld".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "J")
		assert.Equal(fakeTerminal.Text(), "Jello\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello\nWorld", and then CURSOR_TO_HOME+"J"
	Then the terminal should store "Candy\nWorld".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy\nWorld")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "OVERWRITE LINE\nLine2".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "OVERWRITE LINE\nLine2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "OVERWRITE LINE\nLine2\nLine3".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "OVERWRITE LINE\nLine2\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nOVERWRITE LINE\nLine3\nLine4".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nOVERWRITE LINE\nLine3\nLine4")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(1), then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3\nLine4\nLine5", then MoveCursorUpNRows(2), then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
		fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
		fakeTerminal.Print("OVERWRITE LINE")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	}, t)
}

func TestPrintMoveCursorLeft(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "⏳"+MoveCursorLeftNCols(1)+"✅".
	Then the terminal should store "✅".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("⏳" + ansi_escape.MoveCursorLeftNCols(1) + "✅")
		assert.Equal(fakeTerminal.Text(), "✅")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "H"+MoveCursorLeftNCols(1)+"A"
	Then the terminal should store "A".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("H" + ansi_escape.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "H"+MoveCursorLeftNCols(2)+"A"
	Then the terminal should store "H".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("H" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(n>>1)+"A" ,
	Then the terminal should store "AR".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")

		fakeTerminal = terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(10000) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(1)+"A"
	Then the terminal should store "RA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "RA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"A"
	Then the terminal should store "RAR".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "RAR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"AA"
	Then the terminal should store "RAA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "AA")
		assert.Equal(fakeTerminal.Text(), "RAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RRR"+MoveCursorLeftNCols(2)+"AAA"
	Then the terminal should store "RAAA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RRR" + ansi_escape.MoveCursorLeftNCols(2) + "AAA")
		assert.Equal(fakeTerminal.Text(), "RAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "RR"+MoveCursorLeftNCols(0)+"A" 
	Then the terminal should store "RRA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("RR" + ansi_escape.MoveCursorLeftNCols(0) + "A")
		assert.Equal(fakeTerminal.Text(), "RRA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorLeftNCols(1)+"a"
	Then the terminal should store "Hella".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(1) + "a")
		assert.Equal(fakeTerminal.Text(), "Hella")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorLeftNCols(5)+"Juice"
	Then the terminal should store "Juice".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(5) + "Juice")
		assert.Equal(fakeTerminal.Text(), "Juice")
	}, t)
}

func TestPrintMoveCursorRight(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorRightNCols(1)
	Then the terminal should store  "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorRightNCols(3)
	Then the terminal should store "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorRightNCols(3))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorRightNCols(1)+"R"
	Then the terminal should store "R R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "R R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+MoveCursorRightNCols(3)+"World"
	Then the terminal should store "Hello   World".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorRightNCols(3) + "World")
		assert.Equal(fakeTerminal.Text(), "Hello   World")
	}, t)
}

func TestPrintMoveCursorLeftAndRight(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print  MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)
	Then the terminal should store "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1),
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(2)
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)+"R"
	Then the terminal should store "RR".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "RR")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(4)
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorLeftNCols(1) + ansi_escape.MoveCursorRightNCols(4))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+ MoveCursorLeftNCols(2) + MoveCursorRightNCols(1) + "p"
	Then the terminal should store "Hellp".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(2) + ansi_escape.MoveCursorRightNCols(1) + "p")
		assert.Equal(fakeTerminal.Text(), "Hellp")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "Hello"+ MoveCursorLeftNCols(2) + "ix" + MoveCursorRightNCols(1) + "shaped" 
	Then the terminal should store "Helix shaped".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("Hello" + ansi_escape.MoveCursorLeftNCols(2) + "ix" + ansi_escape.MoveCursorRightNCols(1) + "shaped")
		assert.Equal(fakeTerminal.Text(), "Helix shaped")
	}, t)
}

func TestPrintMoveCursorUp(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorUpNRows(1)
	Then the terminal should store "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorUpNRows(10)
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorUpNRows(10))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nB"+MoveCursorUpNRows(1)
	Then the terminal should store "A\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nB" + ansi_escape.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(1)+"B"
	Then the terminal should store "B\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(2)+"B"
	Then the terminal should store "B\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(10)+"B"
	Then the terminal should store "B\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(10) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB\n"+MoveCursorUpNRows(1)+"CD"
	Then the terminal should store "CD\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB\n" + ansi_escape.MoveCursorUpNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "CD\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AB\n"+MoveCursorUpNRows(1)+"CDE"
	Then the terminal should store "CDE\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AB\n" + ansi_escape.MoveCursorUpNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "CDE\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+MoveCursorUpNRows(1)
	Then the terminal should store "\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+MoveCursorUpNRows(1) + "A"
	Then the terminal should store "A\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n" + MoveCursorUpNRows + "AAA"
	Then the terminal should store "AAA\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.MoveCursorUpNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "AAA\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(1) + "B"
	Then the terminal should store " B\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(2) + "B"
	Then the terminal should store " B\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA" + MoveCursorUpNRows(5) + "B"
	Then the terminal should store " B\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.MoveCursorUpNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA" + MoveCursorUpNRows(2) + "B"
	Then the terminal should store " B\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA" + ansi_escape.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "  B\nAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA" + MoveCursorUpNRows(2) + "BBB"
	Then the terminal should store "     BBB\nAAAAA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA" + ansi_escape.MoveCursorUpNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "     BBB\nAAAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA\nA" + MoveCursorUpNRows(1) + "B"
	Then the terminal should store "\nAB\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA\nA" + ansi_escape.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nAB\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA\nBBBBB" + MoveCursorUpNRows(1) + "C"
	Then the terminal should store "\nAA   C\nBBBBB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA\nBBBBB" + ansi_escape.MoveCursorUpNRows(1) + "C")
		assert.Equal(fakeTerminal.Text(), "\nAA   C\nBBBBB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nA D\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nA D\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\n  D\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\n  D\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAAD\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAADA\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(3) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(3) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(4) + "D"
	Then the terminal should store "\nAADAA\nB\nCC".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAAAAA\nB\nCC" + ansi_escape.MoveCursorUpNRows(4) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2", and then MoveCursorUpNRows(1), and then "HELLO"
	Then the terminal should store "OVERWRITE LINE\nLine2".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("HELLO")
		assert.Equal(fakeTerminal.Text(), "Line1HELLO\nLine2")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(1), and then "HELLO"
	Then the terminal should store "Line1\nLine2HELLO\nLine3".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(1))
		fakeTerminal.Print("HELLO")
		assert.Equal(fakeTerminal.Text(), "Line1\nLine2HELLO\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(2), and then "HELLO"
	Then the terminal should store "Line1HELLO\nLine2\nLine3".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(2))
		fakeTerminal.Print("HELLO")
		assert.Equal(fakeTerminal.Text(), "Line1HELLO\nLine2\nLine3")
	}, t)

	Test(`
	Given that there is a terminal with an height=3 and width=infinity
	When we print "Line1\nLine2\nLine3", and then MoveCursorUpNRows(3), and then "HELLO"
	Then the terminal should store "Line1HELLO\nLine2\nLine3".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 5)
		fakeTerminal.Print("Line1\nLine2\nLine3")
		fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(3))
		fakeTerminal.Print("HELLO")
		assert.Equal(fakeTerminal.Text(), "Line1HELLO\nLine2\nLine3")
	}, t)

	// Test(`
	// Given that there is a terminal with an height=3 and width=infinity
	// When we print "Line1\nLine2\nLine3\nLine4", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	// Then the terminal should store "Line1\nOVERWRITE LINE\nLine3\nLine4".`, func(t *testing.T) {
	// 	fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
	// 	fakeTerminal.Print("Line1\nLine2\nLine3\nLine4")
	// 	fakeTerminal.Print(ansi_escape.MoveCursorUpNRows(3))
	// 	fakeTerminal.Print("HELLO")
	// 	assert.Equal(fakeTerminal.Text(), "Line1\nLine2HELLO\nLine3\nLine4")
	// }, t)

	// Test(`
	// Given that there is a terminal with an height=3 and width=infinity
	// When we print "Line1\nLine2\nLine3\nLine4\nLine5", and then CURSOR_TO_HOME, and then "OVERWRITE LINE"
	// Then the terminal should store "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5".`, func(t *testing.T) {
	// 	fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, 3)
	// 	fakeTerminal.Print("Line1\nLine2\nLine3\nLine4\nLine5")
	// 	fakeTerminal.Print(ansi_escape.CURSOR_TO_HOME)
	// 	fakeTerminal.Print("OVERWRITE LINE")
	// 	assert.Equal(fakeTerminal.Text(), "Line1\nLine2\nOVERWRITE LINE\nLine4\nLine5")
	// }, t)
}

func TestPrintMoveCursorDown(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is a terminal with an infinite height and width
	When we print MoveCursorDownNRows(1) 
	Then the terminal should store "".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(ansi_escape.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorDownNRows(1)
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "R"+MoveCursorDownNRows(2)
	Then the terminal should store "R".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("R" + ansi_escape.MoveCursorDownNRows(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n"+MoveCursorUpNRows(1)+MoveCursorDwonNRows(1)+"B"
	Then the terminal should store "A\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.MoveCursorUpNRows(1) + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"B"
	Then the terminal should store "B\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(2)+"B"
	Then the terminal should store "\nA\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nA"+ CURSOR_TO_HOME + MoveCursorDownNRows(2)+"B"
	Then the terminal should store "B\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nA" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\n\n\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAB"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"CD"
	Then the terminal should store "\nCD".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAB" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "\nCD")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\nAB"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)+"CDE"
	Then the terminal should store "\nCDE".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\nAB" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "\nCDE")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+ CURSOR_TO_HOME + MoveCursorDownNRows(1)
	Then the terminal should store "\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n"+ CURSOR_TO_HOME + MoveCursorDownNRows(1) + "A"
	Then the terminal should store "\nA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "\nA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "AAA"
	Then the terminal should store "\nAAA".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "\nAAA")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(2) + "B"
	Then the terminal should store "A\n\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\n" + CURSOR_TO_HOME + MoveCursorDownNRows(5) + "B"
	Then the terminal should store "A\n\n\n\nB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n\n\n\nB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A" + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\n B".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A" + MoveCursorDownNRows(2) + "B"
	Then the terminal should store "A\n\n B".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A" + ansi_escape.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AA" + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "AA\n  B".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AA" + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n  B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AA" + MoveCursorDownNRows(3) + "B"
	Then the terminal should store "AA\n\n\n  B".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AA" + ansi_escape.MoveCursorDownNRows(3) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n\n\n  B")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "AAAAA" + MoveCursorDownNRows(2) + "BBB"
	Then the terminal should store "AAAAA\n\n     BBB".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("AAAAA" + ansi_escape.MoveCursorDownNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "AAAAA\n\n     BBB")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "A\nA\n" + CURSOR_TO_HOME + MoveCursorDownNRows(1) + "B"
	Then the terminal should store "A\nB\n".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print("A\nA\n" + ansi_escape.CURSOR_TO_HOME + ansi_escape.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB\n")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(1) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA C".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(1) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(2) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA  C".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(2) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA  C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(3) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA  C".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(3) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA   C")
	}, t)

	Test(`
	Given that there is a terminal with an infinite height and width
	When we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(6) MoveCursorDownNRows(1) + "C"
	Then the terminal should store "BBBBB\nAA C".`, func(t *testing.T) {
		fakeTerminal := terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
		fakeTerminal.Print(
			"BBBBB\nAA" +
				ansi_escape.MoveCursorUpNRows(1) +
				ansi_escape.MoveCursorRightNCols(6) +
				ansi_escape.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA      C")
	}, t)
}
