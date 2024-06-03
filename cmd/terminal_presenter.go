package cmd

import (
	"fmt"
)

type TerminalPresenter struct {
}

func NewTerminalPresenter() TerminalPresenter {
	return TerminalPresenter{}
}
func (presenter TerminalPresenter) FirstCtestOfPackageStartedRunning(testName string, packageName string) {
	fmt.Printf("📦⏳ %s\n", packageName)
	fmt.Printf("\t⏳ %s\n\n", testName)
}

func (presenter TerminalPresenter) FirstCtestOfPackagePassed(testName string, packageName string, testDuration float64) {
	fmt.Printf("📦⏳ %s\n", packageName)
	fmt.Printf("\t✅ %s\n\n %f\n", testName, testDuration)
}

func (presenter TerminalPresenter) FirstCtestOfPackageFailed(testName string, packageName string, testDuration float64) {
	fmt.Printf("📦⏳ %s\n", packageName)
	fmt.Printf("\t❌ %s\n\n %f\n", testName, testDuration)
}

func (pressenter TerminalPresenter) CtestStartedRunning(testName string) {
	fmt.Printf("\t⏳ %s\n\n", testName)
}

func (presenter TerminalPresenter) CtestPassed(testName string, testDuration float64) {
	fmt.Printf("\t✅ %s\n\n %f\n", testName, testDuration)
}

func (presenter TerminalPresenter) CtestFailed(testName string, testDuration float64) {
	fmt.Printf("\t❌ %s\n\n %f\n", testName, testDuration)
}

func (presenter TerminalPresenter) CtestOutput(testName string, packageName string, output string) {
	fmt.Printf("\t %s\n\n %s\n", testName, output)
}
