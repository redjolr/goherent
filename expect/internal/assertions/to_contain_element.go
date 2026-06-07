package assertions

import (
	"fmt"
	"reflect"

	"github.com/redjolr/goherent/expect/internal"
)

// ToContainElement asserts that the given slice, array or map contains the
// specified element (compared by value). Unlike ToContain, it does not do
// substring matching on strings — the container must be a collection. For maps,
// the element is matched against the map's values (use ToHaveKey to match keys).
//
//	ToContainElement([]string{"Hello", "World"}, "World")
//	ToContainElement(map[string]int{"a": 1}, 1)
func ToContainElement(container, element any) error {
	if container == nil {
		return fmt.Errorf("%#v is not a slice, array or map", container)
	}
	kind := reflect.TypeOf(container).Kind()
	if kind != reflect.Slice && kind != reflect.Array && kind != reflect.Map {
		return fmt.Errorf("%#v is not a slice, array or map", container)
	}

	containerValue := reflect.ValueOf(container)
	if kind == reflect.Map {
		iter := containerValue.MapRange()
		for iter.Next() {
			if internal.ObjectsAreEqual(iter.Value().Interface(), element) {
				return nil
			}
		}
		return fmt.Errorf("%#v does not contain element %#v", container, element)
	}

	for i := 0; i < containerValue.Len(); i++ {
		if internal.ObjectsAreEqual(containerValue.Index(i).Interface(), element) {
			return nil
		}
	}
	return fmt.Errorf("%#v does not contain element %#v", container, element)
}
