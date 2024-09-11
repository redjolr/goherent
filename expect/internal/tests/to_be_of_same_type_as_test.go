package tests_test

import (
	"fmt"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeOfSameTypeAs(t *testing.T) {
	var tests = []struct {
		type1          any
		type2          any
		assertionFails bool
	}{
		// bool
		{type1: true, type2: true, assertionFails: false},
		{type1: true, type2: false, assertionFails: false},
		{type1: false, type2: true, assertionFails: false},
		{type1: true, type2: 5, assertionFails: true},
		{type1: 5, type2: false, assertionFails: true},
		{type1: false, type2: 3, assertionFails: true},
		{type1: false, type2: 3.6, assertionFails: true},
		{type1: false, type2: "str", assertionFails: true},
		{type1: true, type2: [2]int{1, 2}, assertionFails: true},
		{type1: true, type2: []int{1, 2}, assertionFails: true},
		{type1: true, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: true, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// int
		{type1: 5, type2: 6, assertionFails: false},
		{type1: int32(5), type2: int64(5), assertionFails: true},
		{type1: int32(5), type2: uint32(5), assertionFails: true},
		{type1: 4, type2: 3.4, assertionFails: true},
		{type1: 8, type2: true, assertionFails: true},
		{type1: 5, type2: false, assertionFails: true},
		{type1: 5, type2: "str", assertionFails: true},
		{type1: 2, type2: [2]int{1, 2}, assertionFails: true},
		{type1: 3, type2: []int{1, 2}, assertionFails: true},
		{type1: 2, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: 4, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// float64
		{type1: 5.5, type2: 9.2, assertionFails: false},
		{type1: float32(3.2), type2: float64(3.2), assertionFails: true},
		{type1: 3.2, type2: 2, assertionFails: true},
		{type1: 5.2, type2: true, assertionFails: true},
		{type1: 5.6, type2: "str", assertionFails: true},
		{type1: 2.2, type2: [2]int{1, 2}, assertionFails: true},
		{type1: 3.5, type2: []int{1, 2}, assertionFails: true},
		{type1: 2.7, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: 4.7, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// string
		{type1: "str", type2: "other str", assertionFails: false},
		{type1: "abc", type2: 2, assertionFails: true},
		{type1: "abc", type2: 2.6, assertionFails: true},
		{type1: "abc", type2: true, assertionFails: true},
		{type1: "abc", type2: [2]int{1, 2}, assertionFails: true},
		{type1: "abc", type2: []int{1, 2}, assertionFails: true},
		{type1: "abc", type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: "abc", type2: struct{ f string }{f: "val"}, assertionFails: true},

		// array
		{type1: [1]int{7}, type2: [1]int{12}, assertionFails: false},
		{type1: [1]int{7}, type2: [2]int{7, 5}, assertionFails: true},
		{type1: [1]int{7}, type2: 2, assertionFails: true},
		{type1: [1]int{7}, type2: 2.6, assertionFails: true},
		{type1: [1]int{7}, type2: true, assertionFails: true},
		{type1: [1]int{7}, type2: "str", assertionFails: true},
		{type1: [2]int{1, 2}, type2: []int{1, 2}, assertionFails: true},
		{type1: [1]int{7}, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: [1]int{7}, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// slice
		{type1: []int{7, 2}, type2: []int{8, 3}, assertionFails: false},
		{type1: []int{7, 2}, type2: []int{8, 3, 5, 6}, assertionFails: false},
		{type1: []int{7}, type2: [2]int{7, 5}, assertionFails: true},
		{type1: []int{7}, type2: 2, assertionFails: true},
		{type1: []int{7}, type2: 2.6, assertionFails: true},
		{type1: []int{7}, type2: true, assertionFails: true},
		{type1: []int{7}, type2: "str", assertionFails: true},
		{type1: []int{7}, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: []int{7}, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// map
		{type1: map[int]string{12: "str"}, type2: map[int]string{25: "asd"}, assertionFails: false},
		{type1: map[int]string{12: "str"}, type2: [2]int{7, 5}, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: 2, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: 2.6, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: true, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: "str", assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: []int{1, 2}, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: map[int]string{12: "str"}, type2: struct{ f string }{f: "val"}, assertionFails: true},

		// struct
		{type1: struct{ f string }{f: "abc"}, type2: struct{ f string }{f: "str"}, assertionFails: false},
		{type1: struct{ f string }{f: "abc"}, type2: [2]int{7, 5}, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: 2, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: 2.6, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: true, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: "str", assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: []int{1, 2}, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: map[string]int{"k": 1, "k2": 2}, assertionFails: true},
		{type1: struct{ f string }{f: "abc"}, type2: struct{ x string }{x: "val"}, assertionFails: true},
	}

	for _, test := range tests {
		var testName string
		if test.assertionFails {
			testName = fmt.Sprintf(
				"it should fail the assertion, if we assert that type of %v is same as type %v",
				test.type1,
				test.type2,
			)
		} else {
			testName = fmt.Sprintf(
				"it should not fail the assertion, if we assert that type of %v is same as type %v",
				test.type1,
				test.type2,
			)
		}
		t.Run(testName, func(t *testing.T) {
			assertionErr := assertions.ToBeOfSameTypeAs(test.type1, test.type2)
			if (assertionErr == nil) == test.assertionFails {
				t.Errorf("%v", assertionErr)
			}
		})
	}
}
