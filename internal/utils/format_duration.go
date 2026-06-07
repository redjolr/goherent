package utils

import "fmt"

// FormatDuration renders a test duration (given in seconds) compactly for
// display next to a test result:
//
//	>= 1s     -> "1.23s"
//	>= 1ms    -> "12ms"
//	otherwise -> "<1ms"
//
// Go's test2json reports elapsed time in seconds, which is the unit this expects.
func FormatDuration(seconds float64) string {
	if seconds >= 1 {
		return fmt.Sprintf("%.2fs", seconds)
	}
	ms := seconds * 1000
	if ms >= 1 {
		return fmt.Sprintf("%.0fms", ms)
	}
	return "<1ms"
}
