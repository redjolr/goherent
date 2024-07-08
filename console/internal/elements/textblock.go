package elements

import (
	"regexp"
	"strings"

	"github.com/redjolr/goherent/console/coordinates"
)

type Textblock struct {
	id            string
	lines         []string
	renderChanges []RenderChange
	rendered      bool
}

func NewTextBlock(id string, text string) Textblock {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)

	var renderChanges []RenderChange = []RenderChange{}
	renderCoordinates := coordinates.New(0, -1)

	for _, line := range lines {
		renderCoordinates.MoveDown(1)
		renderChanges = append(renderChanges, RenderChange{
			After:      line,
			Coords:     coordinates.New(len(line), renderCoordinates.Y),
			IsAnUpdate: false,
		})
	}

	return Textblock{
		id:            id,
		lines:         lines,
		renderChanges: renderChanges,
		rendered:      false,
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
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(text, -1)
	tb.lines = lines
	newWidth := tb.Width()
	tb.padWithWhiteSpaces(newWidth)
}

func (tb *Textblock) RenderUpdates() []RenderChange {
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

func (tb *Textblock) padWithWhiteSpaces(width int) {
	for i, line := range tb.lines {
		if len(line) < width {
			tb.lines[i] = line + strings.Repeat(" ", width-len(line))
		}
	}
}
