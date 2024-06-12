package console

type Element interface {
	render()
	isRendered() bool
}
