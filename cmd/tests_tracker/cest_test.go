package tests_tracker

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCestIsRunning(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the Cest instance has only a TestRanEvent", func(t *testing.T) {
		cest := NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.True(cest.IsRunning())
	}, t)
}
