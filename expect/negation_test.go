package expect

import "testing"

// spyT is a minimal tFailer that records whether Fail was called, so we can test
// the negation logic without failing the real test.
type spyT struct{ failed bool }

func (s *spyT) Fail() { s.failed = true }

func newExpectation(value any) (*expectation, *spyT) {
	spy := &spyT{}
	return &expectation{t: spy, checkExpectationAgainst: value}, spy
}

// TestNegationUniformity checks that Not() inverts every matcher consistently:
// a negated matcher fails exactly when its positive form would have passed.
func TestNegationUniformity(t *testing.T) {
	cases := []struct {
		name       string
		act        func() *spyT
		wantFailed bool
	}{
		// Positive matchers keep working.
		{"ToEqual passes on equal", func() *spyT {
			e, s := newExpectation(3)
			e.ToEqual(3)
			return s
		}, false},
		{"ToEqual fails on unequal", func() *spyT {
			e, s := newExpectation(3)
			e.ToEqual(4)
			return s
		}, true},

		// Not().ToEqual
		{"Not().ToEqual passes on unequal", func() *spyT {
			e, s := newExpectation(3)
			e.Not().ToEqual(4)
			return s
		}, false},
		{"Not().ToEqual fails on equal", func() *spyT {
			e, s := newExpectation(3)
			e.Not().ToEqual(3)
			return s
		}, true},

		// Double negation is the positive matcher again.
		{"Not().Not().ToEqual fails on unequal", func() *spyT {
			e, s := newExpectation(3)
			e.Not().Not().ToEqual(4)
			return s
		}, true},

		// Not().ToBeNil
		{"Not().ToBeNil passes on non-nil", func() *spyT {
			e, s := newExpectation("x")
			e.Not().ToBeNil()
			return s
		}, false},
		{"Not().ToBeNil fails on nil", func() *spyT {
			e, s := newExpectation(nil)
			e.Not().ToBeNil()
			return s
		}, true},

		// Not().ToHaveKey
		{"Not().ToHaveKey passes on missing key", func() *spyT {
			e, s := newExpectation(map[string]int{"a": 1})
			e.Not().ToHaveKey("b")
			return s
		}, false},
		{"Not().ToHaveKey fails on present key", func() *spyT {
			e, s := newExpectation(map[string]int{"a": 1})
			e.Not().ToHaveKey("a")
			return s
		}, true},

		// Not().ToContainElement
		{"Not().ToContainElement passes when absent", func() *spyT {
			e, s := newExpectation([]int{1, 2})
			e.Not().ToContainElement(3)
			return s
		}, false},
		{"Not().ToContainElement fails when present", func() *spyT {
			e, s := newExpectation([]int{1, 2})
			e.Not().ToContainElement(1)
			return s
		}, true},

		// Not().ToBeGreaterThan
		{"Not().ToBeGreaterThan passes when not greater", func() *spyT {
			e, s := newExpectation(5)
			e.Not().ToBeGreaterThan(10)
			return s
		}, false},
		{"Not().ToBeGreaterThan fails when greater", func() *spyT {
			e, s := newExpectation(5)
			e.Not().ToBeGreaterThan(2)
			return s
		}, true},

		// Not().ToMatch
		{"Not().ToMatch passes when no match", func() *spyT {
			e, s := newExpectation("hello")
			e.Not().ToMatch("^world$")
			return s
		}, false},
		{"Not().ToMatch fails when matches", func() *spyT {
			e, s := newExpectation("hello")
			e.Not().ToMatch("^hello$")
			return s
		}, true},

		// Not().ToPanic
		{"Not().ToPanic passes when no panic", func() *spyT {
			e, s := newExpectation(func() {})
			e.Not().ToPanic()
			return s
		}, false},
		{"Not().ToPanic fails when it panics", func() *spyT {
			e, s := newExpectation(func() { panic("boom") })
			e.Not().ToPanic()
			return s
		}, true},

		// Not().ToBeCloseTo
		{"Not().ToBeCloseTo passes when far", func() *spyT {
			e, s := newExpectation(3.0)
			e.Not().ToBeCloseTo(5.0, 0.1)
			return s
		}, false},
		{"Not().ToBeCloseTo fails when close", func() *spyT {
			e, s := newExpectation(3.0)
			e.Not().ToBeCloseTo(3.0, 0.1)
			return s
		}, true},

		// Not().ToBeTrue
		{"Not().ToBeTrue passes on false", func() *spyT {
			e, s := newExpectation(false)
			e.Not().ToBeTrue()
			return s
		}, false},
		{"Not().ToBeTrue fails on true", func() *spyT {
			e, s := newExpectation(true)
			e.Not().ToBeTrue()
			return s
		}, true},

		// The hand-written NotToEqual alias routes through the same path.
		{"NotToEqual passes on unequal", func() *spyT {
			e, s := newExpectation(3)
			e.NotToEqual(4)
			return s
		}, false},
		{"NotToEqual fails on equal", func() *spyT {
			e, s := newExpectation(3)
			e.NotToEqual(3)
			return s
		}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			spy := c.act()
			if spy.failed != c.wantFailed {
				t.Errorf("Fail() called = %v, want %v", spy.failed, c.wantFailed)
			}
		})
	}
}
