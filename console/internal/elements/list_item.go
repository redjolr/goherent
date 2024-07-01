package elements

import "github.com/redjolr/goherent/console/coordinates"

type ListItem struct {
	order        int
	text         string
	rendered     bool
	renderChange RenderChange
}

func NewListItem(order int, text string) ListItem {
	return ListItem{
		order:        order,
		text:         text,
		rendered:     false,
		renderChange: RenderChange{},
	}
}

func (li *ListItem) Text() string {
	return li.text
}

func (li *ListItem) Edit(newText string) {
	li.rendered = false
	li.renderChange = RenderChange{
		Before: li.text,
		After:  newText,
		Coords: coordinates.Origin(),
	}
	li.text = newText

}

func (li *ListItem) Render() RenderChange {
	renderChange := li.renderChange
	li.rendered = true
	li.renderChange = RenderChange{}
	return renderChange
}

func (li *ListItem) ReRender() RenderChange {
	return RenderChange{
		Before: li.text,
		After:  li.text,
		Coords: coordinates.Origin(),
	}
}

func (li *ListItem) IsRendered() bool {
	return li.rendered
}
