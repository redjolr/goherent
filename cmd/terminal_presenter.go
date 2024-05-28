package cmd

import "fmt"

type TerminalPresenter struct {
}

func NewTerminalPresenter() TerminalPresenter {
	return TerminalPresenter{}
}

func (presenter TerminalPresenter) FirstCtestOfPackagePassed(testName string, packageName string, testDuration float64) {
	fmt.Printf("📦 %s\n", packageName)
	fmt.Printf("\t✅ %s\n\n %f\n", testName, testDuration)

}

func (presenter TerminalPresenter) CtestPassed(testName string, testDuration float64) {
	fmt.Printf("\t✅ %s\n\n %f\n", testName, testDuration)
}

func (presenter TerminalPresenter) CtestFailed(testName string, testDuration float64) {
	fmt.Printf("\t❌ %s\n\n %f\n", testName, testDuration)
}
