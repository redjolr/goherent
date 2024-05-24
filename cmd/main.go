package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/test_failed_event"
	"github.com/redjolr/goherent/cmd/events/test_passed_event"
)

func Main() int {
	baseCommand := "go test -json"
	extraCmdArgs := os.Args[1:]

	commandArgs := append(strings.Split(baseCommand, " "), extraCmdArgs...)
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
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		var jsonEvt events.JsonTestEvent
		err := json.Unmarshal([]byte(m), &jsonEvt)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		var evt events.Event
		if jsonEvt.Action == "pass" {
			evt = test_passed_event.NewFromJsonTestEvent(jsonEvt)
		}
		if jsonEvt.Action == "fail" {
			evt = test_failed_event.NewFromJsonTestEvent(jsonEvt)
		}
		fmt.Printf("%s %s\n%f\n\n", evt.Pictogram(), evt.Message(), evt.Duration())

	}
	cmd.Wait()
	return 0
}
