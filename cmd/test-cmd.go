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

	cmd     *exec.Cmd
	scanner *bufio.Scanner

	startTime  time.Time
	endTime    time.Time
	hasStarted bool
	hasEnded   bool
}

func NewTestCmd(args []string) TestCmd {
	return TestCmd{
		staticArgs: []string{"test", "-json"},
		args:       args,
		cmd:        nil,
		scanner:    nil,
		startTime:  time.Time{},
		endTime:    time.Time{},
		hasStarted: false,
		hasEnded:   false,
	}
}

func (t *TestCmd) RunsTestsConcurrently() bool {
	pArgumentIndex := slices.Index(t.args, "-p")
	return !(pArgumentIndex != -1 && len(t.args) >= pArgumentIndex+2 && t.args[pArgumentIndex+1] == "1")
}

func (t *TestCmd) Exec() *TestCmd {
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
	return t
}

func (t *TestCmd) Wait() {
	t.cmd.Wait()
	t.hasEnded = true
	t.endTime = time.Now()
}

func (t *TestCmd) Output() string {
	return t.scanner.Text()
}

func (t *TestCmd) IsRunning() bool {
	return t.scanner.Scan()
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

func (t *TestCmd) NonVerbose() *TestCmd {
	verboseArgInd := slices.Index(t.args, "-v")
	if verboseArgInd != -1 {
		if verboseArgInd == len(t.args)-1 {
			t.args = t.args[0:verboseArgInd]
		} else {
			t.args = slices.Concat(
				t.args[0:verboseArgInd],
				t.args[verboseArgInd+1:],
			)
		}
	}
	return t
}

func (t *TestCmd) ExitCode() int {
	return t.cmd.ProcessState.ExitCode()
}
