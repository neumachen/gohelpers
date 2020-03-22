package paratils

import (
	"fmt"
	"math"
	"strings"
)

const decimillsExponent = 4
const decimillsMax = 9999

// ToDecimills ...
func ToDecimills(rate string) (int64, error) {
	f := strings.Split(rate, ".")
	if len(f[1]) > decimillsExponent {
		return int64(0), Errx(fmt.Errorf("decimals can not be greater than %d digits", decimillsExponent))
	}
	var integer int64
	var decimills int64

	_, err := fmt.Sscanf(rate, "%d.%d", &integer, &decimills)
	if err != nil {
		return int64(0), Errx(err)
	}

	decimills = func(i int64, s string) int64 {
		maxExponent := decimillsExponent
		maxExponent = maxExponent - len(f[1])
		if maxExponent == 0 {
			return i
		}
		result := math.Pow(float64(10), float64(maxExponent))
		return i * int64(result)
	}(decimills, rate)

	if decimills > int64(decimillsMax) {
		return int64(0), Errx(fmt.Errorf("decimills can not be greater than %d", decimillsMax))
	}

	if integer >= 0 {
		return integer*10000 + (int64(decimills)), nil
	}
	return integer*10000 - (int64(decimills)), nil
}

// FromDecimills ...
func DecimillsToFloat64(decimills int64) float64 {
	return float64(decimills) / 10000
}
