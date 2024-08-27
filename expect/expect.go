package expect

import "testing"

func expect(t *testing.T) func(val any) *expectation {
	return func(value any) *expectation {
		theExpectation := expectation{
			t:        t,
			expected: value,
		}
		return &theExpectation
	}
}

func New(t *testing.T) func(value any) *expectation {
	return expect(t)
}
