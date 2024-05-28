package cmd

type OutputPort interface {
	FirstCtestOfPackagePassed(testName string, packageName string, testDuration float64)
	CtestPassed(testName string, testDuration float64)
	CtestFailed(testName string, testDuration float64)
}
