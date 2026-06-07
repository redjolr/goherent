package expect

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/redjolr/goherent/expect/internal/assertions"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

// tFailer is the slice of *testing.T that an expectation needs: a way to mark the
// test as failed. Storing the dependency as an interface (rather than *testing.T)
// lets the negation logic be unit-tested with a spy.
type tFailer interface {
	Fail()
}

type expectation struct {
	t                       tFailer
	checkExpectationAgainst any
	negated                 bool
}

// Not returns an expectation whose matchers assert the inverse: the test fails
// when the underlying matcher would have passed, and passes when it would have
// failed. Every matcher is negatable through this single path, so a new matcher
// gets its inverse for free — there is no need to hand-write a Not* method.
//
//	Expect(x).Not().ToEqual(y)
//	Expect(m).Not().ToHaveKey("k")
func (e *expectation) Not() *expectation {
	return &expectation{
		t:                       e.t,
		checkExpectationAgainst: e.checkExpectationAgainst,
		negated:                 !e.negated,
	}
}

// report applies the (possibly negated) outcome of a matcher. err is the matcher
// assertion's result (nil == satisfied). matcher is a short description of the
// matcher (e.g. `equal 3`) used only to phrase the message when a negated matcher
// unexpectedly held.
func (e *expectation) report(matcher string, err error) {
	satisfied := err == nil
	// Fail when the outcome doesn't match what was expected: a positive matcher
	// that wasn't satisfied, or a negated matcher that was.
	if e.negated != satisfied {
		return
	}
	file, line := callerOutsideExpectation()
	e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
	if e.negated {
		e.print(fmt.Sprintf("expected %#v not to %s", e.checkExpectationAgainst, matcher), 6)
	} else {
		e.print(err.Error(), 6)
	}
	e.t.Fail()
}

// callerOutsideExpectation walks up the stack to the first frame that is not in
// this file, so the reported location is the test's call site regardless of how
// many internal wrappers (Not, report, a Not* alias) sit in between.
func callerOutsideExpectation() (string, int) {
	for skip := 1; skip < 32; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		if !strings.HasSuffix(file, "/expect/expectation.go") {
			return file, line
		}
	}
	return "", 0
}

func (e *expectation) print(str string, leftPadWhitespace int) {
	lines := utils.SplitStringByNewLine(str)
	for _, line := range lines {
		fmt.Println(strings.Repeat(" ", leftPadWhitespace) + line)
	}
}

func (e *expectation) ToEqual(actual any) {
	e.report(fmt.Sprintf("equal %#v", actual), assertions.ToEqual(e.checkExpectationAgainst, actual))
}

func (e *expectation) ToContain(containee any) {
	e.report(fmt.Sprintf("contain %#v", containee), assertions.ToContain(e.checkExpectationAgainst, containee))
}

func (e *expectation) ToContainElement(element any) {
	e.report(fmt.Sprintf("contain element %#v", element), assertions.ToContainElement(e.checkExpectationAgainst, element))
}

func (e *expectation) ToHaveKey(key any) {
	e.report(fmt.Sprintf("have key %#v", key), assertions.ToHaveKey(e.checkExpectationAgainst, key))
}

func (e *expectation) ToMatch(pattern string) {
	e.report(fmt.Sprintf("match pattern %q", pattern), assertions.ToMatch(e.checkExpectationAgainst, pattern))
}

func (e *expectation) ToPanic() {
	e.report("panic", assertions.ToPanic(e.checkExpectationAgainst))
}

func (e *expectation) ToBeCloseTo(target any, tolerance any) {
	e.report(fmt.Sprintf("be within %v of %v", tolerance, target), assertions.ToBeCloseTo(e.checkExpectationAgainst, target, tolerance))
}

func (e *expectation) ToBeError() {
	e.report("be an error", assertions.ToBeError(e.checkExpectationAgainst))
}

func (e *expectation) ToBeTrue() {
	e.report("be true", assertions.ToBeTrue(e.checkExpectationAgainst))
}

func (e *expectation) ToBeFalse() {
	e.report("be false", assertions.ToBeFalse(e.checkExpectationAgainst))
}

func (e *expectation) ToBeNil() {
	e.report("be nil", assertions.ToBeNil(e.checkExpectationAgainst))
}

func (e *expectation) ToBeOfSameTypeAs(compareVal any) {
	e.report(fmt.Sprintf("be of the same type as %#v", compareVal), assertions.ToBeOfSameTypeAs(e.checkExpectationAgainst, compareVal))
}

func (e *expectation) ToBeString() {
	e.report("be a string", assertions.ToBeString(e.checkExpectationAgainst))
}

func (e *expectation) ToBeGreaterThan(checkIfGreaterAgainst any) {
	e.report(fmt.Sprintf("be greater than %v", checkIfGreaterAgainst), assertions.ToBeGreaterThan(e.checkExpectationAgainst, checkIfGreaterAgainst))
}

func (e *expectation) ToBeGreaterThanOrEqualTo(checkIfGreaterOrEqualAgainst any) {
	e.report(fmt.Sprintf("be greater than or equal to %v", checkIfGreaterOrEqualAgainst), assertions.ToBeGreaterThanOrEqualTo(e.checkExpectationAgainst, checkIfGreaterOrEqualAgainst))
}

func (e *expectation) ToBeLessThan(checkIfLessAgainst any) {
	e.report(fmt.Sprintf("be less than %v", checkIfLessAgainst), assertions.ToBeLessThan(e.checkExpectationAgainst, checkIfLessAgainst))
}

func (e *expectation) ToBeLessThanOrEqualTo(checkIfLessAgainst any) {
	e.report(fmt.Sprintf("be less than or equal to %v", checkIfLessAgainst), assertions.ToBeLessThanOrEqualTo(e.checkExpectationAgainst, checkIfLessAgainst))
}

func (e *expectation) ToBePositive() {
	e.report("be positive", assertions.ToBePositive(e.checkExpectationAgainst))
}

func (e *expectation) ToBeNegative() {
	e.report("be negative", assertions.ToBeNegative(e.checkExpectationAgainst))
}

func (e *expectation) ToHaveLength(length int) {
	e.report(fmt.Sprintf("have length %d", length), assertions.ToHaveLength(e.checkExpectationAgainst, length))
}

func (e *expectation) ToHaveLengthGreaterThan(length int) {
	e.report(fmt.Sprintf("have length greater than %d", length), assertions.ToHaveLengthGreaterThan(e.checkExpectationAgainst, length))
}

func (e *expectation) ToHaveLengthLessThan(length int) {
	e.report(fmt.Sprintf("have length less than %d", length), assertions.ToHaveLengthLessThan(e.checkExpectationAgainst, length))
}

// The Not* methods below are kept for backwards compatibility; each is just the
// corresponding matcher routed through the uniform Not() path.

func (e *expectation) NotToEqual(actual any) { e.Not().ToEqual(actual) }

func (e *expectation) NotToBeError() { e.Not().ToBeError() }

func (e *expectation) NotToBeNil() { e.Not().ToBeNil() }
