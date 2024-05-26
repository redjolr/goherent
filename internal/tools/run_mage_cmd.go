package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunMageCmd(command string) *exec.Cmd {
	commandArgs := strings.Split(command, " ")

	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error opening StdoutPipe: %v\n", err)
		os.Exit(1)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	return cmd
}
