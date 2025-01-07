package assertions

import (
	"fmt"
	"reflect"
)

func ToBeNil(val any) error {
	if val == nil {
		return nil
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		if reflect.ValueOf(val).IsNil() {
			return nil
		}
	}
	return fmt.Errorf("value should be nil, but it is %v", val)
}
