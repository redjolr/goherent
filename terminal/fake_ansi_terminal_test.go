package terminal_test

import (
	"testing"

	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal"
	"github.com/stretchr/testify/assert"
)

func setup() terminal.FakeAnsiTerminal {
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal()
	return fakeAnsiTerminal
}

func TestNewFakeAnsiTerminal(t *testing.T) {
	assert := assert.New(t)
	Test("it should return an instance of type FakeAnsiTerminal", func(t *testing.T) {
		fakeTerminal := setup()
		assert.IsType(terminal.FakeAnsiTerminal{}, fakeTerminal)
	}, t)
}

func TestPrintBasic(t *testing.T) {
	assert := assert.New(t)

	Test(`It should store the string "A ", if we print "A "	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A ")
		assert.Equal(fakeTerminal.Text(), "A ")
	}, t)

	Test(`It should store the string "🚀", if we print "🚀"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("🚀")
		assert.Equal(fakeTerminal.Text(), "🚀")
	}, t)

	Test(`It should store the string "\n🚀", if we print "\n🚀"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n🚀")
		assert.Equal(fakeTerminal.Text(), "\n🚀")
	}, t)

	Test(`It should store the string "🚀A", if we print "🚀A"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("🚀A")
		assert.Equal(fakeTerminal.Text(), "🚀A")
	}, t)

	Test(`It should store the string "🚀🚀A", if we print "🚀🚀A"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("🚀🚀A")
		assert.Equal(fakeTerminal.Text(), "🚀🚀A")
	}, t)

	Test(`It should store the string "A🚀🚀A", if we print "A🚀🚀A"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A🚀🚀A")
		assert.Equal(fakeTerminal.Text(), "A🚀🚀A")
	}, t)

	Test(`It should store the string "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n", 
		if we print "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n"	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
		assert.Equal(fakeTerminal.Text(), "\n🚀 Starting... 2006-01-02 15:04:05.000\n\n")
	}, t)

	Test(`it should store the string "Hello", if we print "Hello".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test(`it should store the string "Hello World",
		if we print "Hello " and then "World".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello ")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello World")
	}, t)

	Test(`it should store the string "Hello\nWorld",
		if we print "Hello\nWorld".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello\nWorld")
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`it should store the string "Hello\nWorld",
		if we print "Hello" and then "\n" and then "World".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print("\n")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`it should print "AAA\nBAA"
		if we print "AAA\n", then "AAA", then MoveCursorUpNRowsAnsiSequence(1), then "\n", then "B"
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AAA\n")
		fakeTerminal.Print("AAA")
		fakeTerminal.Print(terminal.MoveCursorUpNRowsAnsiSequence(1))
		fakeTerminal.Print("\n")
		fakeTerminal.Print("B")
		assert.Equal(fakeTerminal.Text(), "AAA\nBAA")
	}, t)
}

func TestPrintWithANSI_CURSOR_TO_HOME(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "Jello",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME + "J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Jello",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME, and then"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME)
		fakeTerminal.Print("J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME + "Cond".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME, and then "Cond".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME)
		fakeTerminal.Print("Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME + "Candy".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Candy",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME, and then "Candy".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME)
		fakeTerminal.Print("Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME + "Granny".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then ANSI_CURSOR_TO_HOME, and then "Granny".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME)
		fakeTerminal.Print("Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Jello\nWorld",
		if we print "Hello\nWor;d", and then ANSI_CURSOR_TO_HOME+"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "J")
		assert.Equal(fakeTerminal.Text(), "Jello\nWorld")
	}, t)

	Test(`it should store the string "Candy\nWorld",
		if we print "Hello\nWor;d", and then ANSI_CURSOR_TO_HOME+"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(terminal.ANSI_CURSOR_TO_HOME + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy\nWorld")
	}, t)
}

func TestPrintMoveCursorLeft(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "✅",
		if we print "⏳"+MoveCursorLeftNColsAnsiSequence(1)+"✅".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("⏳" + terminal.MoveCursorLeftNColsAnsiSequence(1) + "✅")
		assert.Equal(fakeTerminal.Text(), "✅")
	}, t)

	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNColsAnsiSequence(1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("H" + terminal.MoveCursorLeftNColsAnsiSequence(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNColsAnsiSequence(2)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("H" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "AR",
		if we print "RR"+MoveCursorLeftNColsAnsiSequence(n>>1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")

		fakeTerminal = setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNColsAnsiSequence(10000) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")
	}, t)

	Test(`it should store the string "RA",
		if we print "RR"+MoveCursorLeftNColsAnsiSequence(1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNColsAnsiSequence(1) + "A")
		assert.Equal(fakeTerminal.Text(), "RA")
	}, t)

	Test(`it should store the string "RAR",
		if we print "RRR"+MoveCursorLeftNColsAnsiSequence(2)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "A")
		assert.Equal(fakeTerminal.Text(), "RAR")
	}, t)

	Test(`it should store the string "RAA",
		if we print "RRR"+MoveCursorLeftNColsAnsiSequence(2)+"AA".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "AA")
		assert.Equal(fakeTerminal.Text(), "RAA")
	}, t)

	Test(`it should store the string "RAAA",
		if we print "RRR"+MoveCursorLeftNColsAnsiSequence(2)+"AAA".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "AAA")
		assert.Equal(fakeTerminal.Text(), "RAAA")
	}, t)

	Test(`it should store the string "RRA",
		if we print "RR"+MoveCursorLeftNColsAnsiSequence(0)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNColsAnsiSequence(0) + "A")
		assert.Equal(fakeTerminal.Text(), "RRA")
	}, t)

	Test(`it should store the string "Hella",
		if we print "Hello"+MoveCursorLeftNColsAnsiSequence(1)+"a".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNColsAnsiSequence(1) + "a")
		assert.Equal(fakeTerminal.Text(), "Hella")
	}, t)

	Test(`it should store the string "Juice",
		if we print "Hello"+MoveCursorLeftNColsAnsiSequence(5)+"Juice".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNColsAnsiSequence(5) + "Juice")
		assert.Equal(fakeTerminal.Text(), "Juice")
	}, t)
}

func TestPrintMoveCursorRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty string),
		if we print MoveCursorRightNColsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorRightNColsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "" empty string,
		if we print MoveCursorRightNColsAnsiSequence(3).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorRightNColsAnsiSequence(3))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R R",
		if we print "R"+MoveCursorRightNColsAnsiSequence(1)+"R".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorRightNColsAnsiSequence(1) + "R")
		assert.Equal(fakeTerminal.Text(), "R R")
	}, t)

	Test(`it should store the string "Hello   World",
		if we print "Hello"+MoveCursorRightNColsAnsiSequence(3)+"World".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorRightNColsAnsiSequence(3) + "World")
		assert.Equal(fakeTerminal.Text(), "Hello   World")
	}, t)
}

func TestPrintMoveCursorLeftAndRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
			if we print MoveCursorLeftNColsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorLeftNColsAnsiSequence(1) + terminal.MoveCursorRightNColsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNColsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNColsAnsiSequence(1) + terminal.MoveCursorRightNColsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNColsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(2).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNColsAnsiSequence(1) + terminal.MoveCursorRightNColsAnsiSequence(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "RR",
		if we print "R"+ MoveCursorLeftNColsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(1)+"R".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNColsAnsiSequence(1) + terminal.MoveCursorRightNColsAnsiSequence(1) + "R")
		assert.Equal(fakeTerminal.Text(), "RR")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNColsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(4).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNColsAnsiSequence(1) + terminal.MoveCursorRightNColsAnsiSequence(4))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "Hellp" ,
			if we print "Hello"+ MoveCursorLeftNColsAnsiSequence(2) + MoveCursorRightNColsAnsiSequence(1) + "p".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNColsAnsiSequence(2) + terminal.MoveCursorRightNColsAnsiSequence(1) + "p")
		assert.Equal(fakeTerminal.Text(), "Hellp")
	}, t)

	Test(`it should store the string "Helix shaped" ,
		if we print "Hello"+ MoveCursorLeftNColsAnsiSequence(2) + "ix" + MoveCursorRightNColsAnsiSequence(1) + "shaped".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNColsAnsiSequence(2) + "ix" + terminal.MoveCursorRightNColsAnsiSequence(1) + "shaped")
		assert.Equal(fakeTerminal.Text(), "Helix shaped")
	}, t)
}

func TestPrintMoveCursorUp(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorUpNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorUpNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorUpNRowsAnsiSequence(10).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorUpNRowsAnsiSequence(10))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\nB"+MoveCursorUpNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\nB" + terminal.MoveCursorUpNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRowsAnsiSequence(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRowsAnsiSequence(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRowsAnsiSequence(10)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRowsAnsiSequence(10) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "CD\n",
		if we print "AB\n"+MoveCursorUpNRowsAnsiSequence(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AB\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "CD\n")
	}, t)

	Test(`it should store the string "CDE\n",
		if we print "AB\n"+MoveCursorUpNRowsAnsiSequence(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AB\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "CDE\n")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+MoveCursorUpNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "A\n",
		if we print "\n"+MoveCursorUpNRowsAnsiSequence(1) + "A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A\n")
	}, t)

	Test(`it should store the string "AAA\n",
		if we print "\n" + MoveCursorUpNRowsAnsiSequence + "AAA"`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "AAA\n")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRowsAnsiSequence(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRowsAnsiSequence(5) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRowsAnsiSequence(5) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nAA" + MoveCursorUpNRowsAnsiSequence(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), "  B\nAA")
	}, t)

	Test(`it should store the string "     BBB\nAAAAA",
		if we print "\nAAAAA" + MoveCursorUpNRowsAnsiSequence(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "     BBB\nAAAAA")
	}, t)

	Test(`it should store the string "\nAB\nA",
		if we print "\nA\nA" + MoveCursorUpNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA\nA" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nAB\nA")
	}, t)

	Test(`it should store the string "\nAA   C\nBBBBB",
		if we print "\nAA\nBBBBB" + MoveCursorUpNRowsAnsiSequence(1) + "C".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA\nBBBBB" + terminal.MoveCursorUpNRowsAnsiSequence(1) + "C")
		assert.Equal(fakeTerminal.Text(), "\nAA   C\nBBBBB")
	}, t)

	Test(`it should store the string "\nA D\nB\nCC",
		if we print "\nA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nA D\nB\nCC")
	}, t)

	Test(`it should store the string "\n  D\nB\nCC",
		if we print "\n\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\n  D\nB\nCC")
	}, t)

	Test(`it should store the string "\nAAD\nB\nCC",
		if we print "\nAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(3) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(3) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRowsAnsiSequence(4) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRowsAnsiSequence(4) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)
}

func TestPrintMoveCursorDown(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorDownNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorDownNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorDownNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRowsAnsiSequence(2).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorDownNRowsAnsiSequence(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n"+MoveCursorUpNRowsAnsiSequence(1)+MoveCursorDwonNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRowsAnsiSequence(1) + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nB")
	}, t)

	Test(`it should store the string "\nA\nB",
		if we print "\nA"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(5) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\n\n\n\nB")
	}, t)

	Test(`it should store the string "\nCD",
		if we print "\nAB"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAB" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "\nCD")
	}, t)

	Test(`it should store the string "\nCDE",
		if we print "\nAB"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAB" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "\nCDE")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "\nA",
		if we print "\n"+ ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1) + "A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "A")
		assert.Equal(fakeTerminal.Text(), "\nA")
	}, t)

	Test(`it should store the string "\nAAA",
		if we print "\n" + ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1) + "AAA"`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "\nAAA")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n" + ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "A\n\nB",
		if we print "A\n" + ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\nB")
	}, t)

	Test(`it should store the string "A\n\n\n\nB",
		if we print "A\n" + ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(5) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(5) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n\n\n\nB")
	}, t)

	Test(`it should store the string "A\n B",
		if we print "A" + MoveCursorDownNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A" + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n B")
	}, t)

	Test(`it should store the string "A\n\n B",
		if we print "A" + MoveCursorDownNRowsAnsiSequence(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A" + terminal.MoveCursorDownNRowsAnsiSequence(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n B")
	}, t)

	Test(`it should store the string "AA\n  B",
		if we print "AA" + MoveCursorDownNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AA" + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n  B")
	}, t)

	Test(`it should store the string "AA\n\n\n  B",
		if we print "AA" + MoveCursorDownNRowsAnsiSequence(3) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AA" + terminal.MoveCursorDownNRowsAnsiSequence(3) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n\n\n  B")
	}, t)

	Test(`it should store the string "AAAAA\n\n     BBB",
		if we print "AAAAA" + MoveCursorDownNRowsAnsiSequence(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AAAAA" + terminal.MoveCursorDownNRowsAnsiSequence(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "AAAAA\n\n     BBB")
	}, t)

	Test(`it should store the string "A\nB\n",
		if we print "A\nA\n" + ANSI_CURSOR_TO_HOME + MoveCursorDownNRowsAnsiSequence(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\nA\n" + terminal.ANSI_CURSOR_TO_HOME + terminal.MoveCursorDownNRowsAnsiSequence(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB\n")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRowsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(1) MoveCursorDownNRowsAnsiSequence(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRowsAnsiSequence(1) +
				terminal.MoveCursorRightNColsAnsiSequence(1) +
				terminal.MoveCursorDownNRowsAnsiSequence(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRowsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(2) MoveCursorDownNRowsAnsiSequence(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRowsAnsiSequence(1) +
				terminal.MoveCursorRightNColsAnsiSequence(2) +
				terminal.MoveCursorDownNRowsAnsiSequence(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA  C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRowsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(3) MoveCursorDownNRowsAnsiSequence(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRowsAnsiSequence(1) +
				terminal.MoveCursorRightNColsAnsiSequence(3) +
				terminal.MoveCursorDownNRowsAnsiSequence(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA   C")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRowsAnsiSequence(1) + MoveCursorRightNColsAnsiSequence(6) MoveCursorDownNRowsAnsiSequence(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRowsAnsiSequence(1) +
				terminal.MoveCursorRightNColsAnsiSequence(6) +
				terminal.MoveCursorDownNRowsAnsiSequence(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA      C")
	}, t)
}
