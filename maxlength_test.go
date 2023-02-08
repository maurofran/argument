package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxLength(t *testing.T) {
	tests := map[string]struct {
		value string
		err   error
	}{
		"ShorterThanMaxLength": {"f", nil},
		"ExactlyMaxLength":     {"foo", nil},
		"LongerThanMaxLength":  {"foobaz", argument.ErrInvalid},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.MaxLength("fixture", test.value, 3)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
