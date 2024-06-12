package console

import (
	"slices"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/terminal"
)

type alignedElement struct {
	coords  *coordinates.Coordinates
	element Element
}

type Console struct {
	terminal        terminal.Terminal
	alignedElements []*alignedElement
	cursor          *coordinates.Coordinates
}

func NewConsole(terminal terminal.Terminal) Console {
	origin := coordinates.Origin()
	return Console{
		terminal:        terminal,
		alignedElements: []*alignedElement{},
		cursor:          &origin,
	}
}

func (c *Console) NewTextBlock(text string) *Textblock {
	textBlock := NewTextBlock(text)
	textBlockElement := alignedElement{
		coords: &coordinates.Coordinates{
			X: c.cursor.X,
			Y: c.cursor.Y,
		},
		element: &textBlock,
	}

	c.alignedElements = append(c.alignedElements, &textBlockElement)

	c.MoveDown(textBlock.height() - 1)
	c.MoveRight(textBlock.width())

	return &textBlock
}

func (c *Console) Render() {
	if c.IsRendered() {
		return
	}
	for _, alignedElement := range c.alignedElements {
		if alignedElement.element.hasChangedWithSameWidth() {
			if c.cursor.X >= alignedElement.coords.X {
				c.MoveLeft(c.cursor.X - alignedElement.coords.X)
			} else {
				c.MoveRight(alignedElement.coords.X - c.cursor.X)
			}
			if c.cursor.Y >= alignedElement.coords.Y {
				c.MoveUp(c.cursor.Y - alignedElement.coords.Y)
			} else {
				c.MoveDown(alignedElement.coords.Y - c.cursor.Y)
			}

			renderText := alignedElement.element.render()
			c.terminal.Print(renderText)
			c.cursor.MoveRight(alignedElement.element.width())
			c.cursor.MoveDown(alignedElement.element.height())
		}
	}
}

func (c *Console) IsRendered() bool {
	atLeastOneElementUnrendered := slices.ContainsFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return alignedElement.element.hasChangedWithSameWidth()
	})
	return !atLeastOneElementUnrendered
}

func (c *Console) MoveLeft(n int) {
	c.terminal.Print(terminal.MoveCursorLeftNCols(n))
	c.cursor.MoveLeft(n)
}

func (c *Console) MoveRight(n int) {
	c.terminal.Print(terminal.MoveCursorRightNCols(n))
	c.cursor.MoveRight(n)
}

func (c *Console) MoveDown(n int) {
	c.terminal.Print(terminal.MoveCursorDownNRows(n))
	c.cursor.MoveDown(n)
}

func (c *Console) MoveUp(n int) {
	c.terminal.Print(terminal.MoveCursorUpNRows(n))
	c.cursor.MoveUp(n)
}
