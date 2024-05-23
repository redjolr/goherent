package terminal

type Grid struct {
	height         int
	width          int
	cursorPosition Coordinate
}

func EmptyCanvas() Grid {
	return Grid{
		height:         0,
		width:          0,
		cursorPosition: Origin(),
	}
}
