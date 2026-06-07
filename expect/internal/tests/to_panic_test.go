package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToPanic(t *testing.T) {
	var tests = []struct {
		name           string
		input          any
		assertionFails bool
	}{
		{name: "a function that panics with a string", input: func() { panic("boom") }, assertionFails: false},
		{name: "a function that panics with an error", input: func() { panic(fmt.Errorf("boom")) }, assertionFails: false},
		{name: "a function that panics with nil", input: func() { panic(nil) }, assertionFails: false},
		{name: "a function that does not panic", input: func() {}, assertionFails: true},
		{name: "a function that takes arguments", input: func(x int) {}, assertionFails: true},
		{name: "a non-function value", input: 42, assertionFails: true},
		{name: "a nil value", input: nil, assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, for %s", test.name)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, for %s", test.name)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToPanic(test.input)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
