package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	tests := map[string]struct {
		value string
		err   error
	}{
		"EmptyString":      {"", argument.ErrInvalid},
		"BlankString":      {"   ", argument.ErrInvalid},
		"StringWithLength": {"foo", nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.NotEmpty("fixture", test.value)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
