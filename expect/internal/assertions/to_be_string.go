package assertions

import "errors"

func ToBeString(value interface{}) error {
	_, ok := value.(string)
	if !ok {
		return errors.New("value is not a string")
	}
	return nil
}
