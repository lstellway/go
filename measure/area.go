package recipe

import (
	"fmt"
	"math"
)

var (
	// Ratios
	SquareMetreToSquareFootRatio = math.Pow(MetreToFootRatio, 2)

	// International System of Units (meters)
	SquareMetre      = Unit{Singular: fmt.Sprintf("Square %s", Metre.Singular), Plural: fmt.Sprintf("Square %s", Metre.Plural), Abbreviation: fmt.Sprintf("%s2", Metre.Abbreviation), Ratio: math.Pow(Metre.Ratio, 2)}
	SquareExametre   = Unit{Singular: fmt.Sprintf("Square %s", Exametre.Singular), Plural: fmt.Sprintf("Square %s", Exametre.Plural), Abbreviation: fmt.Sprintf("%s2", Exametre.Abbreviation), Ratio: math.Pow(Exametre.Ratio, 2)}
	SquarePetametre  = Unit{Singular: fmt.Sprintf("Square %s", Petametre.Singular), Plural: fmt.Sprintf("Square %s", Petametre.Plural), Abbreviation: fmt.Sprintf("%s2", Petametre.Abbreviation), Ratio: math.Pow(Petametre.Ratio, 2)}
	SquareTerametre  = Unit{Singular: fmt.Sprintf("Square %s", Terametre.Singular), Plural: fmt.Sprintf("Square %s", Terametre.Plural), Abbreviation: fmt.Sprintf("%s2", Terametre.Abbreviation), Ratio: math.Pow(Terametre.Ratio, 2)}
	SquareGigametre  = Unit{Singular: fmt.Sprintf("Square %s", Gigametre.Singular), Plural: fmt.Sprintf("Square %s", Gigametre.Plural), Abbreviation: fmt.Sprintf("%s2", Gigametre.Abbreviation), Ratio: math.Pow(Gigametre.Ratio, 2)}
	SquareMegametre  = Unit{Singular: fmt.Sprintf("Square %s", Megametre.Singular), Plural: fmt.Sprintf("Square %s", Megametre.Plural), Abbreviation: fmt.Sprintf("%s2", Megametre.Abbreviation), Ratio: math.Pow(Megametre.Ratio, 2)}
	SquareKilometre  = Unit{Singular: fmt.Sprintf("Square %s", Kilometre.Singular), Plural: fmt.Sprintf("Square %s", Kilometre.Plural), Abbreviation: fmt.Sprintf("%s2", Kilometre.Abbreviation), Ratio: math.Pow(Kilometre.Ratio, 2)}
	SquareHectometre = Unit{Singular: fmt.Sprintf("Square %s", Hectometre.Singular), Plural: fmt.Sprintf("Square %s", Hectometre.Plural), Abbreviation: fmt.Sprintf("%s2", Hectometre.Abbreviation), Ratio: math.Pow(Hectometre.Ratio, 2)}
	SquareDecametre  = Unit{Singular: fmt.Sprintf("Square %s", Decametre.Singular), Plural: fmt.Sprintf("Square %s", Decametre.Plural), Abbreviation: fmt.Sprintf("%s2", Decametre.Abbreviation), Ratio: math.Pow(Decametre.Ratio, 2)}
	SquareDecimetre  = Unit{Singular: fmt.Sprintf("Square %s", Decimetre.Singular), Plural: fmt.Sprintf("Square %s", Decimetre.Plural), Abbreviation: fmt.Sprintf("%s2", Decimetre.Abbreviation), Ratio: math.Pow(Decimetre.Ratio, 2)}
	SquareCentimetre = Unit{Singular: fmt.Sprintf("Square %s", Centimetre.Singular), Plural: fmt.Sprintf("Square %s", Centimetre.Plural), Abbreviation: fmt.Sprintf("%s2", Centimetre.Abbreviation), Ratio: math.Pow(Centimetre.Ratio, 2)}
	SquareMillimetre = Unit{Singular: fmt.Sprintf("Square %s", Millimetre.Singular), Plural: fmt.Sprintf("Square %s", Millimetre.Plural), Abbreviation: fmt.Sprintf("%s2", Millimetre.Abbreviation), Ratio: math.Pow(Millimetre.Ratio, 2)}
	SquareMicrometre = Unit{Singular: fmt.Sprintf("Square %s", Micrometre.Singular), Plural: fmt.Sprintf("Square %s", Micrometre.Plural), Abbreviation: fmt.Sprintf("%s2", Micrometre.Abbreviation), Ratio: math.Pow(Micrometre.Ratio, 2)}
	SquareNanometre  = Unit{Singular: fmt.Sprintf("Square %s", Nanometre.Singular), Plural: fmt.Sprintf("Square %s", Nanometre.Plural), Abbreviation: fmt.Sprintf("%s2", Nanometre.Abbreviation), Ratio: math.Pow(Nanometre.Ratio, 2)}
	SquarePicometre  = Unit{Singular: fmt.Sprintf("Square %s", Picometre.Singular), Plural: fmt.Sprintf("Square %s", Picometre.Plural), Abbreviation: fmt.Sprintf("%s2", Picometre.Abbreviation), Ratio: math.Pow(Picometre.Ratio, 2)}
	SquareFemtometre = Unit{Singular: fmt.Sprintf("Square %s", Femtometre.Singular), Plural: fmt.Sprintf("Square %s", Femtometre.Plural), Abbreviation: fmt.Sprintf("%s2", Femtometre.Abbreviation), Ratio: math.Pow(Femtometre.Ratio, 2)}
	SquareAttometre  = Unit{Singular: fmt.Sprintf("Square %s", Attometre.Singular), Plural: fmt.Sprintf("Square %s", Attometre.Plural), Abbreviation: fmt.Sprintf("%s2", Attometre.Abbreviation), Ratio: math.Pow(Attometre.Ratio, 2)}

	// Imperial
	SquareFoot = Unit{Singular: fmt.Sprintf("Square %s", Foot.Singular), Plural: fmt.Sprintf("Square %s", Foot.Plural), Abbreviation: fmt.Sprintf("sq %s", Foot.Abbreviation), Ratio: math.Pow(Foot.Ratio, 2)}
	SquareInch = Unit{Singular: fmt.Sprintf("Square %s", Inch.Singular), Plural: fmt.Sprintf("Square %s", Inch.Plural), Abbreviation: fmt.Sprintf("sq %s", Inch.Abbreviation), Ratio: math.Pow(Inch.Ratio, 2)}
	SquareYard = Unit{Singular: fmt.Sprintf("Square %s", Yard.Singular), Plural: fmt.Sprintf("Square %s", Yard.Plural), Abbreviation: fmt.Sprintf("sq %s", Yard.Abbreviation), Ratio: math.Pow(Yard.Ratio, 2)}
	SquareMile = Unit{Singular: fmt.Sprintf("Square %s", Mile.Singular), Plural: fmt.Sprintf("Square %s", Mile.Plural), Abbreviation: fmt.Sprintf("sq %s", Mile.Abbreviation), Ratio: math.Pow(Mile.Ratio, 2)}

	// Land measurements
	SquareFurlong      = Unit{Singular: fmt.Sprintf("Square %s", Furlong.Singular), Plural: fmt.Sprintf("Square %s", Furlong.Plural), Abbreviation: fmt.Sprintf("sq %s", Furlong.Abbreviation), Ratio: math.Pow(Furlong.Ratio, 2)}
	SquareChain        = Unit{Singular: fmt.Sprintf("Square %s", Chain.Singular), Plural: fmt.Sprintf("Square %s", Chain.Plural), Abbreviation: fmt.Sprintf("sq %s", Chain.Abbreviation), Ratio: math.Pow(Chain.Ratio, 2)}
	SquareLink         = Unit{Singular: fmt.Sprintf("Square %s", Link.Singular), Plural: fmt.Sprintf("Square %s", Link.Plural), Abbreviation: fmt.Sprintf("sq %s", Link.Abbreviation), Ratio: math.Pow(Link.Ratio, 2)}
	SquareRod          = Unit{Singular: fmt.Sprintf("Square %s", Rod.Singular), Plural: fmt.Sprintf("Square %s", Rod.Plural), Abbreviation: fmt.Sprintf("sq %s", Rod.Abbreviation), Ratio: math.Pow(Rod.Ratio, 2)}
	SquareRamsdenChain = Unit{Singular: fmt.Sprintf("Square %s", RamsdenChain.Singular), Plural: fmt.Sprintf("Square %s", RamsdenChain.Plural), Abbreviation: fmt.Sprintf("sq %s", RamsdenChain.Abbreviation), Ratio: math.Pow(RamsdenChain.Ratio, 2)}

	// Nautical units
	SquareNauticalMile = Unit{Singular: fmt.Sprintf("Square %s", NauticalMile.Singular), Plural: fmt.Sprintf("Square %s", NauticalMile.Plural), Abbreviation: fmt.Sprintf("sq %s", NauticalMile.Abbreviation), Ratio: math.Pow(NauticalMile.Ratio, 2)}
	SquareCable        = Unit{Singular: fmt.Sprintf("Square %s", Cable.Singular), Plural: fmt.Sprintf("Square %s", Cable.Plural), Abbreviation: fmt.Sprintf("sq %s", Cable.Abbreviation), Ratio: math.Pow(Cable.Ratio, 2)}
	SquareFathom       = Unit{Singular: fmt.Sprintf("Square %s", Fathom.Singular), Plural: fmt.Sprintf("Square %s", Fathom.Plural), Abbreviation: fmt.Sprintf("sq %s", Fathom.Abbreviation), Ratio: math.Pow(Fathom.Ratio, 2)}
	SquareShackle      = Unit{Singular: fmt.Sprintf("Square %s", Shackle.Singular), Plural: fmt.Sprintf("Square %s", Shackle.Plural), Abbreviation: fmt.Sprintf("sq %s", Shackle.Abbreviation), Ratio: math.Pow(Shackle.Ratio, 2)}

	// Unofficial
	SquareLeague = Unit{Singular: fmt.Sprintf("Square %s", League.Singular), Plural: fmt.Sprintf("Square %s", League.Plural), Abbreviation: fmt.Sprintf("sq %s", League.Abbreviation), Ratio: math.Pow(League.Ratio, 2)}
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
	Measurement Measurement
}

// SquareMetre converts an area to its equivalent in square metres
func (a *Area) SquareMetre() (error, float64) {
	switch a.Measurement.Unit {
	case SquareMetre:
		// Value is already in Square Metres
	case SquareExametre, SquarePetametre, SquareTerametre, SquareGigametre, SquareMegametre, SquareKilometre, SquareHectometre, SquareDecametre, SquareDecimetre, SquareCentimetre, SquareMillimetre, SquareMicrometre, SquareNanometre, SquarePicometre, SquareFemtometre, SquareAttometre:
		a.Measurement.Value = a.Measurement.Unit.OneToRatio(a.Measurement.Value)
	case SquareFoot, SquareInch, SquareYard, SquareMile, SquareFurlong, SquareChain, SquareLink, SquareRod, SquareRamsdenChain, SquareNauticalMile, SquareCable, SquareFathom, SquareShackle, SquareLeague:
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

// ConvertToFoot is a helper to convert a length value from Metres to its equivalent in Feet
func (a *Area) ConvertToSquareFeet(unit Unit) (error, float64) {
	return a.ConvertFromSquareMetre(func(value float64) float64 {
		return unit.OneToRatio(SquareMetresToSquareFeet(value))
	})
}

// SquareFoot converts an area to its equivalent in Square Feet
func (a *Area) SquareFoot() (error, float64) {
	return a.ConvertToSquareFeet(SquareFoot)
}

// SquareInch converts an area to its equivalent in Square Inches
func (a *Area) SquareInch() (error, float64) {
	return a.ConvertToSquareFeet(SquareInch)
}

// SquareYard converts an area to its equivalent in Square Yards
func (a *Area) SquareYard() (error, float64) {
	return a.ConvertToSquareFeet(SquareYard)
}

// SquareMile converts an area to its equivalent in Square Miles
func (a *Area) SquareMile() (error, float64) {
	return a.ConvertToSquareFeet(SquareMile)
}

// SquareFurlong converts an area to its equivalent in Square Furlongs
func (a *Area) SquareFurlong() (error, float64) {
	return a.ConvertToSquareFeet(SquareFurlong)
}

// SquareChain converts an area to its equivalent in Square Chains
func (a *Area) SquareChain() (error, float64) {
	return a.ConvertToSquareFeet(SquareChain)
}

// SquareLink converts an area to its equivalent in Square Links
func (a *Area) SquareLink() (error, float64) {
	return a.ConvertToSquareFeet(SquareLink)
}

// SquareRod converts an area to its equivalent in Square Rods
func (a *Area) SquareRod() (error, float64) {
	return a.ConvertToSquareFeet(SquareRod)
}

// SquareRamsdenChain converts an area to its equivalent in Square Ramsden Chains
func (a *Area) SquareRamsdenChain() (error, float64) {
	return a.ConvertToSquareFeet(SquareRamsdenChain)
}

// SquareNauticalMile converts an area to its equivalent in Square Nautical Miles
func (a *Area) SquareNauticalMile() (error, float64) {
	return a.ConvertToSquareFeet(SquareNauticalMile)
}

// SquareCable converts an area to its equivalent in Square Cables
func (a *Area) SquareCable() (error, float64) {
	return a.ConvertToSquareFeet(SquareCable)
}

// SquareFathom converts an area to its equivalent in Square Fathoms
func (a *Area) SquareFathom() (error, float64) {
	return a.ConvertToSquareFeet(SquareFathom)
}

// SquareShackle converts an area to its equivalent in Square Shackles
func (a *Area) SquareShackle() (error, float64) {
	return a.ConvertToSquareFeet(SquareShackle)
}

// SquareLeague converts an area to its equivalent in Square Leagues
func (a *Area) SquareLeague() (error, float64) {
	return a.ConvertToSquareFeet(SquareLeague)
}
