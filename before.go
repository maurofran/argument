package argument

import "time"

// Before validates that a given argument occurs before a given reference.
func Before(argument string, value, reference time.Time) error {
	if !value.Before(reference) {
		return beforeError{argument: argument, reference: reference}
	}
	return nil
}
