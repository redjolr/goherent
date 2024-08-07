package cmd

import (
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type EventsMapper struct {
}

func NewEventsMapper() EventsMapper {
	return EventsMapper{}
}

func (evtMapper EventsMapper) JsonTestEvt2CtestPassedEvt(jsonEvt events.JsonEvent) events.CtestPassedEvent {
	return events.NewCtestPassedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestRanEvt(jsonEvt events.JsonEvent) events.CtestRanEvent {
	return events.NewCtestRanEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestOutputEvt(jsonEvt events.JsonEvent) events.CtestOutputEvent {
	return events.NewCtestOutputEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  internal.DecodeGoherentTestName(jsonEvt.Output),
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestFailedEvt(jsonEvt events.JsonEvent) events.CtestFailedEvent {
	return events.NewCtestFailedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
			Output:  jsonEvt.Output,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2CtestSkippedEvt(jsonEvt events.JsonEvent) events.CtestSkippedEvent {
	return events.NewCtestSkippedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackagePassedEvt(jsonEvt events.JsonEvent) events.PackagePassedEvent {
	return events.NewPackagePassedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageFailedEvt(jsonEvt events.JsonEvent) events.PackageFailedEvent {
	return events.NewPackageFailedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageStartedEvt(jsonEvt events.JsonEvent) events.PackageStartedEvent {
	return events.NewPackageStartedEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2NoPackTestsFoundEvent(jsonEvt events.JsonEvent) events.NoPackageTestsFoundEvent {
	return events.NewNoPackageTestsFoundEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}
