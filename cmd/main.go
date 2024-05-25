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
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
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
		var jsonEvt events.JsonEvent
		err := json.Unmarshal([]byte(m), &jsonEvt)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		var ctestEvt events.CtestEvent
		// var packEvt events.PackageEvent
		if jsonEvt.Test != nil {
			if jsonEvt.Action == "pass" {
				ctestEvt = ctest_passed_event.NewFromJsonTestEvent(
					events.JsonTestEvent{
						Time:    jsonEvt.Time,
						Action:  jsonEvt.Action,
						Package: jsonEvt.Package,
						Test:    *jsonEvt.Test,
						Elapsed: jsonEvt.Elapsed,
						Output:  jsonEvt.Output,
					},
				)
			}
			if jsonEvt.Action == "fail" {
				ctestEvt = ctest_failed_event.NewFromJsonTestEvent(
					events.JsonTestEvent{
						Time:    jsonEvt.Time,
						Action:  jsonEvt.Action,
						Package: jsonEvt.Package,
						Test:    *jsonEvt.Test,
						Elapsed: jsonEvt.Elapsed,
						Output:  jsonEvt.Output,
					},
				)
			}
			fmt.Printf("%s %s\n\n", ctestEvt.Pictogram(), ctestEvt.CtestName())

		}

	}
	cmd.Wait()
	return 0
}
