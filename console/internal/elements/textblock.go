package elements

import (
	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/internal/utils"
)

type Textblock struct {
	id            string
	text          string
	renderChanges []RenderChange
	rendered      bool
}

func NewTextBlock(id string, text string) Textblock {

	renderChange := RenderChange{
		Before:     "",
		After:      text,
		Coords:     coordinates.Origin(),
		IsAnUpdate: false,
	}

	return Textblock{
		id:            id,
		text:          text,
		renderChanges: []RenderChange{renderChange},
		rendered:      false,
	}
}

func (tb *Textblock) HasId(id string) bool {
	return tb.id == id
}

func (tb *Textblock) Height() int {
	return utils.StrLinesCount(tb.text)
}

func (tb *Textblock) Width() int {
	return len(tb.longestLine())
}

func (tb *Textblock) longestLine() string {
	lines := utils.SplitStringByNewLine(tb.text)
	if len(lines) == 0 {
		return ""
	}

	longest := lines[0]
	for _, line := range lines {
		if len(line) > len(longest) {
			longest = line
		}
	}
	return longest
}

func (tb *Textblock) Edit(text string) {
	tb.rendered = false
	tb.renderChanges = []RenderChange{{
		Before:     tb.text,
		After:      text,
		Coords:     coordinates.Origin(),
		IsAnUpdate: true,
	}}
	tb.text = text
}

func (tb *Textblock) HasChanged() bool {
	return !tb.rendered
}

func (tb *Textblock) Render() []RenderChange {
	if tb.rendered {
		return []RenderChange{}
	}
	renderChanges := make([]RenderChange, len(tb.renderChanges))
	copy(renderChanges, tb.renderChanges)
	tb.rendered = true
	tb.renderChanges = []RenderChange{}
	return renderChanges
}

func (tb *Textblock) IsRendered() bool {
	return tb.rendered
}
