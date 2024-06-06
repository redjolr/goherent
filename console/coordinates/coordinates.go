package coordinates

type Coordinates struct {
	X int
	Y int
}

func New(x int, y int) Coordinates {
	return Coordinates{
		X: x,
		Y: y,
	}
}

func Origin() Coordinates {
	return Coordinates{
		X: 0,
		Y: 0,
	}
}

func (c *Coordinates) OffsetX(x int) {
	c.X += x
}

func (c *Coordinates) MoveRight(x int) {
	c.X += x
}

func (c *Coordinates) MoveLeft(x int) {
	c.X -= x
}

func (c *Coordinates) OffsetY(y int) {
	c.Y += y
}

func (c *Coordinates) GoToOrigin() {
	c.X = 0
	c.Y = 0
}
