package paratils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveEmptyString(t *testing.T) {
	tests := []struct {
		actual   []string
		expected []string
	}{
		{
			actual:   []string{"a", "b", "  ", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			actual:   []string{"", "b", "  ", "c"},
			expected: []string{"b", "c"},
		},
	}

	for i := range tests {
		RemoveEmptyString(&tests[i].actual)
		require.Equal(t, tests[i].expected, tests[i].actual)
	}
}
