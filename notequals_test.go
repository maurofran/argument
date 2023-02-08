package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotEquals(t *testing.T) {
	tests := map[string]struct {
		other string
		err   error
	}{
		"Equals":    {"aValue", argument.ErrInvalid},
		"NotEquals": {"otherValue", nil},
	}
	fixture := "aValue"
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.NotEquals("fixture", fixture, test.other)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
