package cmd

import (
	"encoding/json"
	"log"

	"github.com/redjolr/goherent/cmd/concurrent_events"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/internal/consolesize"
	"github.com/redjolr/goherent/terminal"
)

func Main(extraCmdArgs []string) int {
	router := setup()

	testCmd := NewTestCmd(extraCmdArgs)
	testCmd.
		NonVerbose().
		Exec()
	router.RouteTestingStartedEvent(testCmd.RunsTestsConcurrently())
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
	router.RouteTestingFinishedEvent(testCmd.RunsTestsConcurrently())
	return testCmd.ExitCode()
}

func setup() *Router {
	consoleWidth, consoleHeight := consolesize.GetConsoleSize()
	var ansiTerminal terminal.AnsiTerminal
	if consoleHeight != 0 {
		ansiTerminal = terminal.NewBoundedAnsiTerminal(consoleWidth, consoleHeight)
	} else {
		ansiTerminal = terminal.NewUnboundedAnsiTerminal()
	}
	sequentialEventsRouter := sequential_events.Setup(&ansiTerminal)
	concurrentEventsRouter := concurrent_events.Setup(&ansiTerminal)
	router := NewRouter(sequentialEventsRouter, concurrentEventsRouter)
	return &router
}
