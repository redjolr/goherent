package tests

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToHaveLength(t *testing.T) {
	tests := []struct {
		name      string
		value     any
		length    int
		expectErr bool
	}{
		// Valid cases
		{"Array with correct length", [3]int{1, 2, 3}, 3, false},
		{"Slice with correct length", []int{1, 2, 3}, 3, false},
		{"String with correct length", "abc", 3, false},
		{"Map with correct length", map[string]int{"a": 1, "b": 2}, 2, false},
		{"Channel with correct length", func() chan int { ch := make(chan int, 2); ch <- 1; ch <- 2; return ch }(), 2, false},

		// Invalid cases
		{"Array with incorrect length", [3]int{1, 2, 3}, 2, true},
		{"Slice with incorrect length", []int{1, 2, 3}, 2, true},
		{"String with incorrect length", "abc", 2, true},
		{"Map with incorrect length", map[string]int{"a": 1, "b": 2}, 1, true},
		{"Channel with incorrect length", func() chan int { ch := make(chan int, 2); ch <- 1; return ch }(), 2, true},

		// Edge cases
		{"Nil value", nil, 0, true},
		{"Non-length type (int)", 123, 0, true},
		{"Non-length type (struct)", struct{}{}, 0, true},
		{"Empty slice", []int{}, 0, false},
		{"Empty string", "", 0, false},
		{"Empty map", map[string]int{}, 0, false},
		{"Empty channel", make(chan int), 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := assertions.ToHaveLength(tt.value, tt.length)
			if (err != nil) != tt.expectErr {
				t.Errorf("ToHaveLength(%v, %d) error = %v, expectErr %v", tt.value, tt.length, err, tt.expectErr)
			}
		})
	}
}
