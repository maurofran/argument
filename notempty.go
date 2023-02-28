package argument

import (
	"fmt"
	"strings"
)

// NotEmpty check that the provided string value is not empty.
func NotEmpty(argument string, value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return notEmptyError{argument: argument}
	}
	return nil
}

type notEmptyError struct {
	argument string
}

func (err notEmptyError) Error() string {
	return fmt.Sprintf("the argument %s is empty", err.argument)
}

func (err notEmptyError) Unwrap() error {
	return ErrInvalid
}
