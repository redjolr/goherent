package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/concurrent_events/bounded_terminal_handler"
	"github.com/redjolr/goherent/terminal"
)

type Router struct {
	boundedTerminalRouter *bounded_terminal_handler.Router
	ansiTerminal          *terminal.AnsiTerminal
}

func NewRouter(
	boundedTerminalRouter *bounded_terminal_handler.Router,
	ansiTerminal *terminal.AnsiTerminal,
) Router {
	return Router{
		boundedTerminalRouter: boundedTerminalRouter,
		ansiTerminal:          ansiTerminal,
	}
}

func (r *Router) Route(evt any) {
	r.boundedTerminalRouter.Route(evt)
}
