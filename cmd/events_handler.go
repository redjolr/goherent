package cmd

import (
	"fmt"

	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
)

type EventsHandler struct {
}

func NewEventsHandler() EventsHandler {
	return EventsHandler{}
}

func (handler EventsHandler) HandleCtestPassedEvt(evt ctest_passed_event.CtestPassedEvent) {
	fmt.Printf("✅ %s\n\n", evt.CtestName())
}

func (handler EventsHandler) HandleCtestFailedEvt(evt ctest_failed_event.CtestFailedEvent) {
	fmt.Printf("❌ %s\n\n", evt.CtestName())
}
