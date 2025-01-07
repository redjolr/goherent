package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

type someStruct struct{}

var someStructNil *someStruct = nil

func TestToBeNil(t *testing.T) {
	var tests = []struct {
		input          any
		assertionFails bool
	}{
		{input: nil, assertionFails: false},
		{input: true, assertionFails: true},
		{input: false, assertionFails: true},
		{input: 0, assertionFails: true},
		{input: 2, assertionFails: true},
		{input: 2.14, assertionFails: true},
		{input: 0.0, assertionFails: true},
		{input: "", assertionFails: true},
		{input: "some str", assertionFails: true},
		{input: "nil", assertionFails: true},
		{input: [0]int{}, assertionFails: true},
		{input: [1]int{2}, assertionFails: true},
		{input: [1]bool{true}, assertionFails: true},
		{input: [1]bool{false}, assertionFails: true},
		{input: [1]string{"nil"}, assertionFails: true},
		{input: [2]string{"nil", "str"}, assertionFails: true},
		{input: [0]float64{}, assertionFails: true},
		{input: [1]float64{2.4}, assertionFails: true},
		{input: []int{}, assertionFails: true},
		{input: []int{2}, assertionFails: true},
		{input: []bool{true}, assertionFails: true},
		{input: []bool{false}, assertionFails: true},
		{input: []any{nil}, assertionFails: true},
		{input: []any{nil}, assertionFails: true},
		{input: []string{"nil"}, assertionFails: true},
		{input: []string{"nil", "str"}, assertionFails: true},
		{input: []float64{}, assertionFails: true},
		{input: []float64{0.0}, assertionFails: true},
		{input: []float64{2.3}, assertionFails: true},
		{input: map[string]int{}, assertionFails: true},
		{input: map[string]bool{"true": true}, assertionFails: true},
		{input: map[string]bool{"true": false}, assertionFails: true},
		{input: map[bool]int{}, assertionFails: true},
		{input: map[bool]bool{true: true}, assertionFails: true},
		{input: map[bool]bool{true: false}, assertionFails: true},
		{input: map[any]any{nil: nil}, assertionFails: true},
		{input: struct{}{}, assertionFails: true},
		{input: struct{ f string }{f: "true"}, assertionFails: true},
		{input: struct{ f bool }{f: true}, assertionFails: true},
		{input: someStructNil, assertionFails: false},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf("it should fail the assertion, if we pass %v.", test.input)
		} else {
			testName = fmt.Sprintf("it should not fail the assertion, if we pass %v.", test.input)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeNil(test.input)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
