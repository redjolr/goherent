package sequential_events_handler_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_output_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
	"github.com/redjolr/goherent/cmd/events/testing_finished_event"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	"github.com/redjolr/goherent/cmd/sequential_events_handler"
	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/stretchr/testify/assert"
)

func setup() (*sequential_events_handler.EventsHandler, *terminal.FakeAnsiTerminal, *ctests_tracker.CtestsTracker) {
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal()
	fakeAnsiTerminalPresenter := sequential_events_handler.NewTerminalPresenter(&fakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := sequential_events_handler.NewEventsHandler(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &fakeAnsiTerminal, &ctestTracker
}

func TestCtestRanEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have happened
	When a CtestRanEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()

		// When
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ⏳",
		)
	}, t)

	Test(`
		Given that no events have happened
		When 2 CtestRanEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
		Then the second CtestRanEvent should produce an error
		And an error should be displayed in the terminal.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()

		// When
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "testName1",
			},
		)
		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(
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
		When a CtestRanEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should be informed only once that the given test from the given package is running.
		`, func(t *testing.T) {
		eventsHandler, terminal, _ := setup()

		// Given
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(
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
			"\n\n📦 somePackage\n\n   • testName    ⏳",
		)
	}, t)
}

func TestCtestPassedEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		When a CtestPassedEvent of the same test/package occurs
		Then the user should be informed that the test has passed.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		testPassedElapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ✅\n",
		)
	}, t)

	Test(`
		Given that no events have happened
		When a CtestPassedEvent occurs with test name "testName" from "packageName"
		Then the HandleCtestPassedEvt should produce an error
		And an error should be displayed in the terminal.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		// When
		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
				Output:  "Some output",
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
		Given that a CtestRanEvent and CtestPassedEvent have occurred with test name "testName" of package "somePackage"
		When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should not be informed only once that the test has passed
		`, func(t *testing.T) {
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		// Given
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// When
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ✅\n",
		)
	}, t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		When a CtestPassedEvent of a different package "somePackage 2" occurs
		Then the HandleCtestPassedEvt should produce an error
		And an error should be displayed in the terminal.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		testPassedElapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage 2",
				Elapsed: &testPassedElapsedTime,
				Output:  "Some output",
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

func TestCtestFailedEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		When a CtestFailedEvent of the same test/package occurs
		Then the user should be informed that the test for the "somePackage" package have started
		And then that the Ctest with name "testName" has started running
		And that the Ctest with name "testName" has failed
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
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
			"\n\n📦 somePackage\n\n   • testName    ❌\n",
		)
	}, t)

	Test(`
		Given that no events have happened
		When a CtestFailedEvent occurs with test name "testName" from "packageName"
		Then the HandleCtestFailedEvt should produce an error
		And an error should be displayed in the terminal.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		// When
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
				Output:  "Some output",
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
		And later a CtestFailedEvent occurrs with test name "testName" of package "somePackage"
		When a CtestFailedEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should not be informed about the second failure, when the second event occurs
		`, func(t *testing.T) {
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		// Given
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		// When
		ctestFailedEvt2 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ❌\n",
		)
	}, t)

	Test(`
		Given that 2 CtestOutputEvent for Ctest with name "testName" of package "somePackage" have occurred
		When a CtestFailedEvent of the same test/package occurs
		Then a user should be informed that the Ctest has failed
		And the output from the CtestOutputEvents should be presented
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		ctestOutputEvt1 := ctest_output_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "This is output 1.",
			},
		)

		ctestOutputEvt2 := ctest_output_event.NewFromJsonTestEvent(
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
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
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
		When a CtestFailedEvent of the same test/package occurs
		Then a user should be informed that the Ctest has failed
		And the output from the CtestOutputEvents should be presented
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt := ctest_output_event.NewFromJsonTestEvent(
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
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
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
			"\n\n📦 somePackage\n\n   • testName    ❌\nThis is some output.",
		)
	}, t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		And two CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
		When a CtestFailedEvent of the same test/package occurs
		Then a user should be informed that the Ctest has failed
		And the output from the CtestOutputEvents should be presented
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt1 := ctest_output_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "Some output 1.",
			},
		)

		ctestOutputEvt2 := ctest_output_event.NewFromJsonTestEvent(
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
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
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
			"\n\n📦 somePackage\n\n   • testName    ❌\nSome output 1.\nSome output 2.",
		)
	}, t)
}

func TestCtestSkippedEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		When a CtestSkippedEvent of the same test/package occurs
		Then the user should be informed that the test for the "somePackage" package have started
		And then that the Ctest with name "testName" is skipped
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    "+ansi_escape.YELLOW_CIRCLE+"\n",
		)
	}, t)

	Test(`
		Given that no events have happened
		When a CtestSkippedEvent occurs with test name "testName" from "packageName"
		Then the HandleCtestSkippedEvt should produce an error
		And an error should be displayed in the terminal.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		// When
		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
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
		And later a CtestSkippedEvent occurrs with test name "testName" of package "somePackage"
		When a CtestSkippedEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should not be informed about the second skip
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 2.3

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestSkipped1Evt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkipped1Evt)

		// When
		ctestSkipped2Evt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkipped2Evt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    "+ansi_escape.YELLOW_CIRCLE+"\n",
		)
	}, t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		And a CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
		When a CtestSkippedEvent of the same test/package occurs
		Then a user should be informed that the Ctest has been skipped
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt := ctest_output_event.NewFromJsonTestEvent(
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
		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    "+ansi_escape.YELLOW_CIRCLE+"\n",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	And a CtestPassedEvent for Ctest with name "testName" of package "somePackage" has also occurred
	When a CtestSkippedEvent of the same test/package occurs
	Then the HandleCtestSkippedEvt should produce an error
	And an error should be displayed in the terminal.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// When
		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
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

func TestCtestOutputEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
		Given that there are no events
		When a CtestOutputEvent occurs for the test "testName" of package "somePackage" with output "Some output"
		Then a new package under test should be created with the the test testName
		And a new Ctest with that name should exist
		And that Ctest should have the output "Some output" stored
		`, func(t *testing.T) {
		// Given
		eventsHandler, _, ctestsTracker := setup()

		// When
		ctestOutputEvt := ctest_output_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "testName",
				Package: "somePackage",
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)

		//Then
		ctest := ctestsTracker.FindCtestWithNameInPackage("testName", "somePackage")
		assert.NotNil(ctest)
		assert.Equal(ctest.Output(), "Some output")
	}, t)
}

func TestHandleTestingStarted(t *testing.T) {
	assert := assert.New(t)
	Test("User should be informed, that the testing has started", func(t *testing.T) {
		eventsHandler, terminal, _ := setup()
		now := time.Now()
		testingStartedEvt := testing_started_event.NewTestingStartedEvent(now)
		eventsHandler.HandleTestingStarted(testingStartedEvt)

		assert.Equal(
			terminal.Text(),
			fmt.Sprintf("\n🚀 Starting... %s", now.Format("2006-01-02 15:04:05.000")),
		)
	}, t)
}

func TestHandleTestingFinished(t *testing.T) {
	assert := assert.New(t)

	Test(`
     Given that no test events have occurred
     When a TestingFinishedEvent with a duration of 1.2 seconds occurs
     Then a test summary should be presented
     And that summary should present that 0 packages have been tested, 0 tests have been run
     And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" 0 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
     Given that a Ctest with name "testName" in package "somePackage" has passed
     When a TestingFinishedEvent with a duration of 1.2 seconds occurs
     Then a test summary should be presented
     And that summary should present that there was 1 tested package in total, 1 has passed
     And 1 test was run in total and 1 has passed
     And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ✅\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that there are two Ctests with names "testName 1" and "testName 2" from the package "somePackage"
	And both those tests have passed
	When a TestingFinishedEvent with a duration of 1.2 seconds occurs
	Then a test summary should be presented
	And that summary should present that there was 1 tested package in total, 1 has passed
	And 2 tests were run in total and 2 have passed
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestPassedEvt1 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 1",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 2",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName 1    ✅\n\n   • testName 2    ✅\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are is a Ctest with names "testName 1" from the package "somePackage 1"
			And there are is a Ctest with names "testName 2" from the package "somePackage 2"
			And both those tests have passed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there were 2 tested packages in total, 2 have passed
			And 2 tests were run in total and 2 have passed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
			Output:  "Some output",
		})

		ctestPassedEvt1 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
			Output:  "Some output",
		})

		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    ✅\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    ✅\n"+

				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that a Ctest with name "testName" in package "somePackage" has failed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there was 1 tested package in total, 1 has failed
			And 1 test was run in total and 1 has failed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    ❌\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are two Ctests with names "testName 1" and "testName 2" from the package "somePackage"
			And both those tests have failed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there was 1 tested package in total, 1 has failed
			And 2 tests were run in total and 2 have failed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestFailedEvt2 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 2",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName 1    ❌\n\n   • testName 2    ❌\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are is a Ctest with names "testName 1" from the package "somePackage 1"
			And there are is a Ctest with names "testName 2" from the package "somePackage 2"
			And both those tests have passed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there were 2 tested packages in total, 2 have passed
			And 2 tests were run in total and 2 have passed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
			Output:  "Some output",
		})

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
			Output:  "Some output",
		})

		ctestFailedEvt2 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    ❌\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    ❌\n"+

				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are two Ctests with names "testName 1" and "testName 2" from the package "somePackage"
			And the first Ctest has passed and the second has failed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there was 1 tested package in total, 1 has failed
			And 2 tests were run in total, 1 has passed and 1 has failed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctest1RanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctest1PassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 1",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctest2RanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctest2FailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 2",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctest1RanEvt)
		eventsHandler.HandleCtestPassedEvt(ctest1PassedEvt)
		eventsHandler.HandleCtestRanEvt(ctest2RanEvt)
		eventsHandler.HandleCtestFailedEvt(ctest2FailedEvt)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName 1    ✅\n\n   • testName 2    ❌\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are two Ctests: "testName 1" from package "somePackage 1" and "testName 2" from package "somePackage 2"
			And the first Ctest has passed and the second has failed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there were 2 tested package in total, 1 has failed, 1 has passed
			And 2 tests were run in total, 1 has passed and 1 has failed
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		elapsedTime := 1.2
		ctest1RanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
			Output:  "Some output",
		})

		ctest1PassedEvt := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)

		ctest2RanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
			Output:  "Some output",
		})

		ctest2FailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctest1RanEvt)
		eventsHandler.HandleCtestPassedEvt(ctest1PassedEvt)
		eventsHandler.HandleCtestRanEvt(ctest2RanEvt)
		eventsHandler.HandleCtestFailedEvt(ctest2FailedEvt)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    ✅\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    ❌\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that a Ctest with name "testName" in package "somePackage" is skipped
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there was 1 tested package in total, 1 was skipped
			And 1 test was run in total and 1 was skipped
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are two Ctests with names "testName 1" and "testName 2" from the package "somePackage"
			And both those tests are skipped
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there was 1 tested package in total, 1 was skipped
			And 2 tests were run in total and 2 were skipped
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage",
			Output:  "Some output",
		})

		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage",
		})

		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage\n\n   • testName 1    "+ansi_escape.YELLOW_CIRCLE+"\n\n   • testName 2    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
			Given that there are is a Ctest with names "testName 1" from the package "somePackage 1"
			And there are is a Ctest with names "testName 2" from the package "somePackage 2"
			And both those tests have passed
			When a TestingFinishedEvent with a duration of 1.2 seconds occurs
			Then a test summary should be presented
			And that summary should present that there were 2 tested packages in total, 2 were skipped
			And 2 tests were run in total and 2 were skipped
			And the tests execution time was 1.2 seconds
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
		})

		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage 1",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
		})

		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage 2",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
		Given that there are is a Ctest with names "testName 1" from the package "somePackage 1" that is skipped
		And there are is a Ctest with names "testName 2" from the package "somePackage 2" that has passed
		When a TestingFinishedEvent with a duration of 1.2 seconds occurs
		Then a test summary should be presented
		And that summary should present that there were 2 tested packages in total, 1 was skipped and 1 has passed
		And 2 tests were run in total, of which 1 was skipped and 1 has passed
		And the tests execution time was 1.2 seconds
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		testPassedElapsed := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
		})

		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage 1",
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
		})

		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &testPassedElapsed,
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    ✅\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
		Given that there are is a Ctest with names "testName 1" from the package "somePackage 1" that has failed
		And there are is a Ctest with names "testName 2" from the package "somePackage 1" that has passed
		And there are is a Ctest with names "testName 3" from the package "somePackage 2" that is skipped
		When a TestingFinishedEvent with a duration of 1.2 seconds occurs
		Then a test summary should be presented
		And that summary should present that there were 3 tested packages in total, 1 failed, 1 skipped, and 1 passed
		And 3 tests were run in total, of which 1 failed, 1 was skipped, and 1 passed
		And the tests execution time was 1.2 seconds
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		testElapsed := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
		})

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &testElapsed,
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
		})

		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &testElapsed,
			},
		)

		ctestRanEvt3 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 3",
			Package: "somePackage 3",
		})
		ctestSkippedEvt3 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 3",
				Package: "somePackage 3",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt3)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt3)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    ❌\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    ✅\n"+
				"\n\n📦 somePackage 3\n\n   • testName 3    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 3 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 3 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
		Given that there are is a Ctest with names "testName 1" from the package "somePackage 1" that has failed
		And there are is a Ctest with names "testName 2" from the package "somePackage 2" that is skipped
		When a TestingFinishedEvent with a duration of 1.2 seconds occurs
		Then a test summary should be presented
		And that summary should present that there were 2 tested packages in total, 1 failed and 1 was skipped
		And 2 tests were run in total, of which 1 failed and 1 was skipped
		And the tests execution time was 1.2 seconds
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		testFailedElapsed := 1.2
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 1",
			Package: "somePackage 1",
		})

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &testFailedElapsed,
			},
		)

		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName 2",
			Package: "somePackage 2",
		})

		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage 2",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n\n📦 somePackage 1\n\n   • testName 1    ❌\n"+
				"\n\n📦 somePackage 2\n\n   • testName 2    "+ansi_escape.YELLOW_CIRCLE+"\n"+
				ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)
}
