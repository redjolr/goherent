package cmd

type OutputPort interface {
	FirstCtestOfPackageStartedRunning(testName string, packageName string)
	FirstCtestOfPackageFailed(testName string, packageName string, testDuration float64)
	PackageTestsStartedRunning(packageName string)
	CtestPassed(testName string, testDuration float64)
	CtestStartedRunning(testName string)
	CtestFailed(testName string, testDuration float64)
	CtestOutput(testName string, packageName string, output string)
}
