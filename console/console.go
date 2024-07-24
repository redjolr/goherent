package console

import (
	"slices"
	"strings"

	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/cursor"
	"github.com/redjolr/goherent/console/elements"
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
	renderedLines   []string
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
	goUp := c.cursor.Coordinates().Y
	goLeft := c.cursor.Coordinates().X
	c.terminal.MoveLeft(goLeft)
	c.terminal.MoveUp(goUp)
	c.cursor.GoToOrigin()

	overallLineIndex := 0
	for _, alignedElement := range c.alignedElements {

		if alignedElement.element.HasChanged() {
			lines := alignedElement.element.Render()
			for _, line := range lines {
				if overallLineIndex != 0 {
					c.terminal.Print("\n")
					c.cursor.MoveDown(1)
					c.cursor.MoveAtBeginningOfLine()
				}
				if overallLineIndex > len(c.renderedLines)-1 {
					c.renderedLines = append(c.renderedLines, line)
					c.terminal.Print(line)
					c.cursor.MoveRight(len(line))
				} else if line != c.renderedLines[overallLineIndex] {
					if len(line) < len(c.renderedLines[overallLineIndex]) {
						c.renderedLines[overallLineIndex] = utils.StrRightPad(line, " ", len(c.renderedLines[overallLineIndex]))
						c.terminal.Print(c.renderedLines[overallLineIndex])
						c.cursor.MoveRight(len(c.renderedLines[overallLineIndex]))
					} else if len(line) >= len(c.renderedLines[overallLineIndex]) {
						c.renderedLines[overallLineIndex] = line
						c.terminal.Print(line)
						c.cursor.MoveRight(len(line))

					}
				}
				overallLineIndex += 1
			}
		} else {
			overallLineIndex += alignedElement.element.Height()
		}
	}
	for ; overallLineIndex < len(c.renderedLines); overallLineIndex++ {
		goUp := c.cursor.Coordinates().Y
		goLeft := c.cursor.Coordinates().X
		c.terminal.MoveLeft(goLeft)
		c.terminal.MoveUp(goUp)
		c.terminal.MoveDown(overallLineIndex)
		c.cursor.GoToOrigin()
		c.cursor.MoveDown(overallLineIndex)
		whitespacesOverwrite := strings.Repeat(" ", len(c.renderedLines[overallLineIndex]))
		c.renderedLines[overallLineIndex] = whitespacesOverwrite
		c.terminal.Print(whitespacesOverwrite)
		c.cursor.MoveRight(len(whitespacesOverwrite))
		c.cursor.MoveDown(1)
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

func (c *Console) GetElementWithId(id string) Element {
	idx := slices.IndexFunc(c.alignedElements, func(alignedElement *alignedElement) bool {
		return alignedElement.element.HasId(id)
	})

	if idx == -1 {
		return nil
	}
	return c.alignedElements[idx].element
}
