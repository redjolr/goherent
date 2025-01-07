package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeGreaterThanOrEqualTo(t *testing.T) {
	var tests = []struct {
		isGreaterOrEqualCandidate    any
		checkIfGreaterOrEqualAgainst any
		assertionFails               bool
	}{
		// Assertion should fail for different types
		{isGreaterOrEqualCandidate: 2, checkIfGreaterOrEqualAgainst: 1.2, assertionFails: true},
		{isGreaterOrEqualCandidate: 2, checkIfGreaterOrEqualAgainst: false, assertionFails: true},
		{isGreaterOrEqualCandidate: 3, checkIfGreaterOrEqualAgainst: "", assertionFails: true},
		{isGreaterOrEqualCandidate: 3, checkIfGreaterOrEqualAgainst: struct{ a float32 }{a: 1.3}, assertionFails: true},
		{isGreaterOrEqualCandidate: int(3), checkIfGreaterOrEqualAgainst: int64(1), assertionFails: true},

		// Invalid comparisons (booleans)
		{isGreaterOrEqualCandidate: true, checkIfGreaterOrEqualAgainst: false, assertionFails: true},
		{isGreaterOrEqualCandidate: false, checkIfGreaterOrEqualAgainst: true, assertionFails: true},
		{isGreaterOrEqualCandidate: true, checkIfGreaterOrEqualAgainst: true, assertionFails: true},

		// Valid comparisons (integers)
		{isGreaterOrEqualCandidate: 3, checkIfGreaterOrEqualAgainst: 2, assertionFails: false},
		{isGreaterOrEqualCandidate: 2, checkIfGreaterOrEqualAgainst: 3, assertionFails: true},
		{isGreaterOrEqualCandidate: 2, checkIfGreaterOrEqualAgainst: 2, assertionFails: false},

		// Valid comparisons (floats)
		{isGreaterOrEqualCandidate: 3.5, checkIfGreaterOrEqualAgainst: 2.1, assertionFails: false},
		{isGreaterOrEqualCandidate: 2.1, checkIfGreaterOrEqualAgainst: 3.5, assertionFails: true},
		{isGreaterOrEqualCandidate: 2.1, checkIfGreaterOrEqualAgainst: 2.1, assertionFails: false},

		// Valid comparisons (strings)
		{isGreaterOrEqualCandidate: "b", checkIfGreaterOrEqualAgainst: "a", assertionFails: false},
		{isGreaterOrEqualCandidate: "a", checkIfGreaterOrEqualAgainst: "b", assertionFails: true},
		{isGreaterOrEqualCandidate: "a", checkIfGreaterOrEqualAgainst: "a", assertionFails: false},

		// Invalid comparisons (non-comparable types)
		{isGreaterOrEqualCandidate: []int{1, 2}, checkIfGreaterOrEqualAgainst: []int{3, 4}, assertionFails: true},                       // slices
		{isGreaterOrEqualCandidate: map[string]int{"a": 1}, checkIfGreaterOrEqualAgainst: map[string]int{"b": 2}, assertionFails: true}, // maps
		{isGreaterOrEqualCandidate: make(chan int), checkIfGreaterOrEqualAgainst: make(chan int), assertionFails: true},                 // channels

		// Edge cases
		{isGreaterOrEqualCandidate: nil, checkIfGreaterOrEqualAgainst: nil, assertionFails: true},
		{isGreaterOrEqualCandidate: nil, checkIfGreaterOrEqualAgainst: 0, assertionFails: true},
		{isGreaterOrEqualCandidate: 0, checkIfGreaterOrEqualAgainst: nil, assertionFails: true},
		{isGreaterOrEqualCandidate: 0, checkIfGreaterOrEqualAgainst: 0, assertionFails: false},
		{isGreaterOrEqualCandidate: "0", checkIfGreaterOrEqualAgainst: 0, assertionFails: true},
		{isGreaterOrEqualCandidate: "", checkIfGreaterOrEqualAgainst: nil, assertionFails: true},
		{isGreaterOrEqualCandidate: struct{}{}, checkIfGreaterOrEqualAgainst: struct{}{}, assertionFails: true},

		// Lexigographic comparison of slices
		{isGreaterOrEqualCandidate: []byte{1, 2, 3}, checkIfGreaterOrEqualAgainst: []byte{1, 2}, assertionFails: false},    // longer slice
		{isGreaterOrEqualCandidate: []byte{1, 2, 3}, checkIfGreaterOrEqualAgainst: []byte{1, 2, 3}, assertionFails: false}, // equal slices
		{isGreaterOrEqualCandidate: []byte{1, 3}, checkIfGreaterOrEqualAgainst: []byte{1, 2}, assertionFails: false},       // lexicographically greater
		{isGreaterOrEqualCandidate: []byte{1, 1, 1}, checkIfGreaterOrEqualAgainst: []byte{1, 2}, assertionFails: true},     // lexicographically less
		{isGreaterOrEqualCandidate: []byte{}, checkIfGreaterOrEqualAgainst: []byte{1}, assertionFails: true},               // empty slice vs non-empty
		{isGreaterOrEqualCandidate: []byte{1}, checkIfGreaterOrEqualAgainst: []byte{}, assertionFails: false},              // non-empty slice vs empty
		{isGreaterOrEqualCandidate: []byte{1, 2, 3}, checkIfGreaterOrEqualAgainst: nil, assertionFails: true},              // slice vs nil
		{isGreaterOrEqualCandidate: nil, checkIfGreaterOrEqualAgainst: []byte{1, 2, 3}, assertionFails: true},              // nil vs slice

		// Valid comparisons with uintptr
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: uintptr(50), assertionFails: false},
		{isGreaterOrEqualCandidate: uintptr(50), checkIfGreaterOrEqualAgainst: uintptr(100), assertionFails: true},
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: uintptr(100), assertionFails: false},

		// Comparisons involving zero values
		{isGreaterOrEqualCandidate: uintptr(0), checkIfGreaterOrEqualAgainst: uintptr(100), assertionFails: true},  // zero less than non-zero
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: uintptr(0), assertionFails: false}, // non-zero greater than zero
		{isGreaterOrEqualCandidate: uintptr(0), checkIfGreaterOrEqualAgainst: uintptr(0), assertionFails: false},   // zero vs zero

		// Comparisons with nil (uintptr cannot be nil but testing edge cases)
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: nil, assertionFails: true}, // uintptr vs nil
		{isGreaterOrEqualCandidate: nil, checkIfGreaterOrEqualAgainst: uintptr(100), assertionFails: true}, // nil vs uintptr

		// Invalid comparisons (uintptr with incompatible types)
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: 100, assertionFails: true},        // uintptr vs int
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: 100.0, assertionFails: true},      // uintptr vs float
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: "100", assertionFails: true},      // uintptr vs string
		{isGreaterOrEqualCandidate: uintptr(100), checkIfGreaterOrEqualAgainst: struct{}{}, assertionFails: true}, // uintptr vs struct

		// Edge cases with large values
		{isGreaterOrEqualCandidate: uintptr(^uintptr(0)), checkIfGreaterOrEqualAgainst: uintptr(0), assertionFails: false},           // max uintptr vs zero
		{isGreaterOrEqualCandidate: uintptr(0), checkIfGreaterOrEqualAgainst: uintptr(^uintptr(0)), assertionFails: true},            // zero vs max uintptr
		{isGreaterOrEqualCandidate: uintptr(^uintptr(0)), checkIfGreaterOrEqualAgainst: uintptr(^uintptr(0)), assertionFails: false}, // max vs max
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is greater than or equal to %v(%v)",
				test.isGreaterOrEqualCandidate,
				reflect.TypeOf(test.isGreaterOrEqualCandidate),
				test.checkIfGreaterOrEqualAgainst,
				reflect.TypeOf(test.checkIfGreaterOrEqualAgainst),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is greater than or equal to %v(%v)",
				test.isGreaterOrEqualCandidate,
				reflect.TypeOf(test.isGreaterOrEqualCandidate),
				test.checkIfGreaterOrEqualAgainst,
				reflect.TypeOf(test.checkIfGreaterOrEqualAgainst),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeGreaterThanOrEqualTo(test.isGreaterOrEqualCandidate, test.checkIfGreaterOrEqualAgainst)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
