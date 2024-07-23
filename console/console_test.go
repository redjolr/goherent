package console_test

import (
	"testing"

	"github.com/redjolr/goherent/console"
	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/cursor"
	"github.com/redjolr/goherent/console/terminal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func setup() (console.Console, *terminal.FakeAnsiTerminal, *cursor.Cursor) {
	terminalOrigin := coordinates.Origin()
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal(&terminalOrigin)
	cursor := cursor.NewCursor(&fakeAnsiTerminal, &terminalOrigin)
	return console.NewConsole(&fakeAnsiTerminal, &cursor), &fakeAnsiTerminal, &cursor
}

func TestIsConsoleRendered(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the console has no elements", func(t *testing.T) {
		console, _, _ := setup()
		assert.True(console.IsRendered())
	}, t)

	Test("it should return false, if the console has a Textblock element and it is not rendered.", func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		assert.False(console.IsRendered())
	}, t)

	Test("it should return false, if the console has an UnorderedList element and it is not rendered.", func(t *testing.T) {
		console, _, _ := setup()
		console.NewUnorderedList("id1", "Unordered list name")
		assert.False(console.IsRendered())
	}, t)

	Test("it should return true, if the console has a Textblock element and it is rendered.", func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.Render()
		assert.True(console.IsRendered())
	}, t)

	Test(`it should return true, if the console has a Textblock element and an UnorderedList
			and the console is rendered.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.NewUnorderedList("list1", "List name")
		console.Render()
		assert.True(console.IsRendered())
	}, t)

	Test(`it should return false, if the console has a Textblock element, it is rendered,
			then we add an UnorderedList element.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("id1", "Hello There")
		console.Render()
		console.NewUnorderedList("list1", "List name")
		assert.False(console.IsRendered())
	}, t)
}

func TestRenderingUnorderedList(t *testing.T) {
	assert := assert.New(t)
	Test(`The terminal should print "Some unordered list",
			if we create an UnorderedList with that name and render it.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewUnorderedList("list1", "Some unordered list")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Some unordered list")
	}, t)

	Test(`The terminal should print "Some unordered list",
			if we perform these actions in the given sequence:
			- create an unordered list named "list"
			- render the console
			- edit the list name to "Some unordered list"
			- render the console again
			if we create an UnorderedList with name "list", edit.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		list := console.NewUnorderedList("list1", "list")
		console.Render()
		list.EditName("Some unordered list")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Some unordered list")
	}, t)

	Test(`The terminal should print "Some unordered list",
			if we perform these actions in the given sequence:
			- create an unordered list named "list"
			- edit the list name to "Some unordered list"
			- render the console
			if we create an UnorderedList with name "list", edit.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		list := console.NewUnorderedList("list1", "list")
		list.EditName("Some unordered list")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Some unordered list")
	}, t)

	Test(`The terminal should print "Undordered\nList",
			if we create an UnorderedList with that name and render it.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewUnorderedList("list1", "Unordered\nList")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered\nList")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1",
			if we create an UnorderedList with that name, add an item to the list and render it.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2",
			if we create an UnorderedList with that name, add two items to the list and render it.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 0",
			if we perform these actions in the given sequence:
			- create an UnorderedList with that name
			- add one item with name "List item 1"
			- edit the list item and change its name to "List item 0"
			- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		listItem := unorderedList.NewItem("id1", "List item 1")
		listItem.Edit("List item 0")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 0")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 0",
			if we perform these actions in the given sequence:
			- create an UnorderedList with that name
			- add one item with name "List item 1"
			- render the console
			- edit the list item and change its name to "List item 0"
			- render the console again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		listItem := unorderedList.NewItem("id1", "List item 1")
		console.Render()
		listItem.Edit("List item 0")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 0")
	}, t)

	Test(`The terminal should print "Unordered List\n\tFirst list item\n\tList item 2\n\tList item 3",
			if we perform these actions in the given sequence:
			- create an UnorderedList with the name "Unordered List"
			- add three items with names "List item 1", "List item 2", "List item 3"
			- render the console
			- edit the firs list item and change its name to "First list item"
			- render the console again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		firstListItem := unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		console.Render()
		firstListItem.Edit("First list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tFirst list item\n\tList item 2\n\tList item 3")
	}, t)

	Test(`The terminal should print "Unordered List\n\tFirst list item\n\tList item 2\n\tList item 3",
			if we perform these actions in the given sequence:
			- create an UnorderedList with the name "Unordered List"
			- add three items with names "List item 1", "List item 2", "List item 3"
			- edit the firs list item and change its name to "First list item"
			- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		firstListItem := unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		firstListItem.Edit("First list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tFirst list item\n\tList item 2\n\tList item 3")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tSecond list item\n\tList item 3",
			if we perform these actions in the given sequence:
			- create an UnorderedList with the name "Unordered List"
			- add three items with names "List item 1", "List item 2", "List item 3"
			- render the console
			- edit the second list item and change its name to "Second list item"
			- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		secondListItem := unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		console.Render()
		secondListItem.Edit("Second list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tSecond list item\n\tList item 3")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tSecond list item\n\tList item 3",
			if we perform these actions in the given sequence:
			- create an UnorderedList with the name "Unordered List"
			- add three items with names "List item 1", "List item 2", "List item 3"
			- edit the second list item and change its name to "Second list item"
			- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		secondListItem := unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		secondListItem.Edit("Second list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tSecond list item\n\tList item 3")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2\n\tThird list item",
		if we perform these actions in the given sequence:
		- create an UnorderedList with the name "Unordered List"
		- add three items with names "List item 1", "List item 2", "List item 3"
		- render the console
		- edit the third list item and change its name to "Third list item"
		- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		thirdListItem := unorderedList.NewItem("id3", "List item 3")
		console.Render()
		thirdListItem.Edit("Third list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2\n\tThird list item")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2\n\tThird list item",
		if we perform these actions in the given sequence:
		- create an UnorderedList with the name "Unordered List"
		- add three items with names "List item 1", "List item 2", "List item 3"
		- edit the third list item and change its name to "Second list item"
		- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "Unordered List")
		unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		thirdListItem := unorderedList.NewItem("id3", "List item 3")
		thirdListItem.Edit("Third list item")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2\n\tThird list item")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2\n\tList item 3",
		if we perform these actions in the given sequence:
		- create an UnorderedList with the name "list"
		- add three items with names "List item 1", "List item 2", "List item 3"
		- render the console
		- edit the list name to "Unordered List"
		- render the console again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "list")
		unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		console.Render()
		unorderedList.EditName("Unordered List")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2\n\tList item 3")
	}, t)

	Test(`The terminal should print "Unordered List\n\tList item 1\n\tList item 2\n\tList item 3",
		if we perform these actions in the given sequence:
		- create an UnorderedList with the name "list"
		- add three items with names "List item 1", "List item 2", "List item 3"
		- edit the list name to "Unordered List"
		- render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		unorderedList := console.NewUnorderedList("list1", "list")
		unorderedList.NewItem("id1", "List item 1")
		unorderedList.NewItem("id2", "List item 2")
		unorderedList.NewItem("id3", "List item 3")
		unorderedList.EditName("Unordered List")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Unordered List\n\tList item 1\n\tList item 2\n\tList item 3")
	}, t)
}

func TestTextBlockRender(t *testing.T) {
	assert := assert.New(t)

	Test("The terminal should print a single letter.", func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "A")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A")
	}, t)

	Test("The terminal should print a word.", func(t *testing.T) {
		console, fakeTerminal, _ := setup()

		console.NewTextBlock("id1", "Hello")
		console.Render()

		assert.Equal(fakeTerminal.Text(), "Hello")
	}, t)

	Test(`The terminal should print Hello\nWorld, if we create a Textblock "Hello\nWorld" and render the console.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "Hello\nWorld")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`The terminal should print Hello\nWorld,
		if we create a Textblock "A", and then edit it with "Hello\nWorld" and render the console.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "A")
		tb.Edit("Hello\nWorld")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`The terminal should print Hello\nWorld,
		if we create a Textblock "A" render the console,
		and then edit it with "Hello\nWorld" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "A")
		console.Render()
		tb.Edit("Hello\nWorld")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`The terminal should print "A ",
		if we create a Textblock "BC" render the console,
		and then edit it with "A" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "BC")
		console.Render()
		tb.Edit("A")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A ")
	}, t)

	Test(`The terminal should print " ",
		if we create a Textblock "A" render the console,
		and then edit it with "" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "A")
		console.Render()
		tb.Edit("")
		console.Render()
		assert.Equal(fakeTerminal.Text(), " ")
	}, t)

	Test(`The terminal should print "Hellp\n     ",
		if we create a Textblock "Hello\nWorld" and render the console,
		and then edit it with "Hellp" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "Hello\nWorld")
		console.Render()
		tb.Edit("Hellp")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hellp\n     ")
	}, t)

	Test(`The terminal should print "Hello\n     ",
		if we create a Textblock "Hello\nWorld" and render the console,
		and then edit it with "Hello" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "Hello\nWorld")
		console.Render()
		tb.Edit("Hello")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\n     ")
	}, t)

	Test(`The terminal should print "A \nB ",
		if we create a Textblock "AA\nBB" and render the console,
		and then edit it with "A\nB" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "AA\nBB")
		console.Render()
		tb.Edit("A\nB")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A \nB ")
	}, t)

	Test(`The terminal should print "AA\nBB",
		if we create a Textblock "A\nB" and render the console,
		and then edit it with "AA\nBB" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "A\nB")
		console.Render()
		tb.Edit("AA\nBB")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "AA\nBB")
	}, t)

	Test(`The terminal should print "AA  \nBB \n  ",
		if we create a Textblock "AAAA\nBBB\nCC" and render the console,
		and then edit it with "AA\nBB" and render the console again.`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb := console.NewTextBlock("id1", "AAAA\nBBB\nCC")
		console.Render()
		tb.Edit("AA\nBB")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "AA  \nBB \n  ")
	}, t)

}

func TestMultipleTextblocksRender(t *testing.T) {
	assert := assert.New(t)

	Test(`The terminal should print "\n", if we create two textblocks "" and "" and render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "")
		console.NewTextBlock("id2", "")
		console.Render()

		assert.Equal(fakeTerminal.Text(), "\n")
	}, t)

	Test(`The terminal should print "\n\n", if we create three "" textblocks and render the console`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "")
		console.NewTextBlock("id2", "")
		console.NewTextBlock("id3", "")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "\n\n")
	}, t)

	Test(`The terminal should print Hello\nWorld, if we create two textblocks "Hello" and "World"`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "Hello")
		console.NewTextBlock("id2", "World")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hello\nWorld")
	}, t)

	Test(`The terminal should print "Hellp\nWorld",
		if we create two textblocks "Hello" and "World", render them and then modify the first with "Hellp"
		and render them again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb1 := console.NewTextBlock("id1", "Hello")
		console.NewTextBlock("id2", "World")
		console.Render()
		tb1.Edit("Hellp")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Hellp\nWorld")
	}, t)

	Test(`The terminal should print "Help \nWorld",
		if we create two textblocks "Hello" and "World", render them and then modify the first with "Help "
		and render them again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb1 := console.NewTextBlock("id1", "Hello")
		console.NewTextBlock("id2", "World")
		console.Render()
		tb1.Edit("Help")

		console.Render()
		assert.Equal(fakeTerminal.Text(), "Help \nWorld")
	}, t)

	Test(`The terminal should print "BBB\nBBB",
		if we create two textblocks "AAA" and "AAA", render them and then modify them both with "BBB"
		and render them again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb1 := console.NewTextBlock("id1", "AAA")
		tb2 := console.NewTextBlock("id2", "AAA")
		console.Render()
		tb1.Edit("BBB")
		tb2.Edit("BBB")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "BBB\nBBB")
	}, t)

	Test(`The terminal should print "AAA\nBBB",
		if we create one textblock "AAA", render the console
		and then create a new textblock "BBB"
		and render the console again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		console.NewTextBlock("id1", "AAA")
		console.Render()
		console.NewTextBlock("id2", "BBB")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "AAA\nBBB")
	}, t)

	Test(`The terminal should print "A  \nBBBB\nC",
		if we create one textblock "AAA", render the console
		and then create two new textblock "BBBB" and "C"
		and edit the first textblock to "A"
		and render the console again`, func(t *testing.T) {
		console, fakeTerminal, _ := setup()
		tb1 := console.NewTextBlock("id1", "AAA")
		console.Render()
		console.NewTextBlock("id2", "BBBB")
		console.NewTextBlock("id3", "C")
		tb1.Edit("A")
		console.Render()
		assert.Equal(fakeTerminal.Text(), "A  \nBBBB\nC")
	}, t)
}

func TestHasElementWithId(t *testing.T) {
	assert := assert.New(t)

	Test("it should return false, if the console has no elements.", func(t *testing.T) {
		console, _, _ := setup()
		assert.False(console.HasElementWithId("someId"))
	}, t)

	Test(`it should return false,
		if the console has an unordered list with id 'list1' and we search for 'list2'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		assert.False(console.HasElementWithId("list2"))
	}, t)

	Test(`it should return false,
		if the console has a a textblock with id 'textblock1' and we search for 'textblock2'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("textblock2", "Some textblock")
		assert.False(console.HasElementWithId("textblock1"))
	}, t)

	Test(`it should return true,
		if the console has an unordered list with id 'list1' and we search for 'list1'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		assert.True(console.HasElementWithId("list1"))
	}, t)

	Test(`it should return true,
		if the console has two unordered list with id 'list1', 'list2' and we search for 'list1'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		console.NewUnorderedList("list2", "Some other list")
		assert.True(console.HasElementWithId("list1"))
	}, t)

	Test(`it should return true,
		if the console has two unordered list with id 'list1', 'list2' and we search for 'list2'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewUnorderedList("list1", "Some list")
		console.NewUnorderedList("list2", "Some other list")
		assert.True(console.HasElementWithId("list2"))
	}, t)

	Test(`it should return true,
		if the console has 1 textblockwith id 'textBlock1' and we search for 'textBlock1'.`, func(t *testing.T) {
		console, _, _ := setup()
		console.NewTextBlock("textBlock1", "Some textblock")
		assert.True(console.HasElementWithId("textBlock1"))
	}, t)
}
