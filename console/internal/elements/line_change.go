package elements

import (
	"github.com/redjolr/goherent/console/internal/utils"
)

type LineChange struct {
	Before     string
	After      string
	IsAnUpdate bool
}

func (rc LineChange) HasLineCountChanged() bool {
	lineLengthBefore := len(utils.SplitStringByNewLine(rc.Before))
	lineLengthAfter := len(utils.SplitStringByNewLine(rc.After))
	return lineLengthAfter != lineLengthBefore
}
