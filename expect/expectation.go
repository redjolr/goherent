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
