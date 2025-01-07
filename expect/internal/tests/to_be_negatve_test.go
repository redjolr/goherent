package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeNegative(t *testing.T) {
	var tests = []struct {
		isNegativeCandidate any
		assertionFails      bool
	}{
		{isNegativeCandidate: 1, assertionFails: true},
		{isNegativeCandidate: 0, assertionFails: true},
		{isNegativeCandidate: -1, assertionFails: false},
		{isNegativeCandidate: int16(1), assertionFails: true},
		{isNegativeCandidate: int16(0), assertionFails: true},
		{isNegativeCandidate: int16(-1), assertionFails: false},
		{isNegativeCandidate: 2.3, assertionFails: true},
		{isNegativeCandidate: 0.0, assertionFails: true},
		{isNegativeCandidate: -0.1, assertionFails: false},
		{isNegativeCandidate: true, assertionFails: true},
		{isNegativeCandidate: false, assertionFails: true},
		{isNegativeCandidate: "b", assertionFails: true},
		{isNegativeCandidate: "-1", assertionFails: true},
		{isNegativeCandidate: "", assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that %v(%v) is not negative",
				test.isNegativeCandidate,
				reflect.TypeOf(test.isNegativeCandidate),
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that %v(%v) is negative",
				test.isNegativeCandidate,
				reflect.TypeOf(test.isNegativeCandidate),
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeNegative(test.isNegativeCandidate)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
