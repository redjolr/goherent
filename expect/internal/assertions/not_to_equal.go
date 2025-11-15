package assertions

import (
	"fmt"

	"github.com/redjolr/goherent/expect/internal"
)

func NotToEqual(expected, actual any) error {
	if expected == nil && actual == nil {
		return fmt.Errorf("values should not be equal, but both are nil")
	}
	if internal.IsFunction(expected) || internal.IsFunction(actual) {
		return fmt.Errorf("invalid operation: %#v != %#v (cannot take func type as argument)", expected, actual)
	}

	if internal.ObjectsAreEqual(expected, actual) {
		expected, actual = internal.FormatUnequalValues(expected, actual)
		return fmt.Errorf("values should not be equal, but both are:\n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual,
		)
	}
	return nil
}
