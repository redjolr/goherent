package expect

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

type expectation struct {
	t        *testing.T
	expected any
}

func (e *expectation) ToEqual(actual any) {
	if err := assertions.ToEqual(e.expected, actual); err != nil {
		e.t.Errorf(err.Error())
	}
}
