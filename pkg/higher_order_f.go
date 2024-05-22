package goherent

import (
	"testing"

	"github.com/redjolr/goherent/internal"
)

func Test(name string, testClosure func(t *testing.T), t *testing.T) {
	testName := internal.EncodeGoherentTestName(name)
	t.Run(testName, testClosure)
}
