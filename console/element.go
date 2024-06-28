package console

type Element interface {
	HasId(id string) bool
	Render() string
	HasChangedWithSameWidth() bool
	Width() int
	Height() int
}
