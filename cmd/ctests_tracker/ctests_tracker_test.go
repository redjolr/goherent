package ctests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
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
	Then the CtestsTracker should contain that PackageUnderTest
	`, func(t *testing.T) {
		tracker := ctests_tracker.NewCtestsTracker()

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

func TestInsertCtest(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that we have an empty TestsTracker
	When we call the InsertCtest() method with a Ctest as an argument
	The test will be added and returned from the method
	And a package with the name of the Ctest's packageName will be added to the tracker
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		// When
		ctest := ctests_tracker.NewCtest("testName", "packageName")
		ctestReturned := tracker.InsertCtest(ctest)

		// Then
		assert.Equal(ctestReturned, ctest)
		assert.True(tracker.ContainsCtestWithName("testName"))
		assert.True(tracker.ContainsPackageUtWithName("packageName"))
	}, t)
}
