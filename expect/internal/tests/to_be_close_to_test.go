package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeCloseTo(t *testing.T) {
	var tests = []struct {
		value          any
		target         any
		tolerance      any
		assertionFails bool
	}{
		{value: 3.14159, target: 3.14, tolerance: 0.01, assertionFails: false},
		{value: 3.2, target: 3.14, tolerance: 0.01, assertionFails: true},
		// The tolerance is inclusive.
		{value: 10.0, target: 11.0, tolerance: 1.0, assertionFails: false},
		{value: 10.0, target: 12.0, tolerance: 1.0, assertionFails: true},
		{value: 5.0, target: 5.0, tolerance: 0.0, assertionFails: false},
		// Mixed numeric types are converted to float64.
		{value: 10, target: 10, tolerance: 0, assertionFails: false},
		{value: float32(1.0), target: 1.05, tolerance: 0.1, assertionFails: false},
		// Negative differences are handled via absolute value.
		{value: 2.0, target: 2.5, tolerance: 1.0, assertionFails: false},
		// Non-numeric arguments fail.
		{value: "x", target: 1.0, tolerance: 1.0, assertionFails: true},
		{value: 1.0, target: "y", tolerance: 1.0, assertionFails: true},
		{value: 1.0, target: 1.0, tolerance: "z", assertionFails: true},
		{value: nil, target: 1.0, tolerance: 1.0, assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail, if %#v is within %#v of %#v", test.value, test.tolerance, test.target)
		} else {
			testName = fmt.Sprintf("it should not fail, if %#v is within %#v of %#v", test.value, test.tolerance, test.target)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeCloseTo(test.value, test.target, test.tolerance)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
