package argument

import "fmt"

func InRange[T Ordered](argument string, value, min, max T) error {
	if value < min || value > max {
		return inRangeError{argument: argument, min: min, max: max}
	}
	return nil
}

type inRangeError struct {
	argument string
	min      any
	max      any
}

func (err inRangeError) Error() string {
	return fmt.Sprintf("the argument %s must be between %v and %v", err.argument, err.min, err.max)
}

func (err inRangeError) Unwrap() error {
	return ErrInvalid
}
