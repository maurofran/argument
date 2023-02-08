package argument

import "strings"

// NotEmpty check that the provided string value is not empty.
func NotEmpty(argument string, value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return requiredError{argument: argument}
	}
	return nil
}
