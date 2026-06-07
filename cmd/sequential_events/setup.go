package sequential_events

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
)

func Setup(ansiTerminal *terminal.AnsiTerminal) *Router {
	var sequentialEventsOutputPort OutputPort
	if ansiTerminal.IsBounded() {
		// Interactive terminal: live footer that updates in place.
		sequentialEventsOutputPort = NewLiveTerminalPresenter(ansiTerminal)
	} else {
		// Piped / non-TTY output: plain sequential printing, no cursor control.
		sequentialEventsOutputPort = NewUnboundedTerminalPresenter(ansiTerminal)
	}
	ctestsTracker := ctests_tracker.NewCtestsTracker()

	sequentialEventsInteractor := NewInteractor(sequentialEventsOutputPort, &ctestsTracker)
	router := NewRouter(&sequentialEventsInteractor)
	return &router
}
