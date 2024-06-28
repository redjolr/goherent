package cmd_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/console"
	"github.com/redjolr/goherent/console/terminal"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func setupE2e() (cmd.EventsHandler, terminal.FakeAnsiTerminal) {
	ctestTracker := ctests_tracker.NewCtestsTracker()
	fakeAnsiTerminal := terminal.NewFakeAnsiTerminal()

	outputConsole := console.NewConsole(&fakeAnsiTerminal)
	terminalPresenter := cmd.NewTerminalPresenter(&outputConsole)
	eventHandler := cmd.NewEventsHandler(&terminalPresenter, &ctestTracker)

	return eventHandler, fakeAnsiTerminal
}

func TestE2eCtestRanEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that no events have happened
	When a CtestRanEvent occurs with test name "testName" from "packageName"
	Then the user should be informed that the testing of a new package started and
	that the first test of that package started running
	`, func(t *testing.T) {
		// Given
		eventsHandler, ansiTerminal := setupE2e()

		// When
		ctestRanEvt := ctest_ran_event.NewFromJsonTestEvent(events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Test:    "testName",
			Package: "somePackage",
		})
		eventsHandler.HandleCtestRanEvt(ctestRanEvt)
		fmt.Println("\n\n\n HELLO", ansiTerminal.Text())

		// Then
		assert.Equal(
			ansiTerminal.Text(),
			"📦⏳ somePackage\n\t⏳ testName\n\n",
		)
	}, t)

}
