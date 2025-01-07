package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeLessThanOrEqualTo(t *testing.T) {
	var tests = []struct {
		isLessOrEqualCandidate    any
		checkIfLessOrEqualAgainst any
		assertionFails            bool
	}{
		// Assertion should fail for different types
		{isLessOrEqualCandidate: 2, checkIfLessOrEqualAgainst: 1.2, assertionFails: true},
		{isLessOrEqualCandidate: 2, checkIfLessOrEqualAgainst: false, assertionFails: true},
		{isLessOrEqualCandidate: 3, checkIfLessOrEqualAgainst: "", assertionFails: true},
		{isLessOrEqualCandidate: 3, checkIfLessOrEqualAgainst: struct{ a float32 }{a: 1.3}, assertionFails: true},
		{isLessOrEqualCandidate: int(3), checkIfLessOrEqualAgainst: int64(1), assertionFails: true},

		// Invalid comparisons (booleans)
		{isLessOrEqualCandidate: true, checkIfLessOrEqualAgainst: false, assertionFails: true},
		{isLessOrEqualCandidate: false, checkIfLessOrEqualAgainst: true, assertionFails: true},
		{isLessOrEqualCandidate: true, checkIfLessOrEqualAgainst: true, assertionFails: true},

		// Valid comparisons (integers)
		{isLessOrEqualCandidate: 3, checkIfLessOrEqualAgainst: 2, assertionFails: true},
		{isLessOrEqualCandidate: 2, checkIfLessOrEqualAgainst: 3, assertionFails: false},
		{isLessOrEqualCandidate: 2, checkIfLessOrEqualAgainst: 2, assertionFails: false},

		// Valid comparisons (floats)
		{isLessOrEqualCandidate: 3.5, checkIfLessOrEqualAgainst: 2.1, assertionFails: true},
		{isLessOrEqualCandidate: 2.1, checkIfLessOrEqualAgainst: 3.5, assertionFails: false},
		{isLessOrEqualCandidate: 2.1, checkIfLessOrEqualAgainst: 2.1, assertionFails: false},

		// Valid comparisons (strings)
		{isLessOrEqualCandidate: "b", checkIfLessOrEqualAgainst: "a", assertionFails: true},
		{isLessOrEqualCandidate: "a", checkIfLessOrEqualAgainst: "b", assertionFails: false},
		{isLessOrEqualCandidate: "a", checkIfLessOrEqualAgainst: "a", assertionFails: false},

		// Invalid comparisons (non-comparable types)
		{isLessOrEqualCandidate: []int{1, 2}, checkIfLessOrEqualAgainst: []int{3, 4}, assertionFails: true},                       // slices
		{isLessOrEqualCandidate: map[string]int{"a": 1}, checkIfLessOrEqualAgainst: map[string]int{"b": 2}, assertionFails: true}, // maps
		{isLessOrEqualCandidate: make(chan int), checkIfLessOrEqualAgainst: make(chan int), assertionFails: true},                 // channels

		// Edge cases
		{isLessOrEqualCandidate: nil, checkIfLessOrEqualAgainst: nil, assertionFails: true},               // nil vs nil
		{isLessOrEqualCandidate: nil, checkIfLessOrEqualAgainst: 0, assertionFails: true},                 // nil vs int
		{isLessOrEqualCandidate: 0, checkIfLessOrEqualAgainst: nil, assertionFails: true},                 // int vs nil
		{isLessOrEqualCandidate: 0, checkIfLessOrEqualAgainst: 0, assertionFails: false},                  // zero equality
		{isLessOrEqualCandidate: "0", checkIfLessOrEqualAgainst: 0, assertionFails: true},                 // string vs int
		{isLessOrEqualCandidate: "", checkIfLessOrEqualAgainst: nil, assertionFails: true},                // empty string vs nil
		{isLessOrEqualCandidate: struct{}{}, checkIfLessOrEqualAgainst: struct{}{}, assertionFails: true}, // empty structs

		// Lexigographic comparison of slices
		{isLessOrEqualCandidate: []byte{1, 2, 3}, checkIfLessOrEqualAgainst: []byte{1, 2}, assertionFails: true},     // longer slice
		{isLessOrEqualCandidate: []byte{1, 2, 3}, checkIfLessOrEqualAgainst: []byte{1, 2, 3}, assertionFails: false}, // equal slices
		{isLessOrEqualCandidate: []byte{1, 3}, checkIfLessOrEqualAgainst: []byte{1, 2}, assertionFails: true},        // lexicographically greater
		{isLessOrEqualCandidate: []byte{1, 1, 1}, checkIfLessOrEqualAgainst: []byte{1, 2}, assertionFails: false},    // lexicographically less
		{isLessOrEqualCandidate: []byte{}, checkIfLessOrEqualAgainst: []byte{1}, assertionFails: false},              // empty slice vs non-empty
		{isLessOrEqualCandidate: []byte{1}, checkIfLessOrEqualAgainst: []byte{}, assertionFails: true},               // non-empty slice vs empty
		{isLessOrEqualCandidate: []byte{1, 2, 3}, checkIfLessOrEqualAgainst: nil, assertionFails: true},              // slice vs nil
		{isLessOrEqualCandidate: nil, checkIfLessOrEqualAgainst: []byte{1, 2, 3}, assertionFails: true},              // nil vs slice

		// Valid comparisons with uintptr
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: uintptr(50), assertionFails: true},
		{isLessOrEqualCandidate: uintptr(50), checkIfLessOrEqualAgainst: uintptr(100), assertionFails: false},
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: uintptr(100), assertionFails: false},

		// Comparisons involving zero values
		{isLessOrEqualCandidate: uintptr(0), checkIfLessOrEqualAgainst: uintptr(100), assertionFails: false}, // zero less than non-zero
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: uintptr(0), assertionFails: true},  // non-zero greater than zero
		{isLessOrEqualCandidate: uintptr(0), checkIfLessOrEqualAgainst: uintptr(0), assertionFails: false},   // zero vs zero

		// Comparisons with nil (uintptr cannot be nil but testing edge cases)
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: nil, assertionFails: true}, // uintptr vs nil
		{isLessOrEqualCandidate: nil, checkIfLessOrEqualAgainst: uintptr(100), assertionFails: true}, // nil vs uintptr

		// Invalid comparisons (uintptr with incompatible types)
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: 100, assertionFails: true},        // uintptr vs int
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: 100.0, assertionFails: true},      // uintptr vs float
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: "100", assertionFails: true},      // uintptr vs string
		{isLessOrEqualCandidate: uintptr(100), checkIfLessOrEqualAgainst: struct{}{}, assertionFails: true}, // uintptr vs struct

		// Edge cases with large values
		{isLessOrEqualCandidate: uintptr(^uintptr(0)), checkIfLessOrEqualAgainst: uintptr(0), assertionFails: true},            // max uintptr vs zero
		{isLessOrEqualCandidate: uintptr(0), checkIfLessOrEqualAgainst: uintptr(^uintptr(0)), assertionFails: false},           // zero vs max uintptr
		{isLessOrEqualCandidate: uintptr(^uintptr(0)), checkIfLessOrEqualAgainst: uintptr(^uintptr(0)), assertionFails: false}, // max vs max
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is not less than or equal to %v(%v)",
				test.isLessOrEqualCandidate,
				reflect.TypeOf(test.isLessOrEqualCandidate),
				test.checkIfLessOrEqualAgainst,
				reflect.TypeOf(test.checkIfLessOrEqualAgainst),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is greater than %v(%v)",
				test.isLessOrEqualCandidate,
				reflect.TypeOf(test.isLessOrEqualCandidate),
				test.checkIfLessOrEqualAgainst,
				reflect.TypeOf(test.checkIfLessOrEqualAgainst),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeLessThanOrEqualTo(test.isLessOrEqualCandidate, test.checkIfLessOrEqualAgainst)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
