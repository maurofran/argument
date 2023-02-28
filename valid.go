package argument

import "fmt"

// SelfValidating is the interface implemented by types that validate themselves.
type SelfValidating interface {
	Valid() bool
}

// Valid check if provided SelfValidating value is valid.
func Valid[T SelfValidating](argument string, value T) error {
	if !value.Valid() {
		return notValidError{argument: argument}
	}
	return nil
}

type notValidError struct {
	argument string
}

func (err notValidError) Error() string {
	return fmt.Sprintf("the argument %s is not valid", err.argument)
}

func (err notValidError) Unwrap() error {
	return ErrInvalid
}
