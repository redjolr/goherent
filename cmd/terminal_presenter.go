package cmd

import "fmt"

type TerminalPresenter struct {
}

func NewTerminalPresenter() TerminalPresenter {
	return TerminalPresenter{}
}

func (presenter TerminalPresenter) CtestPassed(testName string, timeElapsed float64) {
	fmt.Printf("✅ %s\n\n %f\n", testName, timeElapsed)
}

func (presenter TerminalPresenter) CtestFailed(testName string, timeElapsed float64) {
	fmt.Printf("❌ %s\n\n %f\n", testName, timeElapsed)
}
