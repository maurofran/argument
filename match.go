package argument

import (
	"fmt"
	"regexp"
)

// Match check that given string value matches the provided pattern.
func Match(argument, value string, pattern *regexp.Regexp) error {
	if !pattern.MatchString(value) {
		return matchError{argument: argument, pattern: pattern}
	}
	return nil
}

type matchError struct {
	argument string
	pattern  *regexp.Regexp
}

func (err matchError) Error() string {
	return fmt.Sprintf("the argument %s must match the pattern %q", err.argument, err.pattern.String())
}

func (err matchError) Unwrap() error {
	return ErrInvalid
}
