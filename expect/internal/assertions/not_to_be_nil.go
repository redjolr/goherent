package assertions

import (
	"errors"
	"reflect"
)

func NotToBeNil(val any) error {
	if val == nil {
		return errors.New("value should not be nil, but it is")
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		if reflect.ValueOf(val).IsNil() {
			return errors.New("value should not be nil, but it is")
		}
	}
	return nil
}
