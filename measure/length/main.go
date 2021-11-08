package length

import (
	"github.com/lstellway/go/measure"
)

var (
	// Ratios
	MetreToFootRatio = 39.37 / 12
)

// FeetToMetres converts a given value in Feet to its equivalent in Metres
func FeetToMetres(value float64) float64 {
	return value / MetreToFootRatio
}

// MetresToFeet converts a given value in Metres to its equivalent Feet
func MetresToFeet(value float64) float64 {
	return value * MetreToFootRatio
}

type Length struct {
	Measurement measure.Measurement
}
