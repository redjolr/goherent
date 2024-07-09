package cursor

import (
	"github.com/redjolr/goherent/console/coordinates"
)

type Cursor struct {
	coords *coordinates.Coordinates
}

func NewCursor() Cursor {
	origin := coordinates.Origin()
	return Cursor{
		coords: &origin,
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
	c.coords.OffsetX(-n)
}

func (c *Cursor) MoveAtBeginningOfLine() {
	c.MoveLeft(c.coords.X)
}

func (c *Cursor) MoveRight(n int) {
	c.coords.OffsetX(n)
}

func (c *Cursor) MoveDown(n int) {
	c.coords.OffsetY(n)
}

func (c *Cursor) MoveUp(n int) {
	c.coords.OffsetY(-n)
}
