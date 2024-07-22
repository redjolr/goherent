package elements_test

import (
	"testing"

	"github.com/redjolr/goherent/console/elements"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	assert := assert.New(t)
	Test("it should add a new item to a UnorderedList and the output should be a pointer to the item.", func(t *testing.T) {
		list := elements.NewUnorderedList("list id", "List name")
		item := list.NewItem("item1")

		assert.IsType(&elements.ListItem{}, item)
	}, t)
}

func TestFindItemByOrder(t *testing.T) {
	assert := assert.New(t)

	Test("it should return nil, if the UnorderedList is empty.", func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		item := list.FindItemByOrder(0)
		assert.Nil(item)
	}, t)

	Test(`it should return nil,
	if the UnorderedList has 1 item and you try to find the item with id 1.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("some text")
		item := list.FindItemByOrder(1)
		assert.Nil(item)
	}, t)

	Test(`it should return a ListItem,
	if the UnorderedList has 1 item and you try to find the item with id 0.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("some text")
		item := list.FindItemByOrder(0)
		assert.IsType(&elements.ListItem{}, item)
	}, t)

	Test(`Given that the ListTiem has two items: "First Item" and "Second Item"
	When you want to find the item with id 0
	Then the item "First Item" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First Item")
		list.NewItem("Second Item")

		// When
		item := list.FindItemByOrder(0)

		// Then
		assert.Equal(item.Text(), "First Item")
	}, t)

	Test(`Given that the ListTiem has two items: "First Item" and "Second Item"
	When you want to find the item with id 1
	Then the item "Second Item" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First Item")
		list.NewItem("Second Item")

		// When
		item := list.FindItemByOrder(1)

		// Then
		assert.Equal(item.Text(), "Second Item")
	}, t)

	Test(`Given that the ListTiem has two items: "First Item" and "Second Item"
	When you want to find the item with id 2
	Then nil will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First Item")
		list.NewItem("Second Item")

		// When
		item := list.FindItemByOrder(2)

		// Then
		assert.Nil(item)
	}, t)

	Test(`Given that the ListTiem has 6 items: "First", "Second", "Third", "Fourth", "Fifth" and "Sixth"
	When you want to find the item with id 0
	Then the item "First" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First")
		list.NewItem("Second")
		list.NewItem("Third")
		list.NewItem("Fourth")
		list.NewItem("Fifth")
		list.NewItem("Sixth")

		// When
		item := list.FindItemByOrder(0)

		// Then
		assert.Equal(item.Text(), "First")
	}, t)

	Test(`Given that the ListTiem has 6 items: "First", "Second", "Third", "Fourth", "Fifth" and "Sixth"
	When you want to find the item with id 2
	Then the item "Third" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First")
		list.NewItem("Second")
		list.NewItem("Third")
		list.NewItem("Fourth")
		list.NewItem("Fifth")
		list.NewItem("Sixth")

		// When
		item := list.FindItemByOrder(2)

		// Then
		assert.Equal(item.Text(), "Third")
	}, t)

	Test(`Given that the ListTiem has 6 items: "First", "Second", "Third", "Fourth", "Fifth" and "Sixth"
	When you want to find the item with id 4
	Then the item "Third" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First")
		list.NewItem("Second")
		list.NewItem("Third")
		list.NewItem("Fourth")
		list.NewItem("Fifth")
		list.NewItem("Sixth")

		// When
		item := list.FindItemByOrder(4)

		// Then
		assert.Equal(item.Text(), "Fifth")
	}, t)

	Test(`Given that the ListTiem has 6 items: "First", "Second", "Third", "Fourth", "Fifth" and "Sixth"
	When you want to find the item with id 5
	Then the item "Third" will be returned
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("First")
		list.NewItem("Second")
		list.NewItem("Third")
		list.NewItem("Fourth")
		list.NewItem("Fifth")
		list.NewItem("Sixth")

		// When
		item := list.FindItemByOrder(5)

		// Then
		assert.Equal(item.Text(), "Sixth")
	}, t)
}

func TestIsRendered(t *testing.T) {
	assert := assert.New(t)

	Test(`it should return false, if:
		the UnorderedList is created, it has no items and has NOT been rendered. `, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")

		assert.False(list.IsRendered())
	}, t)

	Test(`it should return false, if:
		the UnorderedList is created, it has one item and has NOT been rendered.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Some text")
		assert.False(list.IsRendered())
	}, t)

	Test(`it should return true, if:
		the UnorderedList is created, it has no items and has been rendered.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.Render()
		assert.True(list.IsRendered())
	}, t)

	Test(`it should return true, if:
		the UnorderedList is created, it has one item and has been rendered.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Some text")
		list.Render()
		assert.True(list.IsRendered())
	}, t)

	Test(`it should return false, if:
		the UnorderedList is created, it has one item and has been rendered and then another item is added.`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Some text")
		list.Render()
		list.NewItem("Some other text")
		assert.False(list.IsRendered())
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Multi \n line", "Item 3", "Item 4"
		When we check to see if the list is rendered
		Then we should see that the list has been rendered
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("This \n is \n multi \n line")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()
		// When
		isRendered := list.IsRendered()

		// Then
		assert.True(isRendered)
	}, t)
}

func TestRender(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that we create a list with header text "List name"
		And the list does not have any list items
		When we render the changes
		The output should contain a line "List name"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")

		// When
		lines := list.Render()

		// Then
		assert.Equal(lines, []string{"List name"})
	}, t)

	Test(`
		Given that we create a list with header text "List name"
		And the list has one item with text "Item 1"
		When we render the changes
		The output should contain two lines: "List name" and "\tItem1"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")

		// When
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1"})
	}, t)

	Test(`
		Given that we create a list with header text "List name"
		And the list has two items: "Item 1" and "Item 2"
		When we render the changes
		The output should contain three lines:
		- "List name"
		- "\tItem 1"
		- "\tItem 2"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")

		// When
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1", "\tItem 2"})

	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and two items: "Item 1" and "Item 2
		And we add a third item "Item 3"
		When we render the changes
		The output should contain 4 lines:
		- "List name"
		- "\tItem 1"
		- "\tItem 2"
		- "\tItem 3"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")
		list.Render()
		list.NewItem("Item 3")

		// When
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1", "\tItem 2", "\tItem 3"})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 3 items: "Item 1", "Item 2" and "Item 3"
		When we edit the second item to be named "Second item" and render the changes
		The output should contain 4 lines:
		- "List name"
		- "\tItem 1"
		- "\tItem 2"
		- "\tItem 3"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("Item 2")
		list.NewItem("Item 3")
		list.Render()

		// When
		item2.Edit("Second item")
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1", "\tSecond item", "\tItem 3"})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Multi \n line", "Item 3", "Item 4"
		When we edit the second item to  "Item 2"
		The output should contain 5 lines:
		- "List name"
		- "\tItem 1"
		- "\tItem 2"
		- "\tItem 3"
		- "\tItem 4"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("Multi \n line")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()

		// When
		item2.Edit("Item 2")
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1", "\tItem 2", "\tItem 3", "\tItem 4"})
	}, t)
}

func TestListRenderWithMultilineHeaderAndItems(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that we have a list with header text "Line 1\nLine 2" and no items
		When we render the changes
		The output should contain two lines: "Line 1" and "Line 2"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "Line 1\nLine 2")
		// When
		renderLines := list.Render()
		// Then
		assert.Equal(renderLines, []string{"Line 1", "Line 2"})
	}, t)

	Test(`
		Given that we have a list with header text "Line 1\nLine 2" and one item "Item 1"
		When we render the changes
		The output should contain 3 lines
		- "Line 1"
		- "Line 2"
		- "\tItem 1"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "Line 1\nLine 2")
		list.NewItem("Item 1")
		// When
		renderLines := list.Render()
		// Then
		assert.Equal(renderLines, []string{"Line 1", "Line 2", "\tItem 1"})

	}, t)

	Test(`
		Given that we have a list with header text "Line 1\nLine 2\nLine 3\nLine 4" and one item "Item 1"
		When we render the changes
		The output should contain the following lines: "Line 1", "Line 2", "Line 3", "Line 4", "\tItem 1"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "Line 1\nLine 2\nLine 3\nLine 4")
		list.NewItem("Item 1")

		// When
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{"Line 1", "Line 2", "Line 3", "Line 4", "\tItem 1"})
	}, t)

	Test(`
		Given that we have a list with header text "List name" and two items: "Item 1 Line1\nLine2" and "Item 2"
		When we render the changes
		The output should contain the following 4 lines: 
		"List name", "Item 1 Line1", "Line2", "Item 2"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1 Line1\nLine2")
		list.NewItem("Item 2")
		// When
		renderLines := list.Render()
		// Then
		assert.Equal(renderLines, []string{"List name", "\tItem 1 Line1", "\tLine2", "\tItem 2"})
	}, t)

	Test(`
		Given that we have a list with header text "List L1\nL2" and three items:
		- "Item 1 L1\nL2\nL3"
		- "Item 2 L1\nL2\nL3\nL4"
		- "Item 3"
		When we render the changes
		The output should contain the following lines: 
		"List L1", "L2", "\tItem 1 L1", "\tL2", "\tL3", "\tItem 2 L1", "\tL2", "\tL3", "\tL4", "\tItem 3"
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List L1\nL2")
		list.NewItem("Item 1 L1\nL2\nL3")
		list.NewItem("Item 2 L1\nL2\nL3\nL4")
		list.NewItem("Item 3")
		// When
		renderLines := list.Render()
		// Then
		assert.Equal(renderLines, []string{
			"List L1", "L2", "\tItem 1 L1", "\tL2", "\tL3", "\tItem 2 L1", "\tL2", "\tL3", "\tL4", "\tItem 3",
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Item 2", "Item 3", "Item 4"
		When we edit the first item to a multi line text: "This\nis\nthe\nfirst\nitem" and we render the changes
		The output should contain the following lines: 
		"List name", "\tThis", "\tis", "\tthe", "\tfirst", "\titem", "\tItem 2", "\tItem 3", "\tItem 4", 
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		item1 := list.NewItem("Item 1")
		list.NewItem("Item 2")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()

		// When
		item1.Edit("This\nis\nthe\nfirst\nitem")
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{
			"List name", "\tThis", "\tis", "\tthe", "\tfirst", "\titem", "\tItem 2", "\tItem 3", "\tItem 4",
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Item 2", "Item 3", "Item 4"
		When we edit the second item to a multi line text: "This\nis\nthe\nsecond\nitem" and we render the changes
		The output should contain the following lines: 
		"List name", "\tItem 1", "\tThis", "\tis", "\tthe", "\tsecond", "\titem", "\tItem 3", "\tItem 4", 
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("Item 2")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()

		// When
		item2.Edit("This\nis\nthe\nsecond\nitem")
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{
			"List name", "\tItem 1", "\tThis", "\tis", "\tthe", "\tsecond", "\titem", "\tItem 3", "\tItem 4",
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Item 2", "Item 3", "Item 4"
		When we edit the last item to a multi line text: "This\nis\nthe\nlast\nitem" and we render the changes
		The output should contain the following 10 lines: 
		"List name", "\tItem 1", "\tItem 2", "\tItem 3", "\tThis", "\tis", "\tthe", "\tlast", "\titem",
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")
		list.NewItem("Item 3")
		lastItem := list.NewItem("Item 4")
		list.Render()

		// When
		lastItem.Edit("This\nis\nthe\nlast\nitem")
		renderLines := list.Render()

		// Then
		assert.Equal(renderLines, []string{
			"List name", "\tItem 1", "\tItem 2", "\tItem 3", "\tThis", "\tis", "\tthe", "\tlast", "\titem",
		})
	}, t)
}
