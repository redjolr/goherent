package terminal_test

import (
	"testing"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/terminal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func setup() terminal.FakeAnsiTerminal {
	origin := coordinates.Origin()
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal(&origin)
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
		if we print "AAA\n", then "AAA", then MoveCursorUpNRows(1), then "\n", then "B"
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AAA\n")
		fakeTerminal.Print("AAA")
		fakeTerminal.Print(terminal.MoveCursorUpNRows(1))
		fakeTerminal.Print("\n")
		fakeTerminal.Print("B")
		assert.Equal(fakeTerminal.Text(), "AAA\nBAA")
	}, t)
}

func TestPrintWithCursorToHomePosEscapeCode(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "Jello",
		if we print "Hello", and then CursorToHomePosEscapeCode + "J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Jello",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Cond".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Cond".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Candy".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Candy",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Candy".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Granny".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Granny".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Jello\nWorld",
		if we print "Hello\nWor;d", and then CursorToHomePosEscapeCode+"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "J")
		assert.Equal(fakeTerminal.Text(), "Jello\nWorld")
	}, t)

	Test(`it should store the string "Candy\nWorld",
		if we print "Hello\nWor;d", and then CursorToHomePosEscapeCode+"J".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(terminal.CursorToHomePosEscapeCode + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy\nWorld")
	}, t)
}

func TestPrintMoveCursorLeft(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNCols(1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("H" + terminal.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNCols(2)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("H" + terminal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "AR",
		if we print "RR"+MoveCursorLeftNCols(n>>1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")

		fakeTerminal = setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNCols(10000) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")
	}, t)

	Test(`it should store the string "RA",
		if we print "RR"+MoveCursorLeftNCols(1)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "RA")
	}, t)

	Test(`it should store the string "RAR",
		if we print "RRR"+MoveCursorLeftNCols(2)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "RAR")
	}, t)

	Test(`it should store the string "RAA",
		if we print "RRR"+MoveCursorLeftNCols(2)+"AA".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNCols(2) + "AA")
		assert.Equal(fakeTerminal.Text(), "RAA")
	}, t)

	Test(`it should store the string "RAAA",
		if we print "RRR"+MoveCursorLeftNCols(2)+"AAA".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RRR" + terminal.MoveCursorLeftNCols(2) + "AAA")
		assert.Equal(fakeTerminal.Text(), "RAAA")
	}, t)

	Test(`it should store the string "RRA",
		if we print "RR"+MoveCursorLeftNCols(0)+"A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("RR" + terminal.MoveCursorLeftNCols(0) + "A")
		assert.Equal(fakeTerminal.Text(), "RRA")
	}, t)

	Test(`it should store the string "Hella",
		if we print "Hello"+MoveCursorLeftNCols(1)+"a".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNCols(1) + "a")
		assert.Equal(fakeTerminal.Text(), "Hella")
	}, t)

	Test(`it should store the string "Juice",
		if we print "Hello"+MoveCursorLeftNCols(5)+"Juice".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNCols(5) + "Juice")
		assert.Equal(fakeTerminal.Text(), "Juice")
	}, t)
}

func TestPrintMoveCursorRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty string),
		if we print MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "" empty string,
		if we print MoveCursorRightNCols(3).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorRightNCols(3))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R R",
		if we print "R"+MoveCursorRightNCols(1)+"R".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "R R")
	}, t)

	Test(`it should store the string "Hello   World",
		if we print "Hello"+MoveCursorRightNCols(3)+"World".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorRightNCols(3) + "World")
		assert.Equal(fakeTerminal.Text(), "Hello   World")
	}, t)
}

func TestPrintMoveCursorLeftAndRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
			if we print MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorLeftNCols(1) + terminal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNCols(1) + terminal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(2).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNCols(1) + terminal.MoveCursorRightNCols(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "RR",
		if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)+"R".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNCols(1) + terminal.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "RR")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(4).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorLeftNCols(1) + terminal.MoveCursorRightNCols(4))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "Hellp" ,
			if we print "Hello"+ MoveCursorLeftNCols(2) + MoveCursorRightNCols(1) + "p".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNCols(2) + terminal.MoveCursorRightNCols(1) + "p")
		assert.Equal(fakeTerminal.Text(), "Hellp")
	}, t)

	Test(`it should store the string "Helix shaped" ,
		if we print "Hello"+ MoveCursorLeftNCols(2) + "ix" + MoveCursorRightNCols(1) + "shaped".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("Hello" + terminal.MoveCursorLeftNCols(2) + "ix" + terminal.MoveCursorRightNCols(1) + "shaped")
		assert.Equal(fakeTerminal.Text(), "Helix shaped")
	}, t)
}

func TestPrintMoveCursorUp(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorUpNRows(10).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorUpNRows(10))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\nB"+MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\nB" + terminal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(10)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRows(10) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "CD\n",
		if we print "AB\n"+MoveCursorUpNRows(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AB\n" + terminal.MoveCursorUpNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "CD\n")
	}, t)

	Test(`it should store the string "CDE\n",
		if we print "AB\n"+MoveCursorUpNRows(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AB\n" + terminal.MoveCursorUpNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "CDE\n")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "A\n",
		if we print "\n"+MoveCursorUpNRows(1) + "A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A\n")
	}, t)

	Test(`it should store the string "AAA\n",
		if we print "\n" + MoveCursorUpNRows + "AAA"`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.MoveCursorUpNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "AAA\n")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(5) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.MoveCursorUpNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nAA" + MoveCursorUpNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA" + terminal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "  B\nAA")
	}, t)

	Test(`it should store the string "     BBB\nAAAAA",
		if we print "\nAAAAA" + MoveCursorUpNRows(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA" + terminal.MoveCursorUpNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "     BBB\nAAAAA")
	}, t)

	Test(`it should store the string "\nAB\nA",
		if we print "\nA\nA" + MoveCursorUpNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA\nA" + terminal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nAB\nA")
	}, t)

	Test(`it should store the string "\nAA   C\nBBBBB",
		if we print "\nAA\nBBBBB" + MoveCursorUpNRows(1) + "C".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA\nBBBBB" + terminal.MoveCursorUpNRows(1) + "C")
		assert.Equal(fakeTerminal.Text(), "\nAA   C\nBBBBB")
	}, t)

	Test(`it should store the string "\nA D\nB\nCC",
		if we print "\nA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nA D\nB\nCC")
	}, t)

	Test(`it should store the string "\n  D\nB\nCC",
		if we print "\n\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\n  D\nB\nCC")
	}, t)

	Test(`it should store the string "\nAAD\nB\nCC",
		if we print "\nAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAA\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAA\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAA\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(3) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRows(3) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(4) + "D".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + terminal.MoveCursorUpNRows(4) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)
}

func TestPrintMoveCursorDown(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(terminal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRows(2).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("R" + terminal.MoveCursorDownNRows(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n"+MoveCursorUpNRows(1)+MoveCursorDwonNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.MoveCursorUpNRows(1) + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nB")
	}, t)

	Test(`it should store the string "\nA\nB",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nA" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\n\n\n\nB")
	}, t)

	Test(`it should store the string "\nCD",
		if we print "\nAB"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAB" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "\nCD")
	}, t)

	Test(`it should store the string "\nCDE",
		if we print "\nAB"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\nAB" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "\nCDE")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "\nA",
		if we print "\n"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "A".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "\nA")
	}, t)

	Test(`it should store the string "\nAAA",
		if we print "\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "AAA"`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "\nAAA")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "A\n\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\nB")
	}, t)

	Test(`it should store the string "A\n\n\n\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(5) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n\n\n\nB")
	}, t)

	Test(`it should store the string "A\n B",
		if we print "A" + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A" + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n B")
	}, t)

	Test(`it should store the string "A\n\n B",
		if we print "A" + MoveCursorDownNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A" + terminal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n B")
	}, t)

	Test(`it should store the string "AA\n  B",
		if we print "AA" + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AA" + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n  B")
	}, t)

	Test(`it should store the string "AA\n\n\n  B",
		if we print "AA" + MoveCursorDownNRows(3) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AA" + terminal.MoveCursorDownNRows(3) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n\n\n  B")
	}, t)

	Test(`it should store the string "AAAAA\n\n     BBB",
		if we print "AAAAA" + MoveCursorDownNRows(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("AAAAA" + terminal.MoveCursorDownNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "AAAAA\n\n     BBB")
	}, t)

	Test(`it should store the string "A\nB\n",
		if we print "A\nA\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print("A\nA\n" + terminal.CursorToHomePosEscapeCode + terminal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB\n")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(1) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRows(1) +
				terminal.MoveCursorRightNCols(1) +
				terminal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(2) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRows(1) +
				terminal.MoveCursorRightNCols(2) +
				terminal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA  C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(3) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRows(1) +
				terminal.MoveCursorRightNCols(3) +
				terminal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA   C")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(6) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := setup()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				terminal.MoveCursorUpNRows(1) +
				terminal.MoveCursorRightNCols(6) +
				terminal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA      C")
	}, t)
}
