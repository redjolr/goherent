package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

func ToHaveLength(v any, length int) error {
	if v == nil {
		return errors.New("a nil value cannot be checked if it has length, but it is")
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		valueLength := reflect.ValueOf(v).Len()
		if valueLength == length {
			return nil
		}
		return fmt.Errorf("value %v does not have length %d", v, length)
	default:
		typeName := reflect.TypeOf(v).Name()
		return fmt.Errorf("type %v of value %v does not have a length", typeName, v)
	}
}
