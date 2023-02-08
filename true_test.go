package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrue(t *testing.T) {
	tests := map[string]struct {
		value bool
		err   error
	}{
		"True":  {true, nil},
		"False": {false, argument.ErrInvalid},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.True("fixture", test.value)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
