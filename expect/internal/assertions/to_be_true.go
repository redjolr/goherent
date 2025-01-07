package assertions

import (
	"fmt"
)

func ToBeTrue(val any) error {
	if val != true {
		return fmt.Errorf("value should be true, but it is %v", val)
	}
	return nil
}
