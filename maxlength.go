package argument

import (
	"fmt"
	"strings"
)

// MaxLength check that the provided value has a maximum length of maxLength.
func MaxLength(argument, value string, maxLength int) error {
	length := len(strings.TrimSpace(value))
	if length > maxLength {
		return maxLengthError{argument: argument, maxLength: maxLength}
	}
	return nil
}

type maxLengthError struct {
	argument  string
	maxLength int
}

func (err maxLengthError) Error() string {
	return fmt.Sprintf("the argument %s must be long at most %d characters", err.argument, err.maxLength)
}

func (err maxLengthError) Unwrap() error {
	return ErrInvalid
}
