package elements

type ListItem struct {
	order    int
	text     string
	rendered bool
}

func NewListItem(order int, text string) ListItem {
	return ListItem{
		order:    order,
		text:     text,
		rendered: false,
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
