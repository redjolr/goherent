package console

type Element interface {
	HasId(id string) bool
	HasChanged() bool
	Render() []string
	IsRendered() bool
	Width() int
	Height() int
}
