package tests_tracker

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
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
		packageUnderTest := NewPackageUnderTest(packageName)
		testRanEvt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)
		assert.True(packageUnderTest.HasCest(NewCest(testName)))
		cest := packageUnderTest.Cest(testName)
		assert.True(cest.HasEvent(testRanEvt))
	}, t)

	// Test(`
	// Given that we have a PackageUnderTest
	// And that test has received a TestRanEvent for a test with name "someTestName"
	// When a TestPausedEvent is received for the same test
	// Then the PackageUnderTest should have only one Cest with those two events.
	// `, func(t *testing.T) {
	// 	// Given
	// 	testName := "someTestName"
	// 	packageName := "somePackageName"
	// 	packageUnderTest := tests_tracker.NewPackageUnderTest(packageName)
	// 	testRanEvt := test_ran_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Package: packageName,
	// 			Test:    testName,
	// 		},
	// 	)
	// 	packageUnderTest.NewTestRanEvent(testRanEvt)

	// 	// When
	// 	testPausedEvt := test_passed_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Package: packageName,
	// 			Test:    testName,
	// 			Elapsed: 1.2,
	// 		},
	// 	)
	// 	packageUnderTest.NewTestRanEvent(testPausedEvt)
	// }, t)
}
