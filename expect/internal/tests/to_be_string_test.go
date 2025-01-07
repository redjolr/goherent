package tests_test

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestToBeString(t *testing.T) {
	tests := []struct {
		name           string      // Description of the test case
		input          interface{} // Input value to test
		assertionFails bool        // Whether an error is expected
	}{
		{"Empty string input", "", false},
		{"String input", "test string", false},
		{"Integer input", 42, true},
		{"Float input", 3.14, true},
		{"Boolean true", true, true},
		{"Boolean false", false, true},
		{"Nil input", nil, true},
		{"Slice of strings", []string{"hello", "world"}, true},
		{"Slice of integers", []int{1, 2, 3}, true},
		{"Map with string keys", map[string]int{"key": 1}, true},
		{"Struct input", struct{}{}, true},
		{"Pointer to string", new(string), true},
		{"Pointer to int", new(int), true},
		{"Array of integers", [3]int{1, 2, 3}, true},
		{"Array of strings", [2]string{"hello", "world"}, true},
		{"Interface type holding a string", interface{}("interface string"), false},
		{"Interface type holding an int", interface{}(123), true},
		{"Channel input", make(chan int), true},
		{"Function input", func() {}, true},
		{"Empty string", "", false},
		{"Byte slice", []byte("byte slice"), true},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := assertions.ToBeString(test.input)
			if (err == nil) == test.assertionFails {
				t.Errorf("CheckIfString(%v) = %v, expected error: %v", test.input, err, test.assertionFails)
			}
		})
	}
}
