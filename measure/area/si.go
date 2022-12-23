package area

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

// International System of Units (meters)
var (
	SquareMetre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Metre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Metre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Metre.Abbreviation),
		Ratio: math.Pow(length.Metre.Ratio, 2),
	}

	SquareYottametre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Yottametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Yottametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Yottametre.Abbreviation),
		Ratio: math.Pow(length.Yottametre.Ratio, 2),
	}

	SquareZettametre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Zettametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Zettametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Zettametre.Abbreviation),
		Ratio: math.Pow(length.Zettametre.Ratio, 2),
	}

	SquareExametre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Exametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Exametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Exametre.Abbreviation),
		Ratio: math.Pow(length.Exametre.Ratio, 2),
	}

	SquarePetametre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Petametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Petametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Petametre.Abbreviation),
		Ratio: math.Pow(length.Petametre.Ratio, 2),
	}

	SquareTerametre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Terametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Terametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Terametre.Abbreviation),
		Ratio: math.Pow(length.Terametre.Ratio, 2),
	}

	SquareGigametre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Gigametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Gigametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Gigametre.Abbreviation),
		Ratio: math.Pow(length.Gigametre.Ratio, 2),
	}

	SquareMegametre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Megametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Megametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Megametre.Abbreviation),
		Ratio: math.Pow(length.Megametre.Ratio, 2),
	}

	SquareKilometre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Kilometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Kilometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Kilometre.Abbreviation),
		Ratio: math.Pow(length.Kilometre.Ratio, 2),
	}

	SquareHectometre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Hectometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Hectometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Hectometre.Abbreviation),
		Ratio: math.Pow(length.Hectometre.Ratio, 2),
	}

	SquareDecametre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Decametre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Decametre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Decametre.Abbreviation),
		Ratio: math.Pow(length.Decametre.Ratio, 2),
	}

	SquareDecimetre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Decimetre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Decimetre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Decimetre.Abbreviation),
		Ratio: math.Pow(length.Decimetre.Ratio, 2),
	}

	SquareCentimetre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Centimetre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Centimetre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Centimetre.Abbreviation),
		Ratio: math.Pow(length.Centimetre.Ratio, 2),
	}

	SquareMillimetre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Millimetre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Millimetre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Millimetre.Abbreviation),
		Ratio: math.Pow(length.Millimetre.Ratio, 2),
	}

	SquareMicrometre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Micrometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Micrometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Micrometre.Abbreviation),
		Ratio: math.Pow(length.Micrometre.Ratio, 2),
	}

	SquareNanometre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Nanometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Nanometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Nanometre.Abbreviation),
		Ratio: math.Pow(length.Nanometre.Ratio, 2),
	}

	SquarePicometre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Picometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Picometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Picometre.Abbreviation),
		Ratio: math.Pow(length.Picometre.Ratio, 2),
	}

	SquareFemtometre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Femtometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Femtometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Femtometre.Abbreviation),
		Ratio: math.Pow(length.Femtometre.Ratio, 2),
	}

	SquareAttometre  = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Attometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Attometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Attometre.Abbreviation),
		Ratio: math.Pow(length.Attometre.Ratio, 2),
	}

	SquareZeptometre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Zeptometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Zeptometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Zeptometre.Abbreviation),
		Ratio: math.Pow(length.Zeptometre.Ratio, 2),
	}

	SquareYoctometre = measure.Unit{
		Singular: fmt.Sprintf("Square %s", length.Yoctometre.Singular),
		Plural: fmt.Sprintf("Square %s", length.Yoctometre.Plural),
		Abbreviation: fmt.Sprintf("%s2", length.Yoctometre.Abbreviation),
		Ratio: math.Pow(length.Yoctometre.Ratio, 2),
	}
)

// SquareMetre converts an area to its equivalent in square metres
func (a *Area) SquareMetre() (error, float64) {
	switch a.Measurement.Unit {
	// International System of Units (SI)
	case SquareMetre, SquareYottametre, SquareZettametre, SquareExametre, SquarePetametre, SquareTerametre, SquareGigametre, SquareMegametre, SquareKilometre, SquareHectometre, SquareDecametre, SquareDecimetre, SquareCentimetre, SquareMillimetre, SquareMicrometre, SquareNanometre, SquarePicometre, SquareFemtometre, SquareAttometre, SquareZeptometre, SquareYoctometre:
		a.Measurement.Value = a.Measurement.Unit.RatioToOne(a.Measurement.Value)
	// Imperial
	case SquareFoot, SquareInch, SquareYard, SquareMile:
		squareFeet := a.Measurement.Unit.RatioToOne(a.Measurement.Value)
		a.Measurement.Value = SquareFeetToSquareMetres(squareFeet)
	// Land
	case SquareFurlong, SquareChain, SquareLink, SquareRod, SquareRamsdenChain, Rood, Acre:
		squareFeet := a.Measurement.Unit.RatioToOne(a.Measurement.Value)
		a.Measurement.Value = SquareFeetToSquareMetres(squareFeet)
	// Nautical
	case SquareNauticalMile, SquareCable, SquareFathom, SquareShackle, SquareLeague:
		squareFeet := a.Measurement.Unit.RatioToOne(a.Measurement.Value)
		a.Measurement.Value = SquareFeetToSquareMetres(squareFeet)
	default:
		err := fmt.Errorf("Unknown area unit, %v", a.Measurement.Unit)
		return err, 0
	}

	a.Measurement.Unit = SquareMetre

	return nil, a.Measurement.Value
}

// ConvertFromSquareMetre is a helper to convert an area value from its value in Metres
func (a *Area) ConvertFromSquareMetre(handle func(float64) float64) (error, float64) {
	err, squareMetres := a.SquareMetre()
	if err != nil {
		return err, 0
	}

	return nil, handle(squareMetres)
}

// SquareYottametre converts an area to its equivalent in SquareYottametres
func (a *Area) SquareYottametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareYottametre.OneToRatio)
}

// SquareZettametre converts an area to its equivalent in SquareZettametres
func (a *Area) SquareZettametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareZettametre.OneToRatio)
}

// SquareExametre converts an area to its equivalent in SquareExametres
func (a *Area) SquareExametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareExametre.OneToRatio)
}

// SquarePetametre converts an area to its equivalent in SquarePetametres
func (a *Area) SquarePetametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquarePetametre.OneToRatio)
}

// SquareTerametre converts an area to its equivalent in SquareTerametres
func (a *Area) SquareTerametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareTerametre.OneToRatio)
}

// SquareGigametre converts an area to its equivalent in SquareGigametres
func (a *Area) SquareGigametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareGigametre.OneToRatio)
}

// SquareMegametre converts an area to its equivalent in SquareMegametres
func (a *Area) SquareMegametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareMegametre.OneToRatio)
}

// SquareKilometre converts an area to its equivalent in SquareKilometres
func (a *Area) SquareKilometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareKilometre.OneToRatio)
}

// SquareHectometre converts an area to its equivalent in SquareHectometres
func (a *Area) SquareHectometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareHectometre.OneToRatio)
}

// SquareDecametre converts an area to its equivalent in SquareDecametres
func (a *Area) SquareDecametre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareDecametre.OneToRatio)
}

// SquareDecimetre converts an area to its equivalent in SquareDecimetres
func (a *Area) SquareDecimetre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareDecimetre.OneToRatio)
}

// SquareCentimetre converts an area to its equivalent in SquareCentimetres
func (a *Area) SquareCentimetre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareCentimetre.OneToRatio)
}

// SquareMillimetre converts an area to its equivalent in SquareMillimetres
func (a *Area) SquareMillimetre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareMillimetre.OneToRatio)
}

// SquareMicrometre converts an area to its equivalent in SquareMicrometres
func (a *Area) SquareMicrometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareMicrometre.OneToRatio)
}

// SquareNanometre converts an area to its equivalent in SquareNanometres
func (a *Area) SquareNanometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareNanometre.OneToRatio)
}

// SquarePicometre converts an area to its equivalent in SquarePicometres
func (a *Area) SquarePicometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquarePicometre.OneToRatio)
}

// SquareFemtometre converts an area to its equivalent in SquareFemtometres
func (a *Area) SquareFemtometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareFemtometre.OneToRatio)
}

// SquareAttometre converts an area to its equivalent in SquareAttometres
func (a *Area) SquareAttometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareAttometre.OneToRatio)
}

// SquareZeptometre converts an area to its equivalent in SquareZeptometres
func (a *Area) SquareZeptometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareZeptometre.OneToRatio)
}

// SquareYoctometre converts an area to its equivalent in SquareYoctometres
func (a *Area) SquareYoctometre() (error, float64) {
	return a.ConvertFromSquareMetre(SquareYoctometre.OneToRatio)
}
