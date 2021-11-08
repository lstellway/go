package area

import (
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Ratios
	SquareMetreToSquareFootRatio = math.Pow(length.MetreToFootRatio, 2)
)

// SquareFeetToSquareMetres converts a given value in Square Feet to its equivalent in Square Metres
func SquareFeetToSquareMetres(value float64) float64 {
	return value / SquareMetreToSquareFootRatio
}

// SquareMetresToSquareFeet converts a given value in Square Metres to its equivalent Square Feet
func SquareMetresToSquareFeet(value float64) float64 {
	return value * SquareMetreToSquareFootRatio
}

type Area struct {
	Measurement measure.Measurement
}
