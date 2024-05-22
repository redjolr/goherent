package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/redjolr/goherent/internal"
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
		var event TestEvent
		err := json.Unmarshal([]byte(m), &event)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		if event.Action == "pass" {
			var passedTestEvent TestPassEvent
			err := json.Unmarshal([]byte(m), &passedTestEvent)
			if err != nil {
				log.Fatalf("Unable to marshal JSON pass event due to %s", err)
			}
			testName := internal.DecodeGoherentTestName(passedTestEvent.Test)
			fmt.Printf("✅ %s\n%f\n\n", testName, event.Elapsed)
		}
		if event.Action == "fail" {
			var failedTestEvent TestPassEvent
			err := json.Unmarshal([]byte(m), &failedTestEvent)
			if err != nil {
				log.Fatalf("Unable to marshal JSON failed event due to %s", err)
			}
			testName := internal.DecodeGoherentTestName(failedTestEvent.Test)
			fmt.Printf("❌ %s\n%f\n\n", testName, event.Elapsed)
		}
	}
	cmd.Wait()
	return 0
}
