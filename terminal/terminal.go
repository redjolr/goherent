package terminal

type Terminal interface {
	Print(text string)
	Printf(text string, args ...any)
	MoveDown(n int)
	MoveUp(n int)
	MoveLeft(n int)
	MoveRight(n int)
	Height() int
}
