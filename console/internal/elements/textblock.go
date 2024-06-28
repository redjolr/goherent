package elements

import (
	"regexp"
	"strings"
)

type Textblock struct {
	id                   string
	lines                []string
	changedWithSameWidth bool
}

func NewTextBlock(id string, text string) Textblock {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)
	return Textblock{
		id:                   id,
		lines:                lines,
		changedWithSameWidth: true,
	}
}

func (tb *Textblock) HasId(id string) bool {
	return tb.id == id
}

func (tb *Textblock) Lines() []string {
	return tb.lines
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

func (tb *Textblock) Write(text string) {
	oldWidth := tb.Width()
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)
	tb.lines = lines
	newWidth := tb.Width()
	tb.padWithWhiteSpaces(newWidth)
	if oldWidth == newWidth {
		tb.changedWithSameWidth = true
	}
}

func (tb *Textblock) Render() string {
	tb.changedWithSameWidth = false
	return strings.Join(tb.Lines(), "\n")
}

func (tb *Textblock) HasChangedWithSameWidth() bool {
	return tb.changedWithSameWidth
}

func (tb *Textblock) padWithWhiteSpaces(width int) {
	for i, line := range tb.lines {
		if len(line) < width {
			tb.lines[i] = line + strings.Repeat(" ", width-len(line))
		}
	}
}
