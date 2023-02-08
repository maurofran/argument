package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequired_String(t *testing.T) {
	tests := map[string]struct {
		value string
		err   error
	}{
		"EmptyString":      {"", argument.ErrInvalid},
		"StringWithLength": {"foo", nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.Required("fixture", test.value)
			assert.ErrorIs(t, err, test.err)
		})
	}
}

func TestRequired_Pointer(t *testing.T) {
	empty := ""
	valued := "foo"

	tests := map[string]struct {
		value *string
		err   error
	}{
		"Nil":              {nil, argument.ErrInvalid},
		"EmptyString":      {&empty, nil},
		"StringWithLength": {&valued, nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.Required("fixture", test.value)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
