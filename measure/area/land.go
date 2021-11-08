package area

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Land measurements
	SquareFurlong      = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Furlong.Singular), Plural: fmt.Sprintf("Square %s", length.Furlong.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Furlong.Abbreviation), Ratio: math.Pow(length.Furlong.Ratio, 2)}
	SquareChain        = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Chain.Singular), Plural: fmt.Sprintf("Square %s", length.Chain.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Chain.Abbreviation), Ratio: math.Pow(length.Chain.Ratio, 2)}
	SquareLink         = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Link.Singular), Plural: fmt.Sprintf("Square %s", length.Link.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Link.Abbreviation), Ratio: math.Pow(length.Link.Ratio, 2)}
	SquareRod          = measure.Unit{Singular: fmt.Sprintf("Square %s", length.Rod.Singular), Plural: fmt.Sprintf("Square %s", length.Rod.Plural), Abbreviation: fmt.Sprintf("sq %s", length.Rod.Abbreviation), Ratio: math.Pow(length.Rod.Ratio, 2)}
	SquareRamsdenChain = measure.Unit{Singular: fmt.Sprintf("Square %s", length.RamsdenChain.Singular), Plural: fmt.Sprintf("Square %s", length.RamsdenChain.Plural), Abbreviation: fmt.Sprintf("sq %s", length.RamsdenChain.Abbreviation), Ratio: math.Pow(length.RamsdenChain.Ratio, 2)}
	Rood               = measure.Unit{Singular: "Rood", Plural: "Roods", Abbreviation: "rood", Ratio: SquareRod.Ratio * 40}
	Acre               = measure.Unit{Singular: "Rood", Plural: "Roods", Abbreviation: "rood", Ratio: SquareChain.Ratio * 10}
)

// SquareFurlong converts an area to its equivalent in Square Furlongs
func (a *Area) SquareFurlong() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareFurlong)
}

// SquareChain converts an area to its equivalent in Square Chains
func (a *Area) SquareChain() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareChain)
}

// SquareLink converts an area to its equivalent in Square Links
func (a *Area) SquareLink() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareLink)
}

// SquareRod converts an area to its equivalent in Square Rods
func (a *Area) SquareRod() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareRod)
}

// SquareRamsdenChain converts an area to its equivalent in Square Ramsden Chains
func (a *Area) SquareRamsdenChain() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(SquareRamsdenChain)
}

// Rood converts an area to its equivalent in Roods
func (a *Area) Rood() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(Rood)
}

// Acre converts an area to its equivalent in Acres
func (a *Area) Acre() (error, float64) {
	return a.ConvertToUnitFromSquareFeet(Acre)
}
