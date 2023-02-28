package argument

import (
	"fmt"
	"time"
)

// Before validates that a given argument occurs before a given reference.
func Before(argument string, value, reference time.Time) error {
	if !value.Before(reference) {
		return beforeError{argument: argument, reference: reference}
	}
	return nil
}

type beforeError struct {
	argument  string
	reference time.Time
}

func (err beforeError) Error() string {
	return fmt.Sprintf("the argument %s must occurs before %s", err.argument, err.reference.Format(time.RFC3339))
}

func (err beforeError) Unwrap() error {
	return ErrInvalid
}
