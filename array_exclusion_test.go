package paratils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceExclusionStrings(t *testing.T) {
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

	left, right := SliceExclusionStrings(a, b)
	require.Equal(t, []string{"a"}, left)
	require.Equal(t, []string{"d"}, right)
}

func TestSliceExclusionInts(t *testing.T) {
	a := []int{
		1,
		2,
		3,
	}
	b := []int{
		3,
		5,
		6,
	}

	left, right := SliceExclusionInts(a, b)
	require.Equal(t, []int{1, 2}, left)
	require.Equal(t, []int{5, 6}, right)
}
