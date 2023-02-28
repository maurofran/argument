package argument

import (
	"fmt"
	"strings"
)

// Length check that provided value length is between minLength and maxLength.
func Length(argument, value string, minLength, maxLength int) error {
	length := len(strings.TrimSpace(value))
	if length < minLength || length > maxLength {
		return lengthError{argument: argument, minLength: minLength, maxLength: maxLength}
	}
	return nil
}

type lengthError struct {
	argument  string
	minLength int
	maxLength int
}

func (err lengthError) Error() string {
	return fmt.Sprintf("the argument %s must be long between %d and %d characters", err.argument, err.minLength, err.maxLength)
}

func (err lengthError) Unwrap() error {
	return ErrInvalid
}
