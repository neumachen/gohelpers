package paratils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToDecimills(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			String    string
			Decimills int64
		}{
			{
				String:    "12.2200",
				Decimills: 122200,
			},
			{
				String:    "0.2200",
				Decimills: 2200,
			},
			{
				String:    "1.2",
				Decimills: 12000,
			},
			{
				String:    "1.012",
				Decimills: 10120,
			},
			{
				String:    "0.1",
				Decimills: 1000,
			},
		}

		for i := range tests {
			d, err := ToDecimills(tests[i].String)
			require.NoError(t, err)
			// since the default exponent is 4, it's expected that the
			// decimills will have 4 spaces on the fractional part of the
			// number. Example, 12.22 == 12.2200 == 122200
			require.Equal(t, tests[i].Decimills, d.Decimills())
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Run("more than 4 decimal digits", func(t *testing.T) {
			decimills, err := ToDecimills("1.12345")
			require.Error(t, err)
			require.Equal(t, fmt.Sprintf("decimals can not be greater than %d digits", decimillsExponent), err.Error())
			require.Nil(t, decimills)
		})
	})
}

func TestDecimillsToFloat64(t *testing.T) {
	f := DecimillsToFloat64(12000)
	require.Equal(t, 1.2, f)
}
