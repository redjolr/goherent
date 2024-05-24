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

func TestCestIsRunning(t *testing.T) {
	assert := assert.New(t)

	Test("it should return true, if the Cest instance has only a TestRanEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
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

	Test("it should return false, if the Cest instance has a TestRanEvent and a TestPassedEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		cest.NewPassedEvent(
			test_passed_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(cest.IsRunning())
	}, t)

	Test("it should return false, if the Cest instance has a TestRanEvent and a TestFailedEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		cest.NewFailedEvent(
			test_failed_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(cest.IsRunning())
	}, t)

	Test("it should return false, if the Cest instance has a TestRanEvent and a TestPausedEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		cest.NewPausedEvent(
			test_paused_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.False(cest.IsRunning())
	}, t)

	Test("it should return true, if the Cest instance has TestRanEvent, TestPausedEvent, and TestContinuedEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		cest.NewPausedEvent(
			test_paused_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		cest.NewContinuedEvent(
			test_continued_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		assert.True(cest.IsRunning())
	}, t)

	Test("it should return false, if the Cest instance has a TestRanEvent and a TestSkippedEvent", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		cest.NewRanEvent(
			test_ran_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
				},
			),
		)
		cest.NewSkippedEvent(
			test_skipped_event.NewFromJsonTestEvent(
				events.JsonTestEvent{
					Time:    time.Now(),
					Package: "somePackage",
					Test:    "someTestName",
					Elapsed: 1.2,
				},
			),
		)
		assert.False(cest.IsRunning())
	}, t)
}

func TestCestHasEvent(t *testing.T) {
	assert := assert.New(t)
	Test("it should return false, if the Cest instance has no events", func(t *testing.T) {
		cest := tests_tracker.NewCest("cestName")
		evt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		assert.False(cest.HasEvent(evt))
	}, t)

	Test(`
	Given that a Cest instance has a TestRanEvent
	When we check if the Cest instance has a TestPassedEvent with HasEvent(events.Event)
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		cest := tests_tracker.NewCest("cestName")
		testRanEvt := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		cest.NewRanEvent(testRanEvt)

		// When
		checkForTestPassEvt := test_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		)
		hasTestPassEvt := cest.HasEvent(checkForTestPassEvt)

		// Then
		assert.False(hasTestPassEvt)
	}, t)

	Test(`
	Given that a Cest instance has a TestRanEvent
	When we check if the Cest instance has a TestRanEvt with different fields from the first one
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		cest := tests_tracker.NewCest("cestName")
		testRanEvt1 := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		cest.NewRanEvent(testRanEvt1)

		// When
		testRanEvt2 := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName 2",
			},
		)
		hasTestRanEvt2 := cest.HasEvent(testRanEvt2)

		// Then
		assert.False(hasTestRanEvt2)
	}, t)

	Test(`
	Given that a Cest instance has a TestRanEvent
	When we check if the Cest instance has that same TestRanEvt
	Then it should return true.
	`, func(t *testing.T) {
		// Given
		cest := tests_tracker.NewCest("cestName")
		testRanEvt1 := test_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
			},
		)
		cest.NewRanEvent(testRanEvt1)

		// When
		hasTestRanEvt := cest.HasEvent(test_ran_event.NewFromJsonTestEvent(
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
	Given that a Cest instance has a TestPassedEvent
	When we check if the Cest instance has a TestFailedEvent which has the exact same fields as the TestPassedEvt
	Then it should return false.
	`, func(t *testing.T) {
		// Given
		cest := tests_tracker.NewCest("cestName")
		testPassedvt := test_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Test:    "someTestName",
				Elapsed: 1.2,
			},
		)
		cest.NewPassedEvent(testPassedvt)

		// When
		hasTestFailedEvt := cest.HasEvent(test_failed_event.NewFromJsonTestEvent(
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
	Given that a Cest instance has a TestRanEvt, a TestPausedEvt, a TestContinuedEvt
	When we check if the Cest instance has the aforementioned events
	Then it should return true for all.
	`, func(t *testing.T) {
		// Given
		cest := tests_tracker.NewCest("cestName")
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
		cest.NewRanEvent(testRanEvt)
		cest.NewPausedEvent(testPausedEvt)
		cest.NewContinuedEvent(testContinuedEvt)

		// When
		hasTestRanEvt := cest.HasEvent(testRanEvt)
		hasTestPausedEvt := cest.HasEvent(testPausedEvt)
		hasTestContinuedEvt := cest.HasEvent(testContinuedEvt)

		// Then
		assert.True(hasTestRanEvt)
		assert.True(hasTestPausedEvt)
		assert.True(hasTestContinuedEvt)
	}, t)
}
