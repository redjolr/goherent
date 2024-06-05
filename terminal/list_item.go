package terminal

type ListItem struct {
	id   int
	text string
}

func (li *ListItem) Text() string {
	return li.text
}
