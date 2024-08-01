package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

func Main(extraCmdArgs []string) int {
	baseCommand := "go test -json"
	commandArgs := append(strings.Split(baseCommand, " "), extraCmdArgs...)

	pArgumentIndex := slices.Index(commandArgs, "-p")
	testsRunConcurrently := true

	if pArgumentIndex != -1 && len(commandArgs) > pArgumentIndex+2 && commandArgs[pArgumentIndex] == "1" {
		testsRunConcurrently = false
	}
	router := NewRouter()
	concurrentEventsRouter := NewConcurrentEventsRouter()

	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Printf("Error opening StdoutPipe: %v\n", err)
		return 1
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return 1
	}
	startTime := time.Now()

	if testsRunConcurrently {
		concurrentEventsRouter.RouteTestingStartedEvent(time.Now())
	} else {
		router.RouteTestingStartedEvent(time.Now())
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		var jsonEvt events.JsonEvent
		err := json.Unmarshal([]byte(m), &jsonEvt)

		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}

		if testsRunConcurrently {
			concurrentEventsRouter.RouteJsonEvent(jsonEvt)
		} else {
			router.RouteJsonEvent(jsonEvt)
		}
	}
	cmd.Wait()
	elapsed := time.Since(startTime)
	if testsRunConcurrently {
		concurrentEventsRouter.RouteTestingFinishedEvent(elapsed)
	} else {
		router.RouteTestingFinishedEvent(elapsed)
	}
	return 0
}
