package cmd

type OutputPort interface {
	CtestPassed(testName string, testDuration float64)
	CtestFailed(testName string, testDuration float64)
}
