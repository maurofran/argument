package argument

// Error is the custom error type for argument package.
type Error string

// Error implements the standard error interface.
func (err Error) Error() string {
	return string(err)
}

// ErrInvalid is the "base" error returned from argument validations.
const ErrInvalid = Error("invalid argument")
