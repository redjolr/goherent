package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

func ToHaveLengthLessThan(v any, greaterThanLength int) error {
	if v == nil {
		return errors.New("a nil value cannot be checked if it has length, but it is")
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		actualLength := reflect.ValueOf(v).Len()
		if actualLength < greaterThanLength {
			return nil
		}
		return fmt.Errorf("length must be grealesster than: %d. actual length: %d ", greaterThanLength, actualLength)
	default:
		typeName := reflect.TypeOf(v).Name()
		return fmt.Errorf("type %v cannot be passed to the len() function, so it does not have a length", typeName)
	}
}
