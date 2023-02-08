package argument

import "time"

// After validates that a given argument occurs after a given reference.
func After(argument string, value, reference time.Time) error {
	if !value.After(reference) {
		return afterError{argument: argument, reference: reference}
	}
	return nil
}
