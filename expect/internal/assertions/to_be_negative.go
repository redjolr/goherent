package assertions

import "reflect"

func ToBeNegative(isNegativeCandidate any) error {
	zero := reflect.Zero(reflect.TypeOf(isNegativeCandidate))
	return compareTwoValues(isNegativeCandidate, zero.Interface(), []compareResult{compareLess}, "\"%v\" is not negative")
}
