package expect

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

type expectation struct {
	t                       *testing.T
	checkExpectationAgainst any
}

func (e *expectation) ToEqual(actual any) {
	if err := assertions.ToEqual(e.checkExpectationAgainst, actual); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToContain(containee any) {
	if err := assertions.ToContain(e.checkExpectationAgainst, containee); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToBeError() {
	if err := assertions.ToBeError(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) NotToBeError() {
	if err := assertions.NotToBeError(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToBeTrue() {
	if err := assertions.ToBeTrue(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToBeFalse() {
	if err := assertions.ToBeFalse(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToBeNil() {
	if err := assertions.ToBeNil(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) NotToBeNil() {
	if err := assertions.NotToBeNil(e.checkExpectationAgainst); err != nil {
		e.t.Errorf(err.Error())
	}
}

func (e *expectation) ToBeOfSameTypeAs(compareVal any) {
	if err := assertions.ToBeOfSameTypeAs(e.checkExpectationAgainst, compareVal); err != nil {
		e.t.Errorf(err.Error())
	}
}
