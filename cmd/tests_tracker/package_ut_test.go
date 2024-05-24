package tests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
	"github.com/redjolr/goherent/cmd/tests_tracker"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPackageUt(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that we have a PackageUnderTest without Cests
	When a TestRanEvent is received
	Then a new Cest should be created that contains that event
	`, func(t *testing.T) {
		packageUnderTest := tests_tracker.NewPackageUnderTest("somePackageName")
		packageUnderTest.NewTestRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.True(packageUnderTest.HasCest("someTestName"))
	}, t)
}
