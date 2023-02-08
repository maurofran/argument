package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFalse(t *testing.T) {
	tests := map[string]struct {
		value bool
		err   error
	}{
		"True":  {true, argument.ErrInvalid},
		"False": {false, nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.False("fixture", test.value)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
