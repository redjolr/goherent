package terminal

type Terminal interface {
	Print(text string)
	MoveDown(n int)
	MoveLeft(n int)
	MoveUp(n int)
	Height() int
}
