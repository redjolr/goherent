package elements_test

import (
	"testing"

	"github.com/redjolr/goherent/console/internal/elements"
	. "github.com/redjolr/goherent/pkg"

	"github.com/stretchr/testify/assert"
)

func TestListItemIsRendered(t *testing.T) {
	assert := assert.New(t)

	Test("it should return false, if ListItem is created but not rendered.", func(t *testing.T) {
		li := elements.NewListItem(0, "Some list item")
		assert.False(li.IsRendered())
	}, t)

	Test("it should return false, if ListItem is created and then rendered.", func(t *testing.T) {
		li := elements.NewListItem(0, "Some list item")
		li.Render()
		assert.True(li.IsRendered())
	}, t)
}
