package assertions

import (
	"errors"
	"fmt"
)

func ToBeNil(val any) error {
	if val != nil {
		return errors.New(fmt.Sprintf("Value should be nil, but it is %v.", val))
	}
	return nil
}
