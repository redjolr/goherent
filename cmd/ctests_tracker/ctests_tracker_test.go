package ctests_tracker

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewCtestRanEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given an empty CtestsTracker
	When a CtestRanEvent for a certain test in a certain package is received
	Then the CtestsTracker should contain that package 
	`, func(t *testing.T) {
		tracker := NewCtestsTracker()

		packageName := "somePackageName"
		ctestRanEvent := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    "someTestName",
			},
		)

		tracker.NewCtestRanEvent(ctestRanEvent)
		assert.True(tracker.ContainsPackageUtWithName(packageName))
	}, t)
}
