package cmd_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/mock"
)

func setup() (cmd.EventsHandler, *cmd.OutputPortMock) {
	outputPortMock := cmd.NewOutputPortMock()
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventHandler := cmd.NewEventsHandler(outputPortMock, &ctestTracker)
	outputPortMock.On("CtestPassed", mock.Anything, mock.Anything).Return()

	return eventHandler, outputPortMock
}

func TestCtestPassedEvent(t *testing.T) {

	Test(`
	Given that no events have happened
	When a CtestPassedEvent occurs with test name "testName" and elapsed time 2.3s
	Then the user should be informed about the test that has passed
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

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
		outputPortMock.AssertCalled(t, "CtestPassed", "testName", elapsedTime)
	}, t)

	Test(`
	Given that no events have happened
	When 2 CtestPassedEvent occur with test names "testName1", "testName2" and elapsed time 2.3s, 1.2s
	Then the user should be informed about both tests that have passed
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

		outputPortMock.AssertCalled(t, "CtestPassed", "testName1", elapsedTime1)
		outputPortMock.AssertCalled(t, "CtestPassed", "testName2", elapsedTime2)

	}, t)

	Test(`
	Given that a CtestPassedEvent has occurred with test name "testName"
	When a CtestPassedEvent occurs with the same test name "testName"
	Then the user should not be informed about the second passing, when the second event occurs
	`, func(t *testing.T) {
		eventsHandler, outputPortMock := setup()
		elapsedTime := 2.3

		// Given
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

		// When
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)
		outputPortMock.AssertCalled(t, "CtestPassed", "testName", elapsedTime)

		// Then
		outputPortMock.AssertNumberOfCalls(t, "CtestPassed", 1)
	}, t)
}
