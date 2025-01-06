package internal_test

import (
	"testing"

	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal/internal"
	. "github.com/redjolr/goherent/test"
)

func TestTextPop(t *testing.T) {
	Test(`
	Given that there is an empty string text
	When we PopLeft the next sequence
	Then it should return an empty string
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText("")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("", true))
		Expect(text.Value()).ToEqual("")

	}, t)

	Test(`
	Given that there is a text "A"
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText("A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text "🚀"
	When we PopLeft the next sequence
	Then it should return "🚀"
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText("🚀")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("🚀", true))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text "AB"
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be "B".`, func(Expect expect.F) {
		text := internal.NewText("AB")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual("B")
	}, t)

	Test(`
	Given that there is a text "A🚀"
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should "🚀".`, func(Expect expect.F) {
		text := internal.NewText("A🚀")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual("🚀")
	}, t)

	Test(`
	Given that there is a text "🚀A"
	When we PopLeft the next sequence
	Then it should return "🚀"
	And the remaining text should "A".`, func(Expect expect.F) {
		text := internal.NewText("🚀A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("🚀", true))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "🚀✅"
	When we PopLeft the next sequence
	Then it should return "🚀"
	And the remaining text should "✅".`, func(Expect expect.F) {
		text := internal.NewText("🚀✅")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("🚀", true))
		Expect(text.Value()).ToEqual("✅")
	}, t)

	Test(`
	Given that there is a text "\nA"
	When we PopLeft the next sequence
	Then it should return "\n"
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText("\nA")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("\n", true))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A\n"
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be "".`, func(Expect expect.F) {
		text := internal.NewText("A\n")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual("\n")
	}, t)

	Test(`
	Given that there is a text ansi_escape.CURSOR_TO_HOME
	When we PopLeft the next sequence
	Then it should return ansi_escape.CURSOR_TO_HOME
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.CURSOR_TO_HOME)
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.CURSOR_TO_HOME, false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.CURSOR_TO_HOME+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.CURSOR_TO_HOME
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.CURSOR_TO_HOME + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.CURSOR_TO_HOME, false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text ansi_escape.ERASE_SCREEN
	When we PopLeft the next sequence
	Then it should return ansi_escape.ERASE_SCREEN
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.ERASE_SCREEN)
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.ERASE_SCREEN, false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.ERASE_SCREEN+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.ERASE_SCREEN
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.ERASE_SCREEN + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.ERASE_SCREEN, false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.ERASE_SCREEN
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.ERASE_SCREEN.`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.ERASE_SCREEN)
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.ERASE_SCREEN)
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorLeftNCols(1)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorLeftNCols(1)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorLeftNCols(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorLeftNCols(1), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorLeftNCols(1)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorLeftNCols(1)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorLeftNCols(1) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorLeftNCols(1), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorLeftNCols(1)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorLeftNCols(1).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorLeftNCols(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorLeftNCols(1))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorLeftNCols(10)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorLeftNCols(10)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorLeftNCols(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorLeftNCols(10), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorLeftNCols(10)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorLeftNCols(10)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorLeftNCols(10) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorLeftNCols(10), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorLeftNCols(10)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorLeftNCols(10).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorLeftNCols(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorLeftNCols(10))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorRightNCols(1)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorRightNCols(1)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorRightNCols(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorRightNCols(1), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorRightNCols(1)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorRightNCols(1)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorRightNCols(1) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorRightNCols(1), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorRightNCols(1)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorRightNCols(1).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorRightNCols(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorRightNCols(1))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorRightNCols(10)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorRightNCols(10)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorRightNCols(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorRightNCols(10), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorRightNCols(10)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorRightNCols(10)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorRightNCols(10) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorRightNCols(10), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorRightNCols(10)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorRightNCols(10).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorRightNCols(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorRightNCols(10))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorUpNRows(1)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorUpNRows(1)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorUpNRows(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorUpNRows(1), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorUpNRows(1)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorUpNRows(1)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorUpNRows(1) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorUpNRows(1), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorUpNRows(1)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorUpNRows(1).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorUpNRows(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorUpNRows(1))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorUpNRows(10)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorUpNRows(10)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorUpNRows(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorUpNRows(10), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorUpNRows(10)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorUpNRows(10)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorUpNRows(10) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorUpNRows(10), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorUpNRows(10)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorUpNRows(10).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorUpNRows(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorUpNRows(10))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorDownNRows(1)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorDownNRows(1)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorDownNRows(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorDownNRows(1), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorDownNRows(1)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorDownNRows(1)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorDownNRows(1) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorDownNRows(1), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorDownNRows(1)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorDownNRows(1).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorDownNRows(1))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorDownNRows(1))
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorDownNRows(10)
	When we PopLeft the next sequence
	Then it should return the escape code of ansi_escape.MoveCursorDownNRows(10)
	And the remaining text should be an empty string.`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorDownNRows(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorDownNRows(10), false))
		Expect(text.Value()).ToEqual("")
	}, t)

	Test(`
	Given that there is a text ansi_escape.MoveCursorDownNRows(10)+"A"
	When we PopLeft the next sequence
	Then it should return ansi_escape.MoveCursorDownNRows(10)
	And the remaining text should be "A".`, func(Expect expect.F) {
		text := internal.NewText(ansi_escape.MoveCursorDownNRows(10) + "A")
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence(ansi_escape.MoveCursorDownNRows(10), false))
		Expect(text.Value()).ToEqual("A")
	}, t)

	Test(`
	Given that there is a text "A"+ansi_escape.MoveCursorDownNRows(10)
	When we PopLeft the next sequence
	Then it should return "A"
	And the remaining text should be ansi_escape.MoveCursorDownNRows(10).`, func(Expect expect.F) {
		text := internal.NewText("A" + ansi_escape.MoveCursorDownNRows(10))
		popped := text.PopLeft()
		Expect(popped).ToEqual(internal.NewSequence("A", true))
		Expect(text.Value()).ToEqual(ansi_escape.MoveCursorDownNRows(10))
	}, t)
}
