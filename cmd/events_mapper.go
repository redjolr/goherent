package cmd

import (
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
)

type EventsMapper struct {
}

func NewEventsMapper() EventsMapper {
	return EventsMapper{}
}

func (evtMapper EventsMapper) JsonTestEvt2CtestPassedEvt(jsonEvt events.JsonEvent) ctest_passed_event.CtestPassedEvent {
	return ctest_passed_event.NewFromJsonTestEvent(
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

func (evtMapper EventsMapper) JsonTestEvt2CtestRanEvt(jsonEvt events.JsonEvent) ctest_ran_event.CtestRanEvent {
	return ctest_ran_event.NewFromJsonTestEvent(
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

func (evtMapper EventsMapper) JsonTestEvt2CtestFailedEvt(jsonEvt events.JsonEvent) ctest_failed_event.CtestFailedEvent {
	return ctest_failed_event.NewFromJsonTestEvent(
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
