package elements

import (
	"slices"

	"github.com/redjolr/goherent/console/internal/utils"
)

type UnorderedList struct {
	id           string
	name         string
	items        []*ListItem
	nameRendered bool
}

func NewUnorderedList(id string, name string) UnorderedList {
	return UnorderedList{
		id:           id,
		name:         name,
		items:        []*ListItem{},
		nameRendered: false,
	}
}

func (ul *UnorderedList) NewItem(id string, text string) *ListItem {
	lines := utils.SplitStringByNewLine(text)
	item := ListItem{
		id:       id,
		order:    len(ul.items),
		lines:    lines,
		rendered: false,
	}

	ul.items = append(ul.items, &item)
	return &item
}

func (ul *UnorderedList) FindItemById(id string) *ListItem {
	if len(ul.items) == 0 {
		return nil
	}

	listItemIndex := slices.IndexFunc(ul.items, func(item *ListItem) bool {
		return item.id == id
	})

	if listItemIndex == -1 {
		return nil
	}
	return ul.items[listItemIndex]
}

func (ul *UnorderedList) FindItemByOrder(order int) *ListItem {
	if len(ul.items) == 0 {
		return nil
	}

	listItemIndex := slices.IndexFunc(ul.items, func(item *ListItem) bool {
		return item.order == order
	})

	if listItemIndex == -1 {
		return nil
	}
	return ul.items[listItemIndex]
}

func (ul *UnorderedList) EditName(newName string) {
	ul.nameRendered = false
	ul.name = newName
}

func (ul *UnorderedList) Render() []string {
	lines := utils.SplitStringByNewLine(ul.name)
	for _, item := range ul.items {
		for _, itemLine := range item.Render() {
			lines = append(lines, "\t"+itemLine)

		}
	}
	ul.nameRendered = true
	return lines
}

func (ul *UnorderedList) HasChanged() bool {
	return !ul.IsRendered()
}

func (ul *UnorderedList) IsRendered() bool {
	atLeastOneItemUnrendered := slices.ContainsFunc(ul.items, func(item *ListItem) bool {
		return !item.IsRendered()
	})
	return !atLeastOneItemUnrendered && ul.nameRendered
}

func (ul *UnorderedList) HasId(id string) bool {
	return ul.id == id
}
func (ul *UnorderedList) HasChangedWithSameWidth() bool {
	return false
}

func (ul *UnorderedList) Height() int {
	return 0
}

func (ul *UnorderedList) Width() int {
	return 0
}

func (ul *UnorderedList) markForRerenderStartingAtItem(rerenderItem *ListItem) {
	itemOrderIndex := slices.Index(ul.items, rerenderItem)
	subsequentItems := ul.items[itemOrderIndex:]
	for _, item := range subsequentItems {
		item.MarkUnrendered()
	}
	rerenderItem.MarkUnrendered()
}
