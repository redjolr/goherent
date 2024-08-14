package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/concurrent_events/bounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/concurrent_events/unbounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/terminal"
)

type Router struct {
	boundedTerminalRouter   *bounded_terminal_handler.Router
	unboundedTerminalRouter *unbounded_terminal_handler.Router
	ansiTerminal            *terminal.AnsiTerminal
}

func NewRouter(
	boundedTerminalRouter *bounded_terminal_handler.Router,
	unboundedTerminalRouter *unbounded_terminal_handler.Router,
	ansiTerminal *terminal.AnsiTerminal,
) Router {
	return Router{
		boundedTerminalRouter:   boundedTerminalRouter,
		unboundedTerminalRouter: unboundedTerminalRouter,
		ansiTerminal:            ansiTerminal,
	}
}

func (r *Router) Route(jsonEvt events.JsonEvent) {
	if r.ansiTerminal.IsBounded() {
		r.boundedTerminalRouter.Route(jsonEvt)
	} else {
		r.unboundedTerminalRouter.Route(jsonEvt)
	}
}
