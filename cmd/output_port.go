package cmd

type OutputPort interface {
	PackageTestsStartedRunning(packageName string)
	CtestPassed(testName string, testDuration float64)
	CtestStartedRunning(testName string)
	CtestFailed(testName string, testDuration float64)
	CtestOutput(testName string, packageName string, output string)
}
