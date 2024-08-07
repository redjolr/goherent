package sequential_events_handler_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events_handler"
	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	"github.com/stretchr/testify/assert"
)

func setupHandlerWithBoundedTerminal(height int) (
	*sequential_events_handler.EventsHandler,
	*fake_ansi_terminal.FakeAnsiTerminal,
	*ctests_tracker.CtestsTracker,
) {
	boundedFakeAnsiTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, height)
	fakeAnsiTerminalPresenter := sequential_events_handler.NewBoundedTerminalPresenter(&boundedFakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := sequential_events_handler.NewEventsHandler(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &boundedFakeAnsiTerminal, &ctestTracker
}

func TestCtestRanEventWithBoundedTerminal(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		// When
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏳ testName",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with test name "Multiline\ntest name" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		// When
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "Multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏳ Multiline...",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 20
	When 2 CtestRanEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	Then the second CtestRanEvent should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(20)

		// When
		ctestRanEvt1 := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "testName1",
			},
		)
		ctestRanEvt2 := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "testName2",
			},
		)
		ctestRanEvt1Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		ctestRanEvt2Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// Then
		assert.NoError(ctestRanEvt1Err)
		assert.Error(ctestRanEvt2Err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)

	Test(`
	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with the same test name "testName" of package "somePackage"
	Then the user should be informed only once that the given test from the given package is running.`, func(t *testing.T) {
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		// Given
		ctestRanEvt := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏳ testName",
		)
	}, t)
}

func TestCtestPassedEventWithBoundedTerminal(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName 1" of package "somePackage" has occurred
	And a CtestPassedEvent with name "testName 1" has occurred
	And a CtestRanEvent with name "testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of with name "testName 2" of package "somePackage" occurs
	Then the user should be informed that the test has passed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 1",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 2",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ testName 1\n✅ testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestPassedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The 1st multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The second multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ The 1st multiline   \ntest name\n✅ The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestPassedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ multiline 1   \ntest name longer\n✅ multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline 1\ntest name" from "packageName"
	And a CtestPassedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestPassedEvent with test name "The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The multiline 1\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The multiline 2\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ The multiline 1\ntest name\n✅ The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestPassedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestPassedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ multiline 1\ntest name longer\n✅ multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName Line1\nLine2\nLine3",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n✅ testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestPassedEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n"+
				"✅ The 1st multiline\nLine2\nLine3   \nLine4\n"+
				"✅ The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 5
	When a CtestPassedEvent occurs with test name "testName" from "packageName"
	Then the HandleCtestPassedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 2.3

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 5
	When a CtestPassedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestPassedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage 2",
				Elapsed: &testPassedElapsedTime,
			},
		)
		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)
}

func TestCtestFailedEventWithBoundedTerminal(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName 1" of package "somePackage" has occurred
	And a CtestFailedEvent with name "testName 1" has occurred
	And a CtestRanEvent with name "testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of with name "testName 2" of package "somePackage" occurs
	Then the user should be informed that the test has failed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 2",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName 1\n❌ testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestFailedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The 1st multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The second multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ The 1st multiline   \ntest name\n❌ The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestFailedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ multiline 1   \ntest name longer\n❌ multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline 1\ntest name" from "packageName"
	And a CtestFailedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent with test name "The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The multiline 1\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The multiline 2\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ The multiline 1\ntest name\n❌ The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestFailedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ multiline 1\ntest name longer\n❌ multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName Line1\nLine2\nLine3",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestFailedEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n"+
				"❌ The 1st multiline\nLine2\nLine3   \nLine4\n"+
				"❌ The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 5
	When a CtestFailedEvent occurs with test name "testName" from "packageName"
	Then the HandleCtestFailedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 2.3

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestFailedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage 2",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)

	Test(`
	Given that 2 CtestOutputEvent for Ctest with name "testName" of package "somePackage" have occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "This is output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "This is output 2.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)
		assert.Error(err)

		// Then
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."))
	}, t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		And a CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
		And we have a bounded terminal with height 5
		When a CtestFailedEvent of the same test/package occurs
		Then a user should be informed that the Ctest has failed
		And the output from the CtestOutputEvents should be presented
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 1.2

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "This is some output.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName\nThis is some output.",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And two CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "Some output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "Some output 2.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ testName\nSome output 1.\nSome output 2.",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "NameLine1\nNameLine2\nNameLine3" of package "somePackage" has occurred
	And two CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "NameLine1\nNameLine2\nNameLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Output:  "Some output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Output:  "Some output 2.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		t.Errorf("Nope")

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n❌ NameLine1\nNameLine2   \nNameLine3\nSome output 1.\nSome output 2.",
		)
	}, t)
}

func TestCtestSkippedEventWithBoundedTerminal(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName 1" of package "somePackage" has occurred
	And a CtestSkippedEvent with name "testName 1" has occurred
	And a CtestRanEvent with name "testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of with name "testName 2" of package "somePackage" occurs
	Then the user should be informed that the test was skipped.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ testName 1\n⏩ testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestSkippedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The 1st multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The second multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ The 1st multiline   \ntest name\n⏩ The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestSkippedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ multiline 1   \ntest name longer\n⏩ multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "The multiline 1\ntest name" from "packageName"
	And a CtestSkippedEvent with test name "The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent with test name "The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The multiline 1\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The multiline 2\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ The multiline 1\ntest name\n⏩ The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "multiline 1\ntest name longer" from "packageName"
	And a CtestSkippedEvent with test name "multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent with test name "multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(2)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline 1\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "multiline 2\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ multiline 1\ntest name longer\n⏩ multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName Line1\nLine2\nLine3",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n⏩ testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestSkippedEvent with test name "The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(3)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n"+
				"⏩ The 1st multiline\nLine2\nLine3   \nLine4\n"+
				"⏩ The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	When a CtestSkippedEvent occurs with test name "testName" from "packageName"
	Then the HandleCtestSkippedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		elapsedTime := 2.3

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	When a CtestSkippedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestSkippedEvt should produce an error
	And an error should be displayed in the terminal.`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage 2",
			},
		)
		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Error(err)
		assert.True(
			strings.Contains(terminal.Text(), "❗ Error."),
		)
	}, t)
}

func TestHandleTestingStartedWithBoundedTerminal(t *testing.T) {
	assert := assert.New(t)
	Test("User should be informed, that the testing has started", func(t *testing.T) {
		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
		now := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(now)
		eventsHandler.HandleTestingStarted(testingStartedEvt)

		assert.Equal(
			terminal.Text(),
			"\n🚀 Starting...",
		)
	}, t)
}
