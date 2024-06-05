package textblock_test

import (
	"testing"

	"github.com/redjolr/goherent/console/textblock"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewTextblockFromString(t *testing.T) {
	assert := assert.New(t)
	Test(`
	it should store the line []string{""},
	if we pass an empty string`, func(t *testing.T) {
		textBlock := textblock.FromString("")
		assert.Equal(textBlock.Lines(), []string{""})
	}, t)

	Test(`
	it should store the line []string{"A"},
	if we pass the string "A"`, func(t *testing.T) {
		textBlock := textblock.FromString("A")
		assert.Equal(textBlock.Lines(), []string{"A"})
	}, t)

	Test(`
	it should store two empty lines,
	if we pass the string "\n"`, func(t *testing.T) {
		textBlock := textblock.FromString("\n")
		assert.Equal(textBlock.Lines(), []string{"", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := textblock.FromString("\n\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store two empty lines,
	if we pass the string "\r\n"`, func(t *testing.T) {
		textBlock := textblock.FromString("\r\n")
		assert.Equal(textBlock.Lines(), []string{"", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\r\n\r\n"`, func(t *testing.T) {
		textBlock := textblock.FromString("\r\n\r\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store 3 empty lines,
	if we pass the string "\n\n"`, func(t *testing.T) {
		textBlock := textblock.FromString("\n\n")
		assert.Equal(textBlock.Lines(), []string{"", "", ""})
	}, t)

	Test(`
	it should store the lines []string{"This is line 1", ""},
	if we pass the string "This is line 1\n" to NewTextBlock()`, func(t *testing.T) {
		textBlock := textblock.FromString("This is line 1\n")
		assert.Equal(textBlock.Lines(), []string{"This is line 1", ""})
	}, t)

	Test(`
	it should store the lines []string{"This is line 1", "This is line 2"},
	if we pass the string "This is line 1\nThis is line 2" to NewTextBlock()`, func(t *testing.T) {
		textBlock := textblock.FromString("This is line 1\nThis is line 2")
		assert.Equal(textBlock.Lines(), []string{"This is line 1", "This is line 2"})
	}, t)
}

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that there is an empty Textblock
	When we print the "X" character
	Then the first element in the first row will be "X"
	`, func(t *testing.T) {
		textBlock := textblock.EmptyTextblock()
		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "X")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "X" character
	Then the textblock should have one line []string{"XY"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "XomeLine")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "7Letter" string
	Then the textblock should have one line []string{"8Letters"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("7Letter")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "7Lettere")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "8Letters" string
	Then the textblock should have one line []string{"8Letters"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("8Letters")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "8Letters")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "9 Letters" string
	Then the textblock should have one line []string{"9 Letters"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("9 Letters")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "9 Letters")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "THIS_IS_A_VERY_LONG_STRING" string
	Then the textblock should have one line []string{"THIS_IS_A_VERY_LONG_STRING"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("THIS_IS_A_VERY_LONG_STRING")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "THIS_IS_A_VERY_LONG_STRING")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 1,0
	When we Write the "THIS_IS_A_VERY_LONG_STRING" string
	Then the textblock should have one line []string{"THIS_IS_A_VERY_LONG_STRING"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorTo(1, 0)

		textBlock.Write("THIS_IS_A_VERY_LONG_STRING")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "STHIS_IS_A_VERY_LONG_STRING")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 7,0
	When we Write the "THIS_IS_A_VERY_LONG_STRING" string
	Then the textblock should have one line []string{"THIS_IS_A_VERY_LONG_STRING"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorTo(7, 0)

		textBlock.Write("THIS_IS_A_VERY_LONG_STRING")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLinTHIS_IS_A_VERY_LONG_STRING")
	}, t)

	Test(`
	Given that there is a Textblock with line "SomeLine"
	And the cursor is at 8,0
	When we Write the "THIS_IS_A_VERY_LONG_STRING" string
	Then the textblock should have one line []string{"THIS_IS_A_VERY_LONG_STRING"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorTo(8, 0)
		textBlock.Write("THIS_IS_A_VERY_LONG_STRING")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLineTHIS_IS_A_VERY_LONG_STRING")
	}, t)

	Test(`
	Given that there is an Textblock with line "SomeLine"
	And the cursor is at 7,0
	When we Write the "X" character
	Then the textblock should have one line []string{"SomeLinX"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorTo(7, 0)

		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLinX")
	}, t)

	Test(`
	Given that there is an Textblock with line "SomeLine"
	And the cursor is at 8,0
	When we Write the "X" character
	Then the textblock should have one line []string{"SomeLineX"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorTo(8, 0)

		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLineX")
	}, t)

	Test(`
	Given that there is an Textblock with line "SomeLine"
	And the cursor is at the end of the line
	When we Write the "X" character
	Then the textblock should have one line []string{"SomeLineX"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLineX")
	}, t)

	Test(`
	Given that there is an Textblock with line "SomeLine"
	And the Write("X") method is called beforehand
	When the Write("Y") method is called 
	Then the textblock should have one line []string{"SomeLineXY"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.Write("X")
		textBlock.Write("Y")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "SomeLineXY")
	}, t)
}

func TestMoveCursorTo(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0,0
	Then it does NOT panic
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.NotPanics(func() {
			textblock.MoveCursorTo(0, 0)
		})
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0,1
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.Panics(func() {
			textblock.MoveCursorTo(0, 1)
		})
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0,2
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.Panics(func() {
			textblock.MoveCursorTo(0, 2)
		})
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 1,0
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.Panics(func() {
			textblock.MoveCursorTo(1, 0)
		})
	}, t)

	Test(`
	Given that there is a Textblock with one line: "X"
	When we move the cursor to 1,0
	Then it does NOT panic
	`, func(t *testing.T) {
		textblock := textblock.FromString("X")
		assert.NotPanics(func() {
			textblock.MoveCursorTo(1, 0)
		})
	}, t)

	Test(`
	Given that there is a Textblock with one line: "X"
	When we move the cursor to 2,0
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.FromString("X")
		assert.Panics(func() {
			textblock.MoveCursorTo(2, 0)
		})
	}, t)

	Test(`
	Given that there is a Textblock with one line: "X"
	When we move the cursor to 3,0
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.FromString("X")
		assert.Panics(func() {
			textblock.MoveCursorTo(3, 0)
		})
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0, -1
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.Panics(func() {
			textblock.MoveCursorTo(0, -1)
		})
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to -1, 0
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		assert.Panics(func() {
			textblock.MoveCursorTo(-1, 0)
		})
	}, t)

	Test(`
	Given that there is a Textblock from string: "SomeLine1\nSomeLine2"
	When we move the cursor to 8, 1
	Then it does NOT panic
	`, func(t *testing.T) {
		textblock := textblock.FromString("SomeLine1\nSomeLine2")
		assert.NotPanics(func() {
			textblock.MoveCursorTo(8, 1)
		})
	}, t)

	Test(`
	Given that there is a Textblock from string: "SomeLine1\nSomeLine2"
	When we move the cursor to 9, 1
	Then it does NOT panic
	`, func(t *testing.T) {
		textblock := textblock.FromString("SomeLine1\nSomeLine2")
		assert.NotPanics(func() {
			textblock.MoveCursorTo(9, 1)
		})
	}, t)

	Test(`
	Given that there is a Textblock from string: "SomeLine1\nSomeLine2"
	When we move the cursor to 10, 1
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.FromString("SomeLine1\nSomeLine2")
		assert.Panics(func() {
			textblock.MoveCursorTo(10, 1)
		})
	}, t)

	Test(`
	Given that there is a Textblock from string: "SomeLine1\nSomeLine2"
	When we move the cursor to 9, 2
	Then it panics
	`, func(t *testing.T) {
		textblock := textblock.FromString("SomeLine1\nSomeLine2")
		assert.Panics(func() {
			textblock.MoveCursorTo(9, 2)
		})
	}, t)
}
