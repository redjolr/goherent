package unbounded_terminal_handler_test

import (
	"math"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events/unbounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"

	. "github.com/redjolr/goherent/pkg"
	"github.com/redjolr/goherent/terminal/fake_ansi_terminal"
	"github.com/stretchr/testify/assert"
)

func setup() (*unbounded_terminal_handler.Interactor, *fake_ansi_terminal.FakeAnsiTerminal, *ctests_tracker.CtestsTracker) {
	fakeAnsiTerminal := fake_ansi_terminal.NewFakeAnsiTerminal(math.MaxInt, math.MaxInt)
	fakeAnsiTerminalPresenter := unbounded_terminal_handler.NewPresenter(&fakeAnsiTerminal)
	ctestTracker := ctests_tracker.NewCtestsTracker()
	interactor := unbounded_terminal_handler.NewInteractor(&fakeAnsiTerminalPresenter, &ctestTracker)
	return &interactor, &fakeAnsiTerminal, &ctestTracker
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
		packStartedEvt := events.NewPackageStartedEvent(
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
			"\n⏳ somePackage",
		)
	}, t)

	Test(`
	 Given that a HandlePackageStartedEvent for package "somePackage 1" has occured
	 When a HandlePackageStartedEvent for package "somePackage 2" occurs
	 Then the user should be informed that the tests for "somePackage 2" are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
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
			"\n⏳ somePackage 1\n⏳ somePackage 2",
		)
	}, t)

	Test(`
     Given that a PackageStartedEvent for package "somePackage 1" has occured
	 And a CtestPassedEvent for test with name "testName" in package "somePackage 1" has occurred
	 And a PackagePassedEvent for package "somePackage 1" has occurred
     When a PackageStartedEvent for package "somePackage 2" occurs
     Then the user should be informed that the tests for "somePackage 2" are running`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt1)

		packagePassedEvt1 := events.NewPackagePassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackagePassed(packagePassedEvt1)

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
			"\n✅ somePackage 1\n⏳ somePackage 2",
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
		packagePassedEvt := events.NewPackagePassedEvent(
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
	 And the user will be informed that an error has occurred.
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		// When
		packagePassedEvt := events.NewPackagePassedEvent(
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
	 And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	 When a PackagePassedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
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
		packagePassedEvt := events.NewPackagePassedEvent(
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

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestSkippedEvent for test with name "testName" in package "somePackage" has occurred
	 When a PackagePassedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		ctestSkippedEvt := events.NewCtestSkippedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "skip",
				Test:    "testName",
				Package: "somePackage",
			},
		)
		eventsHandler.HandleCtestSkippedEvt(ctestSkippedEvt)

		// When
		packagePassedEvt := events.NewPackagePassedEvent(
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
			"\n⏩ somePackage",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent for package "somePackage 1" has occured
	 And a CtestPassedEvent for test "testName" from package "somePackage 2" has occurred
	 And a PackageStartedEvent for package "somePackage 2" has occured
	 When a PackagePassedEvent for package "somePackage 2" occurs again
	 Then the user should be informed that the tests are running for "somePackage 1" and passed for "somePackage 2"`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		ctestPassedEvt2 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage 2",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt2)

		packStartedEvt2 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 2",
			},
		)

		eventsHandler.HandlePackageStartedEvent(packStartedEvt2)

		// When
		packPassedEvt2 := events.NewPackagePassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 2",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackagePassed(packPassedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n⏳ somePackage 1\n✅ somePackage 2",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent for package "somePackage 1" has occured
	 And a PackageStartedEvent for package "somePackage 2" has occured
	 When a PackagePassedEvent for package "somePackage 1" occurs
	 Then the user should be informed that the tests are running for "somePackage 2" and passed for "somePackage 1"`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt1 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 1",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt1)

		ctestPassedEvt1 := events.NewCtestPassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Test:    "testName",
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestPassedEvent(ctestPassedEvt1)

		packStartedEvt2 := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage 2",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt2)

		// When
		packPassedEvt1 := events.NewPackagePassedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackagePassed(packPassedEvt1)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n✅ somePackage 1\n⏳ somePackage 2",
		)
	}, t)
}

func TestHandlePackageFailedEvent(t *testing.T) {
	assert := assert.New(t)
	Test(`
	 Given that no events have occurred
	 When a PackageFailedEvent for package "somePackage" occurs
	 Then an error will be presented to the user.
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2

		// When
		packageFailedEvt := events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		err := eventsHandler.HandlePackageFailed(packageFailedEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 When a PackageFailedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have passed
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		timeElapsed := 1.2
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)

		// When
		packageFailedEvt := events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Elapsed: &timeElapsed,
			},
		)
		err := eventsHandler.HandlePackageFailed(packageFailedEvt)

		// Then
		assert.Error(err)
		assert.Contains(
			fakeTerminal.Text(),
			"❗ Error.",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent has occurred for "somePackage"
	 And a CtestFailedEvent has occurred for test "testName" in package "somePackage"
	 When a PackageFailedEvent for package "somePackage" occurs
	 And the user will be informed that the package tests have failed
	`, func(t *testing.T) {
		// Given
		eventsHandler, fakeTerminal, _ := setup()
		elapsedTime := 1.2
		packStartedEvt := events.NewPackageStartedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "start",
				Package: "somePackage",
			},
		)
		eventsHandler.HandlePackageStartedEvent(packStartedEvt)
		ctestFailedEvt := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandleCtestFailedEvent(ctestFailedEvt)
		// When
		packageFailedEvt := events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage",
				Elapsed: &elapsedTime,
			},
		)
		eventsHandler.HandlePackageFailed(packageFailedEvt)

		// Then
		assert.Equal(
			fakeTerminal.Text(),
			"\n❌ somePackage",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent for package "somePackage 1" has occured
	 And a PackageStartedEvent for package "somePackage 2" has occured
	 And a CtestFailedEvent for test with name "testName" in package "somePackage 2" occurs
	 When a PackageFailedEvent for package "somePackage 2" occurs
	 Then the user should be informed that the tests are running for "somePackage 1" and failed for "somePackage 2"`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		timeElapsed := 1.2
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

		ctestFailedEvt2 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage 2",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestFailedEvent(ctestFailedEvt2)

		// When
		packageFailedEvt2 := events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 2",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackageFailed(packageFailedEvt2)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n⏳ somePackage 1\n❌ somePackage 2",
		)
	}, t)

	Test(`
	 Given that a PackageStartedEvent for package "somePackage 1" has occured
	 And a PackageStartedEvent for package "somePackage 2" has occured
	 And a CtestFailedEvent for test with name "testName" in package "somePackage 1" occurs
	 When a PackagePassedEvent for package "somePackage 1" occurs
	 Then the user should be informed that the tests are running for "somePackage 2" and failed for "somePackage 1"`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
		timeElapsed := 1.2
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

		ctestFailedEvt1 := events.NewCtestFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "fail",
				Test:    "testName",
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandleCtestFailedEvent(ctestFailedEvt1)

		// When
		packageFailedEvt1 := events.NewPackageFailedEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: "somePackage 1",
				Elapsed: &timeElapsed,
			},
		)
		eventsHandler.HandlePackageFailed(packageFailedEvt1)

		// Then
		assert.Equal(
			terminal.Text(),
			"\n❌ somePackage 1\n⏳ somePackage 2",
		)
	}, t)
}

func TestHandleNoPackageTestsFoundEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have occurred
	When a NoPackageTestsFoundEvent for package "somePackage" occurs
	Then the user should see an error in the terminal.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()

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
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should not see anything on the terminal.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
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
			"\n             ",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage 1" has occured
	And a PackageStartedEvent for package "somePackage 2" has occured
	When a NoPackageTestsFoundEvent for packag "somePackage 1" occurs
	Then the user should only see that the the tests for "somePackage 2" are running.
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
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
			"\n⏳ somePackage 2\n               ",
		)
	}, t)

	Test(`
	Given that a PackageStartedEvent for package "somePackage" has occured
	And a CtestPassedEvent for test with name "testName" in package "somePackage" has occurred
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should see an error in the terminal
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
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
	When a NoPackageTestsFoundEvent for the same package occurs
	Then the user should see an error in the terminal
	`, func(t *testing.T) {
		// Given
		eventsHandler, terminal, _ := setup()
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
