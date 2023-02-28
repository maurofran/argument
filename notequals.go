package argument

import "fmt"

// NotEquals compare value and other in order to check for mismatched equality.
func NotEquals[T comparable](argument string, value T, other T) error {
	if value == other {
		return notEqualsError{argument: argument, value: other}
	}
	return nil
}

type notEqualsError struct {
	argument string
	value    any
}

func (err notEqualsError) Error() string {
	return fmt.Sprintf("the argument %q must be equal to %v", err.argument, err.value)
}

func (err notEqualsError) Unwrap() error {
	return ErrInvalid
}
