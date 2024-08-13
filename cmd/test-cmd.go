package cmd

import (
	"bufio"
	"os/exec"
	"slices"
	"time"
)

type TestCmd struct {
	staticArgs []string
	args       []string

	cmd                  *exec.Cmd
	scanner              *bufio.Scanner
	scannerHasMoreToScan bool

	startTime  time.Time
	endTime    time.Time
	hasStarted bool
	hasEnded   bool
}

func NewTestCmd(args []string) TestCmd {
	return TestCmd{
		staticArgs:           []string{"test", "-json"},
		args:                 args,
		cmd:                  nil,
		scanner:              nil,
		scannerHasMoreToScan: true,
		startTime:            time.Time{},
		endTime:              time.Time{},
		hasStarted:           false,
		hasEnded:             false,
	}
}

func (t *TestCmd) RunsTestsConcurrently() bool {
	pArgumentIndex := slices.Index(t.args, "-p")
	return !(pArgumentIndex != -1 && len(t.args) >= pArgumentIndex+2 && t.args[pArgumentIndex+1] == "1")
}

func (t *TestCmd) Exec() {
	t.cmd = exec.Command("go", slices.Concat(t.staticArgs, t.args)...)
	t.hasStarted = true
	t.startTime = time.Now()

	stdout, err := t.cmd.StdoutPipe()

	if err != nil {
		panic("Error opening stdout pipe.")
	}
	err = t.cmd.Start()
	if err != nil {
		panic("Could not start command.")
	}
	t.scanner = bufio.NewScanner(stdout)
	t.scanner.Split(bufio.ScanLines)
}

func (t *TestCmd) Wait() {
	t.cmd.Wait()
	t.hasEnded = true
	t.endTime = time.Now()
}

func (t *TestCmd) NextOutput() string {
	t.scannerHasMoreToScan = t.scanner.Scan()
	if t.scannerHasMoreToScan {
		return t.scanner.Text()
	}
	return ""
}

func (t *TestCmd) IsRunning() bool {
	return t.scannerHasMoreToScan
}

func (t *TestCmd) ExecutionTime() time.Duration {
	if !t.hasStarted {
		return time.Duration(0)
	}
	if !t.hasEnded {
		return time.Since(t.startTime)
	}
	return t.endTime.Sub(t.startTime)
}
