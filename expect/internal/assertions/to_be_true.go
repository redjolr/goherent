package assertions

import (
	"errors"
	"fmt"
)

func ToBeTrue(val any) error {
	if val != true {
		return errors.New(fmt.Sprintf("Value should be true, but it is %v.", val))
	}
	return nil
}
