package elements

import "github.com/redjolr/goherent/console/coordinates"

type RenderChange struct {
	Before string
	After  string
	Coords coordinates.Coordinates
}
