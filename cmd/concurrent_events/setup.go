package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/terminal"
)

func Setup(ansiTerminal *terminal.AnsiTerminal) *Router {
	ctestsTracker := ctests_tracker.NewCtestsTracker()
	presenter := NewPresenter(ansiTerminal)
	interactor := NewInteractor(&presenter, &ctestsTracker)
	router := NewRouter(&interactor)
	return &router
}
