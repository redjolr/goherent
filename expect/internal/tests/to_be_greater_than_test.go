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

		// Lexigographic comparison of slices
		{isGreaterCandidate: []byte{1, 2, 3}, checkIfGreaterAgainst: []byte{1, 2}, assertionFails: false},   // longer slice
		{isGreaterCandidate: []byte{1, 2, 3}, checkIfGreaterAgainst: []byte{1, 2, 3}, assertionFails: true}, // equal slices
		{isGreaterCandidate: []byte{1, 3}, checkIfGreaterAgainst: []byte{1, 2}, assertionFails: false},      // lexicographically greater
		{isGreaterCandidate: []byte{1, 1, 1}, checkIfGreaterAgainst: []byte{1, 2}, assertionFails: true},    // lexicographically less
		{isGreaterCandidate: []byte{}, checkIfGreaterAgainst: []byte{1}, assertionFails: true},              // empty slice vs non-empty
		{isGreaterCandidate: []byte{1}, checkIfGreaterAgainst: []byte{}, assertionFails: false},             // non-empty slice vs empty
		{isGreaterCandidate: []byte{1, 2, 3}, checkIfGreaterAgainst: nil, assertionFails: true},             // slice vs nil
		{isGreaterCandidate: nil, checkIfGreaterAgainst: []byte{1, 2, 3}, assertionFails: true},             // nil vs slice

		// Valid comparisons with uintptr
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: uintptr(50), assertionFails: false}, // greater
		{isGreaterCandidate: uintptr(50), checkIfGreaterAgainst: uintptr(100), assertionFails: true},  // less
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: uintptr(100), assertionFails: true}, // equal

		// Comparisons involving zero values
		{isGreaterCandidate: uintptr(0), checkIfGreaterAgainst: uintptr(100), assertionFails: true},  // zero less than non-zero
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: uintptr(0), assertionFails: false}, // non-zero greater than zero
		{isGreaterCandidate: uintptr(0), checkIfGreaterAgainst: uintptr(0), assertionFails: true},    // zero vs zero

		// Comparisons with nil (uintptr cannot be nil but testing edge cases)
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: nil, assertionFails: true}, // uintptr vs nil
		{isGreaterCandidate: nil, checkIfGreaterAgainst: uintptr(100), assertionFails: true}, // nil vs uintptr

		// Invalid comparisons (uintptr with incompatible types)
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: 100, assertionFails: true},        // uintptr vs int
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: 100.0, assertionFails: true},      // uintptr vs float
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: "100", assertionFails: true},      // uintptr vs string
		{isGreaterCandidate: uintptr(100), checkIfGreaterAgainst: struct{}{}, assertionFails: true}, // uintptr vs struct

		// Edge cases with large values
		{isGreaterCandidate: uintptr(^uintptr(0)), checkIfGreaterAgainst: uintptr(0), assertionFails: false},          // max uintptr vs zero
		{isGreaterCandidate: uintptr(0), checkIfGreaterAgainst: uintptr(^uintptr(0)), assertionFails: true},           // zero vs max uintptr
		{isGreaterCandidate: uintptr(^uintptr(0)), checkIfGreaterAgainst: uintptr(^uintptr(0)), assertionFails: true}, // max vs max
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
