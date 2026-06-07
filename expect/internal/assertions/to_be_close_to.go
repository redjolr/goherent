package assertions

import (
	"fmt"
	"math"
	"reflect"
)

// ToBeCloseTo asserts that a numeric value is within tolerance (inclusive) of
// target. value, target and tolerance are all converted to float64, so mixed
// numeric types are accepted.
//
//	ToBeCloseTo(3.14159, 3.14, 0.01)
func ToBeCloseTo(val, target, tolerance any) error {
	v, ok := numericAsFloat64(val)
	if !ok {
		return fmt.Errorf("%#v is not a number", val)
	}
	tg, ok := numericAsFloat64(target)
	if !ok {
		return fmt.Errorf("%#v is not a number", target)
	}
	tol, ok := numericAsFloat64(tolerance)
	if !ok {
		return fmt.Errorf("%#v is not a number", tolerance)
	}

	if diff := math.Abs(v - tg); diff > tol {
		return fmt.Errorf("%v is not within %v of %v (differs by %v)", v, tol, tg, diff)
	}
	return nil
}

func numericAsFloat64(val any) (float64, bool) {
	if val == nil {
		return 0, false
	}
	reflectVal := reflect.ValueOf(val)
	switch reflectVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflectVal.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(reflectVal.Uint()), true
	case reflect.Float32, reflect.Float64:
		return reflectVal.Float(), true
	}
	return 0, false
}
