package console

import "github.com/redjolr/goherent/console/internal"

type Console struct {
	terminal internal.Terminal
}

func NewConsole(terminal internal.Terminal) Console {
	return Console{
		terminal: terminal,
	}
}
