package internal_test

import (
	"testing"

	"github.com/redjolr/goherent/console/internal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewFakeAnsiTerminal(t *testing.T) {
	assert := assert.New(t)
	Test("it should return an instance of type FakeAnsiTerminal", func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		assert.IsType(internal.FakeAnsiTerminal{}, fakeTerminal)
	}, t)
}

func TestPrintBasic(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "Hello", if we print "Hello".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test(`it should store the string "Hello World",
		if we print "Hello " and then "World".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello ")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello World")
	}, t)

	Test(`it should store the string "Hello\nWorld",
		if we print "Hello" and then "\n" and then "World".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print("\n")
		fakeTerminal.Print("World")
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)
}

func TestPrintWithCursorToHomePosEscapeCode(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "Jello",
		if we print "Hello", and then CursorToHomePosEscapeCode + "J".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Jello",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then"J".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("J")
		assert.Equal(fakeTerminal.Text(), "Jello")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Cond".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Cond".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Cond")
		assert.Equal(fakeTerminal.Text(), "Condo")
	}, t)

	Test(`it should store the string "Condo",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Candy".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Candy",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Candy".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Candy")
		assert.Equal(fakeTerminal.Text(), "Candy")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then CursorToHomePosEscapeCode + "Granny".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Granny",
		if we print "Hello", and then CursorToHomePosEscapeCode, and then "Granny".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode)
		fakeTerminal.Print("Granny")
		assert.Equal(fakeTerminal.Text(), "Granny")
	}, t)

	Test(`it should store the string "Jello\nWorld",
		if we print "Hello\nWor;d", and then CursorToHomePosEscapeCode+"J".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "J")
		assert.Equal(fakeTerminal.Text(), "Jello\nWorld")
	}, t)

	Test(`it should store the string "Candy\nWorld",
		if we print "Hello\nWor;d", and then CursorToHomePosEscapeCode+"J".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello\nWorld")
		fakeTerminal.Print(internal.CursorToHomePosEscapeCode + "Candy")
		assert.Equal(fakeTerminal.Text(), "Candy\nWorld")
	}, t)
}

func TestPrintMoveCursorLeft(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNCols(1)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("H" + internal.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "H",
		if we print "H"+MoveCursorLeftNCols(2)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("H" + internal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test(`it should store the string "AR",
		if we print "RR"+MoveCursorLeftNCols(n>>1)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RR" + internal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")

		fakeTerminal = internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RR" + internal.MoveCursorLeftNCols(10000) + "A")
		assert.Equal(fakeTerminal.Text(), "AR")
	}, t)

	Test(`it should store the string "RA",
		if we print "RR"+MoveCursorLeftNCols(1)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RR" + internal.MoveCursorLeftNCols(1) + "A")
		assert.Equal(fakeTerminal.Text(), "RA")
	}, t)

	Test(`it should store the string "RAR",
		if we print "RRR"+MoveCursorLeftNCols(2)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RRR" + internal.MoveCursorLeftNCols(2) + "A")
		assert.Equal(fakeTerminal.Text(), "RAR")
	}, t)

	Test(`it should store the string "RAA",
		if we print "RRR"+MoveCursorLeftNCols(2)+"AA".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RRR" + internal.MoveCursorLeftNCols(2) + "AA")
		assert.Equal(fakeTerminal.Text(), "RAA")
	}, t)

	Test(`it should store the string "RAAA",
		if we print "RRR"+MoveCursorLeftNCols(2)+"AAA".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RRR" + internal.MoveCursorLeftNCols(2) + "AAA")
		assert.Equal(fakeTerminal.Text(), "RAAA")
	}, t)

	Test(`it should store the string "RRA",
		if we print "RR"+MoveCursorLeftNCols(0)+"A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("RR" + internal.MoveCursorLeftNCols(0) + "A")
		assert.Equal(fakeTerminal.Text(), "RRA")
	}, t)

	Test(`it should store the string "Hella",
		if we print "Hello"+MoveCursorLeftNCols(1)+"a".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello" + internal.MoveCursorLeftNCols(1) + "a")
		assert.Equal(fakeTerminal.Text(), "Hella")
	}, t)

	Test(`it should store the string "Juice",
		if we print "Hello"+MoveCursorLeftNCols(5)+"Juice".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello" + internal.MoveCursorLeftNCols(5) + "Juice")
		assert.Equal(fakeTerminal.Text(), "Juice")
	}, t)
}

func TestPrintMoveCursorRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty string),
		if we print MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "" empty string,
		if we print MoveCursorRightNCols(3).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorRightNCols(3))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R R",
		if we print "R"+MoveCursorRightNCols(1)+"R".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "R R")
	}, t)

	Test(`it should store the string "Hello   World",
		if we print "Hello"+MoveCursorRightNCols(3)+"World".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello" + internal.MoveCursorRightNCols(3) + "World")
		assert.Equal(fakeTerminal.Text(), "Hello   World")
	}, t)
}

func TestPrintMoveCursorLeftAndRight(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
			if we print MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(2).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "RR",
		if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1)+"R".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(1) + "R")
		assert.Equal(fakeTerminal.Text(), "RR")
	}, t)

	Test(`it should store the string "R",
			if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(4).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(4))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "Hellp" ,
			if we print "Hello"+ MoveCursorLeftNCols(2) + MoveCursorRightNCols(1) + "p".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello" + internal.MoveCursorLeftNCols(2) + internal.MoveCursorRightNCols(1) + "p")
		assert.Equal(fakeTerminal.Text(), "Hellp")
	}, t)

	Test(`it should store the string "Helix shaped" ,
	
			if we print "Hello"+ MoveCursorLeftNCols(2) + "ix" + MoveCursorRightNCols(1) + "shaped".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("Hello" + internal.MoveCursorLeftNCols(2) + "ix" + internal.MoveCursorRightNCols(1) + "shaped")
		assert.Equal(fakeTerminal.Text(), "Helix shaped")
	}, t)
}

func TestPrintMoveCursorUp(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorUpNRows(10).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorUpNRows(10))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\nB"+MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\nB" + internal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "B\n",
		if we print "A\n"+MoveCursorUpNRows(10)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.MoveCursorUpNRows(10) + "B")
		assert.Equal(fakeTerminal.Text(), "B\n")
	}, t)

	Test(`it should store the string "CD\n",
		if we print "AB\n"+MoveCursorUpNRows(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("AB\n" + internal.MoveCursorUpNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "CD\n")
	}, t)

	Test(`it should store the string "CDE\n",
		if we print "AB\n"+MoveCursorUpNRows(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("AB\n" + internal.MoveCursorUpNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "CDE\n")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+MoveCursorUpNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.MoveCursorUpNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "A\n",
		if we print "\n"+MoveCursorUpNRows(1) + "A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.MoveCursorUpNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "A\n")
	}, t)

	Test(`it should store the string "AAA\n",
		if we print "\n" + MoveCursorUpNRows + "AAA"`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.MoveCursorUpNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "AAA\n")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nA" + MoveCursorUpNRows(5) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.MoveCursorUpNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), " B\nA")
	}, t)

	Test(`it should store the string " B\nA",
		if we print "\nAA" + MoveCursorUpNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAA" + internal.MoveCursorUpNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "  B\nAA")
	}, t)

	Test(`it should store the string "     BBB\nAAAAA",
		if we print "\nAAAAA" + MoveCursorUpNRows(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAAAA" + internal.MoveCursorUpNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "     BBB\nAAAAA")
	}, t)

	Test(`it should store the string "\nAB\nA",
		if we print "\nA\nA" + MoveCursorUpNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA\nA" + internal.MoveCursorUpNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nAB\nA")
	}, t)

	Test(`it should store the string "\nAA   C\nBBBBB",
		if we print "\nAA\nBBBBB" + MoveCursorUpNRows(1) + "C".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAA\nBBBBB" + internal.MoveCursorUpNRows(1) + "C")
		assert.Equal(fakeTerminal.Text(), "\nAA   C\nBBBBB")
	}, t)

	Test(`it should store the string "\nA D\nB\nCC",
		if we print "\nA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nA D\nB\nCC")
	}, t)

	Test(`it should store the string "\n  D\nB\nCC",
		if we print "\n\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\n  D\nB\nCC")
	}, t)

	Test(`it should store the string "\nAAD\nB\nCC",
		if we print "\nAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAA\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAA\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAAD\nB\nCC")
	}, t)

	Test(`it should store the string "",
		if we print "\nAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAAA\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(2) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + internal.MoveCursorUpNRows(2) + "D")
		assert.Equal(fakeTerminal.Text(), "\nAADAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(3) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + internal.MoveCursorUpNRows(3) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)

	Test(`it should store the string "\nAADAA\nB\nCC",
		if we print "\nAAAAA\nB\nCC" + MoveCursorUpNRows(4) + "D".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAAAAA\nB\nCC" + internal.MoveCursorUpNRows(4) + "D")
		assert.Equal(fakeTerminal.Text(), "  D\nAAAAA\nB\nCC")
	}, t)
}

func TestPrintMoveCursorDown(t *testing.T) {
	assert := assert.New(t)
	Test(`it should store the string "" (empty space),
		if we print MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+MoveCursorDownNRows(2).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorDownNRows(2))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n"+MoveCursorUpNRows(1)+MoveCursorDwonNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.MoveCursorUpNRows(1) + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "\nB")
	}, t)

	Test(`it should store the string "\nA\nB",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\nB")
	}, t)

	Test(`it should store the string "B\n",
		if we print "\nA"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(2)+"B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nA" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "\nA\n\n\n\nB")
	}, t)

	Test(`it should store the string "\nCD",
		if we print "\nAB"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"CD".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAB" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "CD")
		assert.Equal(fakeTerminal.Text(), "\nCD")
	}, t)

	Test(`it should store the string "\nCDE",
		if we print "\nAB"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1)+"CDE".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\nAB" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "CDE")
		assert.Equal(fakeTerminal.Text(), "\nCDE")
	}, t)

	Test(`it should store the string "\n",
		if we print "\n"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1))
		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`it should store the string "\nA",
		if we print "\n"+ CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "A".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "A")
		assert.Equal(fakeTerminal.Text(), "\nA")
	}, t)

	Test(`it should store the string "\nAAA",
		if we print "\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "AAA"`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "AAA")
		assert.Equal(fakeTerminal.Text(), "\nAAA")
	}, t)

	Test(`it should store the string "A\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB")
	}, t)

	Test(`it should store the string "A\n\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\nB")
	}, t)

	Test(`it should store the string "A\n\n\n\nB",
		if we print "A\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(5) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(5) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n\n\n\nB")
	}, t)

	Test(`it should store the string "A\n B",
		if we print "A" + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A" + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n B")
	}, t)

	Test(`it should store the string "A\n\n B",
		if we print "A" + MoveCursorDownNRows(2) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A" + internal.MoveCursorDownNRows(2) + "B")
		assert.Equal(fakeTerminal.Text(), "A\n\n B")
	}, t)

	Test(`it should store the string "AA\n  B",
		if we print "AA" + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("AA" + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n  B")
	}, t)

	Test(`it should store the string "AA\n\n\n  B",
		if we print "AA" + MoveCursorDownNRows(3) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("AA" + internal.MoveCursorDownNRows(3) + "B")
		assert.Equal(fakeTerminal.Text(), "AA\n\n\n  B")
	}, t)

	Test(`it should store the string "AAAAA\n\n     BBB",
		if we print "AAAAA" + MoveCursorDownNRows(2) + "BBB".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("AAAAA" + internal.MoveCursorDownNRows(2) + "BBB")
		assert.Equal(fakeTerminal.Text(), "AAAAA\n\n     BBB")
	}, t)

	Test(`it should store the string "A\nB\n",
		if we print "A\nA\n" + CursorToHomePosEscapeCode + MoveCursorDownNRows(1) + "B".`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("A\nA\n" + internal.CursorToHomePosEscapeCode + internal.MoveCursorDownNRows(1) + "B")
		assert.Equal(fakeTerminal.Text(), "A\nB\n")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(1) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				internal.MoveCursorUpNRows(1) +
				internal.MoveCursorRightNCols(1) +
				internal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(2) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				internal.MoveCursorUpNRows(1) +
				internal.MoveCursorRightNCols(2) +
				internal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA  C")
	}, t)

	Test(`it should store the string "BBBBB\nAA  C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(3) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				internal.MoveCursorUpNRows(1) +
				internal.MoveCursorRightNCols(3) +
				internal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA   C")
	}, t)

	Test(`it should store the string "BBBBB\nAA C",
		if we print "BBBBB\nAA" + MoveCursorUpNRows(1) + MoveCursorRightNCols(6) MoveCursorDownNRows(1) + "C".
	`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(
			"BBBBB\nAA" +
				internal.MoveCursorUpNRows(1) +
				internal.MoveCursorRightNCols(6) +
				internal.MoveCursorDownNRows(1) +
				"C",
		)
		assert.Equal(fakeTerminal.Text(), "BBBBB\nAA      C")
	}, t)
}
