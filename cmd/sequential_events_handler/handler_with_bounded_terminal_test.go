package sequential_events_handler_test

import (
	"math"
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
	that the first test of that package started running
	`, func(t *testing.T) {
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
	`, func(t *testing.T) {
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
	And we have a bounded terminal with height 2
	When a CtestRanEvent occurs with test name "Multiline\ntest name" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	`, func(t *testing.T) {
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

	// Test(`
	// 	Given that no events have happened
	// 	When 2 CtestRanEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	// 	Then the second CtestRanEvent should produce an error
	// 	And an error should be displayed in the terminal.
	// 	`, func(t *testing.T) {
	// 	// Given
	// 	eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

	// 	// When
	// 	ctestRanEvt1 := events.NewCtestRanEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "run",
	// 			Package: "somePackage",
	// 			Test:    "testName1",
	// 		},
	// 	)
	// 	ctestRanEvt2 := events.NewCtestRanEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "run",
	// 			Package: "somePackage",
	// 			Test:    "testName2",
	// 		},
	// 	)
	// 	ctestRanEvt1Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
	// 	ctestRanEvt2Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

	// 	// Then
	// 	assert.NoError(ctestRanEvt1Err)
	// 	assert.Error(ctestRanEvt2Err)
	// 	assert.True(
	// 		strings.Contains(terminal.Text(), "❗ Error."),
	// 	)
	// }, t)

	// Test(`
	// 	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
	// 	When a CtestRanEvent occurs with the same test name "testName" of package "somePackage"
	// 	Then the user should be informed only once that the given test from the given package is running.
	// 	`, func(t *testing.T) {
	// 	eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)

	// 	// Given
	// 	ctestRanEvt := events.NewCtestRanEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "run",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestRanEvt(ctestRanEvt)

	// 	// When
	// 	eventsHandler.HandleCtestRanEvt(ctestRanEvt)

	// 	// Then
	// 	assert.Equal(
	// 		terminal.Text(),
	// 		"\n\n📦 somePackage\n\n   • testName    ⏳",
	// 	)
	// }, t)
}

// func TestCtestPassedEvent(t *testing.T) {
// 	assert := assert.New(t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		When a CtestPassedEvent of the same test/package occurs
// 		Then the user should be informed that the test has passed.
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		testPassedElapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 			Output:  "Some output",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// When
// 		ctestPassedEvt := events.NewCtestPassedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "pass",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &testPassedElapsedTime,
// 				Output:  "Some output",
// 			},
// 		)
// 		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ✅\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that no events have happened
// 		When a CtestPassedEvent occurs with test name "testName" from "packageName"
// 		Then the HandleCtestPassedEvt should produce an error
// 		And an error should be displayed in the terminal.
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		// When
// 		ctestPassedEvt := events.NewCtestPassedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "pass",
// 				Package: "somePackage",
// 				Test:    "testName",
// 				Elapsed: &elapsedTime,
// 				Output:  "Some output",
// 			},
// 		)
// 		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// Then
// 		assert.Error(err)
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."),
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent and CtestPassedEvent have occurred with test name "testName" of package "somePackage"
// 		When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
// 		Then the user should not be informed only once that the test has passed
// 		`, func(t *testing.T) {
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		// Given
// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 			Output:  "Some output",
// 		})

// 		ctestPassedEvt := events.NewCtestPassedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "pass",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 				Output:  "Some output",
// 			},
// 		)
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
// 		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// When
// 		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ✅\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		When a CtestPassedEvent of a different package "somePackage 2" occurs
// 		Then the HandleCtestPassedEvt should produce an error
// 		And an error should be displayed in the terminal.
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		testPassedElapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 			Output:  "Some output",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// When
// 		ctestPassedEvt := events.NewCtestPassedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "pass",
// 				Test:    "testName",
// 				Package: "somePackage 2",
// 				Elapsed: &testPassedElapsedTime,
// 				Output:  "Some output",
// 			},
// 		)
// 		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// Then
// 		assert.Error(err)
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."),
// 		)
// 	}, t)
// }

// func TestCtestFailedEvent(t *testing.T) {
// 	assert := assert.New(t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		When a CtestFailedEvent of the same test/package occurs
// 		Then the user should be informed that the test for the "somePackage" package have started
// 		And then that the Ctest with name "testName" has started running
// 		And that the Ctest with name "testName" has failed
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// When
// 		ctestFailedEvt := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ❌\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that no events have happened
// 		When a CtestFailedEvent occurs with test name "testName" from "packageName"
// 		Then the HandleCtestFailedEvt should produce an error
// 		And an error should be displayed in the terminal.
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		// When
// 		ctestFailedEvt := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Package: "somePackage",
// 				Test:    "testName",
// 				Elapsed: &elapsedTime,
// 				Output:  "Some output",
// 			},
// 		)
// 		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

// 		// Then
// 		assert.Error(err)
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."),
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		And later a CtestFailedEvent occurrs with test name "testName" of package "somePackage"
// 		When a CtestFailedEvent occurs with the same test name "testName" of package "somePackage"
// 		Then the user should not be informed about the second failure, when the second event occurs
// 		`, func(t *testing.T) {
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		// Given
// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestFailedEvt1 := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

// 		// When
// 		ctestFailedEvt2 := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ❌\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that 2 CtestOutputEvent for Ctest with name "testName" of package "somePackage" have occurred
// 		When a CtestFailedEvent of the same test/package occurs
// 		Then a user should be informed that the Ctest has failed
// 		And the output from the CtestOutputEvents should be presented
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		ctestOutputEvt1 := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "This is output 1.",
// 			},
// 		)

// 		ctestOutputEvt2 := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "This is output 2.",
// 			},
// 		)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

// 		// When
// 		ctestFailedEvt := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)
// 		assert.Error(err)

// 		// Then
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."))
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		And a CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
// 		When a CtestFailedEvent of the same test/package occurs
// 		Then a user should be informed that the Ctest has failed
// 		And the output from the CtestOutputEvents should be presented
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 1.2

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestOutputEvt := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "This is some output.",
// 			},
// 		)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)

// 		// When
// 		ctestFailedEvt := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ❌\nThis is some output.",
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		And two CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
// 		When a CtestFailedEvent of the same test/package occurs
// 		Then a user should be informed that the Ctest has failed
// 		And the output from the CtestOutputEvents should be presented
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestOutputEvt1 := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "Some output 1.",
// 			},
// 		)

// 		ctestOutputEvt2 := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "Some output 2.",
// 			},
// 		)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

// 		// When
// 		ctestFailedEvt := events.NewCtestFailedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "fail",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ❌\nSome output 1.\nSome output 2.",
// 		)
// 	}, t)
// }

// func TestCtestSkippedEvent(t *testing.T) {
// 	assert := assert.New(t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		When a CtestSkippedEvent of the same test/package occurs
// 		Then the user should be informed that the test for the "somePackage" package have started
// 		And then that the Ctest with name "testName" is skipped
// 	`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// When
// 		ctestSkippedEvt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ⏩\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that no events have happened
// 		When a CtestSkippedEvent occurs with test name "testName" from "packageName"
// 		Then the HandleCtestSkippedEvt should produce an error
// 		And an error should be displayed in the terminal.
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		// When
// 		ctestSkippedEvt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Package: "somePackage",
// 				Test:    "testName",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

// 		// Then
// 		assert.Error(err)
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."),
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		And later a CtestSkippedEvent occurrs with test name "testName" of package "somePackage"
// 		When a CtestSkippedEvent occurs with the same test name "testName" of package "somePackage"
// 		Then the user should not be informed about the second skip
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 2.3

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestSkipped1Evt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestSkippedEvt(ctestSkipped1Evt)

// 		// When
// 		ctestSkipped2Evt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestSkippedEvt(ctestSkipped2Evt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ⏩\n",
// 		)
// 	}, t)

// 	Test(`
// 		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 		And a CtestOutputEvent for Ctest with name "testName" of package "somePackage" has also occurred
// 		When a CtestSkippedEvent of the same test/package occurs
// 		Then a user should be informed that the Ctest has been skipped
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 1.2

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestOutputEvt := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "This is some output.",
// 			},
// 		)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)

// 		// When
// 		ctestSkippedEvt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

// 		// Then
// 		assert.Equal(
// 			terminal.Text(),
// 			"\n\n📦 somePackage\n\n   • testName    ⏩\n",
// 		)
// 	}, t)

// 	Test(`
// 	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
// 	And a CtestPassedEvent for Ctest with name "testName" of package "somePackage" has also occurred
// 	When a CtestSkippedEvent of the same test/package occurs
// 	Then the HandleCtestSkippedEvt should produce an error
// 	And an error should be displayed in the terminal.
// 	`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		elapsedTime := 1.2

// 		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		ctestPassedEvt := events.NewCtestPassedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "pass",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Elapsed: &elapsedTime,
// 			},
// 		)
// 		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

// 		// When
// 		ctestSkippedEvt := events.NewCtestSkippedEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "skip",
// 				Test:    "testName",
// 				Package: "somePackage",
// 			},
// 		)

// 		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

// 		// Then
// 		assert.Error(err)
// 		assert.True(
// 			strings.Contains(terminal.Text(), "❗ Error."),
// 		)
// 	}, t)

// }

// func TestCtestOutputEvent(t *testing.T) {
// 	assert := assert.New(t)

// 	Test(`
// 		Given that there are no events
// 		When a CtestOutputEvent occurs for the test "testName" of package "somePackage" with output "Some output"
// 		Then a new package under test should be created with the the test testName
// 		And a new Ctest with that name should exist
// 		And that Ctest should have the output "Some output" stored
// 		`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, _, ctestsTracker := setupHandlerWithBoundedTerminal(1)

// 		// When
// 		ctestOutputEvt := events.NewCtestOutputEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "output",
// 				Test:    "testName",
// 				Package: "somePackage",
// 				Output:  "Some output",
// 			},
// 		)
// 		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)

// 		//Then
// 		ctest := ctestsTracker.FindCtestWithNameInPackage("testName", "somePackage")
// 		assert.NotNil(ctest)
// 		assert.Equal(ctest.Output(), "Some output")
// 	}, t)
// }

// func TestHandleTestingStarted(t *testing.T) {
// 	assert := assert.New(t)
// 	Test("User should be informed, that the testing has started", func(t *testing.T) {
// 		eventsHandler, terminal, _ := setupHandlerWithBoundedTerminal(1)
// 		now := time.Now()
// 		testingStartedEvt := events.NewTestingStartedEvent(now)
// 		eventsHandler.HandleTestingStarted(testingStartedEvt)

// 		assert.Equal(
// 			terminal.Text(),
// 			fmt.Sprintf("\n🚀 Starting..."),
// 		)
// 	}, t)
// }
