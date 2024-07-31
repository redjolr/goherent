package sequential_events_handler

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
