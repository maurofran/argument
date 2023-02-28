package argument

import (
	"fmt"
	"time"
)

// After validates that a given argument occurs after a given reference.
func After(argument string, value, reference time.Time) error {
	if !value.After(reference) {
		return afterError{argument: argument, reference: reference}
	}
	return nil
}

type afterError struct {
	argument  string
	reference time.Time
}

func (err afterError) Error() string {
	return fmt.Sprintf("the argument %s must occurs after %s", err.argument, err.reference.Format(time.RFC3339))
}

func (err afterError) Unwrap() error {
	return ErrInvalid
}
