package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBePositive(t *testing.T) {
	var tests = []struct {
		isPositiveCandidate any
		assertionFails      bool
	}{
		{isPositiveCandidate: 1, assertionFails: false},
		{isPositiveCandidate: 0, assertionFails: true},
		{isPositiveCandidate: -1, assertionFails: true},
		{isPositiveCandidate: int16(1), assertionFails: false},
		{isPositiveCandidate: int16(0), assertionFails: true},
		{isPositiveCandidate: int16(-1), assertionFails: true},
		{isPositiveCandidate: 2.3, assertionFails: false},
		{isPositiveCandidate: 0.0, assertionFails: true},
		{isPositiveCandidate: -0.1, assertionFails: true},
		{isPositiveCandidate: true, assertionFails: true},
		{isPositiveCandidate: false, assertionFails: true},
		{isPositiveCandidate: "b", assertionFails: false},
		{isPositiveCandidate: "-1", assertionFails: false},
		{isPositiveCandidate: "", assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is not positive",
				test.isPositiveCandidate,
				reflect.TypeOf(test.isPositiveCandidate),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is positive",
				test.isPositiveCandidate,
				reflect.TypeOf(test.isPositiveCandidate),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBePositive(test.isPositiveCandidate)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
