package volume

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Nautical units
	CubicNauticalMile = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.NauticalMile.Singular), Plural: fmt.Sprintf("Cubic %s", length.NauticalMile.Plural), Abbreviation: fmt.Sprintf("cu %s", length.NauticalMile.Abbreviation), Ratio: math.Pow(length.NauticalMile.Ratio, 3)}
	CubicCable        = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Cable.Singular), Plural: fmt.Sprintf("Cubic %s", length.Cable.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Cable.Abbreviation), Ratio: math.Pow(length.Cable.Ratio, 3)}
	CubicFathom       = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Fathom.Singular), Plural: fmt.Sprintf("Cubic %s", length.Fathom.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Fathom.Abbreviation), Ratio: math.Pow(length.Fathom.Ratio, 3)}
	CubicShackle      = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Shackle.Singular), Plural: fmt.Sprintf("Cubic %s", length.Shackle.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Shackle.Abbreviation), Ratio: math.Pow(length.Shackle.Ratio, 3)}
	CubicLeague       = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.League.Singular), Plural: fmt.Sprintf("Cubic %s", length.League.Plural), Abbreviation: fmt.Sprintf("cu %s", length.League.Abbreviation), Ratio: math.Pow(length.League.Ratio, 3)}
)

// CubicNauticalMile converts a volume to its equivalent value in Cubic Nautical Miles
func (v *Volume) CubicNauticalMile() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicNauticalMile)
}

// CubicCable converts a volume to its equivalent value in Cubic Cables
func (v *Volume) CubicCable() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicCable)
}

// CubicFathom converts a volume to its equivalent value in Cubic Fathoms
func (v *Volume) CubicFathom() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicFathom)
}

// CubicShackle converts a volume to its equivalent value in Cubic Shackles
func (v *Volume) CubicShackle() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicShackle)
}

// CubicLeague converts a volume to its equivalent value in Cubic Leagues
func (v *Volume) CubicLeague() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicLeague)
}
