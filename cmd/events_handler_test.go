package cmd_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/mock"
)

func setup() (cmd.EventsHandler, *cmd.OutputPortMock) {
	outputPortMock := cmd.NewOutputPortMock()
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventHandler := cmd.NewEventsHandler(outputPortMock, &ctestTracker)
	outputPortMock.On("CtestPassed", mock.Anything, mock.Anything).Return()
	outputPortMock.On("CtestFailed", mock.Anything, mock.Anything).Return()
	outputPortMock.On("FirstCtestOfPackagePassed", mock.Anything, mock.Anything, mock.Anything).Return()
	outputPortMock.On("FirstCtestOfPackageFailed", mock.Anything, mock.Anything, mock.Anything).Return()
	outputPortMock.On("FirstCtestOfPackageStartedRunning", mock.Anything, mock.Anything).Return()
	outputPortMock.On("CtestStartedRunning", mock.Anything).Return()

	return eventHandler, outputPortMock
}

func TestCtestPassedEvent(t *testing.T) {
	Test(`
	Given that no events have happened
	When a CtestPassedEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package passed
	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()
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
		outputPortMock.AssertCalled(t, "FirstCtestOfPackagePassed", "testName", "somePackage", elapsedTime)
	}, t)

	Test(`
	Given that no events have happened
	When 2 CtestPassedEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	Then the user should be informed about both tests that have passed
	And that "testName1" was the first test of its package
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
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
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		outputPortMock.AssertCalled(t, "FirstCtestOfPackagePassed", "testName1", "somePackage", elapsedTime1)
		outputPortMock.AssertCalled(t, "CtestPassed", "testName2", elapsedTime2)

	}, t)

	Test(`
	Given that a CtestPassedEvent has occurred with test name "testName" of package "somePackage"
	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
	Then the user should not be informed about the second passing, when the second event occurs
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
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
		outputPortMock.AssertCalled(t, "FirstCtestOfPackagePassed", "testName", "somePackage", elapsedTime)
		outputPortMock.AssertNumberOfCalls(t, "FirstCtestOfPackagePassed", 1)
		outputPortMock.AssertNumberOfCalls(t, "CtestPassed", 0)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed.

	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()
		testPassedElapsedTime := 1.2

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
		outputPortMock.AssertCalled(t, "CtestPassed", "testName", testPassedElapsedTime)

	}, t)
}

func TestCtestFailedEvent(t *testing.T) {
	Test(`
	Given that no events have happened
	When a CtestFailedEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package failed.
	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

		// When
		ctestFailedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageFailed", "testName", "somePackage", elapsedTime)
	}, t)

	Test(`
	Given that no events have happened
	When 2 CtestPassedFailed of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	Then the user should be informed about both tests that have failed
	And that "testName1" was the first test of its package
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime1, elapsedTime2 := 2.3, 1.2

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName1",
				Elapsed: &elapsedTime1,
				Output:  "Some output",
			},
		)
		ctestFailedEvt2 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName2",
				Package: "somePackage",
				Elapsed: &elapsedTime2,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		outputPortMock.AssertCalled(t, "FirstCtestOfPackageFailed", "testName1", "somePackage", elapsedTime1)
		outputPortMock.AssertCalled(t, "CtestFailed", "testName2", elapsedTime2)

	}, t)

	Test(`
	Given that a CtestFailedEvent has occurred with test name "testName" of package "somePackage"
	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
	Then the user should not be informed about the second passing, when the second event occurs
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

		// Given
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
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// When
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageFailed", "testName", "somePackage", elapsedTime)
		outputPortMock.AssertNumberOfCalls(t, "FirstCtestOfPackageFailed", 1)
		outputPortMock.AssertNumberOfCalls(t, "CtestFailed", 0)
	}, t)

	// Here
	Test(`
	Given that a CtestRanEvent with name "testName" of package "somePackage" has occurred
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed.
	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()
		testPassedElapsedTime := 1.2

		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestPassedEvt)

		// Then
		outputPortMock.AssertCalled(t, "CtestFailed", "testName", testPassedElapsedTime)

	}, t)
}

func TestCtestRanEvent(t *testing.T) {
	Test(`
	Given that no events have happened
	When a CtestRanEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()

		// When
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
			Output:  "Some output",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageStartedRunning", "testName", "somePackage")
	}, t)

	Test(`
	Given that no events have happened
	When 2 CtestRanEvent of package "somePackage" occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	Then the user should be informed about both tests that have started running
	And that "testName1" was the first test of its package
	`, func(t *testing.T) {
		// Given
		eventsHandler, outputPortMock := setup()
		elapsedTime1, elapsedTime2 := 2.3, 1.2

		// When
		ctestRanEvt1 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "testName1",
				Elapsed: &elapsedTime1,
				Output:  "Some output",
			},
		)
		ctestRanEvt2 := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "testName2",
				Elapsed: &elapsedTime2,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageStartedRunning", "testName1", "somePackage")
		outputPortMock.AssertCalled(t, "CtestStartedRunning", "testName2")
	}, t)

	Test(`
	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
	Then the user should not be informed about the second run, when the second event occurs
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

		// Given
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageStartedRunning", "testName", "somePackage")
		outputPortMock.AssertNumberOfCalls(t, "FirstCtestOfPackageStartedRunning", 1)
		outputPortMock.AssertNumberOfCalls(t, "CtestStartedRunning", 0)
	}, t)

	Test(`
	Given that a CtestRanEvent has occurred with test name "testName" of package "somePackage"
	And then a CtestPassedEvent has occurred for the same test
	When a CtestPassedEvent occurs with the same test name "testName" of package "somePackage"
	Then the user should not be informed about the second run, when the second event occurs
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

		// Given
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
				Output:  "Some output",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
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
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		outputPortMock.AssertCalled(t, "FirstCtestOfPackageStartedRunning", "testName", "somePackage")
		outputPortMock.AssertNumberOfCalls(t, "FirstCtestOfPackageStartedRunning", 1)
		outputPortMock.AssertNumberOfCalls(t, "CtestStartedRunning", 0)
	}, t)
}
