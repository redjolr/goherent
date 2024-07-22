package cmd_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	"github.com/redjolr/goherent/console"
	"github.com/redjolr/goherent/console/coordinates"
	"github.com/redjolr/goherent/console/cursor"
	"github.com/redjolr/goherent/console/terminal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func setup() (*cmd.EventsHandler, *terminal.FakeAnsiTerminal) {
	terminalOrigin := coordinates.Origin()
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal(&terminalOrigin)
	cursor := cursor.NewCursor(&fakeAnsiTerminal, &terminalOrigin)
	cons := console.NewConsole(&fakeAnsiTerminal, &cursor)

	fakeAnsiTerminalPresenter := cmd.NewTerminalPresenter(&cons)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := cmd.NewEventsHandler(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &fakeAnsiTerminal
}

func TestCtestPassedEvent(t *testing.T) {
	assert := assert.New(t)
	Test(`
		Given that no events have happened
		When a CtestPassedEvent occurs with test name "testName" from "packageName"
		Then the user should be informed that the testing of a new package started and
		that the first test of that package passed
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal := setup()
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
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t✅ testName\n\t  2.30s",
		)
	}, t)

	Test(`
		Given that no events have happened
		When 2 CtestPassedEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
		Then the user should be informed about both tests that have passed
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal := setup()
		elapsedTime1, elapsedTime2 := 2.3, 1.2

		ctestPassedEvt1 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName1",
				Elapsed: &elapsedTime1,
				Output:  "Some output",
			},
		)
		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName2",
				Elapsed: &elapsedTime2,
				Output:  "Some output",
			},
		)

		// When
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t✅ testName1\n\t  2.30s\n\t✅ testName2\n\t  1.20s",
		)
	}, t)

	Test(`
		Given that a CtestPassedEvent has occurred with test name "testName" of package "somePackage"
		When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should not be informed about the second passing, when the second event occurs
		`, func(t *testing.T) {
		eventsHandler, terminal := setup()
		elapsedTime := 2.3

		// Given
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
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// When
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t✅ testName\n\t  2.30s",
		)
	}, t)

	Test(`
		Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
		When a CtestPassedEvent of the same test/package occurs
		Then the user should be informed that the test has passed.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal := setup()
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
			"📦⏳ somePackage\n\t✅ testName\n\t  2.30s",
		)
	}, t)
}

func TestCtestFailedEvent(t *testing.T) {
	assert := assert.New(t)
	Test(`
		Given that no events have happened
		When a CtestFailedEvent occurs with test name "testName" from "packageName"
		Then the user should be informed that the testing of a new package started and
		that the first test of that package failed.
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal := setup()
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
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t❌ testName\n\t  2.30s",
		)
	}, t)

	Test(`
		Given that no events have happened
		When 2 CtestPassedFailed of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
		Then the user should be informed about both tests that have failed
		`, func(t *testing.T) {
		// Given
		eventsHandler, terminal := setup()
		elapsedTime1, elapsedTime2 := 2.3, 1.2

		// When
		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Package: "somePackage",
				Test:    "testName1",
				Elapsed: &elapsedTime1,
			},
		)
		ctestFailedEvt2 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName2",
				Package: "somePackage",
				Elapsed: &elapsedTime2,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t❌ testName1\n\t  2.30s\n\t❌ testName2\n\t  1.20s",
		)
	}, t)

	Test(`
		Given that a CtestFailedEvent has occurred with test name "testName" of package "somePackage"
		When a CtestFailedEvent occurs with the same test name "testName" of package "somePackage"
		Then the user should not be informed about the second failure, when the second event occurs
		`, func(t *testing.T) {
		eventsHandler, terminal := setup()
		elapsedTime := 2.3

		// Given
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

		// When
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"📦⏳ somePackage\n\t❌ testName\n\t  2.30s",
		)
	}, t)

	// Test(`
	// 	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	// 	When a CtestFailedEvent of the same test/package occurs
	// 	Then the user should be informed that the test for the "somePackage" package have started
	// 	And then that the Ctest with name "testName" has started running
	// 	And that the Ctest with name "testName" has failed
	// 	`, func(t *testing.T) {
	// 	// Given
	// 	eventsHandler, outputPortMock, _ := setup()
	// 	elapsedTime := 1.2

	// 	ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
	// 		Time:    time.Now(),
	// 		Action:  "run",
	// 		Test:    "testName",
	// 		Package: "somePackage",
	// 	})
	// 	eventsHandler.HandleCtestRanEvt(ctestRanEvt)

	// 	// When
	// 	ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "fail",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Elapsed: &elapsedTime,
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

	// 	// Then
	// 	outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
	// 	outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName")
	// 	outputPortMock.AssertCalled(t, "CtestFailed", "testName", elapsedTime)
	// }, t)

	// Test(`
	// 	Given that 2 CtestOutputEvent for Ctest with name "testName" of package "somePackage" have occurred
	// 	When a CtestFailedEvent of the same test/package occurs
	// 	Then a user should be informed that the testing of a new package started
	// 	And that the Ctest has failed
	// 	And the output from the CtestOutputEvents should be presented
	// 	`, func(t *testing.T) {
	// 	// Given
	// 	eventsHandler, outputPortMock, _ := setup()
	// 	elapsedTime := 1.2

	// 	ctestOutputEvt1 := ctest_output_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "output",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Output:  "This is output 1.",
	// 		},
	// 	)

	// 	ctestOutputEvt2 := ctest_output_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "output",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Output:  "This is output 2.",
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
	// 	eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

	// 	// When
	// 	ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "fail",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Elapsed: &elapsedTime,
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

	// 	// Then
	// 	outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
	// 	outputPortMock.AssertCalled(t, "CtestFailed", "testName", elapsedTime)
	// 	outputPortMock.AssertCalled(t, "CtestOutput", "testName", "somePackage", "This is output 1.\nThis is output 2.")
	// }, t)

	// Test(`
	// 	Given that 2 CtestOutputEvent for Ctest with name "testName" of package "somePackage" have occurred
	// 	When a CtestFailedEvent of the same test/package occurs
	// 	Then a user should be informed that the testing of a new package started
	// 	And that the Ctest has failed
	// 	And the output from the CtestOutputEvents should be presented
	// 	`, func(t *testing.T) {
	// 	// Given
	// 	eventsHandler, outputPortMock, _ := setup()
	// 	elapsedTime := 1.2

	// 	ctestOutputEvt1 := ctest_output_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "output",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Output:  "This is output 1.",
	// 		},
	// 	)

	// 	ctestOutputEvt2 := ctest_output_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "output",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Output:  "This is output 2.",
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
	// 	eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

	// 	// When
	// 	ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "fail",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Elapsed: &elapsedTime,
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

	// 	// Then
	// 	outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
	// 	outputPortMock.AssertCalled(t, "CtestFailed", "testName", elapsedTime)
	// 	outputPortMock.AssertNumberOfCalls(t, "CtestFailed", 1)
	// 	outputPortMock.AssertCalled(t, "CtestOutput", "testName", "somePackage", "This is output 1.\nThis is output 2.")
	// }, t)

	// Test(`
	// 	Given that no events have happened
	// 	When these events occurr for a Ctest with name "testName" from package "somePackage":
	// 		- 1 CtestRanEvent
	// 		- 1 CtestOutputEvent
	// 		- 1 CtestFailedEvent
	// 	Then the user should be be presented with this information:
	// 		- Test started running
	// 		- First test of package failed
	// 		- The output from the CtestOutputEvent
	// 	`, func(t *testing.T) {
	// 	// Given
	// 	eventsHandler, outputPortMock, _ := setup()
	// 	elapsedTime := 1.2

	// 	// When
	// 	ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
	// 		Time:    time.Now(),
	// 		Action:  "run",
	// 		Test:    "testName",
	// 		Package: "somePackage",
	// 	})
	// 	ctestOutputEvt := ctest_output_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "output",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Output:  "This is some output.",
	// 		},
	// 	)
	// 	ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
	// 		events.JsonTestEvent{
	// 			Time:    time.Now(),
	// 			Action:  "fail",
	// 			Test:    "testName",
	// 			Package: "somePackage",
	// 			Elapsed: &elapsedTime,
	// 		},
	// 	)
	// 	eventsHandler.HandleCtestRanEvt(ctestRanEvt)
	// 	eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)
	// 	eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

	// 	// Then
	// 	outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
	// 	outputPortMock.AssertCalled(t, "CtestFailed", "testName", elapsedTime)
	// 	outputPortMock.AssertCalled(t, "CtestOutput", "testName", "somePackage", "This is some output.")
	// }, t)
}

// func TestCtestRanEvent(t *testing.T) {
// 	Test(`
// 	Given that no events have happened
// 	When a CtestRanEvent occurs with test name "testName" from "packageName"
// 	Then the user should be informed that the testing of a new package started and
// 	that the first test of that package started running
// 	`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, outputPortMock, _ := setup()

// 		// When
// 		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
// 			Time:    time.Now(),
// 			Action:  "run",
// 			Test:    "testName",
// 			Package: "somePackage",
// 		})
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// Then
// 		outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
// 		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName")
// 	}, t)

// 	Test(`
// 	Given that no events have happened
// 	When 2 CtestRanEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
// 	Then the user should be informed about both tests that have started running
// 	And that "testName1" was the first test of its package
// 	`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, outputPortMock, _ := setup()

// 		// When
// 		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "run",
// 				Package: "somePackage",
// 				Test:    "testName1",
// 			},
// 		)
// 		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "run",
// 				Package: "somePackage",
// 				Test:    "testName2",
// 			},
// 		)
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

// 		// Then
// 		outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
// 		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName1")
// 		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName2")
// 	}, t)

// 	Test(`
// 	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
// 	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
// 	Then the user should not be informed about the second run, when the second event occurs
// 	`, func(t *testing.T) {
// 		eventsHandler, outputPortMock, _ := setup()

// 		// Given
// 		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "run",
// 				Test:    "testName",
// 				Package: "somePackage",
// 			},
// 		)
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// When
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// Then
// 		outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
// 		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName")
// 		outputPortMock.AssertNumberOfCalls(t, "CtestStartedRunning", 1)
// 	}, t)

// 	Test(`
// 	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
// 	And then a CtestPassedEvent has occurred for the same test
// 	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
// 	Then the user should not be informed about the second run, when the second event occurs
// 	`, func(t *testing.T) {
// 		eventsHandler, outputPortMock, _ := setup()
// 		elapsedTime := 2.3

// 		// Given
// 		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(
// 			events.JsonTestEvent{
// 				Time:    time.Now(),
// 				Action:  "run",
// 				Test:    "testName",
// 				Package: "somePackage",
// 			},
// 		)
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
// 		ctestPassedEvt := ctest_passed_event.NewFromJsonTestEvent(
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
// 		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

// 		// Then
// 		outputPortMock.AssertCalled(t, "PackageTestsStartedRunning", "somePackage")
// 		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName")
// 		outputPortMock.AssertNumberOfCalls(t, "CtestStartedRunning", 1)
// 	}, t)
// }

// func TestCtestOutputEvent(t *testing.T) {
// 	assert := assert.New(t)

// 	Test(`
// 	Given that there are no events
// 	When a CtestOutputEvent occurs for the test "testName" of package "somePackage"
// 	Then a new package under test should be created with the the test testName
// 	`, func(t *testing.T) {
// 		// Given
// 		eventsHandler, _, ctestTracker := setup()

// 		// When
// 		ctestOutputEvt := ctest_output_event.NewFromJsonTestEvent(
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
// 		ctest := ctestTracker.FindCtestWithNameInPackage("testName", "somePackage")
// 		assert.NotNil(ctest)
// 	}, t)
// }

func TestHandleTestingStarted(t *testing.T) {
	assert := assert.New(t)
	Test("User should be informed, that the testing has started", func(t *testing.T) {
		eventsHandler, terminal := setup()
		now := time.Now()
		testingStartedEvt := testing_started_event.NewTestingStartedEvent(now)
		eventsHandler.HandleTestingStarted(testingStartedEvt)

		assert.Equal(
			terminal.Text(),
			fmt.Sprintf("\n🚀 Starting... %s\n\n", now.Format("2006-01-02 15:04:05.000")),
		)
	}, t)
}
