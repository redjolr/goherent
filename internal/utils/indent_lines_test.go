package utils_test

import (
	"testing"

	"github.com/redjolr/goherent/internal/utils"
)

func TestIndentLines(t *testing.T) {
	cases := []struct {
		name   string
		text   string
		prefix string
		want   string
	}{
		{"multi-line keeps trailing newline", "a\nb\n", "  ", "  a\n  b\n"},
		{"single line", "solo", ">> ", ">> solo"},
		{"blank lines are not indented", "a\n\nb", "  ", "  a\n\n  b"},
		{"empty input", "", "  ", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := utils.IndentLines(c.text, c.prefix); got != c.want {
				t.Errorf("IndentLines(%q, %q) = %q, want %q", c.text, c.prefix, got, c.want)
			}
		})
	}
}
