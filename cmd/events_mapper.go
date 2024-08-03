package cmd

import (
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	"github.com/redjolr/goherent/cmd/events/ctest_skipped_event"
	"github.com/redjolr/goherent/cmd/events/no_package_tests_found_event"
	"github.com/redjolr/goherent/cmd/events/package_failed_event"
	"github.com/redjolr/goherent/cmd/events/package_passed_event"
	"github.com/redjolr/goherent/cmd/events/package_started_event"
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

func (evtMapper EventsMapper) JsonTestEvt2CtestSkippedEvt(jsonEvt events.JsonEvent) ctest_skipped_event.CtestSkippedEvent {
	return ctest_skipped_event.NewFromJsonTestEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
			Test:    *jsonEvt.Test,
			Elapsed: jsonEvt.Elapsed,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackagePassedEvt(jsonEvt events.JsonEvent) package_passed_event.PackagePassedEvent {
	return package_passed_event.NewFromJsonTestEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageFailedEvt(jsonEvt events.JsonEvent) package_failed_event.PackageFailedEvent {
	return package_failed_event.NewFromJsonTestEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Action:  jsonEvt.Action,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2PackageStartedEvt(jsonEvt events.JsonEvent) package_started_event.PackageStartedEvent {
	return package_started_event.NewFromJsonTestEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}

func (evtMapper EventsMapper) JsonTestEvt2NoPackTestsFoundEvent(
	jsonEvt events.JsonEvent,
) no_package_tests_found_event.NoPackageTestsFoundEvent {
	return no_package_tests_found_event.NewFromJsonTestEvent(
		events.JsonTestEvent{
			Time:    jsonEvt.Time,
			Package: jsonEvt.Package,
		},
	)
}
