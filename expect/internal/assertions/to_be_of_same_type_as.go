package assertions

import (
	"fmt"
	"reflect"

	"github.com/redjolr/goherent/expect/internal"
)

func ToBeOfSameTypeAs(checkValue, expectedType any) error {
	if !internal.ObjectsAreEqual(reflect.TypeOf(checkValue), reflect.TypeOf(expectedType)) {
		return fmt.Errorf(
			"object expected to be of type %v, but was %v",
			reflect.TypeOf(expectedType),
			reflect.TypeOf(checkValue),
		)
	}
	return nil
}
