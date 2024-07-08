package console

import "github.com/redjolr/goherent/console/internal/elements"

type Element interface {
	HasId(id string) bool
	Render() []elements.RenderChange
	IsRendered() bool
	Width() int
	Height() int
}
