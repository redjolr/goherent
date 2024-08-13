package testing_started_handler

import (
	"github.com/redjolr/goherent/cmd/events"
)

type EventsHandler struct {
	output OutputPort
}

func NewEventsHandler(output OutputPort) EventsHandler {
	return EventsHandler{
		output: output,
	}
}

func (eh EventsHandler) HandleTestingStarted(evt events.TestingStartedEvent) {
	eh.output.TestingStarted()
}
