package concurrent_events

import (
	"github.com/redjolr/goherent/cmd/concurrent_events/bounded_terminal_handler"
	"github.com/redjolr/goherent/cmd/concurrent_events/unbounded_terminal_handler"
	"github.com/redjolr/goherent/terminal"
)

func Setup(ansiTerminal *terminal.AnsiTerminal) *Router {
	boundedTerminalRouter := bounded_terminal_handler.Setup(ansiTerminal)
	unboundedTerminalRouter := unbounded_terminal_handler.Setup(ansiTerminal)
	router := NewRouter(boundedTerminalRouter, unboundedTerminalRouter, ansiTerminal)
	return &router
}