package tests_test

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestNotToEqual(t *testing.T) {
	type testCase struct {
		name          string
		expected      any
		actual        any
		shouldPass    bool
		expectedError string
	}

	testCases := []testCase{
		// Integer tests
		{
			name:       "different positive integers",
			expected:   1,
			actual:     2,
			shouldPass: true,
		},
		{
			name:       "different integers with zero",
			expected:   0,
			actual:     5,
			shouldPass: true,
		},
		{
			name:       "different negative integers",
			expected:   -5,
			actual:     -10,
			shouldPass: true,
		},
		{
			name:       "positive and negative integer",
			expected:   5,
			actual:     -5,
			shouldPass: true,
		},
		{
			name:          "same positive integers",
			expected:      3,
			actual:        3,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: 3\nactual  : 3",
		},
		{
			name:          "same negative integers",
			expected:      -7,
			actual:        -7,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: -7\nactual  : -7",
		},
		{
			name:          "both zero",
			expected:      0,
			actual:        0,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: 0\nactual  : 0",
		},
		{
			name:       "large integers",
			expected:   999999999,
			actual:     999999998,
			shouldPass: true,
		},

		// String tests
		{
			name:       "different strings",
			expected:   "hello",
			actual:     "world",
			shouldPass: true,
		},
		{
			name:       "different case strings",
			expected:   "Hello",
			actual:     "hello",
			shouldPass: true,
		},
		{
			name:       "empty string vs non-empty",
			expected:   "",
			actual:     "test",
			shouldPass: true,
		},
		{
			name:       "strings with whitespace difference",
			expected:   "test",
			actual:     "test ",
			shouldPass: true,
		},
		{
			name:          "same strings",
			expected:      "test",
			actual:        "test",
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: \"test\"\nactual  : \"test\"",
		},
		{
			name:          "both empty strings",
			expected:      "",
			actual:        "",
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: \"\"\nactual  : \"\"",
		},
		{
			name:       "multiline strings different",
			expected:   "line1\nline2",
			actual:     "line1\nline3",
			shouldPass: true,
		},
		{
			name:       "strings with special characters",
			expected:   "hello\tworld",
			actual:     "hello world",
			shouldPass: true,
		},

		// Float tests
		{
			name:       "different floats",
			expected:   3.14,
			actual:     2.71,
			shouldPass: true,
		},
		{
			name:          "same floats",
			expected:      3.14,
			actual:        3.14,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: 3.14\nactual  : 3.14",
		},
		{
			name:       "float vs zero",
			expected:   0.0,
			actual:     0.1,
			shouldPass: true,
		},
		{
			name:       "negative floats",
			expected:   -1.5,
			actual:     -2.5,
			shouldPass: true,
		},

		// Boolean tests
		{
			name:       "true vs false",
			expected:   true,
			actual:     false,
			shouldPass: true,
		},
		{
			name:          "both true",
			expected:      true,
			actual:        true,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: true\nactual  : true",
		},
		{
			name:          "both false",
			expected:      false,
			actual:        false,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: false\nactual  : false",
		},

		// Nil tests
		{
			name:          "both nil",
			expected:      nil,
			actual:        nil,
			shouldPass:    false,
			expectedError: "values should not be equal, but both are nil",
		},
		{
			name:       "nil vs integer",
			expected:   nil,
			actual:     5,
			shouldPass: true,
		},
		{
			name:       "integer vs nil",
			expected:   5,
			actual:     nil,
			shouldPass: true,
		},
		{
			name:       "nil vs string",
			expected:   nil,
			actual:     "test",
			shouldPass: true,
		},

		// Mixed type tests
		{
			name:       "int vs string",
			expected:   5,
			actual:     "5",
			shouldPass: true,
		},
		{
			name:       "int vs float",
			expected:   5,
			actual:     5.0,
			shouldPass: true,
		},
		{
			name:       "bool vs int",
			expected:   true,
			actual:     1,
			shouldPass: true,
		},
		{
			name:       "string vs bool",
			expected:   "true",
			actual:     true,
			shouldPass: true,
		},

		// Slice tests
		{
			name:       "different slices",
			expected:   []int{1, 2, 3},
			actual:     []int{4, 5, 6},
			shouldPass: true,
		},
		{
			name:       "slices with different lengths",
			expected:   []int{1, 2},
			actual:     []int{1, 2, 3},
			shouldPass: true,
		},
		{
			name:       "slices with different order",
			expected:   []int{1, 2, 3},
			actual:     []int{3, 2, 1},
			shouldPass: true,
		},
		{
			name:          "same slices",
			expected:      []int{1, 2, 3},
			actual:        []int{1, 2, 3},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: []int{1, 2, 3}\nactual  : []int{1, 2, 3}",
		},
		{
			name:       "empty slice vs non-empty",
			expected:   []int{},
			actual:     []int{1},
			shouldPass: true,
		},
		{
			name:          "both empty slices",
			expected:      []int{},
			actual:        []int{},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: []int{}\nactual  : []int{}",
		},
		{
			name:       "nil slice vs empty slice",
			expected:   []int(nil),
			actual:     []int{},
			shouldPass: true,
		},
		{
			name:       "string slices different",
			expected:   []string{"a", "b"},
			actual:     []string{"c", "d"},
			shouldPass: true,
		},

		// Map tests
		{
			name:       "different maps",
			expected:   map[string]int{"a": 1, "b": 2},
			actual:     map[string]int{"c": 3, "d": 4},
			shouldPass: true,
		},
		{
			name:       "maps with different values",
			expected:   map[string]int{"a": 1},
			actual:     map[string]int{"a": 2},
			shouldPass: true,
		},
		{
			name:          "same maps",
			expected:      map[string]int{"a": 1, "b": 2},
			actual:        map[string]int{"a": 1, "b": 2},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: map[string]int{\"a\":1, \"b\":2}\nactual  : map[string]int{\"a\":1, \"b\":2}",
		},
		{
			name:       "empty map vs non-empty",
			expected:   map[string]int{},
			actual:     map[string]int{"a": 1},
			shouldPass: true,
		},
		{
			name:          "both empty maps",
			expected:      map[string]int{},
			actual:        map[string]int{},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: map[string]int{}\nactual  : map[string]int{}",
		},

		// Struct tests
		{
			name:       "different structs",
			expected:   struct{ x int }{x: 1},
			actual:     struct{ x int }{x: 2},
			shouldPass: true,
		},
		{
			name:          "same structs",
			expected:      struct{ x int }{x: 1},
			actual:        struct{ x int }{x: 1},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: struct { x int }{x:1}\nactual  : struct { x int }{x:1}",
		},
		{
			name: "structs with multiple fields different",
			expected: struct {
				name string
				age  int
			}{name: "Alice", age: 30},
			actual: struct {
				name string
				age  int
			}{name: "Bob", age: 25},
			shouldPass: true,
		},
		{
			name: "structs with one field different",
			expected: struct {
				name string
				age  int
			}{name: "Alice", age: 30},
			actual: struct {
				name string
				age  int
			}{name: "Alice", age: 31},
			shouldPass: true,
		},

		// Pointer tests
		{
			name:       "different int pointers",
			expected:   func() *int { i := 1; return &i }(),
			actual:     func() *int { i := 2; return &i }(),
			shouldPass: true,
		},
		{
			name:       "nil pointer vs non-nil",
			expected:   (*int)(nil),
			actual:     func() *int { i := 5; return &i }(),
			shouldPass: true,
		},
		{
			name:       "different string pointers",
			expected:   func() *string { s := "hello"; return &s }(),
			actual:     func() *string { s := "world"; return &s }(),
			shouldPass: true,
		},

		// Byte slice tests
		{
			name:       "different byte slices",
			expected:   []byte{1, 2, 3},
			actual:     []byte{4, 5, 6},
			shouldPass: true,
		},
		{
			name:          "same byte slices",
			expected:      []byte{1, 2, 3},
			actual:        []byte{1, 2, 3},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: []uint8{0x1, 0x2, 0x3}\nactual  : []uint8{0x1, 0x2, 0x3}",
		},

		// Rune tests
		{
			name:       "different runes",
			expected:   'a',
			actual:     'b',
			shouldPass: true,
		},
		{
			name:          "same runes",
			expected:      'x',
			actual:        'x',
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: 120\nactual  : 120",
		},

		// Complex number tests
		{
			name:       "different complex numbers",
			expected:   complex(1, 2),
			actual:     complex(3, 4),
			shouldPass: true,
		},
		{
			name:          "same complex numbers",
			expected:      complex(1, 2),
			actual:        complex(1, 2),
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: (1+2i)\nactual  : (1+2i)",
		},

		// Array tests
		{
			name:       "different arrays",
			expected:   [3]int{1, 2, 3},
			actual:     [3]int{4, 5, 6},
			shouldPass: true,
		},
		{
			name:          "same arrays",
			expected:      [3]int{1, 2, 3},
			actual:        [3]int{1, 2, 3},
			shouldPass:    false,
			expectedError: "values should not be equal, but both are:\nexpected: [3]int{1, 2, 3}\nactual  : [3]int{1, 2, 3}",
		},

		// Edge cases
		{
			name:       "very long strings",
			expected:   string(make([]byte, 1000)),
			actual:     string(make([]byte, 999)),
			shouldPass: true,
		},
		{
			name:       "unicode strings different",
			expected:   "hello 世界",
			actual:     "hello world",
			shouldPass: true,
		},
		{
			name:       "unicode strings with emojis",
			expected:   "test 😀",
			actual:     "test 😁",
			shouldPass: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := assertions.NotToEqual(tc.expected, tc.actual)

			if tc.shouldPass {
				if err != nil {
					t.Errorf("Expected test to pass, but got error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected test to fail, but it passed")
				}
			}
		})
	}
}
