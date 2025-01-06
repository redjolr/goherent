package goherent

import (
	"testing"

	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/internal"
)

func TestSkip(name string, testClosure func(Expect expect.F), t *testing.T) {
	testName := internal.EncodeGoherentTestName(name)
	t.Skip(testName, func(t *testing.T) {
		Expect := expect.New(t)
		testClosure(Expect)
	})
}
