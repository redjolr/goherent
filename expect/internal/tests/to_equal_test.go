package tests_test

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
	. "github.com/redjolr/goherent/pkg"
)

func TestToEqualIntegers(t *testing.T) {
	Test("it should not return an error if we compare 1 and 1.", func(t *testing.T) {
		oneEqualsOneError := assertions.ToEqual(1, 1)
		if oneEqualsOneError != nil {
			t.Error("One equals one, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare 0 and 0.", func(t *testing.T) {
		zeroEqualsZeroError := assertions.ToEqual(0, 0)
		if zeroEqualsZeroError != nil {
			t.Error("Zero equals zero, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare -1 and -1.", func(t *testing.T) {
		neg1EqualsNeg1Error := assertions.ToEqual(-1, -1)
		if neg1EqualsNeg1Error != nil {
			t.Error("-1 equals -1, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare -1 and 0.", func(t *testing.T) {
		zeroEqualsZeroError := assertions.ToEqual(0, 0)
		if zeroEqualsZeroError != nil {
			t.Error("Zero equals zero, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test(`
	Given that we compare 2 and 3
	When we use the ToEqual assertions
	Then it should return an error with this text: "expected: 2\nactual:  3",`, func(t *testing.T) {
		twoEqualsTwoError := assertions.ToEqual(2, 3)
		expectedErrorMsg := "Not equal:\nexpected: 2\nactual  : 3"

		if twoEqualsTwoError == nil {
			t.Error("2 does not equal 3, but ToEqual() assertion says they do.")
		}
		if twoEqualsTwoError.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, twoEqualsTwoError.Error())
		}
	}, t)

	Test(`
	Given that we compare int8(2) and int16(2)
	When we use the ToEqual assertions
	Then it should return an error with this text: 
		"Not equal:\nexpected: int8(2)\nactual  : int16(2)".`, func(t *testing.T) {
		var twoTypeInt8 int8 = 2
		var twoTypeInt16 int16 = 2
		twoEqualsTwoError := assertions.ToEqual(twoTypeInt8, twoTypeInt16)
		expectedErrorMsg := "Not equal:\nexpected: int8(2)\nactual  : int16(2)"
		if twoEqualsTwoError == nil {
			t.Error("int8(2) does not equal int16(2), but ToEqual() assertion says they do.")
		}
		if twoEqualsTwoError.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, twoEqualsTwoError.Error())
		}
	}, t)

	Test(`
	Given that we compare int32(2) and int64(2)
	When we use the ToEqual assertions
	Then it should return an error with this text: 
		"Not equal:\nexpected: int8(2)\nactual  : int16(2)".`, func(t *testing.T) {
		var twoTypeInt32 int32 = 2
		var twoTypeInt64 int64 = 2
		twoEqualsTwoError := assertions.ToEqual(twoTypeInt32, twoTypeInt64)
		expectedErrorMsg := "Not equal:\nexpected: int32(2)\nactual  : int64(2)"
		if twoEqualsTwoError == nil {
			t.Error("int32(2) does not equal int64(2), but ToEqual() assertion says they do.")
		}
		if twoEqualsTwoError.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, twoEqualsTwoError.Error())
		}
	}, t)

	Test(`
	Given that we compare uint8(2) and int8(2)
	When we use the ToEqual assertions
	Then it should return an error with this text: 
		"Not equal:\nexpected: uint8(0x2)\nactual  : int8(2)".`, func(t *testing.T) {
		var twoTypeInt32 uint8 = 2
		var twoTypeInt64 int8 = 2
		twoEqualsTwoError := assertions.ToEqual(twoTypeInt32, twoTypeInt64)
		expectedErrorMsg := "Not equal:\nexpected: uint8(0x2)\nactual  : int8(2)"
		if twoEqualsTwoError == nil {
			t.Error("uint8(0x2) does not equal int8(2), but ToEqual() assertion says they do.")
		}
		if twoEqualsTwoError.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, twoEqualsTwoError.Error())
		}
	}, t)
}

func TestToEqualFloatingPointNumbers(t *testing.T) {
	Test("it should not return an error if we compare float64(3.25) and float64(3.25).", func(t *testing.T) {
		floatsEqualsErr := assertions.ToEqual(3.25, 3.25)
		if floatsEqualsErr != nil {
			t.Error("One equals one, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare float32(3.25) and float32(3.25).", func(t *testing.T) {
		var firstFloat64 float32 = 3.25
		var secondFloat64 float32 = 3.25

		floatsEqualsErr := assertions.ToEqual(firstFloat64, secondFloat64)
		if floatsEqualsErr != nil {
			t.Error("One equals one, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare float64(3.25) and float64(3.250).", func(t *testing.T) {
		floatsEqualsErr := assertions.ToEqual(3.25, 3.250)
		if floatsEqualsErr != nil {
			t.Error("One equals one, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test(`
	Given that we compare float32(3.25) and float64(3.25)
	When we use the ToEqual assertions
	Then it should return an error with this text: 
		"Not equal:\nexpected: float32(3.25)\nactual  : float64(3.25)".`, func(t *testing.T) {
		var firstFloat32 float32 = 3.25
		var secondFloat64 float64 = 3.25

		floatsEqualsErr := assertions.ToEqual(firstFloat32, secondFloat64)
		expectedErrorMsg := "Not equal:\nexpected: float32(3.25)\nactual  : float64(3.25)"
		if floatsEqualsErr == nil {
			t.Error("float32(3.25) does not equal float64(3.25), but ToEqual() assertion says they do.")
		}
		if floatsEqualsErr.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, floatsEqualsErr.Error())
		}
	}, t)

	Test(`
	Given that we compare float32(7) and int(7)
	When we use the ToEqual assertions
	Then it should return an error with this text: 
		"Not equal:\nexpected: float32(7)\nactual  : int(7)".`, func(t *testing.T) {
		var sevenFloat32 float32 = 7
		var sevenInt int = 7

		floatsEqualsErr := assertions.ToEqual(sevenFloat32, sevenInt)
		expectedErrorMsg := "Not equal:\nexpected: float32(7)\nactual  : int(7)"
		if floatsEqualsErr == nil {
			t.Error("float32(7) does not equal int(7), but ToEqual() assertion says they do.")
		}
		if floatsEqualsErr.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, floatsEqualsErr.Error())
		}
	}, t)
}

func TestToEqualBooleans(t *testing.T) {
	Test("it should not return an error if we compare true and true.", func(t *testing.T) {
		boolsEqualErr := assertions.ToEqual(true, true)
		if boolsEqualErr != nil {
			t.Error("true equals true, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test("it should not return an error if we compare false and false.", func(t *testing.T) {
		boolsEqualErr := assertions.ToEqual(false, false)
		if boolsEqualErr != nil {
			t.Error("false equals false, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test(`
	Given that we compare true and false
	When we use the ToEqual assertions
	Then it should return an error with this text: "Not equal:\nexpected: true\nactual  : false".`, func(t *testing.T) {
		boolsEqualErr := assertions.ToEqual(true, false)
		expectedErrorMsg := "Not equal:\nexpected: true\nactual  : false"

		if boolsEqualErr == nil {
			t.Error("true does not equal false, but ToEqual() assertion says they do.")
		}
		if boolsEqualErr.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, boolsEqualErr.Error())
		}
	}, t)

	Test(`
	Given that we compare false and true
	When we use the ToEqual assertions
	Then it should return an error with this text: "Not equal:\nexpected: false\nactual  : true".`, func(t *testing.T) {
		boolsEqualErr := assertions.ToEqual(false, true)
		expectedErrorMsg := "Not equal:\nexpected: false\nactual  : true"

		if boolsEqualErr == nil {
			t.Error("false does not equal true, but ToEqual() assertion says they do.")
		}
		if boolsEqualErr.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s.\n\nIs:\n%s",
				expectedErrorMsg, boolsEqualErr.Error())
		}
	}, t)
}

func TestToEqualStrings(t *testing.T) {
	Test(`it should not return an error if we compare "A" and "A".`, func(t *testing.T) {
		stringsEqualErr := assertions.ToEqual("A", "A")
		if stringsEqualErr != nil {
			t.Error(`"A" equals "A", but ToEqual() assertion says they don't.`)
		}
	}, t)

	Test(`it should not return an error if we compare "ABC" and "ABC".`, func(t *testing.T) {
		stringsEqualErr := assertions.ToEqual("ABC", "ABC")
		if stringsEqualErr != nil {
			t.Error(`"ABC" equals "ABC", but ToEqual() assertion says they don't.`)
		}
	}, t)

	Test(`it should not return an error if we compare "" and "".`, func(t *testing.T) {
		stringsEqualErr := assertions.ToEqual("", "")
		if stringsEqualErr != nil {
			t.Error("Empty string equals empty, but ToEqual() assertion says they don't.")
		}
	}, t)

	Test(`
	Given that we compare "D" and "E"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("D", "E")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"D\"\n" +
			"actual  : \"E\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1 +1 @@\n" +
			"-D\n" +
			"+E\n"

		if err == nil {
			t.Error("A does not equal B, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello" and "Hallo"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello", "Hallo")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\"\n" +
			"actual  : \"Hallo\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1 +1 @@\n" +
			"-Hello\n" +
			"+Hallo\n"

		if err == nil {
			t.Error("Hello does not equal Hallo, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello\nWorld" and "Hallo\World"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello\nWorld", "Hallo\nWorld")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\\nWorld\"\n" +
			"actual  : \"Hallo\\nWorld\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,2 +1,2 @@\n" +
			"-Hello\n" +
			"+Hallo\n" +
			" World\n"

		if err == nil {
			t.Error("Hello\\nWorld does not equal Hallo\\nWorld, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello\nWorld" and "Hallo\nWelt"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello\nWorld", "Hallo\nWelt")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\\nWorld\"\n" +
			"actual  : \"Hallo\\nWelt\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,2 +1,2 @@\n" +
			"-Hello\n" +
			"-World\n" +
			"+Hallo\n" +
			"+Welt\n"

		if err == nil {
			t.Error("Hello\\nWorld does not equal Hallo\\nWelt, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello\nThere\nWorld" and "Hallo\nThere\nWorld"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello\nThere\nWorld", "Hallo\nThere\nWorld")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\\nThere\\nWorld\"\n" +
			"actual  : \"Hallo\\nThere\\nWorld\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,2 +1,2 @@\n" +
			"-Hello\n" +
			"+Hallo\n" +
			" There\n"

		if err == nil {
			t.Error("Hello\\nThere\\nWorld does not equal Hallo\\nThere\\nWorld, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello\nThere\nWorld" and "Hallo\nThere\nWelt"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello\nThere\nWorld", "Hallo\nThere\nWelt")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\\nThere\\nWorld\"\n" +
			"actual  : \"Hallo\\nThere\\nWelt\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			"-Hello\n" +
			"+Hallo\n" +
			" There\n" +
			"-World\n" +
			"+Welt\n"

		if err == nil {
			t.Error("Hello\\nThere\\nWorld does not equal Hallo\\nThere\\nWelt, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Hello\nThere\nWorld" and "Hallo\nDort\nWelt"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("Hello\nThere\nWorld", "Hallo\nDort\nWelt")
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Hello\\nThere\\nWorld\"\n" +
			"actual  : \"Hallo\\nDort\\nWelt\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			"-Hello\n" +
			"-There\n" +
			"-World\n" +
			"+Hallo\n" +
			"+Dort\n" +
			"+Welt\n"

		if err == nil {
			t.Error("Hello\\nThere\\nWorld does not equal Hallo\\nDort\\nWelt, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "Line1\nLine2\nLine3\nLine4\nLine5\nLine6\nLine7" and 
	"Rreshti1\nRreshti2\nRreshti3\nRreshti4\nRreshti5\nRreshti6\nRreshti7"
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual(
			"Line1\nLine2\nLine3\nLine4\nLine5\nLine6\nLine7",
			"Rreshti1\nRreshti2\nRreshti3\nRreshti4\nRreshti5\nRreshti6\nRreshti7",
		)
		expectedErrorMsg := "Not equal:\n" +
			"expected: \"Line1\\nLine2\\nLine3\\nLine4\\nLine5\\nLine6\\nLine7\"\n" +
			"actual  : \"Rreshti1\\nRreshti2\\nRreshti3\\nRreshti4\\nRreshti5\\nRreshti6\\nRreshti7\"\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,7 +1,7 @@\n" +
			"-Line1\n" +
			"-Line2\n" +
			"-Line3\n" +
			"-Line4\n" +
			"-Line5\n" +
			"-Line6\n" +
			"-Line7\n" +
			"+Rreshti1\n" +
			"+Rreshti2\n" +
			"+Rreshti3\n" +
			"+Rreshti4\n" +
			"+Rreshti5\n" +
			"+Rreshti6\n" +
			"+Rreshti7\n"

		if err == nil {
			t.Error("Expected value does not equal actual value, but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "2" and int(2)
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		err := assertions.ToEqual("2", 2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: string(\"2\")\n" +
			"actual  : int(2)"

		if err == nil {
			t.Error("string(\"2\") does not equal int(2), but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that we compare "2" and float64(2)
	When we use the ToEqual assertions
	Then it should return an error that reports that the values are not equal
	And the Diff should be part of the error.`, func(t *testing.T) {
		var twoFloat float64 = 2
		err := assertions.ToEqual("2", twoFloat)
		expectedErrorMsg := "Not equal:\n" +
			"expected: string(\"2\")\n" +
			"actual  : float64(2)"

		if err == nil {
			t.Error("string(\"2\") does not equal int(2), but ToEqual() assertion says they do.")
		}

		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)
}

func TestToEqualArrays(t *testing.T) {
	Test("it should not return an error if we compare two integer arrays of size 0.", func(t *testing.T) {
		arr1 := [0]int{}
		arr2 := [0]int{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[0]int{} equals [0]int{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test("it should not return an error if we compare two integer arrays of size 1 with their 0 values.", func(t *testing.T) {
		arr1 := [1]int{}
		arr2 := [1]int{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]int{} equals [1]int{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two integer arrays of size 1 with an element equal to 2.`, func(t *testing.T) {
		arr1 := [1]int{2}
		arr2 := [1]int{2}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]int{2} equals [1]int{2}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two integer arrays of size 4 with equal elements.`, func(t *testing.T) {
		arr1 := [4]int{4, 6, -2, 3}
		arr2 := [4]int{4, 6, -2, 3}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("integer arrays are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test("it should not return an error if we compare two boolean arrays of size 0.", func(t *testing.T) {
		arr1 := [0]bool{}
		arr2 := [0]bool{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[0]bool{} equals [0]bool{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test("it should not return an error if we compare two bool arrays of size 1 with their 0 values.", func(t *testing.T) {
		arr1 := [1]bool{}
		arr2 := [1]bool{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]bool{} equals [1]bool{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two bool arrays of size 1 with an element equal to false.`, func(t *testing.T) {
		arr1 := [1]bool{false}
		arr2 := [1]bool{false}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]bool{false} equals [1]bool{false}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two bool arrays of size 4 with equal elements.`, func(t *testing.T) {
		arr1 := [4]bool{false, true, true, false}
		arr2 := [4]bool{false, true, true, false}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("boolean arrays are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test("it should not return an error if we compare two float64 arrays of size 0.", func(t *testing.T) {
		arr1 := [0]float64{}
		arr2 := [0]float64{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[0]float64{} equals [0]float64{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test("it should not return an error if we compare two float64 arrays of size 1 with their 0 values.", func(t *testing.T) {
		arr1 := [1]float64{}
		arr2 := [1]float64{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]float64{} equals [1]float64{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two float64 arrays of size 1 with an element equal to 2.`, func(t *testing.T) {
		arr1 := [1]float64{2.3}
		arr2 := [1]float64{2.3}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]float64{2} equals [1]float64{2}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two float64 arrays of size 4 with equal elements.`, func(t *testing.T) {
		arr1 := [4]float64{4.5, 6.3, -11.2, 2.009}
		arr2 := [4]float64{4.5, 6.3, -11.2, 2.009}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]float64{2} equals [1]float64{2}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test("it should not return an error if we compare two string arrays of size 0.", func(t *testing.T) {
		arr1 := [0]string{}
		arr2 := [0]string{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[0]string{} equals [0]string{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test("it should not return an error if we compare two string arrays of size 1 with their 0 values.", func(t *testing.T) {
		arr1 := [1]string{}
		arr2 := [1]string{}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]string{} equals [1]string{}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two string arrays of size 1 with an element equal to 2.`, func(t *testing.T) {
		arr1 := [1]string{"A"}
		arr2 := [1]string{"A"}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("[1]string{2} equals [1]string{2}, but ToEqual() assertion says they do not.")
		}
	}, t)

	Test(`it should not return an error,
	if we compare two string arrays of size 4 with equal elements.`, func(t *testing.T) {
		arr1 := [4]string{"Lorem", "Ipsum", "Dolor", "Sit"}
		arr2 := [4]string{"Lorem", "Ipsum", "Dolor", "Sit"}

		err := assertions.ToEqual(arr1, arr2)
		if err != nil {
			t.Error("string arrays are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error, if we compare two arrays of size 0 and 
	one of them is of type [0]int and the other of type [0]uint32`, func(t *testing.T) {
		arr1 := [0]int{}
		arr2 := [0]uint{}
		err := assertions.ToEqual(arr1, arr2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: [0]int([0]int{})\n" +
			"actual  : [0]uint([0]uint{})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}

	}, t)

	Test(`it should return an error, if we compare these two arrays: [1]int32{7}, [1]int64{7} `, func(t *testing.T) {
		arr1 := [1]int32{7}
		arr2 := [1]int64{7}
		err := assertions.ToEqual(arr1, arr2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: [1]int32([1]int32{7})\n" +
			"actual  : [1]int64([1]int64{7})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error, if we compare these two arrays:
	[2]float32{7.4, 3.2}, [2]float64{7.4, 3.2} `, func(t *testing.T) {
		arr1 := [2]float32{7.4, 3.2}
		arr2 := [2]float64{7.4, 3.2}
		err := assertions.ToEqual(arr1, arr2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: [2]float32([2]float32{7.4, 3.2})\n" +
			"actual  : [2]float64([2]float64{7.4, 3.2})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)
}

func TestToEqualSlices(t *testing.T) {
	Test("it should not return an error if we compare two empty integer slices.", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{}
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two integer slices,
	where both have length 0 and capacity 0.`, func(t *testing.T) {
		slice1 := make([]int, 0, 0)
		slice2 := make([]int, 0, 0)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two integer slices,
	where both have length 0 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]int, 0, 1)
		slice2 := make([]int, 0, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two integer slices,
	where both have length 1 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]int, 1, 1)
		slice2 := make([]int, 1, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two integer slices,
	where one has len=1,cap=1 and the other has len=1,cap=2.`, func(t *testing.T) {
		slice1 := make([]int, 1, 1)
		slice2 := make([]int, 1, 2)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two integer slices,
	where both have len=2,cap=3 and they have the same two elements.`, func(t *testing.T) {
		slice1 := make([]int, 2, 3)
		slice1[0] = 7
		slice1[1] = 9
		slice2 := make([]int, 2, 3)
		slice2[0] = 7
		slice2[1] = 9
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error if we compare two integer slices,
	where both have len=2,cap=3 and they have different int elements.`, func(t *testing.T) {
		slice1 := make([]int, 2, 3)
		slice1[0] = 7
		slice1[1] = 9
		slice2 := make([]int, 2, 3)
		slice2[0] = 7
		slice2[1] = 12
		expectedErrorMsg := "Not equal:\n" +
			"expected: []int{7, 9}\n" +
			"actual  : []int{7, 12}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -2,3 +2,3 @@\n" +
			"  (int) 7,\n" +
			"- (int) 9\n" +
			"+ (int) 12\n" +
			" }\n"
		err := assertions.ToEqual(slice1, slice2)
		if err == nil {
			t.Error("Slices should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test("it should not return an error if we compare two empty float64 slices.", func(t *testing.T) {
		slice1 := []float64{}
		slice2 := []float64{}
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two float64 slices,
	where both have length 0 and capacity 0.`, func(t *testing.T) {
		slice1 := make([]float64, 0, 0)
		slice2 := make([]float64, 0, 0)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two float64 slices,
	where both have length 0 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]float64, 0, 1)
		slice2 := make([]float64, 0, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two float64 slices,
	where both have length 1 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]float64, 1, 1)
		slice2 := make([]float64, 1, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two float64 slices,
	where one has len=1,cap=1 and the other has len=1,cap=2.`, func(t *testing.T) {
		slice1 := make([]float64, 1, 1)
		slice2 := make([]float64, 1, 2)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two float64 slices,
	where both have len=2,cap=3 and they have the same two elements.`, func(t *testing.T) {
		slice1 := make([]float64, 2, 3)
		slice1[0] = 7.502
		slice1[1] = 9.3111
		slice2 := make([]float64, 2, 3)
		slice2[0] = 7.502
		slice2[1] = 9.3111
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error if we compare two float64 slices,
	where both have len=2,cap=3 and they have different int elements.`, func(t *testing.T) {
		slice1 := make([]float64, 2, 3)
		slice1[0] = 7.502
		slice1[1] = 9.3111
		slice2 := make([]float64, 2, 3)
		slice2[0] = 7.333
		slice2[1] = 9.22
		expectedErrorMsg := "Not equal:\n" +
			"expected: []float64{7.502, 9.3111}\n" +
			"actual  : []float64{7.333, 9.22}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,4 +1,4 @@\n" +
			" ([]float64) (len=2) {\n" +
			"- (float64) 7.502,\n" +
			"- (float64) 9.3111\n" +
			"+ (float64) 7.333,\n" +
			"+ (float64) 9.22\n" +
			" }\n"
		err := assertions.ToEqual(slice1, slice2)
		if err == nil {
			t.Error("Slices should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test("it should not return an error if we compare two empty bool slices.", func(t *testing.T) {
		slice1 := []bool{}
		slice2 := []bool{}
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two bool slices,
	where both have length 0 and capacity 0.`, func(t *testing.T) {
		slice1 := make([]bool, 0, 0)
		slice2 := make([]bool, 0, 0)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two bool slices,
	where both have length 0 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]bool, 0, 1)
		slice2 := make([]bool, 0, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two bool slices,
	where both have length 1 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]bool, 1, 1)
		slice2 := make([]bool, 1, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two bool slices,
	where one has len=1,cap=1 and the other has len=1,cap=2.`, func(t *testing.T) {
		slice1 := make([]bool, 1, 1)
		slice2 := make([]bool, 1, 2)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two bool slices,
	where both have len=2,cap=3 and they have the same two elements.`, func(t *testing.T) {
		slice1 := make([]bool, 2, 3)
		slice1[0] = true
		slice1[1] = false
		slice2 := make([]bool, 2, 3)
		slice2[0] = true
		slice2[1] = false
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error if we compare two bool slices,
	where both have len=2,cap=3 and they have different int elements.`, func(t *testing.T) {
		slice1 := make([]bool, 2, 3)
		slice1[0] = true
		slice1[1] = false
		slice2 := make([]bool, 2, 3)
		slice2[0] = false
		slice2[1] = false
		err := assertions.ToEqual(slice1, slice2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: []bool{true, false}\n" +
			"actual  : []bool{false, false}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" ([]bool) (len=2) {\n" +
			"- (bool) true,\n" +
			"+ (bool) false,\n" +
			"  (bool) false\n"
		if err == nil {
			t.Error("Slices should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test("it should not return an error if we compare two empty string slices.", func(t *testing.T) {
		slice1 := []string{}
		slice2 := []string{}
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two string slices,
	where both have length 0 and capacity 0.`, func(t *testing.T) {
		slice1 := make([]string, 0, 0)
		slice2 := make([]string, 0, 0)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two string slices,
	where both have length 0 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]string, 0, 1)
		slice2 := make([]string, 0, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two string slices,
	where both have length 1 and capacity 1.`, func(t *testing.T) {
		slice1 := make([]string, 1, 1)
		slice2 := make([]string, 1, 1)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two string slices,
	where one has len=1,cap=1 and the other has len=1,cap=2.`, func(t *testing.T) {
		slice1 := make([]string, 1, 1)
		slice2 := make([]string, 1, 2)

		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two string slices,
	where both have len=2,cap=3 and they have the same two elements.`, func(t *testing.T) {
		slice1 := make([]string, 2, 3)
		slice1[0] = "Hello"
		slice1[1] = "World"
		slice2 := make([]string, 2, 3)
		slice2[0] = "Hello"
		slice2[1] = "World"
		err := assertions.ToEqual(slice1, slice2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error if we compare two string slices,
	where both have len=2,cap=3 and they have different int elements.`, func(t *testing.T) {
		slice1 := make([]string, 2, 3)
		slice1[0] = "Hello"
		slice1[1] = "World"
		slice2 := make([]string, 2, 3)
		slice2[0] = "Pershendetje"
		slice2[1] = "World"
		expectedErrorMsg := "Not equal:\n" +
			"expected: []string{\"Hello\", \"World\"}\n" +
			"actual  : []string{\"Pershendetje\", \"World\"}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" ([]string) (len=2) {\n" +
			"- (string) (len=5) \"Hello\",\n" +
			"+ (string) (len=12) \"Pershendetje\",\n" +
			"  (string) (len=5) \"World\"\n"
		err := assertions.ToEqual(slice1, slice2)
		if err == nil {
			t.Error("Slices should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error, if we compare two slices of len=0, cap=0 and
	one of them is of type []int and the other of type []uint32`, func(t *testing.T) {
		slice1 := []int{}
		slice2 := []uint{}
		err := assertions.ToEqual(slice1, slice2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: []int([]int{})\n" +
			"actual  : []uint([]uint{})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error, if we compare these two arrays: []int32{7}, []int64{7} `, func(t *testing.T) {
		slice1 := []int32{7}
		slice2 := []int64{7}
		err := assertions.ToEqual(slice1, slice2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: []int32([]int32{7})\n" +
			"actual  : []int64([]int64{7})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error, if we compare these two slices:
	[]float32{7.4, 3.2}, []float64{7.4, 3.2} `, func(t *testing.T) {
		slice1 := []float32{7.4, 3.2}
		slice2 := []float64{7.4, 3.2}
		err := assertions.ToEqual(slice1, slice2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: []float32([]float32{7.4, 3.2})\n" +
			"actual  : []float64([]float64{7.4, 3.2})"
		if err == nil {
			t.Error("string arrays are not equal in type, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is an array arr := [5]int{10, 20, 0, 10, 20}
	And there is a slice created out of that array sl1 := arr[0:2]
	And there is another slice created out of that array sl2 := arr[3:5]
	When we compare the using the ToEqual() assertion
	Then the assertion should not return an error.`, func(t *testing.T) {
		arr := [5]int{10, 20, 0, 10, 20}
		sl1 := arr[0:2]
		sl2 := arr[3:5]

		err := assertions.ToEqual(sl1, sl2)
		if err != nil {
			t.Error("Slices should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)
}

func TestToEqualStructs(t *testing.T) {
	Test(`
	Given that there is a struct type: S{field int}
	And two structs of that type s1 := S{field: 1}, s2 := S{field: 1}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field int
		}
		s1 := S{field: 1}
		s2 := S{field: 1}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field int}
	And two structs of that type s1 := S{field: 1}, s2 := S{field: 2}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field int
		}
		s1 := S{field: 1}
		s2 := S{field: 2}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:1}\n" +
			"actual  : tests_test.S{field:2}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (tests_test.S) {\n" +
			"- field: (int) 1\n" +
			"+ field: (int) 2\n" +
			" }\n"
		err := assertions.ToEqual(s1, s2)
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there are two different struct types but with identical fields: 
	S1{field int}, S2{field int}
	And two structs of that type s1 := S1{field: 1}, s2 := S2{field: 1}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error.`, func(t *testing.T) {
		type S1 struct{ field int }
		type S2 struct{ field int }
		s1 := S1{field: 1}
		s2 := S2{field: 1}
		err := assertions.ToEqual(s1, s2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S1(tests_test.S1{field:1})\n" +
			"actual  : tests_test.S2(tests_test.S2{field:1})"
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field float64}
	And two structs of that type s1 := S{field: 2.3}, s2 := S{field: 2.3}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field float64
		}
		s1 := S{field: 2.3}
		s2 := S{field: 2.3}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field float64}
	And two structs of that type s1 := S{field: 2.3}, s2 := S{field: 2.7}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field float64
		}
		s1 := S{field: 2.3}
		s2 := S{field: 2.7}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:2.3}\n" +
			"actual  : tests_test.S{field:2.7}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (tests_test.S) {\n" +
			"- field: (float64) 2.3\n" +
			"+ field: (float64) 2.7\n" +
			" }\n"
		err := assertions.ToEqual(s1, s2)
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field bool}
	And two structs of that type s1 := S{field: true}, s2 := S{field: true}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field bool
		}
		s1 := S{field: true}
		s2 := S{field: true}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field bool}
	And two structs of that type s1 := S{field: true}, s2 := S{field: false}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field bool
		}
		s1 := S{field: true}
		s2 := S{field: false}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:true}\n" +
			"actual  : tests_test.S{field:false}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (tests_test.S) {\n" +
			"- field: (bool) true\n" +
			"+ field: (bool) false\n" +
			" }\n"
		err := assertions.ToEqual(s1, s2)
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there are two different struct types but with identical fields: 
	S1{field bool}, S2{field bool}
	And two structs of that type s1 := S1{field: true}, s2 := S2{field: true}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error.`, func(t *testing.T) {
		type S1 struct{ field bool }
		type S2 struct{ field bool }
		s1 := S1{field: true}
		s2 := S2{field: true}
		err := assertions.ToEqual(s1, s2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S1(tests_test.S1{field:true})\n" +
			"actual  : tests_test.S2(tests_test.S2{field:true})"
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field string}
	And two structs of that type s1 := S{field: "Hello"}, s2 := S{field: "Hello"}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field string
		}
		s1 := S{field: "Hello"}
		s2 := S{field: "Hello"}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field string}
	And two structs of that type s1 := S{field: "Hello"}, s2 := S{field: "Hi"}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field string
		}
		s1 := S{field: "Hello"}
		s2 := S{field: "Hi"}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:\"Hello\"}\n" +
			"actual  : tests_test.S{field:\"Hi\"}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (tests_test.S) {\n" +
			"- field: (string) (len=5) \"Hello\"\n" +
			"+ field: (string) (len=2) \"Hi\"\n" +
			" }\n"
		err := assertions.ToEqual(s1, s2)
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there are two different struct types but with identical fields:
	S1{field string}, S2{field string}
	And two structs of that type s1 := S1{field: "Hello"}, s2 := S2{field: "Hello"}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error.`, func(t *testing.T) {
		type S1 struct{ field string }
		type S2 struct{ field string }
		s1 := S1{field: "Hello"}
		s2 := S2{field: "Hello"}
		err := assertions.ToEqual(s1, s2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S1(tests_test.S1{field:\"Hello\"})\n" +
			"actual  : tests_test.S2(tests_test.S2{field:\"Hello\"})"
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field [2]int}
	And two structs of that type s1 := S{field: [2]int{10, 20}}, s2 := S{field: [2]int{10, 20}}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field [2]int
		}
		s1 := S{field: [2]int{10, 20}}
		s2 := S{field: [2]int{10, 20}}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field [2]int}
	And two structs of that type s1 := S{field: [2]int{10, 20}}, s2 := S{field: [2]int{13, 17}}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field [2]int
		}
		s1 := S{field: [2]int{10, 20}}
		s2 := S{field: [2]int{13, 17}}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:[2]int{10, 20}}\n" +
			"actual  : tests_test.S{field:[2]int{13, 17}}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -2,4 +2,4 @@\n" +
			"  field: ([2]int) (len=2) {\n" +
			"-  (int) 10,\n" +
			"-  (int) 20\n" +
			"+  (int) 13,\n" +
			"+  (int) 17\n" +
			"  }\n"
		err := assertions.ToEqual(s1, s2)

		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there are two different struct types but with identical fields:
	S1{field [2]int}, S2{field [2]int}
	And two structs of that type s1 := S1{field: [2]int{10, 20}}, s2 := S2{field: [2]int{10, 20}}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error.`, func(t *testing.T) {
		type S1 struct{ field [2]int }
		type S2 struct{ field [2]int }
		s1 := S1{field: [2]int{10, 20}}
		s2 := S2{field: [2]int{10, 20}}
		err := assertions.ToEqual(s1, s2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S1(tests_test.S1{field:[2]int{10, 20}})\n" +
			"actual  : tests_test.S2(tests_test.S2{field:[2]int{10, 20}})"
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field []int}
	And two structs of that type s1 := S{field: []int{10, 20}}, s2 := S{field: []int{10, 20}}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type S struct {
			field []int
		}
		s1 := S{field: []int{10, 20}}
		s2 := S{field: []int{10, 20}}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there is a struct type: S{field []int}
	And two structs of that type s1 := S{field: []int{10, 20}}, s2 := S{field: []int{13, 17}}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field []int
		}
		s1 := S{field: []int{10, 20}}
		s2 := S{field: []int{13, 17}}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:[]int{10, 20}}\n" +
			"actual  : tests_test.S{field:[]int{13, 17}}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -2,4 +2,4 @@\n" +
			"  field: ([]int) (len=2) {\n" +
			"-  (int) 10,\n" +
			"-  (int) 20\n" +
			"+  (int) 13,\n" +
			"+  (int) 17\n" +
			"  }\n"
		err := assertions.ToEqual(s1, s2)

		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there are two different struct types but with identical fields:
	S1{field []int}, S2{field []int}
	And two structs of that type s1 := S1{field: []int{10, 20}}, s2 := S2{field: []int{10, 20}}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error.`, func(t *testing.T) {
		type S1 struct{ field []int }
		type S2 struct{ field []int }
		s1 := S1{field: []int{10, 20}}
		s2 := S2{field: []int{10, 20}}
		err := assertions.ToEqual(s1, s2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S1(tests_test.S1{field:[]int{10, 20}})\n" +
			"actual  : tests_test.S2(tests_test.S2{field:[]int{10, 20}})"
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there a struct type: NestedStruct{field float64} 
	And another struct type: S{nested NestedStruct}
	And two structs of that type s1 := S{nested: NestedStruct{field: 2}}, s2 := S{nested: NestedStruct{field: 2}}
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error`, func(t *testing.T) {
		type NestedStruct struct {
			field int
		}
		type S struct {
			nested NestedStruct
		}
		s1 := S{nested: NestedStruct{field: 2}}
		s2 := S{nested: NestedStruct{field: 2}}
		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`
	Given that there a struct type: NestedStruct{field float64} 
	And another struct type: S{nested NestedStruct}
	And two structs of that type s1 := S{nested: NestedStruct{field: 2}}, s2 := S{nested: NestedStruct{field: 16}}
	When we compare them using the ToEqual assertion
	Then the assertion should return an error and the error should contain the Diff.`, func(t *testing.T) {
		type S struct {
			field int
		}
		s1 := S{field: 2}
		s2 := S{field: 16}
		expectedErrorMsg := "Not equal:\n" +
			"expected: tests_test.S{field:2}\n" +
			"actual  : tests_test.S{field:16}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (tests_test.S) {\n" +
			"- field: (int) 2\n" +
			"+ field: (int) 16\n" +
			" }\n"
		err := assertions.ToEqual(s1, s2)
		if err == nil {
			t.Error("Structs are not equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`
	Given that there a struct type: NestedStruct{field1 float64, field2 int}
	And another struct type: S{
		intField    int
		floatField  float64
		stringField string
		boolField   bool
		arrField    [2]int
		sliceField  []float64
		nested      NestedStruct
	}
	And two structs of type S that have their fields set to equal values
	When we compare them using the ToEqual assertion
	Then the assertion should not return an error
	`, func(t *testing.T) {
		type NestedStruct struct {
			field1 float64
			field2 int
		}

		type S struct {
			intField    int
			floatField  float64
			stringField string
			boolField   bool
			arrField    [2]int
			sliceField  []float64
			nested      NestedStruct
		}
		s1 := S{
			intField:    2,
			floatField:  4.5,
			stringField: "Hello",
			boolField:   true,
			arrField:    [2]int{8, 4},
			sliceField:  []float64{2.4, 9.2},
			nested: NestedStruct{
				field1: 5.6,
				field2: 8,
			},
		}
		s2 := S{
			intField:    2,
			floatField:  4.5,
			stringField: "Hello",
			boolField:   true,
			arrField:    [2]int{8, 4},
			sliceField:  []float64{2.4, 9.2},
			nested: NestedStruct{
				field1: 5.6,
				field2: 8,
			},
		}

		err := assertions.ToEqual(s1, s2)
		if err != nil {
			t.Error("Structs are equal, but ToEqual() assertion says they are not.")
		}
	}, t)
}

func TestToEqualMaps(t *testing.T) {
	Test("it should not return an error if we compare two empty maps of the same type.", func(t *testing.T) {
		map1 := map[string]int{}
		map2 := map[string]int{}
		err := assertions.ToEqual(map1, map2)
		if err != nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two maps m1: map[string]int{"k1": 2},
	m1: map[string]int{"k1": 2}.`, func(t *testing.T) {
		map1 := map[string]int{"k1": 2}
		map2 := map[string]int{"k1": 2}
		err := assertions.ToEqual(map1, map2)
		if err != nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two maps m1: map[int]int{1: 2},
	m1: map[string]int{1: 2}.`, func(t *testing.T) {
		map1 := map[int]int{1: 2}
		map2 := map[int]int{1: 2}
		err := assertions.ToEqual(map1, map2)
		if err != nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two maps m1: map[bool]string{true: "true"},
	m1: map[bool]string{true: "true"}.`, func(t *testing.T) {
		map1 := map[bool]string{true: "true"}
		map2 := map[bool]string{true: "true"}
		err := assertions.ToEqual(map1, map2)
		if err != nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should not return an error if we compare two maps m1: map[string][]int{"k": {2}},
	m1: map[string][]int{"k": {2}.`, func(t *testing.T) {
		map1 := map[string][]int{"k": {2}}
		map2 := map[string][]int{"k": {2}}
		err := assertions.ToEqual(map1, map2)
		if err != nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are not.")
		}
	}, t)

	Test(`it should return an error if we compare two maps m1: map[string]int{"k1": 2},
	m1: map[string]int{"k1": 3}.`, func(t *testing.T) {
		map1 := map[string]int{"k1": 2}
		map2 := map[string]int{"k1": 3}
		err := assertions.ToEqual(map1, map2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: map[string]int{\"k1\":2}\n" +
			"actual  : map[string]int{\"k1\":3}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (map[string]int) (len=1) {\n" +
			"- (string) (len=2) \"k1\": (int) 2\n" +
			"+ (string) (len=2) \"k1\": (int) 3\n" +
			" }\n"
		if err == nil {
			t.Error("Maps should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error if we compare two maps m1: map[string]int{"k1": 2},
	m1: map[string]int{"k2": 2}.`, func(t *testing.T) {
		map1 := map[string]int{"k1": 2}
		map2 := map[string]int{"k2": 2}
		err := assertions.ToEqual(map1, map2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: map[string]int{\"k1\":2}\n" +
			"actual  : map[string]int{\"k2\":2}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (map[string]int) (len=1) {\n" +
			"- (string) (len=2) \"k1\": (int) 2\n" +
			"+ (string) (len=2) \"k2\": (int) 2\n" +
			" }\n"
		if err == nil {
			t.Error("Maps should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error if we compare two maps m1: map[int]int{1: 7},
	m1: map[string]int{1: 2}.`, func(t *testing.T) {
		map1 := map[int]int{1: 7}
		map2 := map[int]int{1: 2}
		err := assertions.ToEqual(map1, map2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: map[int]int{1:7}\n" +
			"actual  : map[int]int{1:2}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (map[int]int) (len=1) {\n" +
			"- (int) 1: (int) 7\n" +
			"+ (int) 1: (int) 2\n" +
			" }\n"
		if err == nil {
			t.Error("Maps should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error if we compare two maps m1: map[bool]string{true: "true"},
	m1: map[bool]string{false: "false"}.`, func(t *testing.T) {
		map1 := map[bool]string{true: "true"}
		map2 := map[bool]string{true: "false"}
		err := assertions.ToEqual(map1, map2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: map[bool]string{true:\"true\"}\n" +
			"actual  : map[bool]string{true:\"false\"}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -1,3 +1,3 @@\n" +
			" (map[bool]string) (len=1) {\n" +
			"- (bool) true: (string) (len=4) \"true\"\n" +
			"+ (bool) true: (string) (len=5) \"false\"\n" +
			" }\n"
		if err == nil {
			t.Error("Maps should be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)

	Test(`it should return an error if we compare two maps m1: map[string][]int{"k": {2}},
	m1: map[string][]int{"k": {3}.`, func(t *testing.T) {
		map1 := map[string][]int{"k": {2}}
		map2 := map[string][]int{"k": {3}}
		err := assertions.ToEqual(map1, map2)

		expectedErrorMsg := "Not equal:\n" +
			"expected: map[string][]int{\"k\":[]int{2}}\n" +
			"actual  : map[string][]int{\"k\":[]int{3}}\n\n" +
			"Diff:\n" +
			"--- Expected\n" +
			"+++ Actual\n" +
			"@@ -2,3 +2,3 @@\n" +
			"  (string) (len=1) \"k\": ([]int) (len=1) {\n" +
			"-  (int) 2\n" +
			"+  (int) 3\n" +
			"  }\n"
		if err == nil {
			t.Error("Maps should not be equal, but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)
}

func TestNil(t *testing.T) {
	Test("it should not return an error if compare two nils.", func(t *testing.T) {
		err := assertions.ToEqual(nil, nil)
		if err != nil {
			t.Error("nil is equal to nil, but ToEqual() assertion says it is not.")
		}
	}, t)

	Test("it should not return an error if we compare two int pointers set to nil.", func(t *testing.T) {
		var ptr1 *int = nil
		var ptr2 *int = nil
		err := assertions.ToEqual(ptr1, ptr2)
		if err != nil {
			t.Error("nil is equal to nil, but ToEqual() assertion says it is not.")
		}
	}, t)

	Test("it should return an error if we compare two pointers of different types set to nil.", func(t *testing.T) {
		var ptr1 *int32 = nil
		var ptr2 *int64 = nil
		err := assertions.ToEqual(ptr1, ptr2)
		expectedErrorMsg := "Not equal:\n" +
			"expected: *int32((*int32)(nil))\n" +
			"actual  : *int64((*int64)(nil))"
		if err == nil {
			t.Error("*int32((*int32)(nil)) is not equal to *int64((*int64)(nil)), but ToEqual() assertion says they are.")
		}
		if err.Error() != expectedErrorMsg {
			t.Errorf(
				"The error message does not have the expected format.\n\nShould be:\n%s\n\nIs:\n%s",
				expectedErrorMsg, err.Error())
		}
	}, t)
}

func TestFunction(t *testing.T) {
	Test("it should return an error if we compare a function to an integer.", func(t *testing.T) {
		f := func() {}
		err := assertions.ToEqual(f, 5)
		if err == nil {
			t.Error("A function cannot be compared to an integer.")
		}
	}, t)

	Test("it should return an error if we compare a function to an integer.", func(t *testing.T) {
		f := func() {}
		err := assertions.ToEqual(6, f)
		if err == nil {
			t.Error("A function cannot be compared to an integer.")
		}
	}, t)

	Test("it should return an error if we compare a function to an integer (#2).", func(t *testing.T) {
		f1 := func() {}
		f2 := func() {}
		err := assertions.ToEqual(f1, f2)
		if err == nil {
			t.Error("A function cannot be compared to a another function.")
		}
	}, t)

	Test("it should return an error if we compare a defined function to nil.", func(t *testing.T) {
		f := func() {}
		err := assertions.ToEqual(f, nil)
		if err == nil {
			t.Error("A defined function is not nil.")
		}
	}, t)

	Test("it should return an error if we compare an undefined function to nil.", func(t *testing.T) {
		var f func()
		err := assertions.ToEqual(nil, f)
		if err == nil {
			t.Error("A function cannot be compared with ToEqual().")
		}
	}, t)

	Test("it should return an error if we compare an undefined function to nil (#2).", func(t *testing.T) {
		var f func()
		err := assertions.ToEqual(f, nil)
		if err == nil {
			t.Error("A function cannot be compared with ToEqual().")
		}
	}, t)
}
