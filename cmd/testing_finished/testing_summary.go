package testing_finished

type TestingSummary struct {
	packagesCount        int
	passedPackagesCount  int
	failedPackagesCount  int
	skippedPackagesCount int

	testsCount        int
	passedTestsCount  int
	failedTestsCount  int
	skippedTestsCount int

	durationS float32
}
