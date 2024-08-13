package cmd

import (
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/redjolr/goherent/cmd/concurrent_events_handler"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events_handler"
	"github.com/redjolr/goherent/cmd/testing_finished_handler"
	"github.com/redjolr/goherent/cmd/testing_started_handler"
	"github.com/redjolr/goherent/internal/consolesize"
	"github.com/redjolr/goherent/terminal"
)

func Main(extraCmdArgs []string) int {
	router := setup()

	testCmd := NewTestCmd(extraCmdArgs)
	testCmd.Exec()
	router.RouteTestingStartedEvent(time.Now())

	for testCmd.IsRunning() {
		var jsonEvt events.JsonEvent
		output := testCmd.NextOutput()
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

	var sequentialEventsOutputPort sequential_events_handler.OutputPort
	var ansiTerminal terminal.AnsiTerminal
	if terminalHeight != 0 {
		ansiTerminal = terminal.NewAnsiTerminal(terminalWidth, terminalHeight)
		sequentialEventsOutputPort = sequential_events_handler.NewBoundedTerminalPresenter(&ansiTerminal)
	} else {
		ansiTerminal = terminal.NewAnsiTerminal(math.MaxInt, math.MaxInt)
		sequentialEventsOutputPort = sequential_events_handler.NewTerminalPresenter(&ansiTerminal)
	}

	concurrentEventsPresenter := concurrent_events_handler.NewTerminalPresenter(&ansiTerminal)
	testingFinishedPresenter := testing_finished_handler.NewTerminalPresenter(&ansiTerminal)
	testingStartedPresenter := testing_started_handler.NewTerminalPresenter(&ansiTerminal)
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	sequentialEventsHandler := sequential_events_handler.NewEventsHandler(sequentialEventsOutputPort, &ctestsTracker)
	concurrentEventsHandler := concurrent_events_handler.NewEventsHandler(&concurrentEventsPresenter, &ctestsTracker)
	testingFinishedHandler := testing_finished_handler.NewEventsHandler(&testingFinishedPresenter, &ctestsTracker)
	testingStartedHandler := testing_started_handler.NewEventsHandler(&testingStartedPresenter)
	sequentialEventsRouter := sequential_events_handler.NewSequentialEventsRouter(&sequentialEventsHandler)
	concurrentEventsRouter := concurrent_events_handler.NewConcurrentEventsRouter(&concurrentEventsHandler)

	return NewRouter(&sequentialEventsRouter, &concurrentEventsRouter, &testingStartedHandler, &testingFinishedHandler)
}
