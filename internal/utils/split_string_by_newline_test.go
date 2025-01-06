package utils_test

import (
	"testing"

	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/internal/utils"
	. "github.com/redjolr/goherent/test"
)

func TestSplitStringsByNewLine(t *testing.T) {
	Test("it should return [''], if you pass an empty string.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("")).ToEqual([]string{""})
	}, t)

	Test("it should return ['H'], if you pass 'H'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("H")).ToEqual([]string{"H"})
	}, t)

	Test("it should return ['Hello'], if you pass 'Hello'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hello")).ToEqual([]string{"Hello"})
	}, t)

	Test("it should return ['Hi', 'There'], if you pass 'Hi\nThere'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hi\nThere")).ToEqual([]string{"Hi", "There"})
	}, t)

	Test("it should return ['Hi', 'There'], if you pass 'Hi\r\nThere'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hi\r\nThere")).ToEqual([]string{"Hi", "There"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\nThere\nWorld'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hi\nThere\nWorld")).ToEqual([]string{"Hi", "There", "World"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\r\nThere\r\nWorld'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hi\r\nThere\r\nWorld")).ToEqual([]string{"Hi", "There", "World"})
	}, t)

	Test("it should return ['Hi', 'There', 'World'], if you pass 'Hi\r\nThere\nWorld'.", func(Expect expect.F) {
		Expect(utils.SplitStringByNewLine("Hi\r\nThere\nWorld")).ToEqual([]string{"Hi", "There", "World"})
	}, t)
}
