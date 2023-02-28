package argument

import "fmt"

// True will check that the argument value is true
func True(argument string, value bool) error {
	if !value {
		return trueError{argument: argument}
	}
	return nil
}

type trueError struct {
	argument string
}

func (err trueError) Error() string {
	return fmt.Sprintf("the argument %q must be true", err.argument)
}

func (err trueError) Unwrap() error {
	return ErrInvalid
}
