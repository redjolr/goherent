package assertions

import "reflect"

func ToBePositive(isPositiveCandidate any) error {
	zero := reflect.Zero(reflect.TypeOf(isPositiveCandidate))
	return compareTwoValues(isPositiveCandidate, zero.Interface(), []compareResult{compareGreater}, "\"%v\" is not positive")
}
