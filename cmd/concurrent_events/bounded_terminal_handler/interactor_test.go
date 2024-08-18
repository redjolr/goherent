package bounded_terminal_handler_test

import (
	"math"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events/bounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	. "github.com/redjolr/goherent/pkg"
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

func TestHandlePackageStartedEvent(t *testing.T) {
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
		// Given
		eventsHandler, terminal, _ := setup(1)
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// When
		packStartedEvt2 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 2",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"⏳ somePackage 1",
		)
	}, t)
}
