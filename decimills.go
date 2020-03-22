package paratils

import (
	"fmt"
	"math"
	"strings"
)

const decimillsExponent = 4
const decimillsMax = 9999

type Decimills interface {
	IntegerPart() int64
	FractionalPart() int64
	Decimills() int64
}

// Decimills ...
type decimills struct {
	integerPart    int64
	fractionalPart int64
	decimills      int64
}

func (d decimills) IntegerPart() int64 {
	return d.integerPart
}

func (d decimills) FractionalPart() int64 {
	return d.fractionalPart
}
func (d decimills) Decimills() int64 {
	return d.decimills
}

// ToDecimills ...
func ToDecimills(rate string) (Decimills, error) {
	f := strings.Split(rate, ".")
	if len(f[1]) > decimillsExponent {
		return nil, Errx(fmt.Errorf("decimals can not be greater than %d digits", decimillsExponent))
	}
	var integer int64
	var fraction int64

	d := &decimills{}

	_, err := fmt.Sscanf(rate, "%d.%d", &integer, &fraction)
	if err != nil {
		return nil, Errx(err)
	}
	d.integerPart = integer

	decimills := func(i int64, s string) int64 {
		maxExponent := decimillsExponent
		maxExponent = maxExponent - len(f[1])
		if maxExponent == 0 {
			return i
		}
		result := math.Pow(float64(10), float64(maxExponent))
		return i * int64(result)
	}(fraction, rate)
	d.fractionalPart = fraction

	if decimills > int64(decimillsMax) {
		return nil, Errx(fmt.Errorf("decimills can not be greater than %d", decimillsMax))
	}

	if integer >= 0 {
		d.decimills = integer*10000 + (int64(decimills))
	} else {

		d.decimills = integer*10000 - (int64(decimills))
	}
	return d, nil
}

// FromDecimills ...
func DecimillsToFloat64(decimills int64) float64 {
	return float64(decimills) / 10000
}
