package events

import "time"

type JsonPackageEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Elapsed float64   `json:"Elapsed"` //seconds
	Output  string    `json:"Output"`
}
