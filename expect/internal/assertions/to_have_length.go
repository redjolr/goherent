package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

func ToHaveLength(v any, expectedLength int) error {
	if v == nil {
		return errors.New("a nil value cannot be checked if it has length, but it is")
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		actualLength := reflect.ValueOf(v).Len()
		if actualLength == expectedLength {
			return nil
		}
		return fmt.Errorf("expected lenght: %d. actual length: %d ", expectedLength, actualLength)
	default:
		typeName := reflect.TypeOf(v).Name()
		return fmt.Errorf("type %v cannot be passed to the len() function, so it does not have a length", typeName)
	}
}
