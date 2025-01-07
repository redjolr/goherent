package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

func ToBeError(err any) error {
	if err == nil {
		return errors.New("an error is expected but it is in fact nil")
	}
	if !isErrorType(err) {
		return fmt.Errorf("the provided value %v does not implement the error interface", err)
	}
	return nil
}

func isErrorType(v any) bool {
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	vType := reflect.TypeOf(v)
	return vType.Implements(errorInterface)
}
