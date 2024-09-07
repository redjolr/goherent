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

	Test(`
	Given that thse events have occurred in this order:
	- 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	- 2 CtestPassedEvent have occurred for "pack 1" and "pack 3"
	- 1 CtestFailedEvent has occurred for "pack 2"
	- 2 PackagePassedEvents for "pack 1" and "pack 3"
	- 1 PackageFailedEvent for "pack 2"
	And we have a bounded terminal with height 5
	When 3 HandlePackageStartedEvents for packages "pack 4", "pack 5", and "pack 6" occur
	And the printed text should be "❌ pack 2\n✅ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packPassedEvents := makePackagePassedEvents("pack 1", "pack 3")
		packFailedEvents := makePackageFailedEvents("pack 2")

		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest3PassedEvt := makeCtestPassedEvent("pack 3", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("pack 2", "testName")

		// Given
		eventsHandler, terminal, _ := setup(5)
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 3"])
		eventsHandler.HandleCtestPassedEvent(ctest1PassedEvt)
		eventsHandler.HandleCtestPassedEvent(ctest3PassedEvt)
		eventsHandler.HandleCtestFailedEvent(ctest2FailedEvt)
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 1"])
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 3"])
		eventsHandler.HandlePackageFailed(packFailedEvents["pack 2"])

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 4"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 5"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 6"])

		// Then
		assert.Equal(
			terminal.Text(),
			"❌ pack 2\n✅ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)

	Test(`
	Given that thse events have occurred in this order:
	- 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	- 1 CtestPassedEvent have occurred for "pack 1"
	- 1 CtestSkippedEvent have occurred for "pack 2"
	- 1 CtestFailedEvent has occurred for "pack 3"
	- 2 PackagePassedEvents for "pack 1" and "pack 2"
	- 1 PackageFailedEvent for "pack 2"
	And we have a bounded terminal with height 5
	When 3 HandlePackageStartedEvents for packages "pack 4", "pack 5" occur
	And the printed text should be "✅ pack 1\n⏩ pack 2\n❌ pack 3\n⏳ pack 4\n⏳ pack 5"`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5")
		packPassedEvents := makePackagePassedEvents("pack 1", "pack 2")
		packFailedEvents := makePackageFailedEvents("pack 3")

		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("pack 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("pack 3", "testName")

		// Given
		eventsHandler, terminal, _ := setup(5)
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 3"])
		eventsHandler.HandleCtestPassedEvent(ctest1PassedEvt)
		eventsHandler.HandleCtestSkippedEvent(ctest2SkippedEvt)
		eventsHandler.HandleCtestFailedEvent(ctest3FailedEvt)
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 1"])
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 2"])
		eventsHandler.HandlePackageFailed(packFailedEvents["pack 3"])

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 4"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 5"])

		// Then
		assert.Equal(
			terminal.Text(),
			"✅ pack 1\n⏩ pack 2\n❌ pack 3\n⏳ pack 4\n⏳ pack 5",
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
	"Packages: 1 running\nTests: 0 running\nTime: 0.000s"`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 6
	When 2 HandlePackageStartedEvent for packages "package 1", and "package 2" occur
	And the printed text should be"⏳ package 1\n⏳ package 2" and the summary of tests:
	"Packages: 2 running\nTests: \nTime: 0.000s"`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that no events have occurred
	And we have a bounded terminal with height 6
	When 2 HandlePackageStartedEvent for packages "package 1", "package 2", "package 3" occur
	And the printed text should be "⏳ package 1\n⏳ package 2" and the summary of tests:
	"Packages: 3 running\nTests: \nTime: 0.000s"`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that these events have occurred in this order:
	- 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	- 2 CtestPassedEvent have occurred for "pack 1" and "pack 2"
	- 1 CtestFailedEvent has occurred for "pack 3"
	- 2 PackagePassedEvents for "pack 1" and "pack 2"
	- 1 PackageFailedEvent for "pack 3"
	And we have a bounded terminal with height 7
	When a HandlePackageStartedEvent for "pack 4" ocurrs
	And the printed text should be "✅ pack 2\n❌ pack 3\n⏳ pack 4" and the summary of tests:
	"Packages: 1 running, 1 failed, 2 passed\nTests: 1 failed, 2 passed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packPassedEvents := makePackagePassedEvents("pack 1", "pack 2")
		packFailedEvents := makePackageFailedEvents("pack 3")

		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("pack 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("pack 3", "testName")

		// Given
		eventsHandler, terminal, _ := setup(7)
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 3"])
		eventsHandler.HandleCtestPassedEvent(ctest1PassedEvt)
		eventsHandler.HandleCtestPassedEvent(ctest2PassedEvt)
		eventsHandler.HandleCtestFailedEvent(ctest3FailedEvt)
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 1"])
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 2"])
		eventsHandler.HandlePackageFailed(packFailedEvents["pack 3"])

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 4"])

		// Then
		assert.Equal(
			terminal.Text(),
			"✅ pack 2\n❌ pack 3\n⏳ pack 4"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that thse events have occurred in this order:
	- 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	- 1 CtestPassedEvent have occurred for "pack 1"
	- 1 CtestSkippedEvent have occurred for "pack 2"
	- 1 CtestFailedEvent has occurred for "pack 3"
	- 2 PackagePassedEvents for "pack 1" and "pack 2"
	- 1 PackageFailedEvent for "pack 2"
	And we have a bounded terminal with height 9
	When 3 HandlePackageStartedEvents for packages "pack 4", "pack 5" occur
	And the printed text should be "✅ pack 1\n⏩ pack 2\n❌ pack 3\n⏳ pack 4\n⏳ pack 5" and the summary of tests:
	"Packages: 2 running, 1 failed, 1 skipped, 1 passed\nTests: 1 failed, 1 skipped, 1 passed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvents := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5")
		packPassedEvents := makePackagePassedEvents("pack 1", "pack 2")
		packFailedEvents := makePackageFailedEvents("pack 3")

		ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("pack 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("pack 3", "testName")

		// Given
		eventsHandler, terminal, _ := setup(9)
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 1"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 2"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 3"])
		eventsHandler.HandleCtestPassedEvent(ctest1PassedEvt)
		eventsHandler.HandleCtestSkippedEvent(ctest2SkippedEvt)
		eventsHandler.HandleCtestFailedEvent(ctest3FailedEvt)
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 1"])
		eventsHandler.HandlePackagePassed(packPassedEvents["pack 2"])
		eventsHandler.HandlePackageFailed(packFailedEvents["pack 3"])

		// When
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 4"])
		eventsHandler.HandlePackageStartedEvent(packStartedEvents["pack 5"])

		// Then
		assert.Equal(
			terminal.Text(),
			"✅ pack 1\n⏩ pack 2\n❌ pack 3\n⏳ pack 4\n⏳ pack 5"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 2 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
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
	 And the user will be informed that an error has occurred.`, func(t *testing.T) {
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

	Test(`
	Given these events have occurred in this order:
	- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	- 1 CtestFailedEvent has occurred for "package 1"
	- 1 CtestPassedEvent has occurred for "package 2"
	- 1 PackageFailedEvent has ocurred for "package 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "package 2" occurrs
	Then this text will be on the terminal "❌ package 1\n✅ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")

		packageFailedEvts := makePackageFailedEvents("package 1")
		packagePassedEvts := makePackagePassedEvents("package 2")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n✅ package 2",
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
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		packPassedEvts := makePackagePassedEvents("somePackage")

		// Given
		eventsHandler, fakeTerminal, _ := setup(6)
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
	 "\n\nPackages: 0 running, 1 passed\nTests: 1 passed\nTime: 0.000s"`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "pack 1" and "pack 2"
	 And a CtestPassedEvent has occurred for "pack 1"
	 And there is a terminal with height 6
	 When a PackagePassedEvent for package "pack 1"
	 Then this text will be on the terminal "✅ package 1\n⏳ package 2" and the summary of tests
	 "\n\nPackages: 1 running, 1 passed\nTests: 1 passed\nTime: 0.000s`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
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
	 "\n\nPackages: 0 running, 2 passed\nTests: 2 passed\nTime: 0.000s`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
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
	 "\n\nPackages: 1 running, 2 passed\nTests: 2 passed\nTime: 0.000s`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"2 passed"+ansi_escape.COLOR_RESET+
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
	"\n\nPackages: 0 running, 3 passed\nTests: 3 passed\nTime: 0.000s.`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"3 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	And a CtestPassedEvent has occurred for packages "pack 1"
	And there is a terminal with height 6
	And a PackagePassedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n" and the summary of tests
	"\n\nPackages: 2 running, 1 passed\nTests: 1 passed\nTime: 0.000s.`, func(t *testing.T) {
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
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given these events have occurred in this order:
	- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	- 1 CtestFailedEvent has occurred for "package 1"
	- 1 CtestPassedEvent has occurred for "package 2"
	- 1 PackageFailedEvent has ocurred for "package 1"
	And there is a terminal with height 6
	When a PackagePassedEvent for package "package 2" occurrs
	Then this text will be on the terminal "❌ package 1\n✅ package 2" and the summary of tests
	"\n\nPackages: 0 running, 1 failed, 1 passed\nTests: 1 failed, 1 passed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2PassedEvt := makeCtestPassedEvent("package 2", "testName")

		packageFailedEvts := makePackageFailedEvents("package 1")
		packagePassedEvts := makePackagePassedEvents("package 2")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestPassedEvent(ctest2PassedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n✅ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)
}

func TestHandleCtestPassedEvent(t *testing.T) {

	Test(`
	Given that no events have occurred
	When a CtestPassedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		//Given
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "someTest")

		// When
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt)
	}, t)

	Test(`
	Given that a CtestOutputEvent for "someTest" in "somePackage" has occurred
	When a CtestPassedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		// Given
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "someTest", "someOutput")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "someTest")
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt)
	}, t)
}

func TestHandleCtestFailedEvent(t *testing.T) {

	Test(`
	Given that no events have occurred
	When a CtestFailedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		//Given
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "someTest")

		// When
		eventsHandler.HandleCtestFailedEvent(ctestFailedEvt)
	}, t)

	Test(`
	Given that a CtestOutputEvent for "someTest" in "somePackage" has occurred
	When a CtestFailedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		// Given
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "someTest", "someOutput")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "someTest")
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		eventsHandler.HandleCtestFailedEvent(ctestFailedEvt)
	}, t)
}

func TestHandleCtestSkippedEvent(t *testing.T) {

	Test(`
	Given that no events have occurred
	When a CtestSkippedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		//Given
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "someTest")

		// When
		eventsHandler.HandleCtestSkippedEvent(ctestSkippedEvt)
	}, t)

	Test(`
	Given that a CtestOutputEvent for "someTest" in "somePackage" has occurred
	When a CtestSkippedEvent for test "someTest" in "somePackage" occurs
	Then the operation will be successful.`, func(t *testing.T) {
		eventsHandler, _, _ := setup(6)
		// Given
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "someTest", "someOutput")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "someTest")
		eventsHandler.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		eventsHandler.HandleCtestSkippedEvent(ctestSkippedEvt)
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

	Test(`
	Given these events have occurred in this order:
	- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	- 1 CtestPassedEvent has occurred for "package 1"
	- 1 CtestFailedEvent has occurred for "package 2"
	- 1 PackagePassedEvent has ocurred for "package 1"
	And there is a terminal with height 5
	When a PackageFailedEvent for package "package 2" occurrs
	Then this text will be on the terminal "✅ package 1\n❌ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1")
		packageFailedEvts := makePackageFailedEvents("package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n❌ package 2",
		)
	}, t)
}

func TestHandlePackageFailedEvent_TerminalHeightGreaterThan5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that no events have occurred
	 And there is a terminal with height 6
	 When a PackageFailedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.`, func(t *testing.T) {
		packageFailedEvts := makePackageFailedEvents("package 1")
		// Given
		eventsHandler, fakeTerminal, _ := setup(6)

		// When
		err := eventsHandler.HandlePackageFailed(packageFailedEvts["package 1"])

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And there is a terminal with height 6
	 When a PackageFailedEvent for package "somePackage" occurs
	 And the user will be informed that an error has occurred.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		packPassedEvts := makePackagePassedEvents("somePackage")

		// Given
		eventsHandler, fakeTerminal, _ := setup(6)
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
	And a CtestFailedEvent for test with name "testName" in package "somePackage" has occurred
	And there is a terminal with height 6
	When a PackageFailedEvent for package "somePackage" occurs
	Then this text will be on the terminal "❌ somePackage" and the summary of tests
	"\n\nPackages: 0 running, 1 failed\nTests: 1 failed\nTime: 0.000s"`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "testName")
		packageFailedEvts := makePackageFailedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestFailedEvent(ctestFailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "pack 1" and "pack 2"
	 And a CtestFailedEvent has occurred for "pack 1"
	 And there is a terminal with height 6
	 When a PackageFailedEvent for package "pack 1"
	 Then this text will be on the terminal "❌ package 1\n⏳ package 2" and the summary of tests
	 "\n\nPackages: 1 running, 1 failed\nTests: 1 failed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		packageFailedEvts := makePackageFailedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n⏳ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	And a CtestFailedEvent has occurred for each of them
	And a PackageFailedEvent for package "package 1" has occurred
	And there is a terminal with height 6
	When a PackageFailedEvent for package "package 2"
	Then this text will be on the terminal "❌ package 1\n❌ package 2" and the summary of tests
	"\n\nPackages: 0 running, 2 failed\nTests: 2 failed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packageFailedEvts := makePackageFailedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
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
			"❌ package 1\n❌ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 3"
	 And a CtestFailedEvent has occurred for packages "package 1"
	 And a PackageFailedEvent for packages "package 1", "package 2"
	 And there is a terminal with height 6
	 When a PackageFailedEvent for package "package 2"
	 Then the printed text will be: "❌ package 2\n⏳ package 3" and the summary of tests
	 "\n\nPackages: 1 running, 2 failed\nTests: 2 failed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3")
		packageFailedEvts := makePackageFailedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 2\n⏳ package 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	And a CtestFailedEvent has occurred for packages "pack 1", ..., "pack 3"
	And a PackageFailedEvent for packages "pack 1", and "pack 2"
	And there is a terminal with height 6
	When a PackageFailedEvent for package "pack 6"
	Then the printed text will be: "❌ pack 2\n❌ pack 3" and the summary of tests
	"\n\nPackages: 0 running, 3 failed\nTests: 0 running\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packageFailedEvts := makePackageFailedEvents("pack 1", "pack 2", "pack 3")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("pack 2", "testName")
		ctest3FailedEvt := makeCtestFailedEvent("pack 3", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandleCtestFailedEvent(ctest3FailedEvt)

		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])
		interactor.HandlePackageFailed(packageFailedEvts["pack 2"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 3"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ pack 2\n❌ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"3 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"3 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given that 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
	And a CtestFailedEvent has occurred for packages "pack 1"
	And there is a terminal with height 6
	And a PackageFailedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n" and the summary of tests
	"\n\nPackages: 2 running, 1 failedd\nTests: 1 failed\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packageFailedEvts := makePackageFailedEvents("pack 1")
		ctest1FailedEvt := makeCtestFailedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestFailedEvent(ctest1FailedEvt)

		// When
		interactor.HandlePackageFailed(packageFailedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 2 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
	Given these events have occurred in this order:
	- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	- 1 CtestPassedEvent has occurred for "package 1"
	- 1 CtestFailedEvent has occurred for "package 2"
	- 1 PackagePassedEvent has ocurred for "package 1"
	And there is a terminal with height 6
	When a PackageFailedEvent for package "package 2" occurrs
	Then this text will be on the terminal "✅ package 1\n❌ package 2" and the summary of tests
	"\n\nPackages: 0 running, 1 failed, 1 passed\nTests: 1 failed, 1 passed\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1")
		packageFailedEvts := makePackageFailedEvents("package 2")
		ctest1PassedEvt := makeCtestPassedEvent("package 1", "testName")
		ctest2FailedEvt := makeCtestFailedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestPassedEvent(ctest1PassedEvt)
		interactor.HandleCtestFailedEvent(ctest2FailedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackageFailed(packageFailedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ package 1\n❌ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)
}

func TestSkippedPackages_TerminalHeightLessThanOrEqualTo5(t *testing.T) {
	assert := assert.New(t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestSkippedEvent for test with name "testName" in package "somePackage" has occurred
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then this text will be on the terminal "⏩ somePackage".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestSkippedEvent(ctestSkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ somePackage",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And a CtestSkippedEvent has occurred for "package 1"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 1"
	 Then this text will be on the terminal "⏩ package 1\n⏳ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏳ package 2",
		)
	}, t)

	Test(`
	 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	 And 2 CtestSkippedEvents have occurred for each of them
	 And a PackagePassedEvent for package "package 1" has occurred
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 2"
	 Then this text will be on the terminal "⏩ package 1\n⏩ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		pack1Ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName 1")
		pack1Ctest2SkippedEvt := makeCtestSkippedEvent("package 1", "testName 2")
		pack2Ctest1SkippedEvt := makeCtestSkippedEvent("package 2", "testName 1")
		pack2Ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName 2")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestSkippedEvent(pack1Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack1Ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack2Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack2Ctest2SkippedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏩ package 2",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestSkippedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackagePassedEvent for packages "package 1",..., "package 3"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 4"
	 Then the printed text will be:
	 	"⏩ package 1\n⏩ package 2\n⏩ package 3\n⏩ package 4\n⏳ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")
		ctest3SkippedEvt := makeCtestSkippedEvent("package 3", "testName")
		ctest4SkippedEvt := makeCtestSkippedEvent("package 4", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest3SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest4SkippedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])
		interactor.HandlePackagePassed(packagePassedEvts["package 3"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 4"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏩ package 2\n⏩ package 3\n⏩ package 4\n⏳ package 5",
		)
	}, t)

	Test(`
	 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 5"
	 And a CtestSkippedEvent has occurred for packages "package 1", ..., "package 4"
	 And a PackagePassedEvent for packages "package 1",..., "package 5"
	 And there is a terminal with height 5
	 When a PackagePassedEvent for package "package 5"
	 Then the printed text will be:
	 	"✅ package 1\n✅ package 2\n✅ package 3\n✅ package 4\n✅ package 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2", "package 3", "package 4", "package 5")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")
		ctest3SkippedEvt := makeCtestSkippedEvent("package 3", "testName")
		ctest4SkippedEvt := makeCtestSkippedEvent("package 4", "testName")
		ctest5SkippedEvt := makeCtestSkippedEvent("package 5", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 5"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest3SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest4SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest5SkippedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])
		interactor.HandlePackagePassed(packagePassedEvts["package 3"])
		interactor.HandlePackagePassed(packagePassedEvts["package 4"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 5"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏩ package 2\n⏩ package 3\n⏩ package 4\n⏩ package 5",
		)
	}, t)

	Test(`
	Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestSkippedEvent has occurred for packages "pack 1", ..., "pack 6"
	And a PackagePassedEvent for packages "pack 1",..., "pack 5"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 6"
	Then the printed text will be: "⏩ pack 2\n⏩ pack 3\n⏩ pack 4\n⏩ pack 5\n⏩ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("pack 2", "testName")
		ctest3SkippedEvt := makeCtestSkippedEvent("pack 3", "testName")
		ctest4SkippedEvt := makeCtestSkippedEvent("pack 4", "testName")
		ctest5SkippedEvt := makeCtestSkippedEvent("pack 5", "testName")
		ctest6SkippedEvt := makeCtestSkippedEvent("pack 6", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest3SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest4SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest5SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest6SkippedEvt)

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
			"⏩ pack 2\n⏩ pack 3\n⏩ pack 4\n⏩ pack 5\n⏩ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 5"
	And a CtestSkippedEvent has occurred for packages "pack 1"
	And a PackagePassedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 1"
	Then the printed text will be:
		"⏩ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5")
		packagePassedEvts := makePackagePassedEvents("pack 1")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ pack 1\n⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestSkippedEvent has occurred for packages "pack 1", "pack 2"
	And a PackagePassedEvent for packages "pack 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "pack 2"
	Then the printed text will be:
		"⏩ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("pack 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)

	Test(`
	Given that 5 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 6"
	And a CtestSkippedEvent has occurred for packages "pack 1"
	And there is a terminal with height 5
	And a PackagePassedEvent for packages "pack 1"
	Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4", "pack 5", "pack 6")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 5"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 6"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3\n⏳ pack 4\n⏳ pack 5\n⏳ pack 6",
		)
	}, t)

	Test(`
	Given these events have occurred in this order:
	- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
	- 1 CtestFailedEvent has occurred for "package 1"
	- 1 CtestSkippedEvent has occurred for "package 2"
	- 1 PackageFailedEvent has ocurred for "package 1"
	And there is a terminal with height 5
	When a PackagePassedEvent for package "package 2" occurrs
	Then this text will be on the terminal "❌ package 1\n⏩ package 2".`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")

		packageFailedEvts := makePackageFailedEvents("package 1")
		packagePassedEvts := makePackagePassedEvents("package 2")

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n⏩ package 2",
		)
	}, t)
}

func TestSkippedPackages_TerminalHeightGreaterThan5(t *testing.T) {
	assert := assert.New(t)

	Test(`
		 Given that a PackageStartedEvent has occurred for "somePackage"
		 And a CtestSkippedEvent for test with name "testName" in package "somePackage" has occurred
		 And there is a terminal with height 6
		 When a PackagePassedEvent for package "somePackage" occurs
		 Then this text will be on the terminal "⏩ somePackage" and the summary of tests
		 "\n\nPackages: 0 running, 1 skipped\nTests: 1 skipped\nTime: 0.000s"`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestSkippedEvent(ctestSkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		 Given that 2 PackageStartedEvent have occurred for packages "pack 1" and "pack 2"
		 And a CtestSkippedEvent has occurred for "pack 1"
		 And there is a terminal with height 6
		 When a PackagePassedEvent for package "pack 1"
		 Then this text will be on the terminal "⏩ package 1\n⏳ package 2" and the summary of tests
		 "\n\nPackages: 1 running, 1 skipped\nTests: 1 skipped\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		packagePassedEvts := makePackagePassedEvents("package 1")
		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏳ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		 Given that 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
		 And a CtestSkippedEvent has occurred for each of them
		 And a PackagePassedEvent for package "package 1" has occurred
		 And there is a terminal with height 6
		 When a PackagePassedEvent for package "package 2"
		 Then this text will be on the terminal "⏩ package 1\n⏩ package 2" and the summary of tests
		 "\n\nPackages: 0 running, 2 passed\nTests: 0 running\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 1\n⏩ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		 Given that 5 PackageStartedEvent have occurred for packages "package 1", ..., "package 3"
		 And a CtestSkippedEvent has occurred for packages "package 1"
		 And a PackagePassedEvent for packages "package 1", "package 2"
		 And there is a terminal with height 6
		 When a PackagePassedEvent for package "package 2"
		 Then the printed text will be: "⏩ package 2\n⏳ package 3" and the summary of tests
		 "\n\nPackages: 1 running, 2 skipped\nTests: 2 skipped\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2", "package 3")
		packagePassedEvts := makePackagePassedEvents("package 1", "package 2")
		ctest1SkippedEvt := makeCtestSkippedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 3"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ package 2\n⏳ package 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 1 running, "+
				ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"2 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		Given that 6 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
		And a CtestSkippedEvent has occurred for packages "pack 1", ..., "pack 3"
		And a PackagePassedEvent for packages "pack 1", and "pack 2"
		And there is a terminal with height 6
		When a PackagePassedEvent for package "pack 6"
		Then the printed text will be: "⏩ pack 2\n⏩ pack 3" and the summary of tests
		"\n\nPackages: 0 running, 3 skipped\nTests: 3 skipped\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packagePassedEvts := makePackagePassedEvents("pack 1", "pack 2", "pack 3")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("pack 2", "testName")
		ctest3SkippedEvt := makeCtestSkippedEvent("pack 3", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(ctest3SkippedEvt)

		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])
		interactor.HandlePackagePassed(packagePassedEvts["pack 2"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 3"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ pack 2\n⏩ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.YELLOW+"3 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"3 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		Given that 3 PackageStartedEvent have occurred for packages "pack 1", ..., "pack 3"
		And a CtestSkippedEvent has occurred for packages "pack 1"
		And there is a terminal with height 6
		And a PackagePassedEvent for packages "pack 1"
		Then the printed text will be: "⏳ pack 2\n⏳ pack 3\n" and the summary of tests
		"\n\nPackages: 2 running, 1 skipped\nTests: 1 skipped\nTime: 0.000s.`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3")
		packagePassedEvts := makePackagePassedEvents("pack 1")
		ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])

		interactor.HandleCtestSkippedEvent(ctest1SkippedEvt)

		// When
		interactor.HandlePackagePassed(packagePassedEvts["pack 1"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏳ pack 2\n⏳ pack 3"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 2 running, "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)

	Test(`
		Given these events have occurred in this order:
		- 2 PackageStartedEvent have occurred for packages "package 1" and "package 2"
		- 1 CtestFailedEvent has occurred for "package 1"
		- 1 CtestSkippedEvent has occurred for "package 2"
		- 1 PackageFailedEvent has ocurred for "package 1"
		And there is a terminal with height 6
		When a PackagePassedEvent for package "package 2" occurrs
		Then this text will be on the terminal "❌ package 1\n⏩ package 2" and the summary of tests
		"\n\nPackages: 0 running, 1 failed, 1 passed\nTests: 0 running\nTime: 0.000s`, func(t *testing.T) {
		packStartedEvts := makePackageStartedEvents("package 1", "package 2")
		ctest1FailedEvt := makeCtestFailedEvent("package 1", "testName")
		ctest2SkippedEvt := makeCtestSkippedEvent("package 2", "testName")

		packageFailedEvts := makePackageFailedEvents("package 1")
		packagePassedEvts := makePackagePassedEvents("package 2")

		// Given
		interactor, fakeTerminal, _ := setup(6)
		interactor.HandlePackageStartedEvent(packStartedEvts["package 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["package 2"])
		interactor.HandleCtestFailedEvent(ctest1FailedEvt)
		interactor.HandleCtestSkippedEvent(ctest2SkippedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["package 1"])

		// When
		interactor.HandlePackagePassed(packagePassedEvts["package 2"])

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ package 1\n⏩ package 2"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" 0 running, "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     0.000s",
		)
	}, t)
}

func TestHandleNoPackageTestsFoundEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have occurred
	And there is a terminal with height 5
	When a NoPackageTestsFoundEvent for package "somePackage" occurs
	Then the user should see an error in the terminal.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(5)

		// When
		noPackTestsFoundEvt := events.NewNoPackageTestsFoundEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
			},
		)
		err := eventsHandler.HandleNoPackageTestsFoundEvent(noPackTestsFoundEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			terminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage" has occured
	And there is a terminal with height 5
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should not see anything on the terminal.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(5)
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		// When
		noPackTestsFoundEvt := events.NewNoPackageTestsFoundEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
			},
		)
		eventsHandler.HandleNoPackageTestsFoundEvent(noPackTestsFoundEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage 1" has occured
	And a PackageStartedEvent for package "somePackage 2" has occured
	And there is a terminal with height 5
	When a NoPackageTestsFoundEvent for packag "somePackage 1" occurs
	Then the user should only see that the the tests for "somePackage 2" are running.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(5)
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		packStartedEvt2 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 2",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt2)

		// When
		noPackTestsFoundEvt1 := events.NewNoPackageTestsFoundEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandleNoPackageTestsFoundEvent(noPackTestsFoundEvt1)
		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ somePackage 2\n",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage" has occured
	And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	And there is a terminal with height 5
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should see an error in the terminal
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(5)
		timeElapsed := 1.2
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		ctestPassedEvt := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt)

		// When
		noPackTestsFoundEvt := events.NewNoPackageTestsFoundEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
			},
		)
		err := eventsHandler.HandleNoPackageTestsFoundEvent(noPackTestsFoundEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			terminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage" has occured
	And a CtestFailedEvent for test with name "testName" in package "somePackage" has occurred
	And there is a terminal with height 5
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should see an error in the terminal
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup(5)
		timeElapsed := 1.2
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		ctestFaileddEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestFailedEvent(ctestFaileddEvt)

		// When
		noPackTestsFoundEvt := events.NewNoPackageTestsFoundEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
			},
		)
		err := eventsHandler.HandleNoPackageTestsFoundEvent(noPackTestsFoundEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			terminal.Text(),
			"❗ Error.",
		)
	}, t)
}

func TestTestingFinishedSummary(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	And a PackagePassedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then this text will be on the terminal "✅ somePackage" and the summary of tests
	"\n\nPackages: 1 passed, 1 total\nTests: 1 passed, 1 totalt1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		testingStartedEvt := events.NewTestingSt\nTime: 1.200s"`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1200))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestPassedEvent(ctestPassedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a CtestSkippedEvent for test with name "testName" in package "somePackage" has occurred
	And a PackagePassedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.372s occurs
	Then this text will be on the terminal "⏩ somePackage" and the summary of tests
	"\n\nPackages: 1 skipped, 1 total\nTests: 1 skipped, 1 total\nTime: 1.372s"`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "testName")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1372))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestSkippedEvent(ctestSkippedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.372s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a CtestFailedEvent for test with name "testName" in package "somePackage" has occurred
	And a PackageFailedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then this text will be on the terminal "❌ somePackage" and the summary of tests
	"\n\nPackages: 1 failed, 1 total\nTests: 1 failed, 1 total\nTime: 1.200s"`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "testName")
		packageFailedEvts := makePackageFailedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1200))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestFailedEvent(ctestFailedEvt)
		interactor.HandlePackageFailed(packageFailedEvts["somePackage"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ somePackage\n\n"+
				"  "+ansi_escape.RED+"● testName"+ansi_escape.COLOR_RESET+"\n"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a CtestOutputEvent for test "testName" of package "somePackage" with out "Some output" has occurred
	And a CtestFailedEvent for test with name "testName" of package "somePackage" has occurred
	And a PackageFailedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then this text will be on the terminal "❌ somePackage" and the summary of tests
	"\n\nPackages: 1 failed, 1 total\nTests: 1 failed, 1 total\nTime: 1.200s"`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "testName")
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "testName", "Some output")
		packageFailedEvts := makePackageFailedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1200))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestFailedEvent(ctestFailedEvt)
		interactor.HandleCtestOutputEvent(ctestOutputEvt)
		interactor.HandlePackageFailed(packageFailedEvts["somePackage"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"❌ somePackage\n\n"+
				"  "+ansi_escape.RED+"● testName"+ansi_escape.COLOR_RESET+"\n\n"+
				"  Some output\n"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"1 failed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a CtestPassedEvent for "test 1" in "somePackage" has occurred
	And a CtestSkippedEvent for "test 2" in "somePackage" has occurred
	And a PackagePassedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then this text will be on the terminal "✅ somePackage" and the summary of tests
	"\n\nPackages: 1 passed, 1 total\nTests: 1 skipped, 1 passed, 2 total\nTime: 1.200s`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "test 1")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "test 2")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1200))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandleCtestPassedEvent(ctestPassedEvt)
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])
		interactor.HandleCtestSkippedEvent(ctestSkippedEvt)

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 2 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that a TestingStartedEvent occured with timestamp t1
	And a PackageStartedEvent has occurred for "somePackage"
	And a PackagePassedEvent for package "somePackage" occurs
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.2s occurs
	Then this text will be on the terminal "✅ somePackage" and the summary of tests
	"\n\nPackages: 1 passed, 1 total\nTests: 1 skipped, 1 passed, 2 total\nTime: 1.200s`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("somePackage")
		packagePassedEvts := makePackagePassedEvents("somePackage")
		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1200))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["somePackage"])
		interactor.HandlePackagePassed(packagePassedEvts["somePackage"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"⏩ somePackage"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", 1 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    0 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.200s\n"+
				"Ran all tests.",
		)
	}, t)

	Test(`
	Given that these events have occurred in this order:
	- 1 TestingStartedEvent with timestamp t1
	- 4 PackageStartedEvent has occurred for "pack 1", ..., "pack 4"
	- 7 CtestSkippedEvents:  2x"pack 1", 1x"pack 2", 1x"pack 3", 3x"pack 4"
	- 4 CtestPassedEvents: 2x"pack 1",  2x"pack 3"
	- 3 CtestFailedEvent:  2x"pack 2", 1x"pack 3"
	- 2 PackagePassedEvent: 1x"pack 1", 1x"pack 4"
	- 2 PackageFailedEvent: 1x"pack 2", 1x"pack 3"
	And there is a terminal with height 5
	When a TestingFinishedEvent with a timestamp of t1+1.372s occurs
	Then this text will be on the terminal "✅ pack 1\n❌pack 2\n❌ pack 3\n⏩ pack 4" and the summary of tests
	"\n\nPackages: 1 skipped\nTests: 1 skipped\nTime: 1.372s"`, func(t *testing.T) {
		t1 := time.Now()
		testingStartedEvt := events.NewTestingStartedEvent(t1)
		packStartedEvts := makePackageStartedEvents("pack 1", "pack 2", "pack 3", "pack 4")
		packPassedEvts := makePackagePassedEvents("pack 1", "pack 4")
		packFailedEvts := makePackageFailedEvents("pack 2", "pack 3")
		pack1Ctest1SkippedEvt := makeCtestSkippedEvent("pack 1", "testName 1")
		pack1Ctest2SkippedEvt := makeCtestSkippedEvent("pack 1", "testName 2")
		pack2Ctest1SkippedEvt := makeCtestSkippedEvent("pack 2", "testName 3")
		pack3Ctest1SkippedEvt := makeCtestSkippedEvent("pack 3", "testName 4")
		pack4Ctest1SkippedEvt := makeCtestSkippedEvent("pack 4", "testName 5")
		pack4Ctest2SkippedEvt := makeCtestSkippedEvent("pack 4", "testName 6")
		pack4Ctest3SkippedEvt := makeCtestSkippedEvent("pack 4", "testName 7")
		pack1Ctest1PassedEvt := makeCtestPassedEvent("pack 1", "testName 8")
		pack1Ctest2PassedEvt := makeCtestPassedEvent("pack 1", "testName 9")
		pack3Ctest1PassedEvt := makeCtestPassedEvent("pack 3", "testName 10")
		pack3Ctest2PassedEvt := makeCtestPassedEvent("pack 3", "testName 11")
		pack2Ctest1FailedEvt := makeCtestFailedEvent("pack 2", "testName 12")
		pack2Ctest2FailedEvt := makeCtestFailedEvent("pack 2", "testName 13")
		pack3Ctest1FailedEvt := makeCtestFailedEvent("pack 3", "testName 14")

		testingFinishedEvt := events.NewTestingFinishedEvent(t1.Add(time.Millisecond * 1372))

		// Given
		interactor, fakeTerminal, _ := setup(5)
		interactor.HandleTestingStarted(testingStartedEvt)
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 1"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 2"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 3"])
		interactor.HandlePackageStartedEvent(packStartedEvts["pack 4"])
		interactor.HandleCtestSkippedEvent(pack1Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack1Ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack2Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack3Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack4Ctest1SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack4Ctest2SkippedEvt)
		interactor.HandleCtestSkippedEvent(pack4Ctest3SkippedEvt)
		interactor.HandleCtestPassedEvent(pack1Ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(pack1Ctest2PassedEvt)
		interactor.HandleCtestPassedEvent(pack3Ctest1PassedEvt)
		interactor.HandleCtestPassedEvent(pack3Ctest2PassedEvt)
		interactor.HandleCtestFailedEvent(pack2Ctest1FailedEvt)
		interactor.HandleCtestFailedEvent(pack2Ctest2FailedEvt)
		interactor.HandleCtestFailedEvent(pack3Ctest1FailedEvt)
		interactor.HandlePackagePassed(packPassedEvts["pack 1"])
		interactor.HandlePackagePassed(packPassedEvts["pack 4"])
		interactor.HandlePackageFailed(packFailedEvts["pack 2"])
		interactor.HandlePackageFailed(packFailedEvts["pack 3"])

		// When
		interactor.HandleTestingFinished(testingFinishedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"✅ pack 1\n"+
				"❌ pack 2\n\n"+
				"  "+ansi_escape.RED+"● testName 12"+ansi_escape.COLOR_RESET+"\n\n"+
				"  "+ansi_escape.RED+"● testName 13"+ansi_escape.COLOR_RESET+"\n\n"+
				"❌ pack 3\n\n"+
				"  "+ansi_escape.RED+"● testName 14"+ansi_escape.COLOR_RESET+"\n\n"+
				"⏩ pack 4"+
				"\n\n"+ansi_escape.BOLD+"Packages:"+ansi_escape.RESET_BOLD+" "+
				ansi_escape.RED+"2 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"1 skipped"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"1 passed"+ansi_escape.COLOR_RESET+", 4 total"+
				"\n"+ansi_escape.BOLD+"Tests:"+ansi_escape.RESET_BOLD+"    "+
				ansi_escape.RED+"3 failed"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.YELLOW+"7 skipped"+ansi_escape.COLOR_RESET+", "+
				ansi_escape.GREEN+"4 passed"+ansi_escape.COLOR_RESET+", 14 total"+
				"\n"+ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     1.372s\n"+
				"Ran all tests.",
		)
	}, t)
}
