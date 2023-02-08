package argument

import "strings"

// MaxLength check that the provided value has a maximum length of maxLength.
func MaxLength(argument, value string, maxLength int) error {
	length := len(strings.TrimSpace(value))
	if length > maxLength {
		return maxLengthError{argument: argument, maxLength: maxLength}
	}
	return nil
}
