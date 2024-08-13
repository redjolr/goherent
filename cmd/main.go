package cmd

import (
	"encoding/json"
	"log"
	"math"
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

func setup() Router {
	terminalWidth, terminalHeight := consolesize.GetConsoleSize()

	var sequentialEventsOutputPort sequential_events.OutputPort
	var ansiTerminal terminal.AnsiTerminal
	if terminalHeight != 0 {
		ansiTerminal = terminal.NewAnsiTerminal(terminalWidth, terminalHeight)
		sequentialEventsOutputPort = sequential_events.NewBoundedTerminalPresenter(&ansiTerminal)
	} else {
		ansiTerminal = terminal.NewAnsiTerminal(math.MaxInt, math.MaxInt)
		sequentialEventsOutputPort = sequential_events.NewTerminalPresenter(&ansiTerminal)
	}

	concurrentEventsPresenter := concurrent_events.NewUnboundedTerminalPresenter(&ansiTerminal)
	testingFinishedPresenter := testing_finished.NewTerminalPresenter(&ansiTerminal)
	testingStartedPresenter := testing_started.NewTerminalPresenter(&ansiTerminal)
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	sequentialEventsHandler := sequential_events.NewHandler(sequentialEventsOutputPort, &ctestsTracker)
	concurrentEventsHandler := concurrent_events.NewHandler(&concurrentEventsPresenter, &ctestsTracker)
	testingFinishedHandler := testing_finished.NewHandler(&testingFinishedPresenter, &ctestsTracker)
	testingStartedHandler := testing_started.NewEventsHandler(&testingStartedPresenter)
	sequentialEventsRouter := sequential_events.NewRouter(&sequentialEventsHandler)
	concurrentEventsRouter := concurrent_events.NewRouter(&concurrentEventsHandler)

	return NewRouter(&sequentialEventsRouter, &concurrentEventsRouter, &testingStartedHandler, &testingFinishedHandler)
}
