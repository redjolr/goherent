package tests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_continued_event"
	"github.com/redjolr/goherent/cmd/events/test_failed_event"
	"github.com/redjolr/goherent/cmd/events/test_passed_event"
	"github.com/redjolr/goherent/cmd/events/test_paused_event"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
	"github.com/redjolr/goherent/cmd/events/test_skipped_event"
	"github.com/redjolr/goherent/cmd/tests_tracker"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCtestIsRunning(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the Ctest instance has only a TestRanEvent", func(t *testing.T) {
		ctest := tests_tracker.NewCtest("ctestName")
		ctest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
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
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPassedEvent(
			test_passed_event.NewFromJsonTestEvent(
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
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewFailedEvent(
			test_failed_event.NewFromJsonTestEvent(
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
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPausedEvent(
			test_paused_event.NewFromJsonTestEvent(
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
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewPausedEvent(
			test_paused_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		ctest.NewContinuedEvent(
			test_continued_event.NewFromJsonTestEvent(
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
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		ctest.NewSkippedEvent(
			test_skipped_event.NewFromJsonTestEvent(
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
		evt := test_ran_event.NewFromJsonTestEvent(
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
		testRanEvt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt)

		// When
		checkForTestPassEvt := test_passed_event.NewFromJsonTestEvent(
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
		testRanEvt1 := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt1)

		// When
		testRanEvt2 := test_ran_event.NewFromJsonTestEvent(
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
		testRanEvt1 := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		ctest.NewRanEvent(testRanEvt1)

		// When
		hasTestRanEvt := ctest.HasEvent(test_ran_event.NewFromJsonTestEvent(
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
		testPassedvt := test_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		)
		ctest.NewPassedEvent(testPassedvt)

		// When
		hasTestFailedEvt := ctest.HasEvent(test_failed_event.NewFromJsonTestEvent(
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
		testRanEvt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		testPausedEvt := test_paused_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now().Add(time.Millisecond),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		testContinuedEvt := test_continued_event.NewFromJsonTestEvent(
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
