package console

type Element interface {
	render() string
	hasChangedWithSameWidth() bool
	width() int
	height() int
}
