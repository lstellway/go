package volume

import (
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Ratios
	LitreToCubicMetreRatio     = float64(1000)
	CubicMetreToCubicFootRatio = math.Pow(length.MetreToFootRatio, 3)
	ImperialGallonToLitreRatio = 4.54609
	USDryGallonToLitreRatio    = 4.40488377086
	USGallonToCubicInchRatio   = float64(231)
)

type Volume struct {
	Measurement measure.Measurement
}

// Convert a value in cubic metres to cubic litres
func CubicMetreToLitre(value float64) float64 {
	return value * LitreToCubicMetreRatio
}

// Convert a value in litres to cubic metres
func LitreToCubicMetre(value float64) float64 {
	return value / LitreToCubicMetreRatio
}

// CubicFeetToCubicMetres converts a given value in Cubic Feet to its equivalent in Cubic Metres
func CubicFeetToCubicMetres(value float64) float64 {
	return value / CubicMetreToCubicFootRatio
}

// CubicMetresToCubicFeet converts a given value in Cubic Metres to its equivalent Cubic Feet
func CubicMetresToCubicFeet(value float64) float64 {
	return value * CubicMetreToCubicFootRatio
}

// LitresToImperialGallons converts a given value in Litres to its equivalent in Imperial Gallons
func LitresToImperialGallons(value float64) float64 {
	return value / ImperialGallonToLitreRatio
}

// ImperialGallonsToLitres converts a given value in Imperial Gallons to its equivalent in Litres
func ImperialGallonsToLitres(value float64) float64 {
	return value * ImperialGallonToLitreRatio
}

// LitresToUSDryGallons converts a given value in Litres to its equivalent in US Dry Gallons
func LitresToUSDryGallons(value float64) float64 {
	return value / USDryGallonToLitreRatio
}

// USDryGallonsToLitres converts a given value in US Dry Gallons to its equivalent in Litres
func USDryGallonsToLitres(value float64) float64 {
	return value * USDryGallonToLitreRatio
}

// CubicInchesToUSGallons converts a given value from Cubic Inches to its equivalent in US Gallons
func CubicInchesToUSGallons(value float64) float64 {
	return value / USGallonToCubicInchRatio
}

// USGallonsToCubicInches converts a given value from US Gallons to its equivalent in Cubic Inches
func USGallonsToCubicInches(value float64) float64 {
	return value * USGallonToCubicInchRatio
}
