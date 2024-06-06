package console

import "slices"

type Terminal struct {
	areas []Area
}

func NewTerminal() Terminal {
	return Terminal{
		areas: []Area{},
	}
}

func (t *Terminal) NewTextBlock(text string) *Textblock {
	textBlock := NewTextBlock(text)
	t.areas = append(t.areas, &textBlock)
	return &textBlock
}

func (t *Terminal) NewUnorderedList(headingText string) *UnorderedList {
	list := NewUnorderedList(headingText)
	t.areas = append(t.areas, &list)
	return &list
}

func (t *Terminal) Render() {
	if t.IsRendered() {
		return
	}
	for _, area := range t.areas {
		area.render()
	}
}

func (t *Terminal) IsRendered() bool {
	atLeastOneAreaUnrendered := slices.ContainsFunc(t.areas, func(area Area) bool {
		return !area.isRendered()
	})
	return !atLeastOneAreaUnrendered
}
