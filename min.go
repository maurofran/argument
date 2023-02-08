package argument

func Min[T Ordered](argument string, value, min T) error {
	if value < min {
		return minError{argument: argument, min: min}
	}
	return nil
}
