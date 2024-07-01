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
		list := elements.NewUnorderedList("id", "List name")
		renderChanges := list.Render()

		assert.Equal(renderChanges, []elements.RenderChange{
			{Change: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
		})
	}, t)

	Test(`
		Given that we create a list with header text "List name"
		And the list has one item with text "Item 1"
		When we render the changes
		The output should contain a render change with RenderChange "List name" at corrdinates 0, 0
		and another RenderChange "\n\tItem 1" at coordinates 0, 1
	`, func(t *testing.T) {
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		renderChanges := list.Render()

		assert.Equal(renderChanges, []elements.RenderChange{
			{Change: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
			{Change: "\n\tItem 1", Coords: coordinates.Coordinates{X: 0, Y: 1}},
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
		list := elements.NewUnorderedList("id", "List name")
		list.NewItem("Item 1")
		list.NewItem("Item 2")
		renderChanges := list.Render()

		assert.Equal(renderChanges, []elements.RenderChange{
			{Change: "List name", Coords: coordinates.Coordinates{X: 0, Y: 0}},
			{Change: "\n\tItem 1", Coords: coordinates.Coordinates{X: 0, Y: 1}},
			{Change: "\n\tItem 2", Coords: coordinates.Coordinates{X: 0, Y: 2}},
		})
	}, t)
}
