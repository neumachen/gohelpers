package paratils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceUniqInts(t *testing.T) {
	slice := []int{
		1,
		1,
		1,
		1,
		1,
	}

	result := SliceUniqInts(slice)
	require.Equal(t, 1, len(result))
}
