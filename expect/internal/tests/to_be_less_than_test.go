package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeLessThan(t *testing.T) {
	var tests = []struct {
		isLessCandidate    any
		checkIfLessAgainst any
		assertionFails     bool
	}{
		// Assertion should fail for different types
		{isLessCandidate: 2, checkIfLessAgainst: 1.2, assertionFails: true},
		{isLessCandidate: 2, checkIfLessAgainst: false, assertionFails: true},
		{isLessCandidate: 3, checkIfLessAgainst: "", assertionFails: true},
		{isLessCandidate: 3, checkIfLessAgainst: struct{ a float32 }{a: 1.3}, assertionFails: true},
		{isLessCandidate: int(3), checkIfLessAgainst: int64(1), assertionFails: true},

		// Invalid comparisons (booleans)
		{isLessCandidate: true, checkIfLessAgainst: false, assertionFails: true},
		{isLessCandidate: false, checkIfLessAgainst: true, assertionFails: true},
		{isLessCandidate: true, checkIfLessAgainst: true, assertionFails: true},

		// Valid comparisons (integers)
		{isLessCandidate: 3, checkIfLessAgainst: 2, assertionFails: true},
		{isLessCandidate: 2, checkIfLessAgainst: 3, assertionFails: false},
		{isLessCandidate: 2, checkIfLessAgainst: 2, assertionFails: true},

		// Valid comparisons (floats)
		{isLessCandidate: 3.5, checkIfLessAgainst: 2.1, assertionFails: true},
		{isLessCandidate: 2.1, checkIfLessAgainst: 3.5, assertionFails: false},
		{isLessCandidate: 2.1, checkIfLessAgainst: 2.1, assertionFails: true},

		// Valid comparisons (strings)
		{isLessCandidate: "b", checkIfLessAgainst: "a", assertionFails: true},
		{isLessCandidate: "a", checkIfLessAgainst: "b", assertionFails: false},
		{isLessCandidate: "a", checkIfLessAgainst: "a", assertionFails: true},

		// Invalid comparisons (non-comparable types)
		{isLessCandidate: []int{1, 2}, checkIfLessAgainst: []int{3, 4}, assertionFails: true},                       // slices
		{isLessCandidate: map[string]int{"a": 1}, checkIfLessAgainst: map[string]int{"b": 2}, assertionFails: true}, // maps
		{isLessCandidate: make(chan int), checkIfLessAgainst: make(chan int), assertionFails: true},                 // channels

		// Edge cases
		{isLessCandidate: nil, checkIfLessAgainst: nil, assertionFails: true},               // nil vs nil
		{isLessCandidate: nil, checkIfLessAgainst: 0, assertionFails: true},                 // nil vs int
		{isLessCandidate: 0, checkIfLessAgainst: nil, assertionFails: true},                 // int vs nil
		{isLessCandidate: 0, checkIfLessAgainst: 0, assertionFails: true},                   // zero equality
		{isLessCandidate: "0", checkIfLessAgainst: 0, assertionFails: true},                 // string vs int
		{isLessCandidate: "", checkIfLessAgainst: nil, assertionFails: true},                // empty string vs nil
		{isLessCandidate: struct{}{}, checkIfLessAgainst: struct{}{}, assertionFails: true}, // empty structs

		// Lexigographic comparison of slices
		{isLessCandidate: []byte{1, 2, 3}, checkIfLessAgainst: []byte{1, 2}, assertionFails: true},    // longer slice
		{isLessCandidate: []byte{1, 2, 3}, checkIfLessAgainst: []byte{1, 2, 3}, assertionFails: true}, // equal slices
		{isLessCandidate: []byte{1, 3}, checkIfLessAgainst: []byte{1, 2}, assertionFails: true},       // lexicographically greater
		{isLessCandidate: []byte{1, 1, 1}, checkIfLessAgainst: []byte{1, 2}, assertionFails: false},   // lexicographically less
		{isLessCandidate: []byte{}, checkIfLessAgainst: []byte{1}, assertionFails: false},             // empty slice vs non-empty
		{isLessCandidate: []byte{1}, checkIfLessAgainst: []byte{}, assertionFails: true},              // non-empty slice vs empty
		{isLessCandidate: []byte{1, 2, 3}, checkIfLessAgainst: nil, assertionFails: true},             // slice vs nil
		{isLessCandidate: nil, checkIfLessAgainst: []byte{1, 2, 3}, assertionFails: true},             // nil vs slice

		// Valid comparisons with uintptr
		{isLessCandidate: uintptr(100), checkIfLessAgainst: uintptr(50), assertionFails: true},
		{isLessCandidate: uintptr(50), checkIfLessAgainst: uintptr(100), assertionFails: false},
		{isLessCandidate: uintptr(100), checkIfLessAgainst: uintptr(100), assertionFails: true},

		// Comparisons involving zero values
		{isLessCandidate: uintptr(0), checkIfLessAgainst: uintptr(100), assertionFails: false}, // zero less than non-zero
		{isLessCandidate: uintptr(100), checkIfLessAgainst: uintptr(0), assertionFails: true},  // non-zero greater than zero
		{isLessCandidate: uintptr(0), checkIfLessAgainst: uintptr(0), assertionFails: true},    // zero vs zero

		// Comparisons with nil (uintptr cannot be nil but testing edge cases)
		{isLessCandidate: uintptr(100), checkIfLessAgainst: nil, assertionFails: true}, // uintptr vs nil
		{isLessCandidate: nil, checkIfLessAgainst: uintptr(100), assertionFails: true}, // nil vs uintptr

		// Invalid comparisons (uintptr with incompatible types)
		{isLessCandidate: uintptr(100), checkIfLessAgainst: 100, assertionFails: true},        // uintptr vs int
		{isLessCandidate: uintptr(100), checkIfLessAgainst: 100.0, assertionFails: true},      // uintptr vs float
		{isLessCandidate: uintptr(100), checkIfLessAgainst: "100", assertionFails: true},      // uintptr vs string
		{isLessCandidate: uintptr(100), checkIfLessAgainst: struct{}{}, assertionFails: true}, // uintptr vs struct

		// Edge cases with large values
		{isLessCandidate: uintptr(^uintptr(0)), checkIfLessAgainst: uintptr(0), assertionFails: true},           // max uintptr vs zero
		{isLessCandidate: uintptr(0), checkIfLessAgainst: uintptr(^uintptr(0)), assertionFails: false},          // zero vs max uintptr
		{isLessCandidate: uintptr(^uintptr(0)), checkIfLessAgainst: uintptr(^uintptr(0)), assertionFails: true}, // max vs max
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is greater than %v(%v)",
				test.isLessCandidate,
				reflect.TypeOf(test.isLessCandidate),
				test.checkIfLessAgainst,
				reflect.TypeOf(test.checkIfLessAgainst),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is greater than %v(%v)",
				test.isLessCandidate,
				reflect.TypeOf(test.isLessCandidate),
				test.checkIfLessAgainst,
				reflect.TypeOf(test.checkIfLessAgainst),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeLessThan(test.isLessCandidate, test.checkIfLessAgainst)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
