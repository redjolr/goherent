package fake_ansi_terminal

type Coordinates struct {
	X int
	Y int
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

func (c *Coordinates) OffsetY(y int) {
	c.Y += y
}

func (c *Coordinates) SetToOrigin() {
	c.X = 0
	c.Y = 0
}
