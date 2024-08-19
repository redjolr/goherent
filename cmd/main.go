package cmd

import (
	"encoding/json"
	"log"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/cmd/testing_finished"
	"github.com/redjolr/goherent/cmd/testing_started"
	"github.com/redjolr/goherent/internal/consolesize"
	"github.com/redjolr/goherent/terminal"
)

func Main(extraCmdArgs []string) int {
	router := setup()

	testCmd := NewTestCmd(extraCmdArgs)
	testCmd.
		NonVerbose().
		Exec()
	router.RouteTestingStartedEvent(time.Now())
	for testCmd.IsRunning() {
		var jsonEvt events.JsonEvent
		output := testCmd.Output()
		err := json.Unmarshal([]byte(output), &jsonEvt)

		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		router.Route(jsonEvt, testCmd.RunsTestsConcurrently())
	}
	testCmd.Wait()
	router.RouteTestingFinishedEvent(testCmd.ExecutionTime())
	return 0
}

func setup() *Router {
	consoleWidth, consoleHeight := consolesize.GetConsoleSize()
	var ansiTerminal terminal.AnsiTerminal
	if consoleHeight != 0 {
		ansiTerminal = terminal.NewBoundedAnsiTerminal(consoleWidth, consoleHeight)
	} else {
		ansiTerminal = terminal.NewUnboundedAnsiTerminal()
	}
	testingFinishedPresenter := testing_finished.NewTerminalPresenter(&ansiTerminal)
	testingStartedPresenter := testing_started.NewTerminalPresenter(&ansiTerminal)
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	testingFinishedInteractor := testing_finished.NewInteractor(&testingFinishedPresenter, &ctestsTracker)
	testingStartedHandler := testing_started.NewEventsHandler(&testingStartedPresenter)
	sequentialEventsRouter := sequential_events.Setup(&ansiTerminal)
	concurrentEventsRouter := concurrent_events.Setup(&ansiTerminal)
	router := NewRouter(sequentialEventsRouter, concurrentEventsRouter, &testingStartedHandler, &testingFinishedInteractor)
	return &router
}
