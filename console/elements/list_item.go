package elements

import (
	"strings"

	"github.com/redjolr/goherent/console/internal/utils"
)

type ListItem struct {
	id       string
	order    int
	lines    []string
	rendered bool
}

func (li *ListItem) Text() string {
	return strings.Join(li.lines, "\n")
}

func (li *ListItem) Edit(newText string) {
	lines := utils.SplitStringByNewLine(newText)
	li.rendered = false // Test it when you edit it with the same text
	li.lines = lines

}

func (li *ListItem) Render() []string {
	li.rendered = true
	return li.lines
}

func (li *ListItem) IsRendered() bool {
	return li.rendered
}

func (li *ListItem) MarkUnrendered() {
	li.rendered = false
}
