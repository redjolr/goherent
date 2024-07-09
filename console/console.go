package console

import (
	"slices"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/cursor"
	"github.com/redjolr/goherent/console/internal/elements"
	"github.com/redjolr/goherent/console/internal/utils"
	"github.com/redjolr/goherent/console/terminal"
)

type alignedElement struct {
	coords  *coordinates.Coordinates
	element Element
}

type Console struct {
	terminal        terminal.Terminal
	alignedElements []*alignedElement
	cursor          *cursor.Cursor
}

func NewConsole(terminal terminal.Terminal, cursor *cursor.Cursor) Console {
	return Console{
		terminal:        terminal,
		alignedElements: []*alignedElement{},
		cursor:          cursor,
	}
}

func (c *Console) NewUnorderedList(id string, headingText string) *elements.UnorderedList {
	unorderedList := elements.NewUnorderedList(id, headingText)
	unorderedListElement := alignedElement{
		coords:  c.cursor.Coordinates(),
		element: &unorderedList,
	}
	c.alignedElements = append(c.alignedElements, &unorderedListElement)

	return &unorderedList
}

func (c *Console) NewTextBlock(id string, text string) *elements.Textblock {
	textBlock := elements.NewTextBlock(id, text)
	textBlockElement := alignedElement{
		coords:  c.cursor.Coordinates(),
		element: &textBlock,
	}

	c.alignedElements = append(c.alignedElements, &textBlockElement)

	return &textBlock
}

func (c *Console) Render() {
	if c.IsRendered() {
		return
	}
	c.cursor.GoToOrigin()

	for _, alignedElement := range c.alignedElements {
		if alignedElement.element.HasChanged() {
			renderChanges := alignedElement.element.Render()
			for _, renderChange := range renderChanges {
				c.terminal.Print(renderChange.After)
				lines := utils.SplitStringByNewLine(renderChange.After)
				lastLine := lines[len(lines)-1]
				c.cursor.MoveDown(utils.StrLinesCount(renderChange.After))
				c.cursor.MoveAtBeginningOfLine()
				c.cursor.MoveRight(len(lastLine))
			}

		}

	}
}

func (c *Console) IsRendered() bool {
	atLeastOneElementUnrendered := slices.ContainsFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return !alignedElement.element.IsRendered()
	})
	return !atLeastOneElementUnrendered
}

func (c *Console) HasElementWithId(id string) bool {
	return slices.ContainsFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return alignedElement.element.HasId(id)
	})
}
