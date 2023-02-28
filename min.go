package argument

import "fmt"

func Min[T Ordered](argument string, value, min T) error {
	if value < min {
		return minError{argument: argument, min: min}
	}
	return nil
}

type minError struct {
	argument string
	min      any
}

func (err minError) Error() string {
	return fmt.Sprintf("the argument %s must be at least %v", err.argument, err.min)
}

func (err minError) Unwrap() error {
	return ErrInvalid
}
