package elements

import (
	"github.com/redjolr/goherent/console/internal/utils"
)

type Textblock struct {
	id       string
	lines    []string
	rendered bool
}

func NewTextBlock(id string, text string) Textblock {
	lines := utils.SplitStringByNewLine(text)

	return Textblock{
		id:       id,
		lines:    lines,
		rendered: false,
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
	tb.rendered = false // Test it when you edit it with the same text
	tb.lines = lines
}

func (tb *Textblock) HasChanged() bool {
	return !tb.rendered
}

func (tb *Textblock) Render() []string {
	tb.rendered = true
	return tb.lines
}

func (tb *Textblock) IsRendered() bool {
	return tb.rendered
}
