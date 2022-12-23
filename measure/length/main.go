package length

import (
	"math/big"
	"github.com/lstellway/go/measure"
)

var (
	MetresPerFoot float64 = 39.37

	MetreToFootRatio = big.NewFloat(0).Quo(
		big.NewFloat(MetresPerFoot),
		big.NewFloat(InchesPerFoot),
	)
)

type Length struct {
	Measurement measure.Measurement
}

// CreateLength builds a Length struct with a specified unit and Length value
func CreateLength(unit measure.Unit, value float64) Length {
	return Length{
		Measurement: measure.Measurement{
			Value: value,
			Unit:  unit,
		},
	}
}

// FeetToMetres converts a given value in Feet to its equivalent in Metres
func FeetToMetres(value float64) float64 {
	bigValue := big.NewFloat(value)
	quotient, _ := big.NewFloat(0).Quo(bigValue, MetreToFootRatio).Float64()
	return quotient
}

// MetresToFeet converts a given value in Metres to its equivalent Feet
func MetresToFeet(value float64) float64 {
	bigValue := big.NewFloat(value)
	product, _ := big.NewFloat(0).Mul(bigValue, MetreToFootRatio).Float64()
	return product
}

