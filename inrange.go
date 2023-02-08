package argument

func InRange[T Ordered](argument string, value, min, max T) error {
	if value < min || value > max {
		return rangeError{argument: argument, min: min, max: max}
	}
	return nil
}
