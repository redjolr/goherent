package elements_test

import (
	"testing"

	"github.com/redjolr/goherent/console/internal/elements"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestTextBlockRender(t *testing.T) {
	assert := assert.New(t)
	Test(`
	it should render LineChange{ Before: "", After: "", coords: 0,0, IsAnUpdate: false },
	if we pass an empty string`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "")
		renderChanges := textBlock.Render()
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "", After: "", IsAnUpdate: false},
		})
	}, t)

	Test(`
	it should render LineChange{ Before: "", After: "A", coords: 0,0, IsAnUpdate: false },
	if we pass the string "A"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "A")
		renderChanges := textBlock.Render()
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "", After: "A", IsAnUpdate: false},
		})
	}, t)

	Test(`
	it should render LineChange{ Before: "", After: "\n", coords: 0,0, IsAnUpdate: false },
	if we pass the string "\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n")
		renderChanges := textBlock.Render()
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "", After: "\n", IsAnUpdate: false},
		})
	}, t)

	Test(`
	it should render LineChange{ Before: "", After: "\n\n", coords: 0,0, IsAnUpdate: false },
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n\n")
		renderChanges := textBlock.Render()
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "", After: "\n\n", IsAnUpdate: false},
		})
	}, t)

	Test(`
	it should render LineChange{ Before: "", After: "Line1 \n Line2", coords: 0,0, IsAnUpdate: false },
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "Line1 \n Line2")
		renderChanges := textBlock.Render()
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "", After: "Line1 \n Line2", IsAnUpdate: false},
		})
	}, t)

	Test(`
		Given that we have a rendered Textblock "First textblock"
		And the textblock is edited to "Second textblock"
		When we render the changes
		Then the output should contain the following render changes:
		- LineChange{ Before: "First textblock", After: "Second textblock", coords: 0,0, IsAnUpdate: true },
	`, func(t *testing.T) {
		// Given
		textBlock := elements.NewTextBlock("id", "First textblock")
		textBlock.Render()
		textBlock.Edit("Second textblock")

		// When
		renderChanges := textBlock.Render()

		// Then
		assert.Equal(renderChanges, []elements.LineChange{
			{Before: "First textblock", After: "Second textblock", IsAnUpdate: true},
		})
	}, t)
}
