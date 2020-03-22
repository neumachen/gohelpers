package paratils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToDecimills(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		type TestCase struct {
			String         string
			Decimills      int64
			IntegerPart    int64
			FractionalPart int64
		}

		testCases := []TestCase{
			TestCase{
				String:         "10.1011",
				Decimills:      101011,
				IntegerPart:    10,
				FractionalPart: 1011,
			},
			TestCase{
				String:         "1.0001",
				Decimills:      10001,
				IntegerPart:    1,
				FractionalPart: 1,
			},
			TestCase{
				String:         "1.0234",
				Decimills:      10234,
				IntegerPart:    1,
				FractionalPart: 234,
			},
			TestCase{
				String:         "1.1234",
				Decimills:      11234,
				IntegerPart:    1,
				FractionalPart: 1234,
			},
			TestCase{
				String:         "1.01",
				Decimills:      10100,
				IntegerPart:    1,
				FractionalPart: 100,
			},
			TestCase{
				String:         "1.1",
				Decimills:      11000,
				IntegerPart:    1,
				FractionalPart: 1000,
			},
			TestCase{
				String:         "0.1111",
				Decimills:      1111,
				IntegerPart:    0,
				FractionalPart: 1111,
			},
			TestCase{
				String:         "0.1",
				Decimills:      1000,
				IntegerPart:    0,
				FractionalPart: 1000,
			},
			TestCase{
				String:         "0.012",
				Decimills:      120,
				IntegerPart:    0,
				FractionalPart: 120,
			},
			TestCase{
				String:         "1.012",
				Decimills:      10120,
				IntegerPart:    1,
				FractionalPart: 120,
			},
		}

		for i := range testCases {
			d, err := ToDecimills(testCases[i].String)
			require.NoError(t, err)
			// since the default exponent is 4, it's expected that the
			// decimills will have 4 spaces on the fractional part of the
			// number. Example, 12.22 == 12.2200 == 122200
			require.Equal(t, testCases[i].Decimills, d.Decimills())
			require.Equal(t, testCases[i].IntegerPart, d.IntegerPart())
			require.Equal(t, testCases[i].FractionalPart, d.FractionalPart())
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
