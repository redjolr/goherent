package sequential_events_test

import (
	"math"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/expect"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	. "github.com/redjolr/goherent/test"
)

func setupInteractorWithBoundedTerminal(height int) (
	*sequential_events.Interactor,
	*fake_ansi_terminal.FakeAnsiTerminal,
	*ctests_tracker.CtestsTracker,
) {
	boundedFakeAnsiTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, height)
	fakeAnsiTerminalPresenter := sequential_events.NewBoundedTerminalPresenter(&boundedFakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := sequential_events.NewInteractor(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &boundedFakeAnsiTerminal, &ctestTracker
}

func makePackageStartedEvents(packageNames ...string) map[string]events.PackageStartedEvent {
	evts := make(map[string]events.PackageStartedEvent)
	for _, packName := range packageNames {
		evts[packName] = events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: packName,
			})
	}
	return evts
}

func makePackagePassedEvents(packageNames ...string) map[string]events.PackagePassedEvent {
	evts := make(map[string]events.PackagePassedEvent)
	timeElapsed := 1.2
	for _, packName := range packageNames {
		evts[packName] = events.NewPackagePassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packName,
				Elapsed: &timeElapsed,
			})
	}
	return evts
}

func makePackageFailedEvents(packageNames ...string) map[string]events.PackageFailedEvent {
	evts := make(map[string]events.PackageFailedEvent)
	timeElapsed := 1.2
	for _, packName := range packageNames {
		evts[packName] = events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packName,
				Elapsed: &timeElapsed,
			})
	}
	return evts
}

func makeCtestFailedEvent(packageName, testName string) events.CtestFailedEvent {
	timeElapsed := 1.2
	return events.NewCtestFailedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "pass",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func makeCtestOutputEvent(packageName, testName, output string) events.CtestOutputEvent {
	return events.NewCtestOutputEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "output",
			Test:    testName,
			Package: packageName,
			Output:  output,
		},
	)
}

func makeCtestSkippedEvent(packageName, testName string) events.CtestSkippedEvent {
	timeElapsed := 1.2
	return events.NewCtestSkippedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "skip",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func makeCtestPassedEvent(packageName, testName string) events.CtestPassedEvent {
	timeElapsed := 1.2
	return events.NewCtestPassedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "pass",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func TestCtestRanEventWithBoundedTerminal(t *testing.T) {
	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		// When
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏳ ParentTest/testName",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with test name "Multiline\ntest name" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		// When
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/Multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏳ ParentTest/Multiline...",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 20
	When 2 CtestRanEvent of package "somePackage" occur with test names "ParentTest/testName1", "ParentTest/testName2" 
		and elapsed time 2.3s, 1.2s
	Then the second CtestRanEvent should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(20)

		// When
		ctestRanEvt1 := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "ParentTest/testName1",
			},
		)
		ctestRanEvt2 := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Package: "somePackage",
				Test:    "ParentTest/testName2",
			},
		)
		ctestRanEvt1Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt1)
		ctestRanEvt2Err := eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// Then
		Expect(ctestRanEvt1Err).NotToBeError()
		Expect(ctestRanEvt2Err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that a CtestRanEvent has occurred with test name "ParentTest/testName" of package "somePackage"
	And we have a bounded terminal with height 1
	When a CtestRanEvent occurs with the same test name "ParentTest/testName" of package "somePackage"
	Then the user should be informed only once that the given test from the given package is running.`, func(Expect expect.F) {
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		// Given
		ctestRanEvt := events.NewCtestRanEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "run",
				Test:    "ParentTest/testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏳ ParentTest/testName",
		)
	}, t)
}

func TestCtestPassedEventWithBoundedTerminal(t *testing.T) {
	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName 1" of package "somePackage" has occurred
	And a CtestPassedEvent with name "ParentTest/testName 1" has occurred
	And a CtestRanEvent with name "ParentTest/testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of with name "ParentTest/testName 2" of package "somePackage" occurs
	Then the user should be informed that the test has passed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName 1",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName 2",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/testName 1\n✅ ParentTest/testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestPassedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The 1st multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The second multiline\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/The 1st multiline   \ntest name\n✅ ParentTest/The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestPassedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/multiline 1   \ntest name longer\n✅ ParentTest/multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline 1\ntest name" from "packageName"
	And a CtestPassedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestPassedEvent with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The multiline 1\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The multiline 2\ntest name",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/The multiline 1\ntest name\n✅ ParentTest/The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestPassedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestPassedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/multiline 1\ntest name longer\n✅ ParentTest/multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName Line1\nLine2\nLine3",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestPassedEvent of the same test/package occurs
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n✅ ParentTest/testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestPassedEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestPassedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has passed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		testPassedElapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestPassedEvt(ctestPassedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n" +
				"✅ ParentTest/The 1st multiline\nLine2\nLine3   \nLine4\n" +
				"✅ ParentTest/The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 5
	When a CtestPassedEvent occurs with test name "ParentTest/testName" from "packageName"
	Then the HandleCtestPassedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 2.3

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "ParentTest/testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 5
	When a CtestPassedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestPassedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName",
				Package: "somePackage 2",
				Elapsed: &testPassedElapsedTime,
			},
		)
		err := eventsHandler.HandleCtestPassedEvt(ctestPassedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)
}

func TestCtestFailedEventWithBoundedTerminal(t *testing.T) {
	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		testPassedElapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Elapsed: &testPassedElapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName 1" of package "somePackage" has occurred
	And a CtestFailedEvent with name "ParentTest/testName 1" has occurred
	And a CtestRanEvent with name "ParentTest/testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of with name "ParentTest/testName 2" of package "somePackage" occurs
	Then the user should be informed that the test has failed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName 1",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName 2",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName 1\n❌ ParentTest/testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestFailedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The 1st multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The second multiline\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/The 1st multiline   \ntest name\n❌ ParentTest/The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestFailedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/multiline 1   \ntest name longer\n❌ ParentTest/multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline 1\ntest name" from "packageName"
	And a CtestFailedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The multiline 1\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The multiline 2\ntest name",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/The multiline 1\ntest name\n❌ ParentTest/The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestFailedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/multiline 1\ntest name longer\n❌ ParentTest/multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName Line1\nLine2\nLine3",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestFailedEvent of the same test/package occurs
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestFailedEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestFailedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test has failed
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)
		elapsedTime := 2.3

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n" +
				"❌ ParentTest/The 1st multiline\nLine2\nLine3   \nLine4\n" +
				"❌ ParentTest/The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	And we have a bounded terminal with height 5
	When a CtestFailedEvent occurs with test name "ParentTest/testName" from "packageName"
	Then the HandleCtestFailedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 2.3

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Package: "somePackage",
				Test:    "ParentTest/testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestFailedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName",
				Package: "somePackage 2",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that 2 CtestOutputEvent for Ctest with name "ParentTest/testName" of package "somePackage" have occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Output:  "This is output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/testName",
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
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)
		Expect(err).ToBeError()

		// Then
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And a CtestOutputEvent for Ctest with name "ParentTest/testName" of package "somePackage" has also occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 1.2

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/testName",
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
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName\nThis is some output.",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And two CtestOutputEvent for Ctest with name "ParentTest/testName" of package "somePackage" has also occurred
	And we have a bounded terminal with height 5
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(5)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Output:  "Some output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Output:  "_Some output 2.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/testName\nSome output 1._Some output 2.",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/NameLine1\nNameLine2\nNameLine3" of package "somePackage" has occurred
	And two CtestOutputEvent for Ctest with name "ParentTest/testName" of package "somePackage" has also occurred
	And we have a bounded terminal with height 2
	When a CtestFailedEvent of the same test/package occurs
	Then a user should be informed that the Ctest has failed
	And the output from the CtestOutputEvents should be presented.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)
		elapsedTime := 2.3

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/NameLine1\nNameLine2\nNameLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestOutputEvt1 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Output:  "Some output 1.",
			},
		)

		ctestOutputEvt2 := events.NewCtestOutputEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Test:    "ParentTest/NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Output:  "_Some output 2.",
			},
		)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt1)
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt2)

		// When
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "ParentTest/NameLine1\nNameLine2\nNameLine3",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvt(ctestFailedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n❌ ParentTest/NameLine1\nNameLine2   \nNameLine3\nSome output 1._Some output 2.",
		)
	}, t)
}

func TestCtestSkippedEventWithBoundedTerminal(t *testing.T) {
	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/testName",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/The multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/multiline   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline\ntest name" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/The multiline\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline\ntest name longer" from "packageName"
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/multiline\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName 1" of package "somePackage" has occurred
	And a CtestSkippedEvent with name "ParentTest/testName 1" has occurred
	And a CtestRanEvent with name "ParentTest/testName 2" of package "somePackage" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of with name "ParentTest/testName 2" of package "somePackage" occurs
	Then the user should be informed that the test was skipped.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 1",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName 1",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName 2",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName 2",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/testName 1\n⏩ ParentTest/testName 2",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\ntest name" from "packageName" has occurred
	And a CtestSkippedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline 2\ntest name" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The 1st multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The second multiline\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/The 1st multiline   \ntest name\n⏩ ParentTest/The second multiline   \ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestSkippedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/multiline 1   \ntest name longer\n⏩ ParentTest/multiline 2   \ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/The multiline 1\ntest name" from "packageName"
	And a CtestSkippedEvent with test name "ParentTest/The multiline 1\ntest name" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent with test name "ParentTest/The multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 1\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The multiline 1\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The multiline 2\ntest name",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The multiline 2\ntest name",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/The multiline 1\ntest name\n⏩ ParentTest/The multiline 2\ntest name",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/multiline 1\ntest name longer" from "packageName"
	And a CtestSkippedEvent with test name "ParentTest/multiline 1\ntest name longer" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/multiline 2\ntest name longer" from "packageName" has occurred
	And we have a bounded terminal with height 2
	When a CtestSkippedEvent with test name "ParentTest/multiline 2\ntest name longer" from "packageName" occurrs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(2)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 1\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline 1\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/multiline 2\ntest name longer",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/multiline 2\ntest name longer",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/multiline 1\ntest name longer\n⏩ ParentTest/multiline 2\ntest name longer",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName Line1\nLine2\nLine3" of package "somePackage" has occurred
	And we have a bounded terminal with height 3
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName Line1\nLine2\nLine3",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/testName Line1\nLine2\nLine3",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent occurs with test name "ParentTest/testName Line1\nLine2\nLine3\nLine4" from "packageName"
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent of the same test/package occurs
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName Line1\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n⏩ ParentTest/testName Line1\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that a CtestRanEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And a CtestSkippedEvent with test name "ParentTest/The 1st multiline\nLine2\nLine3\nLine4" from packag "packageName" has occurred
	And a CtestRanEvent occurs with test name "ParentTest/The second multiline\nLine2\nLine3\nLine4" from "packageName" has occurred
	And we have a bounded terminal with height 1
	When a CtestSkippedEvent with test name "ParentTest/The second multiline\ntest name" from packag "packageName" has occurred
	Then the user should be informed that the test was skipped
	And the printed test name should be truncated so that it can fit in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(3)

		ctestRanEvt1 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt1)

		ctestSkippedEvt1 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The 1st multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt1)

		ctestRanEvt2 := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt2)

		// When
		ctestSkippedEvt2 := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/The second multiline\nLine2\nLine3\nLine4",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt2)

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n" +
				"⏩ ParentTest/The 1st multiline\nLine2\nLine3   \nLine4\n" +
				"⏩ ParentTest/The second multiline\nLine2\nLine3   \nLine4",
		)
	}, t)

	Test(`
	Given that no events have happened
	When a CtestSkippedEvent occurs with test name "ParentTest/testName" from "packageName"
	Then the HandleCtestSkippedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		elapsedTime := 2.3

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Package: "somePackage",
				Test:    "ParentTest/testName",
				Elapsed: &elapsedTime,
			},
		)
		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)

	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	When a CtestSkippedEvent of a different package "somePackage 2" occurs
	Then the HandleCtestSkippedEvt should produce an error
	And an error should be displayed in the terminal.`, func(Expect expect.F) {
		// Given
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)

		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)

		// When
		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "ParentTest/testName",
				Package: "somePackage 2",
			},
		)
		err := eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// Then
		Expect(err).ToBeError()
		Expect(terminal.Text()).ToContain("❗ Error.")
	}, t)
}

func TestHandleTestingStartedWithBoundedTerminal(t *testing.T) {
	Test("User should be informed, that the testing has started", func(Expect expect.F) {
		eventsHandler, terminal, _ := setupInteractorWithBoundedTerminal(1)
		now := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(now)
		eventsHandler.HandleTestingStarted(testingStartedEvt)

		Expect(terminal.Text()).ToEqual(
			"\n🚀 Starting...",
		)
	}, t)
}

func TestHandlePackageFailedEvent(t *testing.T) {
	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And a CtestFailedEvent for test with name "ParentTest/testName" in package "somePackage" has occurred
	And a CtestOutputEvent for test "TestFunc" in package "somePackage" with output "Some package output" has occurred
	And there is a terminal with height 9
	When a PackageFailedEvent occurs
	Then the failing package, failing test, package output will be displayed
	And this summary will be displayed:
	"\n\nPackages: 1 failed, 1 total\nTests: 1 failed, 1 total\nTime: 1.200s"`, func(Expect expect.F) {
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "ParentTest/testName")
		testFuncOutputEvt := makeCtestOutputEvent("somePackage", "TestFunc", "Some package output")
		packFailedEvts := makePackageFailedEvents("somePackage")

		// Given
		interactor, terminal, _ := setupInteractorWithBoundedTerminal(9)
		interactor.HandleCtestRanEvt(ctestRanEvt)

		interactor.HandleCtestFailedEvt(ctestFailedEvt)
		interactor.HandleCtestOutputEvent(testFuncOutputEvt)

		// When
		interactor.HandlePackageFailedEvt(packFailedEvts["somePackage"])

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n" +
				"❌ ParentTest/testName" +
				"\n\nSome package output",
		)
	}, t)

	//
	Test(`
	Given that a CtestRanEvent with name "ParentTest/testName" of package "somePackage" has occurred
	And a CtestFailedEvent for test with name "ParentTest/testName" in package "somePackage" has occurred
	And a PackageFailedEvent for package "somePackage" occurs
	And a CtestOutputEvent for test "TestFunc" in package "somePackage" with output "Some package output 1" has occurred
	And a CtestOutputEvent for test "TestFunc" in package "somePackage" with output "Some package output 2" has occurred
	And another PackageOutputEvent for package "somePackage" with output "Some package output 2" has occurred
	And there is a terminal with height 9
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then the failing package, failing test, package output will be displayed
	And this summary will be displayed:
	"\n\nPackages: 1 failed, 1 total\nTests: 1 failed, 1 total\nTime: 1.200s"`, func(Expect expect.F) {
		ctestRanEvt := events.NewCtestRanEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "ParentTest/testName",
			Package: "somePackage",
		})
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "ParentTest/testName")
		packageFailedEvts := makePackageFailedEvents("somePackage")
		testFuncOutputEvt1 := makeCtestOutputEvent("somePackage", "TestFunc", "Some package output 1")
		testFuncOutputEvt2 := makeCtestOutputEvent("somePackage", "TestFunc", "Some package output 2")

		// Given
		interactor, terminal, _ := setupInteractorWithBoundedTerminal(9)
		interactor.HandleCtestRanEvt(ctestRanEvt)
		interactor.HandleCtestFailedEvt(ctestFailedEvt)
		interactor.HandleCtestOutputEvent(testFuncOutputEvt1)
		interactor.HandleCtestOutputEvent(testFuncOutputEvt2)

		// When
		interactor.HandlePackageFailedEvt(packageFailedEvts["somePackage"])

		// Then
		Expect(terminal.Text()).ToEqual(
			"\n\n📦 somePackage\n\n" +
				"❌ ParentTest/testName" +
				"\n\nSome package output 1Some package output 2",
		)
	}, t)
}
