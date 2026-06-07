package assertions

import (
	"fmt"
	"reflect"

	"github.com/redjolr/goherent/expect/internal"
)

// ToHaveKey asserts that the given map contains the specified key.
//
//	ToHaveKey(map[string]int{"a": 1}, "a")
func ToHaveKey(m, key any) error {
	if m == nil {
		return fmt.Errorf("%#v is not a map", m)
	}
	if reflect.TypeOf(m).Kind() != reflect.Map {
		return fmt.Errorf("%#v is not a map", m)
	}

	mapValue := reflect.ValueOf(m)
	for _, mapKey := range mapValue.MapKeys() {
		if internal.ObjectsAreEqual(mapKey.Interface(), key) {
			return nil
		}
	}
	return fmt.Errorf("%#v does not have key %#v", m, key)
}
