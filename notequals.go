package argument

// NotEquals compare value and other in order to check for mismatched equality.
func NotEquals[T comparable](argument string, value T, other T) error {
	if value == other {
		return notEqualError{argument: argument, value: other}
	}
	return nil
}
