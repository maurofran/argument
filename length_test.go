package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLength(t *testing.T) {
	tests := map[string]struct {
		value string
		err   error
	}{
		"ShorterThanMinLength":   {"f", argument.ErrInvalid},
		"ExactlyMinLength":       {"foo", nil},
		"BetweenMinAndMaxLength": {"foo123", nil},
		"ExactlyMaxLength":       {"foobaz", nil},
		"LongerThanMaxLength":    {"foobarbaz", argument.ErrInvalid},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.Length("fixture", test.value, 3, 6)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
