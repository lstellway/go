package recipe

import (
	"fmt"
	"math"
)

// https://github.com/bcicen/go-units
// https://en.wikipedia.org/wiki/Unit_of_volume
// https://www.asknumbers.com/volumeconverter.aspx
var (
	// Ratios
	LitreToCubicMetreRatio     = float64(1000)
	CubicMetreToCubicFootRatio = math.Pow(MetreToFootRatio, 3)
	ImperialGallonToLitreRatio = 4.54609
	USDryGallonToLitreRatio    = 4.40488377086
	USGallonToCubicInchRatio   = float64(231)
	BushelToUSDryGallonRatio   = float64(8)

	// Litres
	Litre      = Unit{Singular: "Litre", Plural: "Litres", Abbreviation: "L", Ratio: 1} // Base 1
	Exalitre   = Exa.Unit("%slitre", "%sL")
	Petalitre  = Peta.Unit("%slitre", "%sL")
	Teralitre  = Tera.Unit("%slitre", "%sL")
	Gigalitre  = Giga.Unit("%slitre", "%sL")
	Megalitre  = Mega.Unit("%slitre", "%sL")
	Kilolitre  = Kilo.Unit("%slitre", "%sL")
	Hectolitre = Hecto.Unit("%slitre", "%sL")
	Decalitre  = Deca.Unit("%slitre", "%sL")
	Decilitre  = Deci.Unit("%slitre", "%sL")
	Centilitre = Centi.Unit("%slitre", "%sL")
	Millilitre = Milli.Unit("%slitre", "%sL")
	Microlitre = Micro.Unit("%slitre", "%sL")
	Nanolitre  = Nano.Unit("%slitre", "%sL")
	Picolitre  = Pico.Unit("%slitre", "%sL")
	Femtolitre = Femto.Unit("%slitre", "%sL")
	Attolitre  = Atto.Unit("%slitre", "%sL")

	// Cubic Metres
	CubicMetre      = Unit{Singular: fmt.Sprintf("Cubic %s", Metre.Singular), Plural: fmt.Sprintf("Cubic %s", Metre.Plural), Abbreviation: fmt.Sprintf("%s3", Metre.Abbreviation), Ratio: math.Pow(Metre.Ratio, 3)}
	CubicExametre   = Unit{Singular: fmt.Sprintf("Cubic %s", Exametre.Singular), Plural: fmt.Sprintf("Cubic %s", Exametre.Plural), Abbreviation: fmt.Sprintf("%s3", Exametre.Abbreviation), Ratio: math.Pow(Exametre.Ratio, 3)}
	CubicPetametre  = Unit{Singular: fmt.Sprintf("Cubic %s", Petametre.Singular), Plural: fmt.Sprintf("Cubic %s", Petametre.Plural), Abbreviation: fmt.Sprintf("%s3", Petametre.Abbreviation), Ratio: math.Pow(Petametre.Ratio, 3)}
	CubicTerametre  = Unit{Singular: fmt.Sprintf("Cubic %s", Terametre.Singular), Plural: fmt.Sprintf("Cubic %s", Terametre.Plural), Abbreviation: fmt.Sprintf("%s3", Terametre.Abbreviation), Ratio: math.Pow(Terametre.Ratio, 3)}
	CubicGigametre  = Unit{Singular: fmt.Sprintf("Cubic %s", Gigametre.Singular), Plural: fmt.Sprintf("Cubic %s", Gigametre.Plural), Abbreviation: fmt.Sprintf("%s3", Gigametre.Abbreviation), Ratio: math.Pow(Gigametre.Ratio, 3)}
	CubicMegametre  = Unit{Singular: fmt.Sprintf("Cubic %s", Megametre.Singular), Plural: fmt.Sprintf("Cubic %s", Megametre.Plural), Abbreviation: fmt.Sprintf("%s3", Megametre.Abbreviation), Ratio: math.Pow(Megametre.Ratio, 3)}
	CubicKilometre  = Unit{Singular: fmt.Sprintf("Cubic %s", Kilometre.Singular), Plural: fmt.Sprintf("Cubic %s", Kilometre.Plural), Abbreviation: fmt.Sprintf("%s3", Kilometre.Abbreviation), Ratio: math.Pow(Kilometre.Ratio, 3)}
	CubicHectometre = Unit{Singular: fmt.Sprintf("Cubic %s", Hectometre.Singular), Plural: fmt.Sprintf("Cubic %s", Hectometre.Plural), Abbreviation: fmt.Sprintf("%s3", Hectometre.Abbreviation), Ratio: math.Pow(Hectometre.Ratio, 3)}
	CubicDecametre  = Unit{Singular: fmt.Sprintf("Cubic %s", Decametre.Singular), Plural: fmt.Sprintf("Cubic %s", Decametre.Plural), Abbreviation: fmt.Sprintf("%s3", Decametre.Abbreviation), Ratio: math.Pow(Decametre.Ratio, 3)}
	CubicDecimetre  = Unit{Singular: fmt.Sprintf("Cubic %s", Decimetre.Singular), Plural: fmt.Sprintf("Cubic %s", Decimetre.Plural), Abbreviation: fmt.Sprintf("%s3", Decimetre.Abbreviation), Ratio: math.Pow(Decimetre.Ratio, 3)}
	CubicCentimetre = Unit{Singular: fmt.Sprintf("Cubic %s", Centimetre.Singular), Plural: fmt.Sprintf("Cubic %s", Centimetre.Plural), Abbreviation: fmt.Sprintf("%s3", Centimetre.Abbreviation), Ratio: math.Pow(Centimetre.Ratio, 3)}
	CubicMillimetre = Unit{Singular: fmt.Sprintf("Cubic %s", Millimetre.Singular), Plural: fmt.Sprintf("Cubic %s", Millimetre.Plural), Abbreviation: fmt.Sprintf("%s3", Millimetre.Abbreviation), Ratio: math.Pow(Millimetre.Ratio, 3)}
	CubicMicrometre = Unit{Singular: fmt.Sprintf("Cubic %s", Micrometre.Singular), Plural: fmt.Sprintf("Cubic %s", Micrometre.Plural), Abbreviation: fmt.Sprintf("%s3", Micrometre.Abbreviation), Ratio: math.Pow(Micrometre.Ratio, 3)}
	CubicNanometre  = Unit{Singular: fmt.Sprintf("Cubic %s", Nanometre.Singular), Plural: fmt.Sprintf("Cubic %s", Nanometre.Plural), Abbreviation: fmt.Sprintf("%s3", Nanometre.Abbreviation), Ratio: math.Pow(Nanometre.Ratio, 3)}
	CubicPicometre  = Unit{Singular: fmt.Sprintf("Cubic %s", Picometre.Singular), Plural: fmt.Sprintf("Cubic %s", Picometre.Plural), Abbreviation: fmt.Sprintf("%s3", Picometre.Abbreviation), Ratio: math.Pow(Picometre.Ratio, 3)}
	CubicFemtometre = Unit{Singular: fmt.Sprintf("Cubic %s", Femtometre.Singular), Plural: fmt.Sprintf("Cubic %s", Femtometre.Plural), Abbreviation: fmt.Sprintf("%s3", Femtometre.Abbreviation), Ratio: math.Pow(Femtometre.Ratio, 3)}
	CubicAttometre  = Unit{Singular: fmt.Sprintf("Cubic %s", Attometre.Singular), Plural: fmt.Sprintf("Cubic %s", Attometre.Plural), Abbreviation: fmt.Sprintf("%s3", Attometre.Abbreviation), Ratio: math.Pow(Attometre.Ratio, 3)}

	// Imperial
	CubicFoot = Unit{Singular: fmt.Sprintf("Cubic %s", Foot.Singular), Plural: fmt.Sprintf("Cubic %s", Foot.Plural), Abbreviation: fmt.Sprintf("cu %s", Foot.Abbreviation), Ratio: math.Pow(Foot.Ratio, 3)}
	CubicInch = Unit{Singular: fmt.Sprintf("Cubic %s", Inch.Singular), Plural: fmt.Sprintf("Cubic %s", Inch.Plural), Abbreviation: fmt.Sprintf("cu %s", Inch.Abbreviation), Ratio: math.Pow(Inch.Ratio, 3)}
	CubicYard = Unit{Singular: fmt.Sprintf("Cubic %s", Yard.Singular), Plural: fmt.Sprintf("Cubic %s", Yard.Plural), Abbreviation: fmt.Sprintf("cu %s", Yard.Abbreviation), Ratio: math.Pow(Yard.Ratio, 3)}
	CubicMile = Unit{Singular: fmt.Sprintf("Cubic %s", Mile.Singular), Plural: fmt.Sprintf("Cubic %s", Mile.Plural), Abbreviation: fmt.Sprintf("cu %s", Mile.Abbreviation), Ratio: math.Pow(Mile.Ratio, 3)}

	// Land measurements
	CubicFurlong      = Unit{Singular: fmt.Sprintf("Cubic %s", Furlong.Singular), Plural: fmt.Sprintf("Cubic %s", Furlong.Plural), Abbreviation: fmt.Sprintf("cu %s", Furlong.Abbreviation), Ratio: math.Pow(Furlong.Ratio, 3)}
	CubicChain        = Unit{Singular: fmt.Sprintf("Cubic %s", Chain.Singular), Plural: fmt.Sprintf("Cubic %s", Chain.Plural), Abbreviation: fmt.Sprintf("cu %s", Chain.Abbreviation), Ratio: math.Pow(Chain.Ratio, 3)}
	CubicLink         = Unit{Singular: fmt.Sprintf("Cubic %s", Link.Singular), Plural: fmt.Sprintf("Cubic %s", Link.Plural), Abbreviation: fmt.Sprintf("cu %s", Link.Abbreviation), Ratio: math.Pow(Link.Ratio, 3)}
	CubicRod          = Unit{Singular: fmt.Sprintf("Cubic %s", Rod.Singular), Plural: fmt.Sprintf("Cubic %s", Rod.Plural), Abbreviation: fmt.Sprintf("cu %s", Rod.Abbreviation), Ratio: math.Pow(Rod.Ratio, 3)}
	CubicRamsdenChain = Unit{Singular: fmt.Sprintf("Cubic %s", RamsdenChain.Singular), Plural: fmt.Sprintf("Cubic %s", RamsdenChain.Plural), Abbreviation: fmt.Sprintf("cu %s", RamsdenChain.Abbreviation), Ratio: math.Pow(RamsdenChain.Ratio, 3)}

	// Nautical units
	CubicNauticalMile = Unit{Singular: fmt.Sprintf("Cubic %s", NauticalMile.Singular), Plural: fmt.Sprintf("Cubic %s", NauticalMile.Plural), Abbreviation: fmt.Sprintf("cu %s", NauticalMile.Abbreviation), Ratio: math.Pow(NauticalMile.Ratio, 3)}
	CubicCable        = Unit{Singular: fmt.Sprintf("Cubic %s", Cable.Singular), Plural: fmt.Sprintf("Cubic %s", Cable.Plural), Abbreviation: fmt.Sprintf("cu %s", Cable.Abbreviation), Ratio: math.Pow(Cable.Ratio, 3)}
	CubicFathom       = Unit{Singular: fmt.Sprintf("Cubic %s", Fathom.Singular), Plural: fmt.Sprintf("Cubic %s", Fathom.Plural), Abbreviation: fmt.Sprintf("cu %s", Fathom.Abbreviation), Ratio: math.Pow(Fathom.Ratio, 3)}
	CubicShackle      = Unit{Singular: fmt.Sprintf("Cubic %s", Shackle.Singular), Plural: fmt.Sprintf("Cubic %s", Shackle.Plural), Abbreviation: fmt.Sprintf("cu %s", Shackle.Abbreviation), Ratio: math.Pow(Shackle.Ratio, 3)}

	// Unofficial
	CubicLeague = Unit{Singular: fmt.Sprintf("Cubic %s", League.Singular), Plural: fmt.Sprintf("Cubic %s", League.Plural), Abbreviation: fmt.Sprintf("cu %s", League.Abbreviation), Ratio: math.Pow(League.Ratio, 3)}

	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO

	// Imperial
	ImperialGallon     = Unit{Singular: "Imperial Gallon", Plural: "Imperial Gallons", Abbreviation: "imp gal", Ratio: 1}
	ImperialPint       = Unit{Singular: "Imperial Pint", Plural: "Imperial Pints", Abbreviation: "imp pt", Ratio: ImperialGallon.Ratio / 8}
	ImperialQuart      = Unit{Singular: "Imperial Quart", Plural: "Imperial Quarts", Abbreviation: "imp qt", Ratio: ImperialGallon.Ratio / 4}
	ImperialFluidOunce = Unit{Singular: "Imperial Fluid Ounce", Plural: "Imperial Fluid Ounces", Abbreviation: "imp fl oz", Ratio: ImperialGallon.Ratio / 160}

	// Metric
	MetricCup = Unit{Singular: "Metric Cup", Plural: "Metric Cups", Abbreviation: "C"} // 250 milliliters

	// US Liquid
	USGallon     = Unit{Singular: "US Gallon", Plural: "US Gallons", Abbreviation: "US gal", Ratio: 1}
	USPint       = Unit{Singular: "US Pint", Plural: "US Pints", Abbreviation: "US pt", Ratio: USGallon.Ratio / 8}
	USQuart      = Unit{Singular: "US Quart", Plural: "US Quarts", Abbreviation: "US qt", Ratio: USGallon.Ratio / 4}
	USCup        = Unit{Singular: "US Cup", Plural: "US Cups", Abbreviation: "US C", Ratio: USPint.Ratio / 2}
	USFluidOunce = Unit{Singular: "US Fluid Ounce", Plural: "US Fluid Ounces", Abbreviation: "US fl oz", Ratio: USGallon.Ratio / 128}

	// US Dry
	USDryGallon = Unit{Singular: "US Dry Gallon", Plural: "US Dry Gallons", Abbreviation: "usdrygal", Ratio: 1}
	USDryPint   = Unit{Singular: "US Dry Pint", Plural: "US Dry Pints", Abbreviation: "usdrypt", Ratio: USDryGallon.Ratio / 8}
	USDryQuart  = Unit{Singular: "US Dry Quart", Plural: "US Dry Quarts", Abbreviation: "usdryqt", Ratio: USDryGallon.Ratio / 4}
	USBushel    = Unit{Singular: "US Bushel", Plural: "US Bushels", Abbreviation: "bsh", Ratio: BushelToUSDryGallonRatio}

	// Other
	// https://en.wikipedia.org/wiki/Cooking_weights_and_measures
	// Gill = Unit{Singular: "Gill", Plural: "Gills", Abbreviation: "gi"}
	// Barrel     = Unit{Singular: "Barrel", Plural: "Barrels", Abbreviation: "bbl"}
	// Cup        = Unit{Singular: "Cup", Plural: "Cups", Abbreviation: "C"}
	// Teaspoon   = Unit{Singular: "Teaspoon", Plural: "Teaspoons", Abbreviation: "tsp"}
	// Dessertspoon   = Unit{Singular: "Dessertspoon", Plural: "Dessertspoons", Abbreviation: "tsp"}
	// Tablespoon = Unit{Singular: "Tablespoon", Plural: "Tablespoons", Abbreviation: "tbsp"}
	// Vat        = Unit{Singular: "Vat", Abbreviation: "vat"}
	// Cask       = Unit{Singular: "Cask", Abbreviation: "cask"}
	// Keg        = Unit{Singular: "Keg", Abbreviation: "keg"}
)

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

// USDryGallonToBushels converts a given value from US Dry Gallons to its equivalent in Bushels
func USDryGallonsToBushels(value float64) float64 {
	return value * BushelToUSDryGallonRatio
}

// BushelsToUSDryGallon converts a given value from Bushels to its equivalent in US Dry Gallons
func BushelsToUSDryGallons(value float64) float64 {
	return value * BushelToUSDryGallonRatio
}

type Volume struct {
	Measurement Measurement
}

// Litre converts a volume to its equivalent in Litres
func (v *Volume) Litre() (error, float64) {
	switch v.Measurement.Unit {
	case Litre:
		// Measurement is already in Litres
	case Gigalitre, Megalitre, Kilolitre, Hectolitre, Decalitre, Decilitre, Centilitre, Millilitre, Microlitre, Nanolitre, Picolitre:
		v.Measurement.Value = v.Measurement.Unit.OneToRatio(v.Measurement.Value)
	case CubicMetre, CubicGigametre, CubicMegametre, CubicKilometre, CubicHectometre, CubicDecametre, CubicDecimetre, CubicCentimetre, CubicMillimetre, CubicMicrometre, CubicNanometre, CubicPicometre:
		v.Measurement.Value = CubicMetreToLitre(v.Measurement.Unit.RatioToOne(v.Measurement.Value))
	case CubicFoot, CubicInch, CubicYard, CubicMile, CubicFurlong, CubicChain, CubicLink, CubicRod, CubicRamsdenChain, CubicNauticalMile, CubicCable, CubicFathom, CubicShackle, CubicLeague:
		cubicFeet := v.Measurement.Unit.RatioToOne(v.Measurement.Value)
		cubicMetres := CubicFeetToCubicMetres(cubicFeet)
		v.Measurement.Value = CubicMetreToLitre(cubicMetres)
	default:
		err := fmt.Errorf("Unknown volume unit, %v", v.Measurement.Unit)
		return err, 0
	}

	// Update the measurement unit to reflect the new Unit
	v.Measurement.Unit = Litre

	return nil, v.Measurement.Value
}

// ConvertFromLitre is a helper to convert a value from its value in Litres
func (v *Volume) ConvertFromLitre(handle func(float64) float64) (error, float64) {
	err, litre := v.Litre()
	if err != nil {
		return err, 0
	}

	return nil, handle(litre)
}

// Exalitre converts a volume to its equivalent in Exalitres
func (v *Volume) Exalitre() (error, float64) {
	return v.ConvertFromLitre(Exalitre.OneToRatio)
}

// Petalitre converts a volume to its equivalent in Petalitres
func (v *Volume) Petalitre() (error, float64) {
	return v.ConvertFromLitre(Petalitre.OneToRatio)
}

// Teralitre converts a volume to its equivalent in Teralitres
func (v *Volume) Teralitre() (error, float64) {
	return v.ConvertFromLitre(Teralitre.OneToRatio)
}

// Gigalitre converts a volume to its equivalent in Gigalitres
func (v *Volume) Gigalitre() (error, float64) {
	return v.ConvertFromLitre(Gigalitre.OneToRatio)
}

// Megalitre converts a volume to its equivalent in Megalitres
func (v *Volume) Megalitre() (error, float64) {
	return v.ConvertFromLitre(Megalitre.OneToRatio)
}

// Kilolitre converts a volume to its equivalent in Kilolitres
func (v *Volume) Kilolitre() (error, float64) {
	return v.ConvertFromLitre(Kilolitre.OneToRatio)
}

// Hectolitre converts a volume to its equivalent in Hectolitres
func (v *Volume) Hectolitre() (error, float64) {
	return v.ConvertFromLitre(Hectolitre.OneToRatio)
}

// Decalitre converts a volume to its equivalent in Decalitres
func (v *Volume) Decalitre() (error, float64) {
	return v.ConvertFromLitre(Decalitre.OneToRatio)
}

// Decilitre converts a volume to its equivalent in Decilitres
func (v *Volume) Decilitre() (error, float64) {
	return v.ConvertFromLitre(Decilitre.OneToRatio)
}

// Centilitre converts a volume to its equivalent in Centilitres
func (v *Volume) Centilitre() (error, float64) {
	return v.ConvertFromLitre(Centilitre.OneToRatio)
}

// Millilitre converts a volume to its equivalent in Millilitres
func (v *Volume) Millilitre() (error, float64) {
	return v.ConvertFromLitre(Millilitre.OneToRatio)
}

// Microlitre converts a volume to its equivalent in Microlitres
func (v *Volume) Microlitre() (error, float64) {
	return v.ConvertFromLitre(Microlitre.OneToRatio)
}

// Nanolitre converts a volume to its equivalent in Nanolitres
func (v *Volume) Nanolitre() (error, float64) {
	return v.ConvertFromLitre(Nanolitre.OneToRatio)
}

// Picolitre converts a volume to its equivalent in Picolitres
func (v *Volume) Picolitre() (error, float64) {
	return v.ConvertFromLitre(Picolitre.OneToRatio)
}

// Femtolitre converts a volume to its equivalent in Femtolitres
func (v *Volume) Femtolitre() (error, float64) {
	return v.ConvertFromLitre(Femtolitre.OneToRatio)
}

// Attolitre converts a volume to its equivalent in Attolitres
func (v *Volume) Attolitre() (error, float64) {
	return v.ConvertFromLitre(Attolitre.OneToRatio)
}

// ConvertFromLitre is a helper to convert a value to cubic metres
func (v *Volume) ConvertToCubicMetre(unit Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		return unit.OneToRatio(LitreToCubicMetre(value))
	})
}

// CubicMetre converts a volume to its equivalent in Cubic Meters
func (v *Volume) CubicMetre() (error, float64) {
	return v.ConvertToCubicMetre(CubicMetre)
}

// CubicExametre converts a volume to its equivalent in CubicExametres
func (v *Volume) CubicExametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicExametre)
}

// CubicPetametre converts a volume to its equivalent in CubicPetametres
func (v *Volume) CubicPetametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicPetametre)
}

// CubicTerametre converts a volume to its equivalent in CubicTerametres
func (v *Volume) CubicTerametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicTerametre)
}

// CubicGigametre converts a volume to its equivalent in CubicGigametres
func (v *Volume) CubicGigametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicGigametre)
}

// CubicMegametre converts a volume to its equivalent in CubicMegametres
func (v *Volume) CubicMegametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicMegametre)
}

// CubicKilometre converts a volume to its equivalent in CubicKilometres
func (v *Volume) CubicKilometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicKilometre)
}

// CubicHectometre converts a volume to its equivalent in CubicHectometres
func (v *Volume) CubicHectometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicHectometre)
}

// CubicDecametre converts a volume to its equivalent in CubicDecametres
func (v *Volume) CubicDecametre() (error, float64) {
	return v.ConvertToCubicMetre(CubicDecametre)
}

// CubicDecimetre converts a volume to its equivalent in CubicDecimetres
func (v *Volume) CubicDecimetre() (error, float64) {
	return v.ConvertToCubicMetre(CubicDecimetre)
}

// CubicCentimetre converts a volume to its equivalent in CubicCentimetres
func (v *Volume) CubicCentimetre() (error, float64) {
	return v.ConvertToCubicMetre(CubicCentimetre)
}

// CubicMillimetre converts a volume to its equivalent in CubicMillimetres
func (v *Volume) CubicMillimetre() (error, float64) {
	return v.ConvertToCubicMetre(CubicMillimetre)
}

// CubicMicrometre converts a volume to its equivalent in CubicMicrometres
func (v *Volume) CubicMicrometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicMicrometre)
}

// CubicNanometre converts a volume to its equivalent in CubicNanometres
func (v *Volume) CubicNanometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicNanometre)
}

// CubicPicometre converts a volume to its equivalent in CubicPicometres
func (v *Volume) CubicPicometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicPicometre)
}

// CubicFemtometre converts a volume to its equivalent in CubicFemtometres
func (v *Volume) CubicFemtometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicFemtometre)
}

// CubicAttometre converts a volume to its equivalent in CubicAttometres
func (v *Volume) CubicAttometre() (error, float64) {
	return v.ConvertToCubicMetre(CubicAttometre)
}

// ConvertToCubicFeet is a helper to convert a volume to cubic feet
func (v *Volume) ConvertToCubicFeet(unit Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		cubicMetres := LitreToCubicMetre(value)
		cubicFeet := CubicMetresToCubicFeet(cubicMetres)
		return unit.OneToRatio(cubicFeet)
	})
}

// CubicFoot converts a volume to its equivalent value in Cubic Feet
func (v *Volume) CubicFoot() (error, float64) {
	return v.ConvertToCubicFeet(CubicFoot)
}

// CubicInch converts a volume to its equivalent value in Cubic Inches
func (v *Volume) CubicInch() (error, float64) {
	return v.ConvertToCubicFeet(CubicInch)
}

// CubicYard converts a volume to its equivalent value in Cubic Yards
func (v *Volume) CubicYard() (error, float64) {
	return v.ConvertToCubicFeet(CubicYard)
}

// CubicMile converts a volume to its equivalent value in Cubic Miles
func (v *Volume) CubicMile() (error, float64) {
	return v.ConvertToCubicFeet(CubicMile)
}

// CubicFurlong converts a volume to its equivalent value in Cubic Furlongs
func (v *Volume) CubicFurlong() (error, float64) {
	return v.ConvertToCubicFeet(CubicFurlong)
}

// CubicChain converts a volume to its equivalent value in Cubic Chains
func (v *Volume) CubicChain() (error, float64) {
	return v.ConvertToCubicFeet(CubicChain)
}

// CubicLink converts a volume to its equivalent value in Cubic Links
func (v *Volume) CubicLink() (error, float64) {
	return v.ConvertToCubicFeet(CubicLink)
}

// CubicRod converts a volume to its equivalent value in Cubic Rods
func (v *Volume) CubicRod() (error, float64) {
	return v.ConvertToCubicFeet(CubicRod)
}

// CubicRamsdenChain converts a volume to its equivalent value in Cubic Ramsden Chains
func (v *Volume) CubicRamsdenChain() (error, float64) {
	return v.ConvertToCubicFeet(CubicRamsdenChain)
}

// CubicNauticalMile converts a volume to its equivalent value in Cubic Nautical Miles
func (v *Volume) CubicNauticalMile() (error, float64) {
	return v.ConvertToCubicFeet(CubicNauticalMile)
}

// CubicCable converts a volume to its equivalent value in Cubic Cables
func (v *Volume) CubicCable() (error, float64) {
	return v.ConvertToCubicFeet(CubicCable)
}

// CubicFathom converts a volume to its equivalent value in Cubic Fathoms
func (v *Volume) CubicFathom() (error, float64) {
	return v.ConvertToCubicFeet(CubicFathom)
}

// CubicShackle converts a volume to its equivalent value in Cubic Shackles
func (v *Volume) CubicShackle() (error, float64) {
	return v.ConvertToCubicFeet(CubicShackle)
}

// CubicLeague converts a volume to its equivalent value in Cubic Leagues
func (v *Volume) CubicLeague() (error, float64) {
	return v.ConvertToCubicFeet(CubicLeague)
}
