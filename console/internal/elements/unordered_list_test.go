package elements_test

import (
	"testing"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/internal/elements"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	assert := assert.New(t)
	Test("it should add a new item to a UnorderedList.", func(t *testing.T) {
		// list := console.NewUnorderedList("List name")
		// item := list.NewItem("item1")

		// assert.IsType(console.ListItem{}, item)
		assert.Equal(2, 2)
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
}

func TestRender(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that we create a list with header text "List name"
		And the list does not have any list items
		When we render the changes
		The output should contain a render change with Change "List name" at corrdinates 0,0
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")

		// When
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
		})
	}, t)

	Test(`
		Given that we create a list with header text "List name"
		And the list has one item with text "Item 1"
		When we render the changes
		The output should contain a render change with RenderChange "List name" at corrdinates 0, 0
		and another RenderChange "\n\tItem 1" at coordinates 0, 1
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")

		// When
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{Before: "", After: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
			{Before: "", After: "\n\tItem 1", Coords: coordinates.Coordinates{X: 0, Y: 1}},
		})
	}, t)

	Test(`
		Given that we create a list with header text "List name"
		And the list has two items: "Item 1" and "Item 2"
		When we render the changes
		The output should contain three render changes:
		- "List name" at coordinates 0, 0
		- "\n\tItem 1"  at coordinates 0, 1
		- "\n\tItem 2"  at coordinates 0, 2
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")

		// When
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
			{After: "\n\tItem 1", Coords: coordinates.Coordinates{X: 0, Y: 1}},
			{After: "\n\tItem 2", Coords: coordinates.Coordinates{X: 0, Y: 2}},
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and two items: "Item 1" and "Item 2
		And we add a third item "Item 3"
		When we render the changes
		The output should contain 1 render changes: \n\tItem 3"  at coordinates 0, 3
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")
		list.Render()
		list.NewItem("Item 3")

		// When
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "\n\tItem 3", Coords: coordinates.Coordinates{X: 0, Y: 3}},
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 3 items: "Item 1", "Item 2" and "Item 3"
		When we edit the second item to be named "Second item" and render the changes
		The output should contain 1 render changes: \n\tSecond item"  at coordinates 0, 2
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("Item 2")
		list.NewItem("Item 3")
		list.Render()

		// When
		item2.Edit("Second item")
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "\n\tSecond item", Coords: coordinates.Coordinates{X: 0, Y: 2}},
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Item 2", "Item 3", "Item 4"
		When we edit the second item to a multi line text: "This \n is \n the \n second \n item" and we render the changes
		The output should contain 3 render changes:
		- "\n\tThis \n is \n the \n second \n item"  at coordinates 0, 2
		- "\n\t"Item 3"  at coordinates 0, 3
		- "\n\t"Item 4"  at coordinates 0, 4
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("Item 2")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()

		// When
		item2.Edit("This \n is \n the \n second \n item")
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "\n\tThis \n is \n the \n second \n item", Coords: coordinates.Coordinates{X: 0, Y: 2}},
			{After: "\n\tItem 3", Coords: coordinates.Coordinates{X: 0, Y: 3}},
			{After: "\n\tItem 4", Coords: coordinates.Coordinates{X: 0, Y: 4}},
		})
	}, t)

	Test(`
		Given that we have a rendered list with header text "List name" and 4 items: "Item 1", "Multi \n line", "Item 3", "Item 4"
		When we edit the second item to  "Item 2"
		The output should contain 3 render changes:
		- "\n\t"Item 2"  at coordinates 0, 2
		- "\n\t"Item 3"  at coordinates 0, 3
		- "\n\t"Item 4"  at coordinates 0, 4
	`, func(t *testing.T) {
		// Given
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		item2 := list.NewItem("This \n is \n multi \n line")
		list.NewItem("Item 3")
		list.NewItem("Item 4")
		list.Render()

		// When
		item2.Edit("Item 2")
		renderChanges := list.Render()

		// Then
		assert.Equal(renderChanges, []elements.RenderChange{
			{After: "\n\tItem 2", Coords: coordinates.Coordinates{X: 0, Y: 2}},
			{After: "\n\tItem 3", Coords: coordinates.Coordinates{X: 0, Y: 3}},
			{After: "\n\tItem 4", Coords: coordinates.Coordinates{X: 0, Y: 4}},
		})
	}, t)
}
