package console

import (
	"slices"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/internal/elements"
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

func (c *Console) NewUnorderedList(id string, headingText string) *elements.UnorderedList {
	unorderedList := elements.NewUnorderedList(id, headingText)
	unorderedListElement := alignedElement{
		coords: &coordinates.Coordinates{
			X: c.cursor.X,
			Y: c.cursor.Y,
		},
		element: &unorderedList,
	}
	c.alignedElements = append(c.alignedElements, &unorderedListElement)

	return &unorderedList
}

func (c *Console) NewTextBlock(id string, text string) *elements.Textblock {
	textBlock := elements.NewTextBlock(id, text)
	textBlockElement := alignedElement{
		coords: &coordinates.Coordinates{
			X: c.cursor.X,
			Y: c.cursor.Y,
		},
		element: &textBlock,
	}

	c.alignedElements = append(c.alignedElements, &textBlockElement)

	c.MoveDown(textBlock.Height() - 1)
	c.MoveRight(textBlock.Width())

	return &textBlock
}

func (c *Console) Render() {
	if c.IsRendered() {
		return
	}
	// for _, alignedElement := range c.alignedElements {
	// 	renderChanges := alignedElement.element.Render()
	// 	c.terminal.Print(renderChanges)
	// }
}

func (c *Console) IsRendered() bool {
	atLeastOneElementUnrendered := slices.ContainsFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return !alignedElement.element.IsRendered()
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

func (c *Console) HasElementWithId(id string) bool {
	return slices.ContainsFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return alignedElement.element.HasId(id)
	})
}
