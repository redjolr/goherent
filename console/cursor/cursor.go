package cursor

import (
	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/terminal"
)

type Cursor struct {
	terminal terminal.Terminal
	coords   *coordinates.Coordinates
}

func NewCursor(terminal terminal.Terminal, origin *coordinates.Coordinates) Cursor {
	return Cursor{
		terminal: terminal,
		coords:   origin,
	}
}

func (c *Cursor) Coordinates() *coordinates.Coordinates {
	return c.coords
}

func (c *Cursor) GoToOrigin() {
	c.MoveUp(c.coords.Y)
	c.MoveAtBeginningOfLine()
}

func (c *Cursor) MoveLeft(n int) {
	c.coords.X -= n
}

func (c *Cursor) MoveAtBeginningOfLine() {
	c.MoveLeft(c.coords.X)
}

func (c *Cursor) MoveRight(n int) {
	c.coords.X += n
}

func (c *Cursor) MoveDown(n int) {
	c.coords.Y += n
}

func (c *Cursor) MoveUp(n int) {
	c.coords.Y -= n
}
