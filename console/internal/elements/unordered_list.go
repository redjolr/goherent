package elements

import (
	"slices"

	"github.com/redjolr/goherent/console/coordinates"
)

type UnorderedList struct {
	id                  string
	headingText         string
	items               []*ListItem
	headingTextRendered bool
}

func NewUnorderedList(id string, headingText string) UnorderedList {
	return UnorderedList{
		id:                  id,
		headingText:         headingText,
		items:               []*ListItem{},
		headingTextRendered: false,
	}
}

func (ul *UnorderedList) NewItem(text string) *ListItem {
	item := ListItem{
		order:    len(ul.items),
		text:     text,
		rendered: false,
	}

	ul.items = append(ul.items, &item)
	return &item
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

func (ul *UnorderedList) Render() []RenderChange {
	renderChanges := []RenderChange{}

	if !ul.headingTextRendered {
		renderChanges = append(renderChanges, RenderChange{Change: ul.headingText, Coords: coordinates.Coordinates{X: 0, Y: 0}})
	}

	for order, item := range ul.items {
		if !item.IsRendered() {
			renderChanges = append(
				renderChanges,
				RenderChange{
					Change: "\n\t" + item.Render(),
					Coords: coordinates.Coordinates{X: 0, Y: order + 1},
				},
			)
		}
	}
	return renderChanges
}

func (ul *UnorderedList) IsRendered() bool {
	atLeastOneItemUnrendered := slices.ContainsFunc(ul.items, func(item *ListItem) bool {
		return !item.IsRendered()
	})
	return !atLeastOneItemUnrendered && ul.headingTextRendered
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
