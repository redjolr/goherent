package events

type EventsMapper struct {
}

func NewEventsMapper() EventsMapper {
	return EventsMapper{}
}

func (evtMapper EventsMapper) JsonTestEvt2CtestPassedEvt(jsonEvt JsonEvent) CtestPassedEvent {
	return NewCtestPassedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestRanEvt(jsonEvt JsonEvent) CtestRanEvent {
	return NewCtestRanEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestOutputEvt(jsonEvt JsonEvent) CtestOutputEvent {
	return NewCtestOutputEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestFailedEvt(jsonEvt JsonEvent) CtestFailedEvent {
	return NewCtestFailedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestSkippedEvt(jsonEvt JsonEvent) CtestSkippedEvent {
	return NewCtestSkippedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackagePassedEvt(jsonEvt JsonEvent) PackagePassedEvent {
	return NewPackagePassedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageFailedEvt(jsonEvt JsonEvent) PackageFailedEvent {
	return NewPackageFailedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageStartedEvt(jsonEvt JsonEvent) PackageStartedEvent {
	return NewPackageStartedEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2NoPackTestsFoundEvent(jsonEvt JsonEvent) NoPackageTestsFoundEvent {
	return NewNoPackageTestsFoundEvent(
		JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}
