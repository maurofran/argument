package argument

// Required check if provided comparable value is not equal to zero value.
func Required[T comparable](argument string, value T) error {
	var zero T
	if value == zero {
		return requiredError{argument: argument}
	}
	return nil
}
