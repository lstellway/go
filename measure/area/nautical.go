package area

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Nautical units
	SquareNauticalMile = measure.Unit{Singular: fmt.Sprintf("Square %s", length.NauticalMile.Singular), Plural: fmt.Sprintf("Square %s", length.NauticalMile.Plural), Abbreviation: fmt.Sprintf("sq %s", length.NauticalMile.Abbreviation), Ratio: math.Pow(length.NauticalMile.Ratio, 2)}
	SquareCable        = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Cable.Singular), Plural: fmt.Sprintf("Square %s", length.Cable.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Cable.Abbreviation), Ratio: math.Pow(length.Cable.Ratio, 2)}
	SquareFathom       = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Fathom.Singular), Plural: fmt.Sprintf("Square %s", length.Fathom.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Fathom.Abbreviation), Ratio: math.Pow(length.Fathom.Ratio, 2)}
	SquareShackle      = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Shackle.Singular), Plural: fmt.Sprintf("Square %s", length.Shackle.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Shackle.Abbreviation), Ratio: math.Pow(length.Shackle.Ratio, 2)}
	SquareLeague       = measure.Unit{Singular: fmt.Sprintf("Square %s", length.League.Singular), Plural: fmt.Sprintf("Square %s", length.League.Plural), Abbreviation: fmt.Sprintf("sq %s", length.League.Abbreviation), Ratio: math.Pow(length.League.Ratio, 2)}
)

// SquareNauticalMile converts an area to its equivalent in Square Nautical Miles
func (a *Area) SquareNauticalMile() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareNauticalMile)
}

// SquareCable converts an area to its equivalent in Square Cables
func (a *Area) SquareCable() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareCable)
}

// SquareFathom converts an area to its equivalent in Square Fathoms
func (a *Area) SquareFathom() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareFathom)
}

// SquareShackle converts an area to its equivalent in Square Shackles
func (a *Area) SquareShackle() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareShackle)
}

// SquareLeague converts an area to its equivalent in Square Leagues
func (a *Area) SquareLeague() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareLeague)
}
