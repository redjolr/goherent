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
		testName := "someTestName"
		packageName := "somePackageName"
		packageUnderTest := tests_tracker.NewPackageUnderTest(packageName)
		testRanEvt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)
		assert.True(packageUnderTest.HasCest(testName))
		cest := packageUnderTest.Cest(testName)
		assert.True(cest.HasEvent(testRanEvt))
	}, t)
}
