package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToMatch(t *testing.T) {
	var tests = []struct {
		input          any
		pattern        string
		assertionFails bool
	}{
		{input: "goherent v1.2.3", pattern: `^goherent v\d+\.\d+\.\d+$`, assertionFails: false},
		{input: "hello world", pattern: "world", assertionFails: false},
		{input: "hello world", pattern: "^hello", assertionFails: false},
		{input: "hello", pattern: "^world$", assertionFails: true},
		{input: "", pattern: ".+", assertionFails: true},
		// Invalid regex patterns fail rather than match.
		{input: "abc", pattern: "[", assertionFails: true},
		// Non-string values fail.
		{input: 123, pattern: "1", assertionFails: true},
		{input: nil, pattern: ".*", assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if %#v matches %q", test.input, test.pattern)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if %#v matches %q", test.input, test.pattern)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToMatch(test.input, test.pattern)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
