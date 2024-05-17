package goherent

import (
	"testing"

	"github.com/redjolr/goherent/pkg/internal"
)

func Test(message string, testClosure func(t *testing.T), t *testing.T) {
	testName := internal.EncodeGoherentTestMessage(message)
	t.Run(testName, testClosure)
}
