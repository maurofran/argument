package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	tests := map[string]struct {
		other string
		err   error
	}{
		"Equals":    {"aValue", nil},
		"NotEquals": {"otherValue", argument.ErrInvalid},
	}
	fixture := "aValue"
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.Equals("fixture", fixture, test.other)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
