package assertions

import (
	"errors"
)

func NotToBeNil(val any) error {
	if val == nil {
		return errors.New(("Value should not be nil, but it is."))
	}
	return nil
}
