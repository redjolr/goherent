package tests_test

import (
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

func TestStrings(t *testing.T) {
	t.Run("it should return nil, if we check if \"\" contains \"\".", func(t *testing.T) {
		err := assertions.ToContain("", "")
		if err != nil {
			t.Errorf("An empty string contains an empty string, but ToContain assertions says no. Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if \"A\" contains \"\".", func(t *testing.T) {
		err := assertions.ToContain("A", "")
		if err != nil {
			t.Errorf(`\"A\" contains \"\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("", "A")
		if err == nil {
			t.Errorf(`\"\" does not contain \"\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"A\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("A", "A")
		if err != nil {
			t.Errorf(`\"A\" contains \"A\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"a\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("a", "A")
		if err == nil {
			t.Errorf(`\"a\" does not contain \"A\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"A\" contains \"a\".", func(t *testing.T) {
		err := assertions.ToContain("A", "a")
		if err == nil {
			t.Errorf(`\"a\" does not contain \"A\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"AB\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("AB", "A")
		if err != nil {
			t.Errorf(`\"AB\" contains \"A\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"BA\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("BA", "A")
		if err != nil {
			t.Errorf(`\"BA\" contains \"A\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"BA\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("BAB", "A")
		if err != nil {
			t.Errorf(`\"BAB\" contains \"A\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"BA\" contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain("BAB", "A")
		if err != nil {
			t.Errorf(`\"BAB\" contains \"A\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"AB\" contains \"AB\".", func(t *testing.T) {
		err := assertions.ToContain("AB", "AB")
		if err != nil {
			t.Errorf(`\"AB\" contains \"AB\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"ABCDEFG\" contains \"ABCDEFG\".", func(t *testing.T) {
		err := assertions.ToContain("ABCDEFG", "ABCDEFG")
		if err != nil {
			t.Errorf(`\"ABCDEFG\" contains \"ABCDEFG\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"ABCDEFG\" contains \"BCDEFG\".", func(t *testing.T) {
		err := assertions.ToContain("ABCDEFG", "BCDEFG")
		if err != nil {
			t.Errorf(`\"ABCDEFG\" contains \"BCDEFG\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"ABCDEFG\" contains \"ABCDEF\".", func(t *testing.T) {
		err := assertions.ToContain("ABCDEFG", "ABCDEF")
		if err != nil {
			t.Errorf(`\"ABCDEFG\" contains \"ABCDEF\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \" \" contains \" \".", func(t *testing.T) {
		err := assertions.ToContain(" ", " ")
		if err != nil {
			t.Errorf(`\" \" contains \" \", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"  \" contains \" \".", func(t *testing.T) {
		err := assertions.ToContain("  ", " ")
		if err != nil {
			t.Errorf(`\"  \" contains \" \", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"AB\" contains \"BA\".", func(t *testing.T) {
		err := assertions.ToContain("AB", "BA")
		if err == nil {
			t.Errorf(`\"AB\" does contains \"BA\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"BAB\" contains \"AB\".", func(t *testing.T) {
		err := assertions.ToContain("BAB", "AB")
		if err != nil {
			t.Errorf(`\"BAB\" contains \"AB\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return nil, if we check if \"ABC\" contains \"AB\".", func(t *testing.T) {
		err := assertions.ToContain("ABC", "AB")
		if err != nil {
			t.Errorf(`\"ABC\" contains \"AB\", but ToContain assertions says no. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"ABC\" contains \"AC\".", func(t *testing.T) {
		err := assertions.ToContain("ABC", "AC")
		if err == nil {
			t.Errorf(`\"ABC\" does not contain \"AC\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"ABC\" contains \"CB\".", func(t *testing.T) {
		err := assertions.ToContain("ABC", "CB")
		if err == nil {
			t.Errorf(`\"ABC\" does not contain \"CB\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"ABCDE\" contains \"BCZ\".", func(t *testing.T) {
		err := assertions.ToContain("ABCDE", "BCZ")
		if err == nil {
			t.Errorf(`\"ABCDE\" does not contain \"BCZ\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"ABCDE\" contains \"EDCBA\".", func(t *testing.T) {
		err := assertions.ToContain("ABCDE", "EDCBA")
		if err == nil {
			t.Errorf(`\"ABCDE\" does not contain \"EDCBA\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"ABC DE\" contains \"BCD\".", func(t *testing.T) {
		err := assertions.ToContain("ABC DE", "BCD")
		if err == nil {
			t.Errorf(`\"ABC DE\" does not contain \"BCD\", but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"123\" contains 123.", func(t *testing.T) {
		err := assertions.ToContain("123", 123)
		if err == nil {
			t.Errorf(`\"123\" does not contain 123, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"3.14\" contains 3.14.", func(t *testing.T) {
		err := assertions.ToContain("3.14", 3.14)
		if err == nil {
			t.Errorf(`\"3.14\" does not contain 3.14, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"true\" contains true.", func(t *testing.T) {
		err := assertions.ToContain("true", true)
		if err == nil {
			t.Errorf(`\"true\" does not contain true, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"\" contains an empty strings array.", func(t *testing.T) {
		err := assertions.ToContain("", [0]string{})
		if err == nil {
			t.Errorf(`\"\" does not contain [0]string{}, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run(`
	it should return an error, 
	if we check if \"\" contains a strings array with only \"\" as an element`, func(t *testing.T) {
		err := assertions.ToContain("", [1]string{""})
		if err == nil {
			t.Errorf(`\"\" does not contain [1]string{""}, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run(`
	it should return an error, 
	if we check if \"A\" contains the array [1]string{"A"}.`, func(t *testing.T) {
		err := assertions.ToContain("A", [1]string{"A"})
		if err == nil {
			t.Errorf(`\"A\" does not contain [1]string{"A"}, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run(`
	it should return an error, 
	if we check if \"A\" contains the map map[string]string{"A":"A"}.`, func(t *testing.T) {
		err := assertions.ToContain("A", map[string]string{"A": "A"})
		if err == nil {
			t.Errorf(`\"A\" does not contain map[string]string{"A": "A"}, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run(`
	it should return an error, 
	if we check if \"A\" contains the struct A{A:"A"}.`, func(t *testing.T) {
		type A struct{ A string }

		err := assertions.ToContain("A", A{A: "A"})
		if err == nil {
			t.Errorf(`\"A\" does not contain the struct A{A:"A"}, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"\" contains nil.", func(t *testing.T) {
		err := assertions.ToContain("", nil)
		if err == nil {
			t.Errorf(`\"\" does not contain nil, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"nil\" contains nil.", func(t *testing.T) {
		err := assertions.ToContain("nil", nil)
		if err == nil {
			t.Errorf(`\"nil\" does not contain nil, but ToContain assertions says yes. Error: %v`, err)
		}
	})

	t.Run("it should return an error, if we check if \"\" contains a function.", func(t *testing.T) {
		err := assertions.ToContain("", func() {})
		if err == nil {
			t.Errorf(`\"\" does not contain a function, but ToContain assertions says yes. Error: %v`, err)
		}
	})
}

func TestStringArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]string{\"\"} contains \"\".", func(t *testing.T) {
		err := assertions.ToContain([1]string{""}, "")
		if err != nil {
			t.Errorf("[1]string{\"\"} contains \"\", but ToContain assertions says no. Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if []string{} contains \"\".", func(t *testing.T) {
		err := assertions.ToContain([0]string{}, "")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [1]string{\"A\"} contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain([1]string{"A"}, "A")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]string{\"A\", \"B\"} contains \"A\".", func(t *testing.T) {
		err := assertions.ToContain([2]string{"A", "B"}, "A")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]string{\"A\", \"B\"} contains \"B\".", func(t *testing.T) {
		err := assertions.ToContain([2]string{"A", "B"}, "B")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [2]string{\"A\", \"B\"} contains \"AB\".", func(t *testing.T) {
		err := assertions.ToContain([2]string{"A", "B"}, "AB")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]string{\"A\", \"B\", \"C\"} contains \"B\".", func(t *testing.T) {
		err := assertions.ToContain([3]string{"A", "B", "C"}, "B")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]string{\"A\", \"Beeee\", \"C\"} contains \"Beeee\".", func(t *testing.T) {
		err := assertions.ToContain([3]string{"A", "Beeee", "C"}, "Beeee")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]string{\"A\", \"Beeee\", \"C\"} contains \"B\".", func(t *testing.T) {
		err := assertions.ToContain([3]string{"A", "Beeee", "C"}, "B")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]string{\"13\"} contains 13.", func(t *testing.T) {
		err := assertions.ToContain([1]string{"13"}, 13)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]string{\"3.14\"} contains 3.14.", func(t *testing.T) {
		err := assertions.ToContain([1]string{"3.14"}, 3.14)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]string{\"true\"} contains true.", func(t *testing.T) {
		err := assertions.ToContain([1]string{"true"}, true)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]string{\"nil\"} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([1]string{"nil"}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]string{\"\"} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([1]string{""}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]string{} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([0]string{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestIntegerArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]int{0} contains 0.", func(t *testing.T) {
		err := assertions.ToContain([1]int{0}, 0)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if []int{} contains 0.", func(t *testing.T) {
		err := assertions.ToContain([0]int{}, 0)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [1]int{7} contains 7.", func(t *testing.T) {
		err := assertions.ToContain([1]int{7}, 7)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]int{7, 12} contains 7.", func(t *testing.T) {
		err := assertions.ToContain([2]int{7, 12}, 7)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]int{7, 12} contains 12.", func(t *testing.T) {
		err := assertions.ToContain([2]int{7, 12}, 12)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [2]int{7, 12} contains 712.", func(t *testing.T) {
		err := assertions.ToContain([2]int{7, 12}, 712)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]int{4, 9, 10} contains 9.", func(t *testing.T) {
		err := assertions.ToContain([3]int{4, 9, 10}, 9)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]int{4, 9, 10} contains 129.", func(t *testing.T) {
		err := assertions.ToContain([3]int{4, 9, 10}, 129)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]int{77} contains \"77\".", func(t *testing.T) {
		err := assertions.ToContain([1]int{77}, "77")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]int{} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([0]int{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestFloatArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]float64{0} contains 0.0.", func(t *testing.T) {
		err := assertions.ToContain([1]float64{0.0}, 0.0)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if []float64{} contains 0.0.", func(t *testing.T) {
		err := assertions.ToContain([0]float64{}, 0.0)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [1]float64{7.2} contains 7.2.", func(t *testing.T) {
		err := assertions.ToContain([1]float64{7.2}, 7.2)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]float64{7.1, 12.3} contains 7.1.", func(t *testing.T) {
		err := assertions.ToContain([2]float64{7.1, 12.3}, 7.1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]float64{7.1, 12.3} contains 12.3.", func(t *testing.T) {
		err := assertions.ToContain([2]float64{7.1, 12.3}, 12.3)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]float64{4.1, 9.2, 10.3} contains 9.2.", func(t *testing.T) {
		err := assertions.ToContain([3]float64{4.1, 9.2, 10.3}, 9.2)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]float64{4.1, 9.2, 10.3} contains 3.14.", func(t *testing.T) {
		err := assertions.ToContain([3]float64{4.1, 9.2, 10.3}, 3.14)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]float64{7.4} contains \"7.4\".", func(t *testing.T) {
		err := assertions.ToContain([1]float64{7.4}, "7.4")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]float64{} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([0]float64{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestBooleanArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]bool{true} contains true.", func(t *testing.T) {
		err := assertions.ToContain([1]bool{true}, true)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if []bool{} contains false.", func(t *testing.T) {
		err := assertions.ToContain([0]bool{}, false)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [1]bool{false} contains false.", func(t *testing.T) {
		err := assertions.ToContain([1]bool{false}, false)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]bool{true, false} contains true.", func(t *testing.T) {
		err := assertions.ToContain([2]bool{true, false}, true)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]bool{true, false} contains false.", func(t *testing.T) {
		err := assertions.ToContain([2]bool{true, false}, false)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]bool{false, true, false} contains true.", func(t *testing.T) {
		err := assertions.ToContain([3]bool{false, true, false}, true)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]bool{false, false, false} contains true.", func(t *testing.T) {
		err := assertions.ToContain([3]bool{false, false, false}, true)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]bool{true} contains \"true\".", func(t *testing.T) {
		err := assertions.ToContain([1]bool{true}, false)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]bool{} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([0]bool{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestNilArrays(t *testing.T) {
	t.Run("it should not return an error, if we check if [1]any{nil} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([1]any{nil}, nil)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if []any{} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([0]any{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should not return an error, if we check if [2]any{nil, 78} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([2]any{nil, 78}, nil)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should not return an error, if we check if [2]any{23.2, nil} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([2]any{23.2, nil}, nil)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should not return an error, if we check if [3]any{\"str\", nil, 12} contains nil.", func(t *testing.T) {
		err := assertions.ToContain([3]any{"str", nil, 12}, nil)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]any{\"str\", nil, 12} contains 372.", func(t *testing.T) {
		err := assertions.ToContain([3]any{"str", nil, 12}, 372)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]any{nil} contains \"nil\".", func(t *testing.T) {
		err := assertions.ToContain([1]any{nil}, "nil")
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestStructArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]S{{f: 7}} contains S{f: 7}.", func(t *testing.T) {
		type S struct{ f int }
		err := assertions.ToContain([1]S{{f: 7}}, S{f: 7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]S{} contains S{f: 0}.", func(t *testing.T) {
		type S struct{ f int }
		err := assertions.ToContain([0]S{}, S{f: 0})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]S{{f: 7}, {f: 12}} contains S{f: 7}.", func(t *testing.T) {
		type S struct{ f int }
		err := assertions.ToContain([2]S{{f: 7}, {f: 12}}, S{f: 7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]S{{f: 7}, {f: 12}} contains S{f: 12}.", func(t *testing.T) {
		type S struct{ f int }

		err := assertions.ToContain([2]S{{f: 7}, {f: 12}}, S{f: 12})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]S{{f: 4}, {f: 9}, {f: 10}} contains S{f: 9}.", func(t *testing.T) {
		type S struct{ f int }
		err := assertions.ToContain([3]S{{f: 4}, {f: 9}, {f: 10}}, S{f: 9})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]S{{f: 4}, {f: 9}, {f: 10}} contains S{f:639}.", func(t *testing.T) {
		type S struct{ f int }
		err := assertions.ToContain([3]S{{f: 4}, {f: 9}, {f: 10}}, S{f: 639})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if  [1]S{{f: 77}} contains S{f: \"77\"}.", func(t *testing.T) {
		type S struct{ f any }

		err := assertions.ToContain([1]S{{f: 77}}, S{f: "77"})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if  [1]S1{{f: 77}} contains S2{f: 77}.", func(t *testing.T) {
		type S1 struct{ f int }
		type S2 struct{ f int }

		err := assertions.ToContain([1]S1{{f: 77}}, S2{f: 77})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]S{} contains nil.", func(t *testing.T) {
		type S struct{ f int }

		err := assertions.ToContain([0]S{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestArraysOfMaps(t *testing.T) {
	t.Run("it should return nil, if we check if [1]M{{\"f\": 7}} contains M{{\"f\": 7}}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([1]M{{"f": 7}}, M{"f": 7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]M{} contains M{{\"f\": 0}}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([0]M{}, M{"f": 0})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]M{{\"f\": 7}, {\"f\": 12}} contains M{\"f\": 7}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([2]M{{"f": 7}, {"f": 12}}, M{"f": 7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]M{{\"f\": 7}, {\"f\": 12}} contains M{\"f\": 12}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([2]M{{"f": 7}, {"f": 12}}, M{"f": 12})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]M{{\"f\": 4}, {\"f\": 9}, {\"f\": 10}} contains M{\"f\": 9}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([3]M{{"f": 4}, {"f": 9}, {"f": 10}}, M{"f": 9})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]M{{\"f\": 4}, {\"f\": 9}, {\"f\": 10}} contains M{\"f\": 129}.", func(t *testing.T) {
		type M map[string]int
		err := assertions.ToContain([3]M{{"f": 4}, {"f": 9}, {"f": 10}}, M{"f": 129})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]M1{{\"f\": 7}} contains M2{{\"f\": 7}}.", func(t *testing.T) {
		type M1 map[string]int
		type M2 map[string]int

		err := assertions.ToContain([1]M1{{"f": 7}}, M2{"f": 7})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]M{} contains nil.", func(t *testing.T) {
		type M map[string]int

		err := assertions.ToContain([0]M{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestArraysOfArrays(t *testing.T) {
	t.Run("it should return nil, if we check if [1]IntArr{{7}} contains IntArr{7}.", func(t *testing.T) {
		type IntArr [1]int
		err := assertions.ToContain([1]IntArr{{7}}, IntArr{7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]IntArr{} containsIntArr{7}.", func(t *testing.T) {
		type IntArr [1]int
		err := assertions.ToContain([0]IntArr{}, IntArr{7})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]IntArr{{7, 3}, {12, 4}} contains IntArr{7, 3}.", func(t *testing.T) {
		type IntArr [2]int
		err := assertions.ToContain([2]IntArr{{7, 3}, {12, 4}}, IntArr{7, 3})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]IntArr{{7, 3}, {12, 4}} contains IntArr{12, 4}.", func(t *testing.T) {
		type IntArr [2]int
		err := assertions.ToContain([2]IntArr{{7, 3}, {12, 4}}, IntArr{12, 4})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]IntArr{{7, 3}, {12, 4}, {5, 6}} contains IntArr{12, 4}.", func(t *testing.T) {
		type IntArr [2]int
		err := assertions.ToContain([3]IntArr{{7, 3}, {12, 4}, {5, 6}}, IntArr{12, 4})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]IntArr{{7, 3}, {12, 4}, {5, 6}} contains IntArr{12, 100}.", func(t *testing.T) {
		type IntArr [2]int
		err := assertions.ToContain([3]IntArr{{7, 3}, {12, 4}, {5, 6}}, IntArr{12, 100})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]IntArr1{{7, 3}} contains IntArr2{7, 3}.", func(t *testing.T) {
		type IntArr1 [2]int
		type IntArr2 [2]int

		err := assertions.ToContain([1]IntArr1{{7, 3}}, IntArr2{7, 3})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]IntArr{} contains nil.", func(t *testing.T) {
		type IntArr [2]int
		err := assertions.ToContain([0]IntArr{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestArraysOfSlices(t *testing.T) {
	t.Run("it should return nil, if we check if [1]IntSlice{{7}} contains IntSlice{7}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([1]IntSlice{{7}}, IntSlice{7})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]IntSlice{} contains IntSlice{7}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([0]IntSlice{}, IntSlice{7})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]IntSlice{{7, 3, 9}, {12, 4}} contains IntSlice{7, 3, 9}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([2]IntSlice{{7, 3, 9}, {12, 4}}, IntSlice{7, 3, 9})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [2]IntSlice{{7, 3, 9}, {12, 4}} contains IntSlice{12, 4}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([2]IntSlice{{7, 3, 9}, {12, 4}}, IntSlice{12, 4})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return nil, if we check if [3]IntSlice{{7, 3}, {12, 4}, {5, 6}} contains IntSlice{12, 4}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([3]IntSlice{{7, 3}, {12, 4}, {5, 6, 12, 9}}, IntSlice{12, 4})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]IntSlice{{7, 3}, {12, 4, 9}, {5, 6}} contains IntSlice{12, 4}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([3]IntSlice{{7, 3}, {12, 4, 9}, {5, 6}}, IntSlice{12, 4})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [3]IntArr{{7, 3}, {12, 4}, {5, 6}} contains IntArr{12, 100}.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([3]IntSlice{{7, 3}, {12, 4}, {5, 6}}, IntSlice{12, 100})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [1]IntSlice1{{7, 3}} contains IntSlice2{7, 3}.", func(t *testing.T) {
		type IntSlice1 []int
		type IntSlice2 []int

		err := assertions.ToContain([1]IntSlice1{{7, 3}}, IntSlice2{7, 3})
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("it should return an error, if we check if [0]IntSlice{} contains nil.", func(t *testing.T) {
		type IntSlice []int
		err := assertions.ToContain([0]IntSlice{}, nil)
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}
