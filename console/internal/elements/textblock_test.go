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
	it should render the lines: "",
	if we create a new Textblock with an empty string.`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "")
		renderLines := textBlock.Render()
		assert.Equal(renderLines, []string{""})
	}, t)

	Test(`
		it should render the lines: "A"
		if we create a new Textblock with "A" as text.`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "A")
		renderLines := textBlock.Render()
		assert.Equal(renderLines, []string{"A"})

	}, t)

	Test(`
		it should render two empty lines
		if we create a new Textblock with "\n" as text.`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n")
		renderLines := textBlock.Render()
		assert.Equal(renderLines, []string{"", ""})
	}, t)

	Test(`
		it should render three empty lines
		if we create a new Textblock with "\n\n" as text.`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n\n")
		renderLines := textBlock.Render()
		assert.Equal(renderLines, []string{"", "", ""})
	}, t)

	Test(`
		it should render these lines: "Line 1", "Line 2"
		if we create a new Textblock with "Line 1\nLine 2" as text.`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "Line 1\nLine 2")
		renderLines := textBlock.Render()
		assert.Equal(renderLines, []string{"Line 1", "Line 2"})
	}, t)

	Test(`
		Given that we have a rendered Textblock "First textblock"
		And the textblock is edited to "Second textblock"
		When we render the changes
		Then the output should contain the following line: "Second textblock"
	`, func(t *testing.T) {
		// Given
		textBlock := elements.NewTextBlock("id", "First textblock")
		textBlock.Render()
		textBlock.Edit("Second textblock")

		// When
		renderLines := textBlock.Render()

		// Then
		assert.Equal(renderLines, []string{"Second textblock"})
	}, t)
}
