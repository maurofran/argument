package argument

import (
	"fmt"
	"reflect"
)

// NotNil check that provided argument is not nil.
func NotNil(argument string, value any) error {
	if isNil(value) {
		return notNilError{argument: argument}
	}
	return nil
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	}
	return false
}

type notNilError struct {
	argument string
}

func (err notNilError) Error() string {
	return fmt.Sprintf("the argument %s is nil", err.argument)
}

func (err notNilError) Unwrap() error {
	return ErrInvalid
}
