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

func TestIsConsoleRendered(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the console has no elements", func(t *testing.T) {
		console, _ := setup()
		assert.True(console.IsRendered())
	}, t)

	Test("it should return false, if the console has a Textblock element and it is not rendered.", func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		assert.False(console.IsRendered())
	}, t)

	Test("it should return false, if the console has an UnorderedList element and it is not rendered.", func(t *testing.T) {
		console, _ := setup()
		console.NewUnorderedList("id1", "Unordered list name")
		assert.False(console.IsRendered())
	}, t)

	Test("it should return true, if the console has a Textblock element and it is rendered.", func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.Render()
		assert.True(console.IsRendered())
	}, t)

	Test(`it should return true, if the console has a Textblock element and an UnorderedList
		and the console is rendered.`, func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.NewUnorderedList("list1", "List name")
		console.Render()
		assert.True(console.IsRendered())
	}, t)

	Test(`it should return false, if the console has a Textblock element, it is rendered,
		then we add an UnorderedList element.`, func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.Render()
		console.NewUnorderedList("list1", "List name")
		assert.False(console.IsRendered())
	}, t)
}

func TestRenderingTextBlock(t *testing.T) {
	assert := assert.New(t)
	Test("The terminal should print a single letter.", func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("id1", "A")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test("The terminal should print a word.", func(t *testing.T) {
		console, fakeTerminal := setup()

		console.NewTextBlock("id1", "Hello")
		console.Render()

		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test("The terminal should print two lines.", func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("id1", "Hello\nThere")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nThere")
	}, t)
}

func TestRenderingUnorderedList(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print "Some unordered list", 
		if we create an UnorderedList with that name and render it.`, func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewUnorderedList("list1", "Some unordered list")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Some unordered list")
	}, t)

	Test(`The terminal should print "Undordered\nList", 
		if we create an UnorderedList with that name and render it.`, func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewUnorderedList("list1", "Unordered\nList")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered\nList")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1", 
		if we create an UnorderedList with that name, add an item to the list and render it.`, func(t *testing.T) {
		console, fakeTerminal := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("List item 1")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2", 
		if we create an UnorderedList with that name, add two items to the list and render it.`, func(t *testing.T) {
		console, fakeTerminal := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("List item 1")
		unorderedList.NewItem("List item 2")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2")
	}, t)

	// Test(`The terminal should print "Unordered List\n\tList item 0",
	// 	if we perform these actions in the given sequence:
	// 	- create an UnorderedList with that name
	// 	- add one item with name "List item 1"
	// 	- render the console
	// 	- edit the list item and change its name to "List item 0"
	// 	- render the console again`, func(t *testing.T) {
	// 	console, fakeTerminal := setup()
	// 	unorderedList := console.NewUnorderedList("list1", "Unordered List")
	// 	listItem := unorderedList.NewItem("List item 1")
	// 	console.Render()
	// 	listItem.Edit("List item 0")
	// 	console.Render()
	// 	assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 0")
	// }, t)
}

func TestTextBlockWrite(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print Hello\nWorld, if we write "Hello\nWorld".`, func(t *testing.T) {
		console, fakeTerminal := setup()
		tb := console.NewTextBlock("id1", "A")
		tb.Write("Hello\nWorld")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)
}

func TestTwoTextblocks(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print HelloWorld, if we create two textblocks "Hello" and "World"`, func(t *testing.T) {
		console, fakeTerminal := setup()
		console.NewTextBlock("id1", "Hello")
		console.NewTextBlock("id2", "World")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "HelloWorld")
	}, t)

	Test(`The terminal should print "Hellp World",
		if we create two textblocks "Hello " and "World", render them and then modify the first with "Hellp "
		and render them again`, func(t *testing.T) {
		console, fakeTerminal := setup()
		tb1 := console.NewTextBlock("id1", "Hello ")
		console.NewTextBlock("id2", "World")
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
		tb1 := console.NewTextBlock("id1", "Hello ")
		console.NewTextBlock("id2", "World")
		console.Render()
		tb1.Write("Help ")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Help World")
	}, t)
}

func TestHasElementWithId(t *testing.T) {
	assert := assert.New(t)

	Test("it should return false, if the console has no elements.", func(t *testing.T) {
		console, _ := setup()
		assert.False(console.HasElementWithId("someId"))
	}, t)

	Test(`it should return false, 
		if the console has an unordered list with id 'list1' and we search for 'list2'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		assert.False(console.HasElementWithId("list2"))
	}, t)

	Test(`it should return false, 
		if the console has a a textblock with id 'textblock1' and we search for 'textblock2'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("textblock2", "Some textblock")
		assert.False(console.HasElementWithId("textblock1"))
	}, t)

	Test(`it should return true,
		if the console has an unordered list with id 'list1' and we search for 'list1'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		assert.True(console.HasElementWithId("list1"))
	}, t)

	Test(`it should return true,
		if the console has two unordered list with id 'list1', 'list2' and we search for 'list1'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		console.NewUnorderedList("list2", "Some other list")
		assert.True(console.HasElementWithId("list1"))
	}, t)

	Test(`it should return true,
		if the console has two unordered list with id 'list1', 'list2' and we search for 'list2'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		console.NewUnorderedList("list2", "Some other list")
		assert.True(console.HasElementWithId("list2"))
	}, t)

	Test(`it should return true,
		if the console has 1 textblockwith id 'textBlock1' and we search for 'textBlock1'.`, func(t *testing.T) {
		console, _ := setup()
		console.NewTextBlock("textBlock1", "Some textblock")
		assert.True(console.HasElementWithId("textBlock1"))
	}, t)
}
