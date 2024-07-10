package console

import "github.com/redjolr/goherent/console/internal/elements"

type Element interface {
	HasId(id string) bool
	HasChanged() bool
	Render() []elements.LineChange
	IsRendered() bool
	Width() int
	Height() int
}
