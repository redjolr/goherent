package events

import "time"

type JsonEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    *string   `json:"Test"`
	Elapsed *float64  `json:"Elapsed"` //seconds
	Output  string    `json:"Output"`
}
