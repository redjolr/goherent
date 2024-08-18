package testing_started

import (
	"github.com/redjolr/goherent/cmd/events"
)

type Interactor struct {
	output OutputPort
}

func NewEventsHandler(output OutputPort) Interactor {
	return Interactor{
		output: output,
	}
}

func (i Interactor) HandleTestingStarted(evt events.TestingStartedEvent) {
	i.output.TestingStarted()
}
