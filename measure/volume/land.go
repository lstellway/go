package volume

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Land measurements
	CubicFurlong      = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Furlong.Singular), Plural: fmt.Sprintf("Cubic %s", length.Furlong.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Furlong.Abbreviation), Ratio: math.Pow(length.Furlong.Ratio, 3)}
	CubicChain        = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Chain.Singular), Plural: fmt.Sprintf("Cubic %s", length.Chain.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Chain.Abbreviation), Ratio: math.Pow(length.Chain.Ratio, 3)}
	CubicLink         = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Link.Singular), Plural: fmt.Sprintf("Cubic %s", length.Link.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Link.Abbreviation), Ratio: math.Pow(length.Link.Ratio, 3)}
	CubicRod          = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Rod.Singular), Plural: fmt.Sprintf("Cubic %s", length.Rod.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Rod.Abbreviation), Ratio: math.Pow(length.Rod.Ratio, 3)}
	CubicRamsdenChain = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.RamsdenChain.Singular), Plural: fmt.Sprintf("Cubic %s", length.RamsdenChain.Plural), Abbreviation: fmt.Sprintf("cu %s", length.RamsdenChain.Abbreviation), Ratio: math.Pow(length.RamsdenChain.Ratio, 3)}
)

// CubicFurlong converts a volume to its equivalent value in Cubic Furlongs
func (v *Volume) CubicFurlong() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicFurlong)
}

// CubicChain converts a volume to its equivalent value in Cubic Chains
func (v *Volume) CubicChain() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicChain)
}

// CubicLink converts a volume to its equivalent value in Cubic Links
func (v *Volume) CubicLink() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicLink)
}

// CubicRod converts a volume to its equivalent value in Cubic Rods
func (v *Volume) CubicRod() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicRod)
}

// CubicRamsdenChain converts a volume to its equivalent value in Cubic Ramsden Chains
func (v *Volume) CubicRamsdenChain() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicRamsdenChain)
}
