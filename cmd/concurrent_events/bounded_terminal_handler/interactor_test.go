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

func makeNPackageStartedEvents(packageNames ...string) map[string]events.PackageStartedEvent {
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
		packStartedEvents := makeNPackageStartedEvents("somePackage 1", "somePackage 2")

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
	Given that no events have occurred
	And we have a bounded terminal with height 5
	When 3 HandlePackageStartedEvent for packages "package 1", ..., "package 5" occur
	And the printed text should be "⏳ package 1\n⏳ package 2\n⏳ package 3\n⏳ package 4\n⏳ package 5"`, func(t *testing.T) {
		packStartedEvents := makeNPackageStartedEvents("package 1", "package 2", "package 3", "package 4", "package 5")

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
		packStartedEvents := makeNPackageStartedEvents(
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
		packStartedEvents := makeNPackageStartedEvents("package 1", "package 2")
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
		packStartedEvents := makeNPackageStartedEvents("package 1", "package 2", "package 3")
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
