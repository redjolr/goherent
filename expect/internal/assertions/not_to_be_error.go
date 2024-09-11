package assertions

import (
	"errors"
	"fmt"
)

func NotToBeError(err any) error {
	if err == nil {
		return nil
	}
	if isErrorType(err) {
		return errors.New(fmt.Sprintf("The provided value should not be an error, but it is: %v", err))
	}

	return nil
}