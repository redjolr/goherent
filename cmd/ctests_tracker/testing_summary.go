package ctests_tracker

type TestingSummary struct {
	PackagesCount        int
	PassedPackagesCount  int
	FailedPackagesCount  int
	SkippedPackagesCount int

	TestsCount        int
	PassedTestsCount  int
	FailedTestsCount  int
	SkippedTestsCount int

	DurationS float32
}
