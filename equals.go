package argument

import "fmt"

// Equals compare value and other in order to check for equality.
func Equals[T comparable](argument string, value T, other T) error {
	if value != other {
		return equalsError{argument: argument, value: other}
	}
	return nil
}

type equalsError struct {
	argument string
	value    any
}

func (err equalsError) Error() string {
	return fmt.Sprintf("the argument %q cannot be equal to %v", err.argument, err.value)
}

func (err equalsError) Unwrap() error {
	return ErrInvalid
}
