package concurrent_events_handler

import (
	"time"
)

type OutputPort interface {
	TestingStarted(timestamp time.Time)
}
