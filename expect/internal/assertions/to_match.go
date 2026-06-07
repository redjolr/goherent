package assertions

import (
	"fmt"
	"regexp"
)

// ToMatch asserts that the given string matches the regular expression pattern.
//
//	ToMatch("goherent v1.2.3", `^goherent v\d+\.\d+\.\d+$`)
func ToMatch(val any, pattern string) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("%#v is not a string", val)
	}
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return fmt.Errorf("invalid regex pattern %q: %s", pattern, err)
	}
	if !matched {
		return fmt.Errorf("%q does not match pattern %q", str, pattern)
	}
	return nil
}
