package templates

func RunningTestsSummary(packagesSummary, testsSummary, timeSummary string) string {
	return "\n\n" + packagesSummary + "\n" +
		testsSummary + "\n" +
		timeSummary
}
