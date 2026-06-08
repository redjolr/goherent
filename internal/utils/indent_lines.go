package utils

import "strings"

// IndentLines prefixes every non-empty line of text with prefix. Empty lines
// (including the empty segment after a trailing newline) are left untouched, so
// IndentLines("a\nb\n", "  ") yields "  a\n  b\n" rather than indenting the blank
// tail. It is used to keep multi-line blocks (e.g. compiler output under a failed
// package) aligned as a single indented unit.
func IndentLines(text, prefix string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}
