package testing_finished_handler

type OutputPort interface {
	TestingFinishedSummary(summary TestingSummary)
}
