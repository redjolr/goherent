package cmd

import "time"

type OutputPort interface {
	TestingStarted(timestamp time.Time)
	PackageTestsStartedRunning(packageName string)
	CtestPassed(testName string, testDuration float64)
	CtestStartedRunning(testName string)
	CtestFailed(testName string, testDuration float64)
	CtestOutput(testName string, packageName string, output string)
}
