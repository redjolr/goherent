package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/redjolr/goherent/cmd/internal"
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
		var event internal.TestEvent
		err := json.Unmarshal([]byte(m), &event)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		fmt.Println(m)
	}
	cmd.Wait()
	return 0
}
