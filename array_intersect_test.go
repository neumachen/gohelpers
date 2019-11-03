package paratils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceIntersectStrings(t *testing.T) {
	a := []string{
		"a",
		"b",
		"c",
	}
	b := []string{
		"b",
		"c",
		"d",
	}

	result := SliceIntersectStrings(a, b)
	require.True(t, len(result) == 2)
	require.True(t, reflect.DeepEqual([]string{"b", "c"}, result))
}

func TestSliceIntersectInts(t *testing.T) {
	a := []int{
		1,
		2,
		3,
	}
	b := []int{
		2,
		3,
		4,
	}

	result := SliceIntersectInts(a, b)
	require.True(t, len(result) == 2)
	require.True(t, reflect.DeepEqual([]int{2, 3}, result))
}
