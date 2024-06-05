package console

type Terminal struct {
	areas    []Area
	rendered bool
}

func NewTerminal() Terminal {
	return Terminal{
		areas:    []Area{},
		rendered: true,
	}
}

func (t *Terminal) NewTextBlock(text string) Textblock {
	textBlock := NewTextBlock(text)
	t.areas = append(t.areas, &textBlock)
	return textBlock
}

func (t *Terminal) NewUnorderedList(headingText string) UnorderedList {
	list := NewUnorderedList(headingText)
	t.areas = append(t.areas, &list)
	t.rendered = false
	return list
}

func (t *Terminal) Render() {
	if t.rendered {
		return
	}
	for _, area := range t.areas {
		area.render()
	}
	t.rendered = true
}
