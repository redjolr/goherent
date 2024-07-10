package elements

import (
	"github.com/redjolr/goherent/console/internal/utils"
)

type Textblock struct {
	id          string
	lines       []string
	lineChanges []LineChange
	rendered    bool
}

func NewTextBlock(id string, text string) Textblock {
	lines := utils.SplitStringByNewLine(text)
	lineChanges := []LineChange{}
	for _, line := range lines {
		lineChanges = append(lineChanges, LineChange{
			Before:     "",
			After:      line,
			IsAnUpdate: false,
		})
	}

	return Textblock{
		id:          id,
		lines:       lines,
		lineChanges: lineChanges,
		rendered:    false,
	}
}

func (tb *Textblock) HasId(id string) bool {
	return tb.id == id
}

func (tb *Textblock) Height() int {
	return len(tb.lines)
}

func (tb *Textblock) Width() int {
	return len(tb.longestLine())
}

func (tb *Textblock) longestLine() string {
	if len(tb.lines) == 0 {
		return ""
	}

	longest := tb.lines[0]
	for _, line := range tb.lines {
		if len(line) > len(longest) {
			longest = line
		}
	}
	return longest
}

func (tb *Textblock) Edit(text string) {
	lines := utils.SplitStringByNewLine(text)
	lineChanges := []LineChange{}
	for i, line := range lines {
		if i < len(tb.lines) && tb.lines[i] != line {
			lineChanges = append(lineChanges, LineChange{
				Before:     tb.lines[i],
				After:      line,
				IsAnUpdate: true,
			})
		} else if i >= len(tb.lines) {
			lineChanges = append(lineChanges, LineChange{
				Before:     "",
				After:      line,
				IsAnUpdate: false,
			})
		}

	}

	tb.rendered = false
	tb.lineChanges = lineChanges
	tb.lines = lines
}

func (tb *Textblock) HasChanged() bool {
	return !tb.rendered
}

func (tb *Textblock) Render() []LineChange {
	if tb.rendered {
		return []LineChange{}
	}
	renderChanges := make([]LineChange, len(tb.lineChanges))
	copy(renderChanges, tb.lineChanges)
	tb.rendered = true
	tb.lineChanges = []LineChange{}
	return renderChanges
}

func (tb *Textblock) IsRendered() bool {
	return tb.rendered
}
