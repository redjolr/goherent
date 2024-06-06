package console

import "fmt"

type ListItem struct {
	id       int
	text     string
	rendered bool
}

func (li *ListItem) NewListItem(id int, text string) ListItem {
	return ListItem{
		id:       id,
		text:     text,
		rendered: true,
	}
}

func (li *ListItem) Text() string {
	return li.text
}

func (li *ListItem) render() {
	fmt.Println(li.text)
	li.rendered = true
}

func (li *ListItem) isRendered() bool {
	return li.rendered
}
