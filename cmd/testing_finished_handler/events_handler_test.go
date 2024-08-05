package testing_finished_handler_test

import (
	"math"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
	"github.com/redjolr/goherent/cmd/events/testing_finished_event"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	"github.com/stretchr/testify/assert"
)

func setup() (*testing_finished_handler.EventsHandler, *fake_ansi_terminal.FakeAnsiTerminal, *ctests_tracker.CtestsTracker) {
	fakeAnsiTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
	fakeAnsiTerminalPresenter := testing_finished_handler.NewTerminalPresenter(&fakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := testing_finished_handler.NewEventsHandler(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &fakeAnsiTerminal, &ctestTracker
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
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2

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
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt))
		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2
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
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2
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
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2

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
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2

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
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt2))
		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2

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
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2
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
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctest1PassedEvt))
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctest2FailedEvt))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		elapsedTime := 1.2

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
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctest1PassedEvt))
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctest2FailedEvt))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()

		ctestSkippedEvt := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage",
			},
		)
		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage",
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage 1",
			},
		)

		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage 2",
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		testPassedElapsed := 1.2
		ctestSkippedEvt1 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 1",
				Package: "somePackage 1",
			},
		)

		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &testPassedElapsed,
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt2))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		testElapsed := 1.2
		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &testElapsed,
			},
		)
		ctestPassedEvt2 := ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "testName 2",
				Package: "somePackage 2",
				Elapsed: &testElapsed,
			},
		)
		ctestSkippedEvt3 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 3",
				Package: "somePackage 3",
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt2))
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt3))

		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
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
	And the tests execution time was 1.2 seconds`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, ctestsTracker := setup()
		testFailedElapsed := 1.2

		ctestFailedEvt1 := ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName 1",
				Package: "somePackage 1",
				Elapsed: &testFailedElapsed,
			},
		)

		ctestSkippedEvt2 := ctest_skipped_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName 2",
				Package: "somePackage 2",
			},
		)
		ctestsTracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestFailedEvt1))
		ctestsTracker.InsertCtest(ctests_tracker.NewSkippedCtest(ctestSkippedEvt2))
		// When
		testingFinishedEvent := testing_finished_event.NewTestingFinishedEvent(time.Millisecond * 1200)
		eventsHandler.HandleTestingFinished(testingFinishedEvent)

		// Then
		assert.Equal(
			terminal.Text(),
			ansi_escape.BOLD+"\nPackages:"+ansi_escape.RESET_BOLD+" "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 2 total\n"+
				ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)
}
