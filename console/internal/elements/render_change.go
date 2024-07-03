package elements

import (
	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/internal/utils"
)

type RenderChange struct {
	Before string
	After  string
	Coords coordinates.Coordinates
}

func (rc RenderChange) HasLineCountChanged() bool {
	lineLengthBefore := len(utils.SplitStringByNewLine(rc.Before))
	lineLengthAfter := len(utils.SplitStringByNewLine(rc.After))
	return lineLengthAfter != lineLengthBefore
}
