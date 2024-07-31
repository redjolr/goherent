package cmd

type TestingSummary struct {
	packagesCount       int
	passedPackagesCount int
	failedPackagesCount int

	testsCount       int
	passedTestsCount int
	failedTestsCount int

	durationS float32
}
