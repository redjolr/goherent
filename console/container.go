package console

import "slices"

type Container struct {
	areas []Area
}

func NewContainer() Container {
	return Container{
		areas: []Area{},
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
	atLeastOneAreaUnrendered := slices.ContainsFunc(c.areas, func(area Area) bool {
		return !area.isRendered()
	})
	return !atLeastOneAreaUnrendered
}
