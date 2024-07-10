package elements

type ListItem struct {
	order      int
	text       string
	rendered   bool
	lineChange LineChange
}

func (li *ListItem) Text() string {
	return li.text
}

func (li *ListItem) Edit(newText string) {
	li.rendered = false
	li.lineChange = LineChange{
		Before:     li.text,
		After:      newText,
		IsAnUpdate: true,
	}
	li.text = newText

}

func (li *ListItem) Render() LineChange {
	renderChange := li.lineChange
	li.rendered = true
	li.lineChange = LineChange{}
	return renderChange
}

func (li *ListItem) IsRendered() bool {
	return li.rendered
}

func (li *ListItem) MarkUnrendered() {
	li.rendered = false
	li.lineChange = LineChange{
		Before: li.text,
		After:  li.text,
	}
}
