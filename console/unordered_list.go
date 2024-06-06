package console

import (
	"fmt"
	"slices"
)

type UnorderedList struct {
	headingText string
	items       []ListItem
	rendered    bool
}

func NewUnorderedList(headingText string) UnorderedList {
	return UnorderedList{
		headingText: headingText,
		items:       []ListItem{},
		rendered:    false,
	}
}

func (ul *UnorderedList) NewItem(text string) ListItem {
	item := ListItem{
		id:       len(ul.items),
		text:     text,
		rendered: false,
	}

	ul.items = append(ul.items, item)
	return item
}

func (ul *UnorderedList) FindItemById(id int) *ListItem {
	if len(ul.items) == 0 {
		return nil
	}

	listItemIndex := slices.IndexFunc(ul.items, func(item ListItem) bool {
		return item.id == id
	})

	if listItemIndex == -1 {
		return nil
	}
	return &ul.items[listItemIndex]
}

func (ul *UnorderedList) render() {
	fmt.Println(ul.headingText)
	for _, item := range ul.items {
		if !item.isRendered() {
			item.render()
		}
	}
	ul.rendered = true
}

func (ul *UnorderedList) isRendered() bool {
	atLeastOneItemUnrendered := slices.ContainsFunc(ul.items, func(item ListItem) bool {
		return !item.isRendered()
	})
	return !atLeastOneItemUnrendered
}
