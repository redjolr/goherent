package elements

import (
	"slices"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/internal/utils"
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
		renderChange: RenderChange{
			Before: "",
			After:  text,
			Coords: coordinates.Origin(),
		},
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
	currentRenderedHeight := 0
	if !ul.headingTextRendered {
		renderChanges = append(renderChanges, RenderChange{
			After:  ul.headingText,
			Coords: coordinates.Coordinates{X: 0, Y: currentRenderedHeight},
		})
		ul.headingTextRendered = true
	}
	currentRenderedHeight += utils.StrLinesCount(ul.headingText)

	for _, item := range ul.items {
		if !item.IsRendered() {
			renderChange := item.Render()
			renderChanges = append(
				renderChanges,
				RenderChange{
					After:  "\t" + renderChange.After,
					Coords: coordinates.Coordinates{X: 0, Y: currentRenderedHeight},
				},
			)
			if renderChange.IsAnUpdate && renderChange.HasLineCountChanged() {
				ul.markForRerenderStartingAtItem(item)
			}
		}
		currentRenderedHeight += utils.StrLinesCount(item.Text())
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

func (ul *UnorderedList) markForRerenderStartingAtItem(rerenderItem *ListItem) {
	itemOrderIndex := slices.Index(ul.items, rerenderItem)
	subsequentItems := ul.items[itemOrderIndex:]
	for _, item := range subsequentItems {
		item.MarkUnrendered()
	}
	rerenderItem.MarkUnrendered()
}
