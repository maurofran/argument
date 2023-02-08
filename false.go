package argument

// False will check that the argument value is false
func False(argument string, value bool) error {
	if value {
		return notEqualError{argument: argument, value: true}
	}
	return nil
}
