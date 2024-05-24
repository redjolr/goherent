package tests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_continued_event"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_paused_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
	"github.com/redjolr/goherent/cmd/tests_tracker"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCtestIsRunning(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the Ctest instance has only a TestRanEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.True(ctest.IsRunning())
	}, t)

	Test("it should return false, if the Ctest instance has a TestRanEvent and a TestPassedEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPassedEvent(
			ctest_passed_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(ctest.IsRunning())
	}, t)

	Test("it should return false, if the Ctest instance has a TestRanEvent and a TestFailedEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewFailedEvent(
			ctest_failed_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(ctest.IsRunning())
	}, t)

	Test("it should return false, if the Ctest instance has a TestRanEvent and a TestPausedEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPausedEvent(
			ctest_paused_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.False(ctest.IsRunning())
	}, t)

	Test("it should return true, if the Ctest instance has TestRanEvent, TestPausedEvent, and TestContinuedEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPausedEvent(
			ctest_paused_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		ctest.NewContinuedEvent(
			ctest_continued_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.True(ctest.IsRunning())
	}, t)

	Test("it should return false, if the Ctest instance has a TestRanEvent and a TestSkippedEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			ctest_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewSkippedEvent(
			ctest_skipped_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(ctest.IsRunning())
	}, t)
}

func TestCtestHasEvent(t *testing.T) {
	assert := assert.New(t)
	Test("it should return false, if the Ctest instance has no events", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		evt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		assert.False(ctest.HasEvent(evt))
	}, t)

	Test(`
	Given that a Ctest instance has a TestRanEvent
	When we check if the Ctest instance has a TestPassedEvent with HasEvent(events.Event)
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		ctest := tests_tracker.NewCtest("ctestName")
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt)

		// When
		checkForTestPassEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		)
		hasTestPassEvt := ctest.HasEvent(checkForTestPassEvt)

		// Then
		assert.False(hasTestPassEvt)
	}, t)

	Test(`
	Given that a Ctest instance has a TestRanEvent
	When we check if the Ctest instance has a TestRanEvt with different fields from the first one
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		ctest := tests_tracker.NewCtest("ctestName")
		testRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt1)

		// When
		testRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName 2",
			},
		)
		hasTestRanEvt2 := ctest.HasEvent(testRanEvt2)

		// Then
		assert.False(hasTestRanEvt2)
	}, t)

	Test(`
	Given that a Ctest instance has a TestRanEvent
	When we check if the Ctest instance has that same TestRanEvt
	Then it should return true.
	`, func(t *testing.T) {
		// Given
		ctest := tests_tracker.NewCtest("ctestName")
		testRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt1)

		// When
		hasTestRanEvt := ctest.HasEvent(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		))

		// Then
		assert.True(hasTestRanEvt)
	}, t)

	Test(`
	Given that a Ctest instance has a TestPassedEvent
	When we check if the Ctest instance has a TestFailedEvent which has the exact same fields as the TestPassedEvt
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		ctest := tests_tracker.NewCtest("ctestName")
		testPassedvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		)
		ctest.NewPassedEvent(testPassedvt)

		// When
		hasTestFailedEvt := ctest.HasEvent(ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		))

		// Then
		assert.False(hasTestFailedEvt)
	}, t)

	Test(`
	Given that a Ctest instance has a TestRanEvt, a TestPausedEvt, a TestContinuedEvt
	When we check if the Ctest instance has the aforementioned events
	Then it should return true for all.
	`, func(t *testing.T) {
		// Given
		ctest := tests_tracker.NewCtest("ctestName")
		testRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		testPausedEvt := ctest_paused_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now().Add(time.Millisecond),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		testContinuedEvt := ctest_continued_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now().Add(time.Millisecond * 2),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt)
		ctest.NewPausedEvent(testPausedEvt)
		ctest.NewContinuedEvent(testContinuedEvt)

		// When
		hasTestRanEvt := ctest.HasEvent(testRanEvt)
		hasTestPausedEvt := ctest.HasEvent(testPausedEvt)
		hasTestContinuedEvt := ctest.HasEvent(testContinuedEvt)

		// Then
		assert.True(hasTestRanEvt)
		assert.True(hasTestPausedEvt)
		assert.True(hasTestContinuedEvt)
	}, t)
}
