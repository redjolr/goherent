package tests

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToHaveLengthGreaterThan(t *testing.T) {
	tests := []struct {
		name              string
		value             any
		greaterThanLength int
		expectErr         bool
	}{
		// Exact length
		{"Array with exact length", [3]int{1, 2, 3}, 3, true},
		{"Slice with exact length", []int{1, 2, 3}, 3, true},
		{"String with exact length", "abc", 3, true},
		{"Map with exact length", map[string]int{"a": 1, "b": 2}, 2, true},
		{"Channel with exact length", func() chan int { ch := make(chan int, 2); ch <- 1; ch <- 2; return ch }(), 2, true},

		// Length less
		{"Array with length less than expected", [3]int{1, 2, 3}, 4, true},
		{"Slice with length less than expected", []int{1, 2, 3}, 4, true},
		{"String with length less than expected", "abc", 4, true},
		{"Map with length less than expected", map[string]int{"a": 1, "b": 2}, 4, true},
		{"Channel with length less than expected", func() chan int { ch := make(chan int, 2); ch <- 1; ch <- 2; return ch }(), 3, true},

		// Length Greater
		{"Array with length greater than expected", [3]int{1, 2, 3}, 2, false},
		{"Slice with length greater than expected", []int{1, 2, 3}, 2, false},
		{"String with length greater than expected", "abc", 2, false},
		{"Map with length greater than expected", map[string]int{"a": 1, "b": 2}, 1, false},
		{"Channel with length greater than expected", func() chan int { ch := make(chan int, 2); ch <- 1; ch <- 2; return ch }(), 1, false},

		// Edge cases
		{"Nil value", nil, 0, true},
		{"Non-length type (int)", 123, 0, true},
		{"Non-length type (struct)", struct{}{}, 0, true},
		{"Empty slice 1", []int{}, 0, true},
		{"Empty slice 2", []int{}, 1, true},
		{"Empty string", "", 0, true},
		{"Empty map", map[string]int{}, 0, true},
		{"Empty channel", make(chan int), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := assertions.ToHaveLengthGreaterThan(tt.value, tt.greaterThanLength)
			if (err != nil) != tt.expectErr {
				t.Errorf("ToHaveLengthGreaterThan(%v, %d) error = %v, expectErr %v", tt.value, tt.greaterThanLength, err, tt.expectErr)
			}
		})
	}
}
