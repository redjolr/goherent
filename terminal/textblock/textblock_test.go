package textblock_test

import (
	"testing"

	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal/textblock"
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
	Given that there is an Textblock with line "SomeLine"
	And the cursor is at 0,0
	When we Write the "X" character
	Then the textbox should have one line []string{"XY"}
	`, func(t *testing.T) {
		textBlock := textblock.FromString("SomeLine")
		textBlock.MoveCursorToOrigin()

		textBlock.Write("X")
		lines := textBlock.Lines()
		assert.Equal(lines[0], "XomeLine")
	}, t)

	// Test(`
	// Given that there is an Textblock with line "SomeLine"
	// And the cursor is at 0,7
	// When we Write the "X" character
	// Then the textbox should have one line []string{"XY"}
	// `, func(t *testing.T) {
	// 	textBlock := textblock.FromString("SomeLine")
	// 	textBlock.MoveCursorToOrigin()

	// 	textBlock.Write("X")
	// 	lines := textBlock.Lines()
	// 	assert.Equal(lines[0], "XomeLine")
	// }, t)

	Test(`
	Given that there is an Textblock with line "SomeLine"
	And the cursor is at the end of the line
	When we Write the "X" character
	Then the textbox should have one line []string{"SomeLineX"}
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
	Then the textbox should have one line []string{"SomeLineXY"}
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
	Then an error should not be thrown
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		err := textblock.MoveCursorTo(0, 0)
		assert.NoError(err)
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0,1
	Then an error should be thrown
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		err := textblock.MoveCursorTo(0, 1)
		assert.Error(err)
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0,2
	Then an error should be thrown
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		err := textblock.MoveCursorTo(0, 1)
		assert.Error(err)
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to 0, -1
	Then an error should be thrown
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		err := textblock.MoveCursorTo(0, -1)
		assert.Error(err)
	}, t)

	Test(`
	Given that there is an empty Textblock
	When we move the cursor to -1, 0
	Then an error should be thrown
	`, func(t *testing.T) {
		textblock := textblock.EmptyTextblock()
		err := textblock.MoveCursorTo(-1, 0)
		assert.Error(err)
	}, t)
}
