package assertions

import (
	"fmt"
)

func NotToBeError(err any) error {
	if err == nil {
		return nil
	}
	if isErrorType(err) {
		return fmt.Errorf("the provided value should not be an error, but it is: %v", err)
	}

	return nil
}
