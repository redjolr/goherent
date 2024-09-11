package expect

import "testing"

func expect(t *testing.T) func(val any) *expectation {
	return func(value any) *expectation {
		theExpectation := expectation{
			t:                       t,
			checkExpectationAgainst: value,
		}
		return &theExpectation
	}
}

type F func(val any) *expectation

func New(t *testing.T) func(value any) *expectation {
	return expect(t)
}
