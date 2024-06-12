package console

import (
	"regexp"
	"strings"
)

type Textblock struct {
	lines                []string
	changedWithSameWidth bool
}

func NewTextBlock(text string) Textblock {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)
	return Textblock{
		lines:                lines,
		changedWithSameWidth: true,
	}
}

func (tb *Textblock) Lines() []string {
	return tb.lines
}

func (tb *Textblock) height() int {
	return len(tb.lines)
}

func (tb *Textblock) width() int {
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
	oldWidth := tb.width()
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)
	tb.lines = lines
	newWidth := tb.width()
	tb.padWithWhiteSpaces(newWidth)
	if oldWidth == newWidth {
		tb.changedWithSameWidth = true
	}
}

func (tb *Textblock) render() string {
	tb.changedWithSameWidth = false
	return strings.Join(tb.Lines(), "\n")
}

func (ul *Textblock) hasChangedWithSameWidth() bool {
	return ul.changedWithSameWidth
}

func (tb *Textblock) padWithWhiteSpaces(width int) {
	for i, line := range tb.lines {
		if len(line) < width {
			tb.lines[i] = line + strings.Repeat(" ", width-len(line))
		}
	}
}
