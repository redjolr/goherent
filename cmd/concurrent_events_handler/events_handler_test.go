package concurrent_events_handler_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/package_passed_event"
	"github.com/redjolr/goherent/cmd/events/package_started_event"
	"github.com/redjolr/goherent/cmd/events/testing_started_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal"
	"github.com/stretchr/testify/assert"
)

func setup() (*concurrent_events_handler.EventsHandler, *terminal.FakeAnsiTerminal, *ctests_tracker.CtestsTracker) {
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal()
	fakeAnsiTerminalPresenter := concurrent_events_handler.NewTerminalPresenter(&fakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	eventsHandler := concurrent_events_handler.NewEventsHandler(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &eventsHandler, &fakeAnsiTerminal, &ctestTracker
}

func TestHandleTestingStarted(t *testing.T) {
	assert := assert.New(t)
	Test("User should be informed, that the testing has started", func(t *testing.T) {
		eventsHandler, terminal, _ := setup()
		now := time.Now()
		testingStartedEvt := testing_started_event.NewTestingStartedEvent(now)
		eventsHandler.HandleTestingStarted(testingStartedEvt)

		assert.Equal(
			terminal.Text(),
			fmt.Sprintf("\n🚀 Starting... %s", now.Format("2006-01-02 15:04:05.000")),
		)
	}, t)
}

func TestHandlePackageStartedEvent(t *testing.T) {
	assert := assert.New(t)
	Test(`
     Given that no events have occurred
     When a HandlePackageStartedEvent occurs for package "somePackage"
     Then the user should be informed that the tests for that package are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()

		// When
		packStartedEvt := package_started_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n⏳ somePackage",
		)
	}, t)

	Test(`
     Given that a HandlePackageStartedEvent for package "somePackage" has occurred
     When a HandlePackageStartedEvent occurs for package "somePackage"
     Then the user should be informed only once that the tests for the "somePackage" package are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		packStartedEvt := package_started_event.NewFromJsonTestEvent(
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
			"\n⏳ somePackage",
		)
	}, t)

	Test(`
     Given that a HandlePackageStartedEvent for package "somePackage 1" has occured
     When a HandlePackageStartedEvent for package "somePackage 2" occurs
     Then the user should be informed that the tests for that package are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		packStartedEvt1 := package_started_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// When
		packStartedEvt2 := package_started_event.NewFromJsonTestEvent(
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
			"\n⏳ somePackage 1\n⏳ somePackage 2",
		)
	}, t)

	Test(`
     Given that a PackageStartedEvent for package "somePackage 1" has occured
     When a PackageStartedEvent for package "somePackage 2" occurs
     Then the user should be informed that the tests for that package are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		packStartedEvt1 := package_started_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// When
		packStartedEvt2 := package_started_event.NewFromJsonTestEvent(
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
			"\n⏳ somePackage 1\n⏳ somePackage 2",
		)
	}, t)
}

func TestHandlePackagePassedEvent(t *testing.T) {
	assert := assert.New(t)
	Test(`
	 Given that no events have occurred
	 When a PackagePassedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2

		// When
		packagePassedEvt := package_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		err := eventsHandler.HandlePackagePassed(packagePassedEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
     Given that a PackageStartedEvent has occurred for "somePackage"
     When a PackagePassedEvent for package "somePackage" occurs
     Then the current screen will be erased
     And the user will be informed that the package tests have passed
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt1 := package_started_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "output",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// When
		packagePassedEvt := package_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackagePassed(packagePassedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"\n✅ somePackage",
		)
	}, t)
}
