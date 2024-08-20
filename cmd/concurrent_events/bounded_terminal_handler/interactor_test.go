package bounded_terminal_handler_test

import (
	"math"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events/bounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	"github.com/stretchr/testify/assert"
)

func setup(terminalHeight int) (*bounded_terminal_handler.Interactor, *fake_ansi_terminal.FakeAnsiTerminal, *ctests_tracker.CtestsTracker) {
	fakeAnsiTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, terminalHeight)
	fakeAnsiTerminalPresenter := bounded_terminal_handler.NewPresenter(&fakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	interactor := bounded_terminal_handler.NewInteractor(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &interactor, &fakeAnsiTerminal, &ctestTracker
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

func TestHandlePackageStartedEvent_TerminalHeightLessThanOrEqualTo5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that no events have occurred
	 And we have a bounded terminal with height 1
	 When a HandlePackageStartedEvent occurs for package "somePackage"
	 Then the user should be informed that the tests for that package are running`, func(t *testing.T) {
		// Given
		interactor, terminal, _ := setup(1)

		// When
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		interactor.HandlePackageStartedEvent(packStartedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ somePackage",
		)
	}, t)

	Test(`
	 Given that a HandlePackageStartedEvent for package "somePackage" has occurred
	 And we have a bounded terminal with height 1
	 When a HandlePackageStartedEvent occurs for package "somePackage"
	 Then the user should be informed only once that the tests for the "somePackage" package are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(1)
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ somePackage",
		)
	}, t)

	Test(`
	 Given that a HandlePackageStartedEvent for package "somePackage 1" has occured
	 And we have a bounded terminal with height 1
	 When a HandlePackageStartedEvent for package "somePackage 2" occurs
	 And the printed text in the viewport should be "⏳ somePackage 1"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("somePackage 1", "somePackage 2")

		// Given
		eventsHandler, terminal, _ := setup(1)
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["somePackage 1"])

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["somePackage 2"])

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ somePackage 1",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent has occurred for "package 1"
	And a CtestFailedEvent has occurred for "package 1"
	And there is a terminal with height 5
	When a PackageStartedEvent occurrs for package "package 2"
	Then this text will be on the terminal "⏳ package 1\n⏳ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ package 1\n⏳ package 2",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent has occurred for "package 1"
	And a CtestPassedEvent has occurred for "package 1"
	And there is a terminal with height 5
	When a PackageStartedEvent occurrs for package "package 2"
	Then this text will be on the terminal "⏳ package 1\n⏳ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ package 1\n⏳ package 2",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 5
	When 3 HandlePackageStartedEvent for packages "package 1", ..., "package 5" occur
	And the printed text should be "⏳ package 1\n⏳ package 2\n⏳ package 3\n⏳ package 4\n⏳ package 5"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")

		// Given
		eventsHandler, terminal, _ := setup(5)

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 3"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 4"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 5"])

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ package 1\n⏳ package 2\n⏳ package 3\n⏳ package 4\n⏳ package 5",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 5
	When 6 HandlePackageStartedEvent for packages "package 1", ..., "package 6" occur
	And the printed text should be "⏳ package 1\n⏳ package 2\n⏳ package 3\n⏳ package 4\n⏳ package 5"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents(
			"package 1",
			"package 2",
			"package 3",
			"package 4",
			"package 5",
			"package 6",
		)

		// Given
		eventsHandler, terminal, _ := setup(5)

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 3"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 4"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 5"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 6"])

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ package 1\n⏳ package 2\n⏳ package 3\n⏳ package 4\n⏳ package 5",
		)
	}, t)
}

func TestHandlePackageStartedEvent_TerminalHeightGreaterThan5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 6
	When 1 HandlePackageStartedEvent for package "package 1" occur
	And the printed text should be "⏳ package 1" and the summary of tests:
	"<bold>Packages</bold>: 1 running\n<bold>Tests</bold>: 0 running\n<bold>Time</bold>: 0.000s"`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(6)

		// When
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "package 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ package 1"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 6
	When 2 HandlePackageStartedEvent for packages "package 1", and "package 2" occur
	And the printed text should be"⏳ package 1\n⏳ package 2" and the summary of tests:
	"<bold>Packages</bold>: 2 running\n<bold>Tests</bold>: 0 running\n<bold>Time</bold>: 0.000s"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("package 1", "package 2")
		// Given
		eventsHandler, terminal, _ := setup(6)

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 2"])

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ package 1\n⏳ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 2 running"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 6
	When 2 HandlePackageStartedEvent for packages "package 1", "package 2", "package 3" occur
	And the printed text should be "⏳ package 1\n⏳ package 2" and the summary of tests:
	"<bold>Packages</bold>: 3 running\n<bold>Tests</bold>: 0 running\n<bold>Time</bold>: 0.000s"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("package 1", "package 2", "package 3")
		// Given
		eventsHandler, terminal, _ := setup(6)

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["package 3"])

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ package 1\n⏳ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 3 running"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)
}

func TestHandlePackagePassedEvent_TerminalHeightLessThanOrEqualTo5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that no events have occurred
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.`, func(t *testing.T) {
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		eventsHandler, fakeTerminal, _ := setup(5)

		// When
		err := eventsHandler.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		packPassedEvts := makePackagePassedEvents("somePackage")

		// Given
		eventsHandler, fakeTerminal, _ := setup(5)
		eventsHandler.HandlePackageStartedEvent(packStartedEvts["somePackage"])

		// When
		err := eventsHandler.HandlePackagePassed(packPassedEvts["somePackage"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then this text will be on the terminal "✅ somePackage".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestPassedEvent(ctestPassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ somePackage",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestPassedEvent has occurred for "package 1"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 1"
	 Then this text will be on the terminal "✅ package 1\n⏳ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n⏳ package 2",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestPassedEvent has occurred for each of them
	 And a PackagePassedEvent for package "package 1" has occurred
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 2"
	 Then this text will be on the terminal "✅ package 1\n✅ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n✅ package 2",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestPassedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackagePassedEvent for packages "package 1",..., "package 3"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 4"
	 Then the printed text will be:
	 	"✅ package 1\n✅ package 2\n✅ package 3\n✅ package 4\n⏳ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")
		ctest3PassedEvt := makeCtestPassedEvent("package 3", "testName")
		ctest4PassedEvt := makeCtestPassedEvent("package 4", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandleCtestPassedEvent(ctest3PassedEvt)
		interactor.HandleCtestPassedEvent(ctest4PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])
		interactor.HandlePackagePassed(packagePassedEvts["package 3"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 4"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n✅ package 2\n✅ package 3\n✅ package 4\n⏳ package 5",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestPassedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackagePassedEvent for packages "package 1",..., "package 5"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 5"
	 Then the printed text will be:
	 	"✅ package 1\n✅ package 2\n✅ package 3\n✅ package 4\n✅ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")
		ctest3PassedEvt := makeCtestPassedEvent("package 3", "testName")
		ctest4PassedEvt := makeCtestPassedEvent("package 4", "testName")
		ctest5PassedEvt := makeCtestPassedEvent("package 5", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandleCtestPassedEvent(ctest3PassedEvt)
		interactor.HandleCtestPassedEvent(ctest4PassedEvt)
		interactor.HandleCtestPassedEvent(ctest5PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])
		interactor.HandlePackagePassed(packagePassedEvts["package 3"])
		interactor.HandlePackagePassed(packagePassedEvts["package 4"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 5"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n✅ package 2\n✅ package 3\n✅ package 4\n✅ package 5",
		)
	}, t)

	Test(`
	Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestPassedEvent has occurred for packages "pack 1", ..., "pack 6"
	And a PackagePassedEvent for packages "pack 1",..., "pack 5"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 6"
	Then the printed text will be: "✅ pack 2\n✅ pack 3\n✅ pack 4\n✅ pack 5\n✅ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("pack 2", "testName")
		ctest3PassedEvt := makeCtestPassedEvent("pack 3", "testName")
		ctest4PassedEvt := makeCtestPassedEvent("pack 4", "testName")
		ctest5PassedEvt := makeCtestPassedEvent("pack 5", "testName")
		ctest6PassedEvt := makeCtestPassedEvent("pack 6", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandleCtestPassedEvent(ctest3PassedEvt)
		interactor.HandleCtestPassedEvent(ctest4PassedEvt)
		interactor.HandleCtestPassedEvent(ctest5PassedEvt)
		interactor.HandleCtestPassedEvent(ctest6PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 2"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 3"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 4"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 5"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 6"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ pack 2\n✅ pack 3\n✅ pack 4\n✅ pack 5\n✅ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 5"
	And a CtestPassedEvent has occurred for packages "pack 1"
	And a PackagePassedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 1"
	Then the printed text will be:
		"✅ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5")
		packagePassedEvts := makePackagePassedEvents("pack 1")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestPassedEvent has occurred for packages "pack 1", "pack 2"
	And a PackagePassedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 2"
	Then the printed text will be:
		"✅ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("pack 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestPassedEvent has occurred for packages "pack 1"
	And there is a terminal with height 5
	And a PackagePassedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)
}

func TestHandlePackagePassedEvent_TerminalHeightGreaterThan5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that no events have occurred
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.`, func(t *testing.T) {
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		eventsHandler, fakeTerminal, _ := setup(6)

		// When
		err := eventsHandler.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		packPassedEvts := makePackagePassedEvents("somePackage")

		// Given
		eventsHandler, fakeTerminal, _ := setup(5)
		eventsHandler.HandlePackageStartedEvent(packStartedEvts["somePackage"])

		// When
		err := eventsHandler.HandlePackagePassed(packPassedEvts["somePackage"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then this text will be on the terminal "✅ somePackage" and the summary of tests
	 "\n\nPackages: 0 running\nTests: 0 running\nTime: 0.000s"`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestPassedEvent(ctestPassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "pack 1" and "pack 2"
	 And a CtestPassedEvent has occurred for "pack 1"
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "pack 1"
	 Then this text will be on the terminal "✅ package 1\n⏳ package 2" and the summary of tests
	 "\n\nPackages: 0 running\nTests: 0 running\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n⏳ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestPassedEvent has occurred for each of them
	 And a PackagePassedEvent for package "package 1" has occurred
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "package 2"
	 Then this text will be on the terminal "✅ package 1\n✅ package 2" and the summary of tests
	 "\n\nPackages: 0 running\nTests: 0 running\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n✅ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 3"
	 And a CtestPassedEvent has occurred for packages "package 1"
	 And a PackagePassedEvent for packages "package 1", "package 2"
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "package 2"
	 Then the printed text will be: "✅ package 2\n⏳ package 3" and the summary of tests
	 "\n\nPackages: 1 running\nTests: 0 running\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 2\n⏳ package 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	And a CtestPassedEvent has occurred for packages "pack 1", ..., "pack 3"
	And a PackagePassedEvent for packages "pack 1", and "pack 2"
	And there is a terminal with height 6
	When a PackagePassedEvent for package "pack 6"
	Then the printed text will be: "✅ pack 2\n✅ pack 3" and the summary of tests
	"\n\nPackages: 0 running\nTests: 0 running\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2", "pack 3")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("pack 2", "testName")
		ctest3PassedEvt := makeCtestPassedEvent("pack 3", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandleCtestPassedEvent(ctest3PassedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 2"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 3"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ pack 2\n✅ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.GREEN+"3 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	And a CtestPassedEvent has occurred for packages "pack 1"
	And there is a terminal with height 6
	And a PackagePassedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n" and the summary of tests
	"\n\nPackages: 2 running\nTests: 0 running\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packagePassedEvts := makePackagePassedEvents("pack 1")
		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestPassedEvent(ctest1PassedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 2 running, "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 running"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)
}

func TestHandlePackageFailedEvent_TerminalHeightLessThanOrEqualTo5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that no events have occurred
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.`, func(t *testing.T) {
		packagePassedEvts := makePackageFailedEvents("package 1")
		// Given
		eventsHandler, fakeTerminal, _ := setup(5)

		// When
		err := eventsHandler.HandlePackageFailed(packagePassedEvts["package 1"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		packFailedEvts := makePackageFailedEvents("somePackage")

		// Given
		eventsHandler, fakeTerminal, _ := setup(5)
		eventsHandler.HandlePackageStartedEvent(packStartedEvts["somePackage"])

		// When
		err := eventsHandler.HandlePackageFailed(packFailedEvts["somePackage"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestFailedEvent for test with name "testName" in package "somePackage" has occurred
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "somePackage" occurs
	 Then this text will be on the terminal "❌ somePackage".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "testName")
		packageFailedEvts := makePackageFailedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestFailedEvent(ctestFailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ somePackage",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestFailedEvent has occurred for "package 1"
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "package 1"
	 Then this text will be on the terminal "❌ package 1\n⏳ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		packageFailedEvts := makePackageFailedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n⏳ package 2",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestFailedEvent has occurred for each of them
	 And a PackageFailedEvent for package "package 1" has occurred
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "package 2"
	 Then this text will be on the terminal "❌ package 1\n❌ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packageFailedEvts := makePackageFailedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n❌ package 2",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestFailedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackageFailedEvent for packages "package 1",..., "package 3"
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "package 4"
	 Then the printed text will be:
	 	"❌ package 1\n❌ package 2\n❌ package 3\n❌ package 4\n⏳ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packageFailedEvts := makePackageFailedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("package 3", "testName")
		ctest4FailedEvt := makeCtestFailedEvent("package 4", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandleCtestFailedEvent(ctest3FailedEvt)
		interactor.HandleCtestFailedEvent(ctest4FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["package 1"])
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])
		interactor.HandlePackageFailed(packageFailedEvts["package 3"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 4"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n❌ package 2\n❌ package 3\n❌ package 4\n⏳ package 5",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestFailedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackageFailedEvent for packages "package 1",..., "package 5"
	 And there is a terminal with height 5
	 When a PackageFailedEvent for package "package 5"
	 Then the printed text will be:
	 	"❌ package 1\n❌ package 2\n❌ package 3\n❌ package 4\n❌ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packageFailedEvts := makePackageFailedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("package 3", "testName")
		ctest4FailedEvt := makeCtestFailedEvent("package 4", "testName")
		ctest5FailedEvt := makeCtestFailedEvent("package 5", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandleCtestFailedEvent(ctest3FailedEvt)
		interactor.HandleCtestFailedEvent(ctest4FailedEvt)
		interactor.HandleCtestFailedEvent(ctest5FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["package 1"])
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])
		interactor.HandlePackageFailed(packageFailedEvts["package 3"])
		interactor.HandlePackageFailed(packageFailedEvts["package 4"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 5"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n❌ package 2\n❌ package 3\n❌ package 4\n❌ package 5",
		)
	}, t)

	Test(`
	Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestFailedEvent has occurred for packages "pack 1", ..., "pack 6"
	And a PackageFailedEvent for packages "pack 1",..., "pack 5"
	And there is a terminal with height 5
	When a PackageFailedEvent for package "pack 6"
	Then the printed text will be: "❌ pack 2\n❌ pack 3\n❌ pack 4\n❌ pack 5\n❌ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packageFailedEvts := makePackageFailedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("pack 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("pack 3", "testName")
		ctest4FailedEvt := makeCtestFailedEvent("pack 4", "testName")
		ctest5FailedEvt := makeCtestFailedEvent("pack 5", "testName")
		ctest6FailedEvt := makeCtestFailedEvent("pack 6", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandleCtestFailedEvent(ctest3FailedEvt)
		interactor.HandleCtestFailedEvent(ctest4FailedEvt)
		interactor.HandleCtestFailedEvent(ctest5FailedEvt)
		interactor.HandleCtestFailedEvent(ctest6FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])
		interactor.HandlePackageFailed(packageFailedEvts["pack 2"])
		interactor.HandlePackageFailed(packageFailedEvts["pack 3"])
		interactor.HandlePackageFailed(packageFailedEvts["pack 4"])
		interactor.HandlePackageFailed(packageFailedEvts["pack 5"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 6"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ pack 2\n❌ pack 3\n❌ pack 4\n❌ pack 5\n❌ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 5"
	And a CtestFailedEvent has occurred for packages "pack 1"
	And a PackageFailedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackageFailedEvent for package "pack 1"
	Then the printed text will be:
		"❌ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5")
		packageFailedEvts := makePackageFailedEvents("pack 1")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestFailedEvent has occurred for packages "pack 1", "pack 2"
	And a PackageFailedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackageFailedEvent for package "pack 2"
	Then the printed text will be:
		"❌ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packageFailedEvts := makePackageFailedEvents("pack 1", "pack 2")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("pack 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestFailedEvent has occurred for packages "pack 1"
	And there is a terminal with height 5
	And a PackageFailedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packageFailedEvts := makePackageFailedEvents("pack 1", "pack 2")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)
}
