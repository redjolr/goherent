package testing_started

import (
	"github.com/redjolr/goherent/cmd/events"
)

type Handler struct {
	output OutputPort
}

func NewEventsHandler(output OutputPort) Handler {
	return Handler{
		output: output,
	}
}

func (h Handler) HandleTestingStarted(evt events.TestingStartedEvent) {
	h.output.TestingStarted()
}
