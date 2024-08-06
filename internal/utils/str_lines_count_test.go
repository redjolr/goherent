package utils_test

import (
	"testing"

	"github.com/redjolr/goherent/internal/utils"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestStrLinesCount(t *testing.T) {
	assert := assert.New(t)
	Test("it should return 1, if you pass an empty string.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount(""), 1)
	}, t)

	Test("it should return 1, if you pass 'H'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("H"), 1)
	}, t)

	Test("it should return 1, if you pass 'Hello'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hello"), 1)
	}, t)

	Test("it should return 2, if you pass 'Hi\nThere'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hi\nThere"), 2)
	}, t)

	Test("it should return 2, if you pass 'Hi\r\nThere'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hi\r\nThere"), 2)
	}, t)

	Test("it should return 3, if you pass 'Hi\nThere\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hi\nThere\nWorld"), 3)
	}, t)

	Test("it should return 3, if you pass 'Hi\r\nThere\r\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hi\r\nThere\r\nWorld"), 3)
	}, t)

	Test("it should return 3, if you pass 'Hi\r\nThere\nWorld'.", func(t *testing.T) {
		assert.Equal(utils.StrLinesCount("Hi\r\nThere\nWorld"), 3)
	}, t)
}
