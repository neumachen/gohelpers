package paratils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type isNilTest interface {
	Foo()
}

func TestIsNil(t *testing.T) {
	var testInterface isNilTest
	tests := []struct {
		actual   interface{}
		expected bool
	}{
		{
			actual:   nil,
			expected: true,
		},
		{
			actual:   []byte(nil),
			expected: true,
		},
		{
			actual:   testInterface,
			expected: true,
		},
	}

	for i := range tests {
		require.Equal(t, tests[i].expected, IsNil(tests[i].actual))
	}

}
