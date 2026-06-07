package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"slices"
	"sync"

	"github.com/redjolr/goherent/cmd/concurrent_events"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/sequential_events"
	"github.com/redjolr/goherent/internal"
	"github.com/redjolr/goherent/internal/consolesize"
	"github.com/redjolr/goherent/terminal"
)

func Main(extraCmdArgs []string) int {
	if os.Getenv("CI") == "true" {
		args := slices.Concat([]string{"test"}, extraCmdArgs)
		cmd := exec.Command("go", args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stdout pipe: %v\n", err)
			return 1
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stderr pipe: %v\n", err)
			return 1
		}
		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "command failed to start: %v\n", err)
			return 1
		}

		var wg sync.WaitGroup
		wg.Add(2)
		go decodeAndForward(stdout, os.Stdout, &wg)
		go decodeAndForward(stderr, os.Stderr, &wg)
		wg.Wait()

		if err := cmd.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "command failed: %v\n", err)
		}
		return cmd.ProcessState.ExitCode()
	}

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
	isRunnignInCI := os.Getenv("CI") != ""
	var ansiTerminal terminal.AnsiTerminal
	if isRunnignInCI {
		ansiTerminal = terminal.NewUnboundedAnsiTerminal()
	} else if consoleHeight != 0 {
		ansiTerminal = terminal.NewBoundedAnsiTerminal(consoleWidth, consoleHeight)
	} else {
		ansiTerminal = terminal.NewUnboundedAnsiTerminal()
	}
	sequentialEventsRouter := sequential_events.Setup(&ansiTerminal)
	concurrentEventsRouter := concurrent_events.Setup(&ansiTerminal)
	router := NewRouter(sequentialEventsRouter, concurrentEventsRouter)
	return &router
}

func decodeAndForward(src io.Reader, dst io.Writer, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		fmt.Fprintln(dst, internal.DecodeGoherentTestName(scanner.Text()))
	}
}
