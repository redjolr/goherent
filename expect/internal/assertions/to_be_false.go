package assertions

import (
	"fmt"
)

func ToBeFalse(val any) error {
	if val != false {
		return fmt.Errorf("value should be false, but it is %v", val)
	}
	return nil
}
