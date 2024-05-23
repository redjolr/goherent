package terminal

type Coordinate struct {
	x int
	y int
}

func Origin() Coordinate {
	return Coordinate{
		x: 0,
		y: 0,
	}
}
