package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToContainElement(t *testing.T) {
	var tests = []struct {
		container      any
		element        any
		assertionFails bool
	}{
		{container: []int{1, 2, 3}, element: 2, assertionFails: false},
		{container: []int{1, 2, 3}, element: 4, assertionFails: true},
		{container: []string{"a", "b"}, element: "b", assertionFails: false},
		{container: []string{"a", "b"}, element: "c", assertionFails: true},
		{container: [3]int{1, 2, 3}, element: 3, assertionFails: false},
		{container: [3]int{1, 2, 3}, element: 9, assertionFails: true},
		{container: map[string]int{"a": 1, "b": 2}, element: 2, assertionFails: false},
		{container: map[string]int{"a": 1, "b": 2}, element: 3, assertionFails: true},
		// A string is not treated as a collection (use ToContain for substrings).
		{container: "hello", element: "ell", assertionFails: true},
		{container: 42, element: 42, assertionFails: true},
		{container: nil, element: 1, assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if %#v contains %#v", test.container, test.element)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if %#v contains %#v", test.container, test.element)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToContainElement(test.container, test.element)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
