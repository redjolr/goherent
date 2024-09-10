package templates

func TestingFinishedSummary(packagesSummary, testsSummary, timeSummary string) string {
	return "\n\n" + packagesSummary + "\n" +
		testsSummary + "\n" +
		timeSummary + "\n" +
		"Ran all tests."
}
