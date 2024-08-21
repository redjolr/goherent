package ctests_tracker

type TestingSummary struct {
	PackagesCount        int
	PassedPackagesCount  int
	FailedPackagesCount  int
	SkippedPackagesCount int
	RunningPackagesCount int

	TestsCount        int
	PassedTestsCount  int
	FailedTestsCount  int
	SkippedTestsCount int
	RunningTestsCount int

	DurationS float32
}
