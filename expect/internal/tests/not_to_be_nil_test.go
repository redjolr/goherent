package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestNotToBeNil(t *testing.T) {
	var tests = []struct {
		input          any
		assertionFails bool
	}{
		{input: nil, assertionFails: true},
		{input: true, assertionFails: false},
		{input: false, assertionFails: false},
		{input: 0, assertionFails: false},
		{input: 2, assertionFails: false},
		{input: 2.14, assertionFails: false},
		{input: 0.0, assertionFails: false},
		{input: "", assertionFails: false},
		{input: "some str", assertionFails: false},
		{input: "nil", assertionFails: false},
		{input: [0]int{}, assertionFails: false},
		{input: [1]int{2}, assertionFails: false},
		{input: [1]bool{true}, assertionFails: false},
		{input: [1]bool{false}, assertionFails: false},
		{input: [1]string{"nil"}, assertionFails: false},
		{input: [2]string{"nil", "str"}, assertionFails: false},
		{input: [0]float64{}, assertionFails: false},
		{input: [1]float64{2.4}, assertionFails: false},
		{input: []int{}, assertionFails: false},
		{input: []int{2}, assertionFails: false},
		{input: []bool{true}, assertionFails: false},
		{input: []bool{false}, assertionFails: false},
		{input: []any{nil}, assertionFails: false},
		{input: []any{nil}, assertionFails: false},
		{input: []string{"nil"}, assertionFails: false},
		{input: []string{"nil", "str"}, assertionFails: false},
		{input: []float64{}, assertionFails: false},
		{input: []float64{0.0}, assertionFails: false},
		{input: []float64{2.3}, assertionFails: false},
		{input: map[string]int{}, assertionFails: false},
		{input: map[string]bool{"true": true}, assertionFails: false},
		{input: map[string]bool{"true": false}, assertionFails: false},
		{input: map[bool]int{}, assertionFails: false},
		{input: map[bool]bool{true: true}, assertionFails: false},
		{input: map[bool]bool{true: false}, assertionFails: false},
		{input: map[any]any{nil: nil}, assertionFails: false},
		{input: struct{}{}, assertionFails: false},
		{input: struct{ f string }{f: "true"}, assertionFails: false},
		{input: struct{ f bool }{f: true}, assertionFails: false},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if we pass %v.", test.input)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if we pass %v.", test.input)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.NotToBeNil(test.input)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
