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
	Test(`it should store the string " " (empty space),
		if we print MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), " ")
	}, t)

	Test(`it should store the string "   " (3 empty spaces),
		if we print MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorRightNCols(3))
		assert.Equal(fakeTerminal.Text(), "   ")
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
	Test(`it should store the string " " (empty space),
		if we print MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print(internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), " ")
	}, t)

	Test(`it should store the string "R",
		if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(1).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(1))
		assert.Equal(fakeTerminal.Text(), "R")
	}, t)

	Test(`it should store the string "R ",
		if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(2).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(2))
		assert.Equal(fakeTerminal.Text(), "R ")
	}, t)

	Test(`it should store the string "R   " (R + 3 empty spaces),
		if we print "R"+ MoveCursorLeftNCols(1) + MoveCursorRightNCols(4).`, func(t *testing.T) {
		fakeTerminal := internal.NewFakeAnsiTerminal()
		fakeTerminal.Print("R" + internal.MoveCursorLeftNCols(1) + internal.MoveCursorRightNCols(4))
		assert.Equal(fakeTerminal.Text(), "R   ")
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
