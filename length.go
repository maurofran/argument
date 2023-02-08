package argument

import "strings"

// Length check that provided value length is between minLength and maxLength.
func Length(argument, value string, minLength, maxLength int) error {
	length := len(strings.TrimSpace(value))
	if length < minLength || length > maxLength {
		return lengthError{argument: argument, minLength: minLength, maxLength: maxLength}
	}
	return nil
}
