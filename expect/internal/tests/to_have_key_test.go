package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToHaveKey(t *testing.T) {
	var tests = []struct {
		container      any
		key            any
		assertionFails bool
	}{
		{container: map[string]int{"a": 1}, key: "a", assertionFails: false},
		{container: map[string]int{"a": 1}, key: "b", assertionFails: true},
		{container: map[int]string{1: "x", 2: "y"}, key: 2, assertionFails: false},
		{container: map[int]string{1: "x"}, key: 2, assertionFails: true},
		{container: map[string]int{}, key: "a", assertionFails: true},
		// Non-maps always fail.
		{container: []int{1, 2}, key: 1, assertionFails: true},
		{container: "str", key: "s", assertionFails: true},
		{container: nil, key: "a", assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if %#v has key %#v", test.container, test.key)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if %#v has key %#v", test.container, test.key)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToHaveKey(test.container, test.key)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
