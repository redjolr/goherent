package events

import (
	"strings"
	"time"

	"github.com/redjolr/goherent/internal"
)

type CtestOutputEvent struct {
	Time        time.Time
	PackageName string
	TestName    string
	Output      string
}

func NewCtestOutputEvent(jsonEvt JsonTestEvent) CtestOutputEvent {
	return CtestOutputEvent{
		Time:        jsonEvt.Time,
		PackageName: jsonEvt.Package,
		TestName:    internal.DecodeGoherentTestName(jsonEvt.Test),
		Output:      jsonEvt.Output,
	}
}

func (e CtestOutputEvent) IsEventOfAParentTest() bool {
	return !strings.Contains(e.TestName, "/")
}

func (e CtestOutputEvent) IsAGenericRunPassFailOutput() bool {
	return strings.Contains(e.Output, "--- PASS") ||
		strings.Contains(e.Output, "=== RUN") ||
		strings.Contains(e.Output, "--- FAIL")
}
