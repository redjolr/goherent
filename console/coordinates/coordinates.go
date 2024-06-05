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

func (c Coordinates) OffsetX(x int) Coordinates {
	c.X += x
	return c
}
