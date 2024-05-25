package ctests_tracker

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_paused_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPackageUt(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that we have a PackageUnderTest without Ctests
	When a CtestRanEvent is received
	Then a new Ctest should be created that contains that event
	`, func(t *testing.T) {
		testName := "someTestName"
		packageName := "somePackageName"
		packageUnderTest := NewPackageUnderTest(packageName)
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)
		assert.True(packageUnderTest.HasCtest(NewCtest(testName)))
		ctest := packageUnderTest.Ctest(testName)
		assert.True(ctest.HasEvent(testRanEvt))
	}, t)

	Test(`
	Given that we have a PackageUnderTest
	And that test has received a CtestRanEvent for a test with name "someTestName"
	When a CtestPausedEvent is received for the same test
	Then the PackageUnderTest should have only one Ctest with only those two events.
	`, func(t *testing.T) {
		// Given
		testName := "someTestName"
		packageName := "somePackageName"
		packageUnderTest := NewPackageUnderTest(packageName)
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)

		// When
		testPausedEvt := ctest_paused_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestPausedEvent(testPausedEvt)

		// Then
		assert.Equal(packageUnderTest.TestCount(), 1)
		ctest := packageUnderTest.Ctest(testName)
		assert.Equal(ctest.EventCount(), 2)
		assert.True(ctest.HasEvent(testRanEvt))
		assert.True(ctest.HasEvent(testPausedEvt))
	}, t)

	Test(`
	Given that we have a PackageUnderTest
	And that test has received a CtestRanEvent for a test with name "someTestName"
	When a CtestPassedEvent is received for the same test
	Then the PackageUnderTest should have only one Ctest with only those two events.
	`, func(t *testing.T) {
		// Given
		testName := "someTestName"
		packageName := "somePackageName"
		packageUnderTest := NewPackageUnderTest(packageName)
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)

		// When
		testPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
				Elapsed: 1.2,
			},
		)
		packageUnderTest.NewTestPassedEvent(testPassedEvt)

		// Then
		assert.Equal(packageUnderTest.TestCount(), 1)
		ctest := packageUnderTest.Ctest(testName)
		assert.Equal(ctest.EventCount(), 2)
		assert.True(ctest.HasEvent(testRanEvt))
		assert.True(ctest.HasEvent(testPassedEvt))
	}, t)

	Test(`
	Given that we have a PackageUnderTest
	And that test has received a CtestRanEvent for a test with name "someTestName"
	When a CtestFailedEvent is received for the same test
	Then the PackageUnderTest should have only one Ctest with only those two events.
	`, func(t *testing.T) {
		// Given
		testName := "someTestName"
		packageName := "somePackageName"
		packageUnderTest := NewPackageUnderTest(packageName)
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
			},
		)
		packageUnderTest.NewTestRanEvent(testRanEvt)

		// When
		testFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    testName,
				Elapsed: 1.2,
			},
		)
		packageUnderTest.NewTestFailedEvent(testFailedEvt)

		// Then
		assert.Equal(packageUnderTest.TestCount(), 1)
		ctest := packageUnderTest.Ctest(testName)
		assert.Equal(ctest.EventCount(), 2)
		assert.True(ctest.HasEvent(testRanEvt))
		assert.True(ctest.HasEvent(testFailedEvt))
	}, t)
}
