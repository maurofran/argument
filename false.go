package argument

import "fmt"

// False will check that the argument value is false
func False(argument string, value bool) error {
	if value {
		return falseError{argument: argument}
	}
	return nil
}

type falseError struct {
	argument string
}

func (err falseError) Error() string {
	return fmt.Sprintf("the argument %q must be false", err.argument)
}

func (err falseError) Unwrap() error {
	return ErrInvalid
}
