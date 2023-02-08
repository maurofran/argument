package argument_test

import (
	"github.com/maurofran/argument"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBefore(t *testing.T) {
	tests := map[string]struct {
		reference time.Time
		err       error
	}{
		"Before":    {time.Now().Add(1 * time.Hour), nil},
		"NotBefore": {time.Now().Add(-1 * time.Hour), argument.ErrInvalid},
	}
	fixture := time.Now()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := argument.Before("fixture", fixture, test.reference)
			assert.ErrorIs(t, err, test.err)
		})
	}
}
