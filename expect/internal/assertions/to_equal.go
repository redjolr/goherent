package assertions

import (
	"errors"
	"fmt"

	"github.com/redjolr/goherent/expect/internal"
)

func ToEqual(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}
	if internal.IsFunction(expected) || internal.IsFunction(actual) {
		return errors.New(fmt.Sprintf("Invalid operation: %#v == %#v (cannot take func type as argument)", expected, actual))
	}

	if !internal.ObjectsAreEqual(expected, actual) {
		diff := internal.Diff(expected, actual)
		expected, actual = internal.FormatUnequalValues(expected, actual)
		return errors.New(
			fmt.Sprintf("Not equal:\n"+
				"expected: %s\n"+
				"actual  : %s%s", expected, actual, diff),
		)
	}
	return nil
}
