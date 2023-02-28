package argument

import "fmt"

// Required check if provided comparable value is not equal to zero value.
func Required[T comparable](argument string, value T) error {
	var zero T
	if value == zero {
		return requiredError{argument: argument}
	}
	return nil
}

type requiredError struct {
	argument string
}

func (err requiredError) Error() string {
	return fmt.Sprintf("the argument %s is required", err.argument)
}

func (err requiredError) Unwrap() error {
	return ErrInvalid
}
