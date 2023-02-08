package argument

import "reflect"

// NotNil check that provided argument is not nil.
func NotNil(argument string, value any) error {
	if isNil(value) {
		return requiredError{argument: argument}
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
