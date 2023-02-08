package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInRange(t *testing.T) {
	tests := map[string]struct {
		value int
		err   error
	}{
		"LesserThanMin":    {1, argument.ErrInvalid},
		"ExactlyMin":       {2, nil},
		"BetweenMinAndMax": {3, nil},
		"ExactlyMax":       {4, nil},
		"GreaterThanMax":   {5, argument.ErrInvalid},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.InRange("fixture", test.value, 2, 4)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
