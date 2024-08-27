package internal

import "reflect"

func IsFunction(arg any) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}
