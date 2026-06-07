package assertions

import (
	"fmt"
	"reflect"
)

// ToPanic asserts that the given argument is a no-argument function that panics
// when called.
//
//	ToPanic(func() { panic("boom") })
func ToPanic(fn any) error {
	if fn == nil {
		return fmt.Errorf("%#v is not a function", fn)
	}
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return fmt.Errorf("%#v is not a function", fn)
	}
	if fnValue.Type().NumIn() != 0 {
		return fmt.Errorf("the function passed to ToPanic must take no arguments")
	}

	// Infer the panic from the call not completing, rather than from the
	// recovered value — that way panic(nil) counts as a panic on every Go version.
	didPanic := true
	func() {
		defer func() { recover() }()
		fnValue.Call(nil)
		didPanic = false
	}()

	if !didPanic {
		return fmt.Errorf("the function should have panicked, but it did not")
	}
	return nil
}
