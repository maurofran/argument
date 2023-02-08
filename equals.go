package argument

// Equals compare value and other in order to check for equality.
func Equals[T comparable](argument string, value T, other T) error {
	if value != other {
		return equalError{argument: argument, value: other}
	}
	return nil
}
