package cmd

import "time"

// {"Time":"2024-05-17T20:33:05.7278673+02:00","Action":"run","Package":"github.com/redjolr/go-iam/src/core/tests","Test":"TestUuid"}
type TestRunEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
}

// {"Time":"2024-05-17T20:33:05.7278673+02:00","Action":"output","Package":"github.com/redjolr/go-iam/src/core/tests",
// "Test":"TestUuid","Output":"=== RUN   TestUuid\n"}
type TestOutputEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Output  string    `json:"Output"`
}

// {"Time":"2024-05-17T20:33:05.7278673+02:00",
// "Action":"pass","Package":"github.com/redjolr/go-iam/src/core/tests","Test":"TestUuid","Elapsed":0}
type TestPassEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

// {"Time":"2024-05-23T17:04:33.6991398+02:00","Action":"fail","Package":"github.com/redjolr/go-iam/src/core/tests",
//
//	"Test":"TestUuid/equals","Elapsed":0}
type TestFailEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

// {"Time":"2024-05-17T20:33:05.750866+02:00","Action":"pass","Package":"github.com/redjolr/go-iam/src/core/tests","Elapsed":0.389}
type AllTestsInPackagePassedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

// {"Time":"2024-05-17T20:50:40.2662653+02:00","Action":"skip","Package":"github.com/redjolr/go-iam/src/core/tests",
// "Test":"TestUuid/Doing_something","Elapsed":0}
type TestSkippedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
}

// {"Time":"2024-05-17T21:02:20.7568334+02:00","Action":"pause","Package":"github.com/redjolr/go-iam/src/core/tests", "Test":"TestUuid"}
// EVent expected to be succeeded by "cont"
type TestEventPaused struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
}

// {"Time":"2024-05-17T21:02:20.7568334+02:00","Action":"cont","Package":"github.com/redjolr/go-iam/src/core/tests","Test":"TestUuid"}
// Event preceded by "pause"
type TestContinuedEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
}

type TestEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` //seconds
	Output  string    `json:"Output"`
}
