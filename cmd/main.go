package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

func Main(extraCmdArgs []string) int {
	baseCommand := "go test -json"
	commandArgs := append(strings.Split(baseCommand, " "), extraCmdArgs...)

	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)
	stdout, err := cmd.StdoutPipe()
	router := NewRouter()
	if err != nil {
		fmt.Printf("Error opening StdoutPipe: %v\n", err)
		return 1
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return 1
	}
	router.RouteTestingStartedEvent(time.Now())

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		var jsonEvt events.JsonEvent
		err := json.Unmarshal([]byte(m), &jsonEvt)

		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}

		router.RouteJsonEvent(jsonEvt)
	}
	cmd.Wait()
	return 0
}
