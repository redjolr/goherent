package utils_test

import (
	"testing"

	"github.com/redjolr/goherent/console/internal/utils"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	assert := assert.New(t)
	Test("it should return [''], if you pass an empty string.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine(""), []string{""})
	}, t)

	Test("it should return ['H'], if you pass 'H'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("H"), []string{"H"})
	}, t)

	Test("it should return ['Hello'], if you pass 'Hello'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hello"), []string{"Hello"})
	}, t)

	Test("it should return ['Hi', 'There'], if you pass 'Hi\nThere'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hi\nThere"), []string{"Hi", "There"})
	}, t)

	Test("it should return ['Hi', 'There'], if you pass 'Hi\r\nThere'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hi\r\nThere"), []string{"Hi", "There"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\nThere\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hi\nThere\nWorld"), []string{"Hi", "There", "World"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\r\nThere\r\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hi\r\nThere\r\nWorld"), []string{"Hi", "There", "World"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\r\nThere\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.SplitStringByNewLine("Hi\r\nThere\nWorld"), []string{"Hi", "There", "World"})
	}, t)
}
