package cmd

type OutputPort interface {
	FirstCtestOfPackageStartedRunning(testName string, packageName string)
	FirstCtestOfPackagePassed(testName string, packageName string, testDuration float64)
	FirstCtestOfPackageFailed(testName string, packageName string, testDuration float64)
	CtestPassed(testName string, testDuration float64)
	CtestStartedRunning(testName string)
	CtestFailed(testName string, testDuration float64)
}
