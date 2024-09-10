package utils_test

import (
	"testing"

	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/internal/utils"
	. "github.com/redjolr/goherent/pkg"
)

func TestStrLinesCount(t *testing.T) {
	Test("it should return 1, if you pass an empty string.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("")).ToEqual(1)
	}, t)

	Test("it should return 1, if you pass 'H'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("H")).ToEqual(1)
	}, t)

	Test("it should return 1, if you pass 'Hello'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hello")).ToEqual(1)
	}, t)

	Test("it should return 2, if you pass 'Hi\nThere'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hi\nThere")).ToEqual(2)
	}, t)

	Test("it should return 2, if you pass 'Hi\r\nThere'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hi\r\nThere")).ToEqual(2)
	}, t)

	Test("it should return 3, if you pass 'Hi\nThere\nWorld'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hi\nThere\nWorld")).ToEqual(3)
	}, t)

	Test("it should return 3, if you pass 'Hi\r\nThere\r\nWorld'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hi\r\nThere\r\nWorld")).ToEqual(3)
	}, t)

	Test("it should return 3, if you pass 'Hi\r\nThere\nWorld'.", func(Expect expect.F) {
		Expect(utils.StrLinesCount("Hi\r\nThere\nWorld")).ToEqual(3)
	}, t)
}
