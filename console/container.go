package console

import "slices"

type Container struct {
	areas []Element
}

func NewContainer() Container {
	return Container{
		areas: []Element{},
	}
}

func (c *Container) NewTextBlock(text string) *Textblock {
	textBlock := NewTextBlock(text)
	c.areas = append(c.areas, &textBlock)
	return &textBlock
}

func (c *Container) NewUnorderedList(headingText string) *UnorderedList {
	list := NewUnorderedList(headingText)
	c.areas = append(c.areas, &list)
	return &list
}

func (c *Container) Render() {
	if c.IsRendered() {
		return
	}
	for _, area := range c.areas {
		area.render()
	}
}

func (c *Container) IsRendered() bool {
	atLeastOneElementUnrendered := slices.ContainsFunc(c.areas, func(element Element) bool {
		return !element.isRendered()
	})
	return !atLeastOneElementUnrendered
}
