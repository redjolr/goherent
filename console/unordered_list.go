package console

import "slices"

type UnorderedList struct {
	items []ListItem
}

func NewUnorderedList() UnorderedList {
	return UnorderedList{
		items: []ListItem{},
	}
}

func (ul *UnorderedList) NewItem(text string) ListItem {
	item := ListItem{
		id:   len(ul.items),
		text: text,
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
