package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeGreaterThan(t *testing.T) {
	var tests = []struct {
		isGreaterCandidate    any
		checkIfGreaterAgainst any
		assertionFails        bool
	}{
		// Assertion should fail for different types
		{isGreaterCandidate: 2, checkIfGreaterAgainst: 1.2, assertionFails: true},
		{isGreaterCandidate: 2, checkIfGreaterAgainst: false, assertionFails: true},
		{isGreaterCandidate: 3, checkIfGreaterAgainst: "", assertionFails: true},
		{isGreaterCandidate: 3, checkIfGreaterAgainst: struct{ a float32 }{a: 1.3}, assertionFails: true},
		{isGreaterCandidate: int(3), checkIfGreaterAgainst: int64(1), assertionFails: true},

		// Invalid comparisons (booleans)
		{isGreaterCandidate: true, checkIfGreaterAgainst: false, assertionFails: true},
		{isGreaterCandidate: false, checkIfGreaterAgainst: true, assertionFails: true},
		{isGreaterCandidate: true, checkIfGreaterAgainst: true, assertionFails: true},

		// Valid comparisons (integers)
		{isGreaterCandidate: 3, checkIfGreaterAgainst: 2, assertionFails: false}, // greater
		{isGreaterCandidate: 2, checkIfGreaterAgainst: 3, assertionFails: true},  // less
		{isGreaterCandidate: 2, checkIfGreaterAgainst: 2, assertionFails: true},  // equal

		// Valid comparisons (floats)
		{isGreaterCandidate: 3.5, checkIfGreaterAgainst: 2.1, assertionFails: false}, // greater
		{isGreaterCandidate: 2.1, checkIfGreaterAgainst: 3.5, assertionFails: true},  // less
		{isGreaterCandidate: 2.1, checkIfGreaterAgainst: 2.1, assertionFails: true},  // equal

		// Valid comparisons (strings)
		{isGreaterCandidate: "b", checkIfGreaterAgainst: "a", assertionFails: false}, // lexicographical order
		{isGreaterCandidate: "a", checkIfGreaterAgainst: "b", assertionFails: true},  // lexicographical order
		{isGreaterCandidate: "a", checkIfGreaterAgainst: "a", assertionFails: true},  // equal strings

		// Invalid comparisons (non-comparable types)
		{isGreaterCandidate: []int{1, 2}, checkIfGreaterAgainst: []int{3, 4}, assertionFails: true},                       // slices
		{isGreaterCandidate: map[string]int{"a": 1}, checkIfGreaterAgainst: map[string]int{"b": 2}, assertionFails: true}, // maps
		{isGreaterCandidate: make(chan int), checkIfGreaterAgainst: make(chan int), assertionFails: true},                 // channels

		// Edge cases
		{isGreaterCandidate: nil, checkIfGreaterAgainst: nil, assertionFails: true},               // nil vs nil
		{isGreaterCandidate: nil, checkIfGreaterAgainst: 0, assertionFails: true},                 // nil vs int
		{isGreaterCandidate: 0, checkIfGreaterAgainst: nil, assertionFails: true},                 // int vs nil
		{isGreaterCandidate: 0, checkIfGreaterAgainst: 0, assertionFails: true},                   // zero equality
		{isGreaterCandidate: "0", checkIfGreaterAgainst: 0, assertionFails: true},                 // string vs int
		{isGreaterCandidate: "", checkIfGreaterAgainst: nil, assertionFails: true},                // empty string vs nil
		{isGreaterCandidate: struct{}{}, checkIfGreaterAgainst: struct{}{}, assertionFails: true}, // empty structs
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is greater than %v(%v)",
				test.isGreaterCandidate,
				reflect.TypeOf(test.isGreaterCandidate),
				test.checkIfGreaterAgainst,
				reflect.TypeOf(test.checkIfGreaterAgainst),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is greater than %v(%v)",
				test.isGreaterCandidate,
				reflect.TypeOf(test.isGreaterCandidate),
				test.checkIfGreaterAgainst,
				reflect.TypeOf(test.checkIfGreaterAgainst),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeGreaterThan(test.isGreaterCandidate, test.checkIfGreaterAgainst)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
