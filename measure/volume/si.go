package volume

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Litres
	Litre      = measure.Unit{Singular: "Litre", Plural: "Litres", Abbreviation: "L", Ratio: 1} // Base 1
	Yottalitre = measure.Yotta.Unit("%slitre", "%sL")
	Zettalitre = measure.Zetta.Unit("%slitre", "%sL")
	Exalitre   = measure.Exa.Unit("%slitre", "%sL")
	Petalitre  = measure.Peta.Unit("%slitre", "%sL")
	Teralitre  = measure.Tera.Unit("%slitre", "%sL")
	Gigalitre  = measure.Giga.Unit("%slitre", "%sL")
	Megalitre  = measure.Mega.Unit("%slitre", "%sL")
	Kilolitre  = measure.Kilo.Unit("%slitre", "%sL")
	Hectolitre = measure.Hecto.Unit("%slitre", "%sL")
	Decalitre  = measure.Deca.Unit("%slitre", "%sL")
	Decilitre  = measure.Deci.Unit("%slitre", "%sL")
	Centilitre = measure.Centi.Unit("%slitre", "%sL")
	Millilitre = measure.Milli.Unit("%slitre", "%sL")
	Microlitre = measure.Micro.Unit("%slitre", "%sL")
	Nanolitre  = measure.Nano.Unit("%slitre", "%sL")
	Picolitre  = measure.Pico.Unit("%slitre", "%sL")
	Femtolitre = measure.Femto.Unit("%slitre", "%sL")
	Attolitre  = measure.Atto.Unit("%slitre", "%sL")
	Zeptolitre = measure.Zepto.Unit("%slitre", "%sL")
	Yoctolitre = measure.Yocto.Unit("%slitre", "%sL")

	// Cubic Metres
	CubicMetre      = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Metre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Metre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Metre.Abbreviation), Ratio: math.Pow(length.Metre.Ratio, 3)}
	CubicYottametre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Yottametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Yottametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Yottametre.Abbreviation), Ratio: math.Pow(length.Yottametre.Ratio, 3)}
	CubicZettametre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Zettametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Zettametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Zettametre.Abbreviation), Ratio: math.Pow(length.Zettametre.Ratio, 3)}
	CubicExametre   = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Exametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Exametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Exametre.Abbreviation), Ratio: math.Pow(length.Exametre.Ratio, 3)}
	CubicPetametre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Petametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Petametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Petametre.Abbreviation), Ratio: math.Pow(length.Petametre.Ratio, 3)}
	CubicTerametre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Terametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Terametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Terametre.Abbreviation), Ratio: math.Pow(length.Terametre.Ratio, 3)}
	CubicGigametre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Gigametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Gigametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Gigametre.Abbreviation), Ratio: math.Pow(length.Gigametre.Ratio, 3)}
	CubicMegametre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Megametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Megametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Megametre.Abbreviation), Ratio: math.Pow(length.Megametre.Ratio, 3)}
	CubicKilometre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Kilometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Kilometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Kilometre.Abbreviation), Ratio: math.Pow(length.Kilometre.Ratio, 3)}
	CubicHectometre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Hectometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Hectometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Hectometre.Abbreviation), Ratio: math.Pow(length.Hectometre.Ratio, 3)}
	CubicDecametre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Decametre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Decametre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Decametre.Abbreviation), Ratio: math.Pow(length.Decametre.Ratio, 3)}
	CubicDecimetre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Decimetre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Decimetre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Decimetre.Abbreviation), Ratio: math.Pow(length.Decimetre.Ratio, 3)}
	CubicCentimetre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Centimetre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Centimetre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Centimetre.Abbreviation), Ratio: math.Pow(length.Centimetre.Ratio, 3)}
	CubicMillimetre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Millimetre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Millimetre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Millimetre.Abbreviation), Ratio: math.Pow(length.Millimetre.Ratio, 3)}
	CubicMicrometre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Micrometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Micrometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Micrometre.Abbreviation), Ratio: math.Pow(length.Micrometre.Ratio, 3)}
	CubicNanometre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Nanometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Nanometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Nanometre.Abbreviation), Ratio: math.Pow(length.Nanometre.Ratio, 3)}
	CubicPicometre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Picometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Picometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Picometre.Abbreviation), Ratio: math.Pow(length.Picometre.Ratio, 3)}
	CubicFemtometre = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Femtometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Femtometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Femtometre.Abbreviation), Ratio: math.Pow(length.Femtometre.Ratio, 3)}
	CubicAttometre  = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Attometre.Singular), Plural: fmt.Sprintf("Cubic %s", length.Attometre.Plural), Abbreviation: fmt.Sprintf("%s3", length.Attometre.Abbreviation), Ratio: math.Pow(length.Attometre.Ratio, 3)}
)

// Litre converts a volume to its equivalent in Litres
func (v *Volume) Litre() (error, float64) {
	switch v.Measurement.Unit {
	// Litre - International System of Units (SI)
	case Litre, Yottalitre, Zettalitre, Exalitre, Petalitre, Teralitre, Gigalitre, Megalitre, Kilolitre, Hectolitre, Decalitre, Decilitre, Centilitre, Millilitre, Microlitre, Nanolitre, Picolitre, Femtolitre, Attolitre, Zeptolitre, Yoctolitre:
		v.Measurement.Value = v.Measurement.Unit.RatioToOne(v.Measurement.Value)
	// Cubic Metre - International System of Units (SI)
	case CubicMetre, CubicYottametre, CubicZettametre, CubicExametre, CubicPetametre, CubicTerametre, CubicGigametre, CubicMegametre, CubicKilometre, CubicHectometre, CubicDecametre, CubicDecimetre, CubicCentimetre, CubicMillimetre, CubicMicrometre, CubicNanometre, CubicPicometre, CubicFemtometre, CubicAttometre:
		v.Measurement.Value = CubicMetreToLitre(v.Measurement.Unit.RatioToOne(v.Measurement.Value))
	// Metric (based on Litre)
	case MetricCup, MetricTeaspoon, MetricDessertspoon, MetricTablespoon, MetricPint, MetricQuart, MetricGallon:
		v.Measurement.Value = v.Measurement.Unit.RatioToOne(v.Measurement.Value)
	// Imperial
	case CubicFoot, CubicInch, CubicYard, CubicMile:
		v.Measurement.Value = v.ImperialToLitre()
	// Land
	case CubicFurlong, CubicChain, CubicLink, CubicRod, CubicRamsdenChain:
		v.Measurement.Value = v.ImperialToLitre()
	// Nautical
	case CubicNauticalMile, CubicCable, CubicFathom, CubicShackle, CubicLeague:
		v.Measurement.Value = v.ImperialToLitre()
	// U.S. Liquid
	case USGallon, USQuart, USPint, USGill, USJack, USCup, USFluidOunce, USFluidDram, USMinim, USTeaspoon, USDessertspoon, USTablespoon, USOilBarrel, USBeerBarrel, USKeg:
		v.Measurement.Value = v.USGallonToLitre()
	// U.S. Dry
	case USDryGallon, USPeck, USBushel, USDryQuart, USDryPint, USDryBarrel:
		usDryGallons := v.Measurement.Unit.RatioToOne(v.Measurement.Value)
		v.Measurement.Value = USDryGallonsToLitres(usDryGallons)
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

// ImperialToLitre converts an imperial volume to its equivalent in Litres
func (v *Volume) ImperialToLitre() float64 {
	cubicFeet := v.Measurement.Unit.RatioToOne(v.Measurement.Value)
	cubicMetres := CubicFeetToCubicMetres(cubicFeet)
	return CubicMetreToLitre(cubicMetres)
}

// USGallonToLitre converts a given volume expressed in US Gallons to its equivalent in Litres
func (v *Volume) USGallonToLitre() float64 {
	usGallons := v.Measurement.Unit.RatioToOne(v.Measurement.Value)
	cubicInches := USGallonsToCubicInches(usGallons)
	cubicFeet := cubicInches * CubicFoot.RelativeRatio(CubicInch)
	cubicMetres := CubicFeetToCubicMetres(cubicFeet)
	return CubicMetreToLitre(cubicMetres)
}

// Yottalitre converts a volume to its equivalent in Yottalitres
func (v *Volume) Yottalitre() (error, float64) {
	return v.ConvertFromLitre(Yottalitre.OneToRatio)
}

// Zettalitre converts a volume to its equivalent in Zettalitres
func (v *Volume) Zettalitre() (error, float64) {
	return v.ConvertFromLitre(Zettalitre.OneToRatio)
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

// Zeptolitre converts a volume to its equivalent in Zeptolitres
func (v *Volume) Zeptolitre() (error, float64) {
	return v.ConvertFromLitre(Zeptolitre.OneToRatio)
}

// Yoctolitre converts a volume to its equivalent in Yoctolitres
func (v *Volume) Yoctolitre() (error, float64) {
	return v.ConvertFromLitre(Yoctolitre.OneToRatio)
}

// ConvertFromLitre is a helper to convert a value to cubic metres
func (v *Volume) ConvertToUnitFromCubicMetres(unit measure.Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		cubicMetres := LitreToCubicMetre(value)
		return unit.OneToRatio(cubicMetres)
	})
}

// CubicMetre converts a volume to its equivalent in Cubic Metres
func (v *Volume) CubicMetre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicMetre)
}

// CubicYottametre converts a volume to its equivalent in Cubic Yottametres
func (v *Volume) CubicYottametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicYottametre)
}

// CubicZettametre converts a volume to its equivalent in Cubic Zettametres
func (v *Volume) CubicZettametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicZettametre)
}

// CubicExametre converts a volume to its equivalent in Cubic Exametres
func (v *Volume) CubicExametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicExametre)
}

// CubicPetametre converts a volume to its equivalent in Cubic Petametres
func (v *Volume) CubicPetametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicPetametre)
}

// CubicTerametre converts a volume to its equivalent in Cubic Terametres
func (v *Volume) CubicTerametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicTerametre)
}

// CubicGigametre converts a volume to its equivalent in Cubic Gigametres
func (v *Volume) CubicGigametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicGigametre)
}

// CubicMegametre converts a volume to its equivalent in Cubic Megametres
func (v *Volume) CubicMegametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicMegametre)
}

// CubicKilometre converts a volume to its equivalent in Cubic Kilometres
func (v *Volume) CubicKilometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicKilometre)
}

// CubicHectometre converts a volume to its equivalent in Cubic Hectometres
func (v *Volume) CubicHectometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicHectometre)
}

// CubicDecametre converts a volume to its equivalent in Cubic Decametres
func (v *Volume) CubicDecametre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicDecametre)
}

// CubicDecimetre converts a volume to its equivalent in Cubic Decimetres
func (v *Volume) CubicDecimetre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicDecimetre)
}

// CubicCentimetre converts a volume to its equivalent in Cubic Centimetres
func (v *Volume) CubicCentimetre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicCentimetre)
}

// CubicMillimetre converts a volume to its equivalent in Cubic Millimetres
func (v *Volume) CubicMillimetre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicMillimetre)
}

// CubicMicrometre converts a volume to its equivalent in Cubic Micrometres
func (v *Volume) CubicMicrometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicMicrometre)
}

// CubicNanometre converts a volume to its equivalent in Cubic Nanometres
func (v *Volume) CubicNanometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicNanometre)
}

// CubicPicometre converts a volume to its equivalent in Cubic Picometres
func (v *Volume) CubicPicometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicPicometre)
}

// CubicFemtometre converts a volume to its equivalent in Cubic Femtometres
func (v *Volume) CubicFemtometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicFemtometre)
}

// CubicAttometre converts a volume to its equivalent in Cubic Attometres
func (v *Volume) CubicAttometre() (error, float64) {
	return v.ConvertToUnitFromCubicMetres(CubicAttometre)
}
