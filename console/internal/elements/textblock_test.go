package elements_test

import (
	"testing"

	"github.com/redjolr/goherent/console/internal/elements"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewTextblockNewTextBlock(t *testing.T) {
	assert := assert.New(t)
	Test(`
	it should store the line []string{""},
	if we pass an empty string`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "")
		assert.Equal(textBlock.Lines(), []string{""})
	}, t)

	Test(`
	it should store the line []string{"A"},
	if we pass the string "A"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "A")
		assert.Equal(textBlock.Lines(), []string{"A"})
	}, t)

	Test(`
	it should store two empty lines,
	if we pass the string "\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n")
		assert.Equal(textBlock.Lines(), []string{"", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store two empty lines,
	if we pass the string "\r\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\r\n")
		assert.Equal(textBlock.Lines(), []string{"", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\r\n\r\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\r\n\r\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "\n\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store the lines []string{"This is line 1", ""},
	if we pass the string "This is line 1\n" to NewTextBlock()`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "This is line 1\n")
		assert.Equal(textBlock.Lines(), []string{"This is line 1", ""})
	}, t)

	Test(`
	it should store the lines []string{"This is line 1", "This is line 2"},
	if we pass the string "This is line 1\nThis is line 2" to NewTextBlock()`, func(t *testing.T) {
		textBlock := elements.NewTextBlock("id", "This is line 1\nThis is line 2")
		assert.Equal(textBlock.Lines(), []string{"This is line 1", "This is line 2"})
	}, t)
}
