package concurrent_events_handler

import (
	"fmt"
	"time"

	"github.com/redjolr/goherent/terminal"
)

const testsListId string = "testsList"
const startingTestsTextblockId string = "startingTestsTextBlock"

type TerminalPresenter struct {
	terminal terminal.Terminal
}

func NewTerminalPresenter(term terminal.Terminal) TerminalPresenter {
	return TerminalPresenter{
		terminal: term,
	}
}

func (tp *TerminalPresenter) TestingStarted(timestamp time.Time) {
	tp.terminal.Print(fmt.Sprintf("\n🚀 Starting... %s", timestamp.Format("2006-01-02 15:04:05.000")))
}
