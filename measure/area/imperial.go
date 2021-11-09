package area

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Imperial
	SquareFoot = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Foot.Singular), Plural: fmt.Sprintf("Square %s", length.Foot.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Foot.Abbreviation), Ratio: math.Pow(length.Foot.Ratio, 2)}
	SquareInch = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Inch.Singular), Plural: fmt.Sprintf("Square %s", length.Inch.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Inch.Abbreviation), Ratio: math.Pow(length.Inch.Ratio, 2)}
	SquareYard = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Yard.Singular), Plural: fmt.Sprintf("Square %s", length.Yard.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Yard.Abbreviation), Ratio: math.Pow(length.Yard.Ratio, 2)}
	SquareMile = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Mile.Singular), Plural: fmt.Sprintf("Square %s", length.Mile.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Mile.Abbreviation), Ratio: math.Pow(length.Mile.Ratio, 2)}
)

// ConvertToFoot is a helper to convert a length value from Metres to its equivalent in Feet
func (a *Area) ConvertToUnitFromSquareFeet(unit measure.Unit) (error, float64) {
	return a.ConvertFromSquareMetre(func(value float64) float64 {
		squareFeet := SquareMetresToSquareFeet(value)
		return unit.OneToRatio(squareFeet)
	})
}

// SquareFoot converts an area to its equivalent in Square Feet
func (a *Area) SquareFoot() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareFoot)
}

// SquareInch converts an area to its equivalent in Square Inches
func (a *Area) SquareInch() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareInch)
}

// SquareYard converts an area to its equivalent in Square Yards
func (a *Area) SquareYard() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareYard)
}

// SquareMile converts an area to its equivalent in Square Miles
func (a *Area) SquareMile() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareMile)
}
