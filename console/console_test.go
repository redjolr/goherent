package console_test

import (
	"testing"

	"github.com/redjolr/goherent/console"
	"github.com/redjolr/goherent/console/terminal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func setup() (console.Console, *terminal.FakeAnsiTerminal) {
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal()
	return console.NewConsole(&fakeAnsiTerminal), &fakeAnsiTerminal
}

func TestRenderingTextBlock(t *testing.T) {
	assert := assert.New(t)
	Test("The terminal should print a single letter.", func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("A")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test("The terminal should print a word.", func(t *testing.T) {
		console, fakeTerminal := setup()

		console.NewTextBlock("Hello")
		console.Render()

		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test("The terminal should print two lines.", func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("Hello\nThere")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nThere")
	}, t)
}

func TestTextBlockWrite(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print Hello\nWorld, if we write "Hello\nWorld".`, func(t *testing.T) {
		console, fakeTerminal := setup()
		tb := console.NewTextBlock("A")
		tb.Write("Hello\nWorld")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)
}

func TestTwoTextblocks(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print HelloWorld, if we create two textblocks "Hello" and "World"`, func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("Hello")
		console.NewTextBlock("World")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "HelloWorld")
	}, t)

	Test(`The terminal should print "Hellp World",
		if we create two textblocks "Hello " and "World", render them and then modify the first with "Hellp "
		and render them again`, func(t *testing.T) {
		console, fakeTerminal := setup()
		tb1 := console.NewTextBlock("Hello ")
		console.NewTextBlock("World")
		console.Render()
		tb1.Write("Hellp ")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hellp World")
	}, t)

	// Failing
	Test(`The terminal should print "Help World",
		if we create two textblocks "Hello " and "World", render them and then modify the first with "Help "
		and render them again`, func(t *testing.T) {
		console, fakeTerminal := setup()
		tb1 := console.NewTextBlock("Hello ")
		console.NewTextBlock("World")
		console.Render()
		tb1.Write("Help ")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Help World")
	}, t)
}
