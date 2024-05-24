package tests_tracker

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

	Test("it should return false, if the Cest instance has a TestRanEvent and a TestPassedEvent", func(t *testing.T) {
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
