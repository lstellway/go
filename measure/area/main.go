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

type Area struct {
	Measurement measure.Measurement
}

// CreateArea builds a Area struct with a specified unit and Area value
func CreateArea(unit measure.Unit, value float64) Area {
	return Area{
		Measurement: measure.Measurement{
			Value: value,
			Unit:  unit,
		},
	}
}

// SquareFeetToSquareMetres converts a given value in Square Feet to its equivalent in Square Metres
func SquareFeetToSquareMetres(value float64) float64 {
	return value / SquareMetreToSquareFootRatio
}

// SquareMetresToSquareFeet converts a given value in Square Metres to its equivalent Square Feet
func SquareMetresToSquareFeet(value float64) float64 {
	return value * SquareMetreToSquareFootRatio
}
