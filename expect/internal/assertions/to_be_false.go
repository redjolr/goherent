package assertions

import (
	"errors"
	"fmt"
)

func ToBeFalse(val any) error {
	if val != false {
		return errors.New(fmt.Sprintf("Value should be false, but it is %v.", val))
	}
	return nil
}
