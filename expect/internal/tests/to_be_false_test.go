package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeFalse(t *testing.T) {
	var tests = []struct {
		input          any
		assertionFails bool
	}{
		{input: false, assertionFails: false},
		{input: true, assertionFails: true},
		{input: 2, assertionFails: true},
		{input: 2.14, assertionFails: true},
		{input: nil, assertionFails: true},
		{input: "", assertionFails: true},
		{input: "some str", assertionFails: true},
		{input: "true", assertionFails: true},
		{input: [0]int{}, assertionFails: true},
		{input: [1]int{2}, assertionFails: true},
		{input: [1]bool{true}, assertionFails: true},
		{input: [1]bool{false}, assertionFails: true},
		{input: [1]string{"false"}, assertionFails: true},
		{input: [2]string{"false", "str"}, assertionFails: true},
		{input: [0]float64{}, assertionFails: true},
		{input: [1]float64{2.4}, assertionFails: true},
		{input: []int{}, assertionFails: true},
		{input: []int{2}, assertionFails: true},
		{input: []bool{true}, assertionFails: true},
		{input: []bool{false}, assertionFails: true},
		{input: []string{"false"}, assertionFails: true},
		{input: []string{"false", "str"}, assertionFails: true},
		{input: []float64{}, assertionFails: true},
		{input: []float64{2.4}, assertionFails: true},
		{input: map[string]int{}, assertionFails: true},
		{input: map[string]bool{"false": true}, assertionFails: true},
		{input: map[string]bool{"false": false}, assertionFails: true},
		{input: map[bool]int{}, assertionFails: true},
		{input: map[bool]bool{false: true}, assertionFails: true},
		{input: map[bool]bool{false: false}, assertionFails: true},
		{input: struct{}{}, assertionFails: true},
		{input: struct{ f string }{f: "false"}, assertionFails: true},
		{input: struct{ f bool }{f: false}, assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if we pass %v.", test.input)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if we pass %v.", test.input)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeFalse(test.input)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}

}
