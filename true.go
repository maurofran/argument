package argument

// True will check that the argument value is true
func True(argument string, value bool) error {
	if !value {
		return notEqualError{argument: argument, value: true}
	}
	return nil
}
