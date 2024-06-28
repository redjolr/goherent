package elements

type ListItem struct {
	id       int
	text     string
	rendered bool
}

func NewListItem(id int, text string) ListItem {
	return ListItem{
		id:       id,
		text:     text,
		rendered: true,
	}
}

func (li *ListItem) Text() string {
	return li.text
}

func (li *ListItem) Render() string {
	li.rendered = true
	return li.text
}

func (li *ListItem) IsRendered() bool {
	return li.rendered
}
