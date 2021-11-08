package volume

import (
	"fmt"
	"math"

	"github.com/lstellway/go/measure"
	"github.com/lstellway/go/measure/length"
)

var (
	// Imperial
	CubicFoot = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Foot.Singular), Plural: fmt.Sprintf("Cubic %s", length.Foot.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Foot.Abbreviation), Ratio: math.Pow(length.Foot.Ratio, 3)}
	CubicInch = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Inch.Singular), Plural: fmt.Sprintf("Cubic %s", length.Inch.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Inch.Abbreviation), Ratio: math.Pow(length.Inch.Ratio, 3)}
	CubicYard = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Yard.Singular), Plural: fmt.Sprintf("Cubic %s", length.Yard.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Yard.Abbreviation), Ratio: math.Pow(length.Yard.Ratio, 3)}
	CubicMile = measure.Unit{Singular: fmt.Sprintf("Cubic %s", length.Mile.Singular), Plural: fmt.Sprintf("Cubic %s", length.Mile.Plural), Abbreviation: fmt.Sprintf("cu %s", length.Mile.Abbreviation), Ratio: math.Pow(length.Mile.Ratio, 3)}

	ImperialGallon     = measure.Unit{Singular: "Imperial Gallon", Plural: "Imperial Gallons", Abbreviation: "imp gal", Ratio: 1}
	ImperialQuart      = measure.Unit{Singular: "Imperial Quart", Plural: "Imperial Quarts", Abbreviation: "imp qt", Ratio: ImperialGallon.Ratio / 4}
	ImperialPint       = measure.Unit{Singular: "Imperial Pint", Plural: "Imperial Pints", Abbreviation: "imp pt", Ratio: ImperialQuart.Ratio / 2}
	ImperialGill       = measure.Unit{Singular: "Imperial Gill", Plural: "Imperial Gills", Abbreviation: "imp gi", Ratio: ImperialPint.Ratio / 4}
	ImperialJack       = measure.Unit{Singular: "Imperial Jack", Plural: "Imperial Jacks", Abbreviation: "imp jack", Ratio: ImperialGill.Ratio / 2}
	ImperialCup        = measure.Unit{Singular: "Imperial Cup", Plural: "Imperial Cups", Abbreviation: "imp C", Ratio: ImperialPint.Ratio / 2}
	ImperialFluidOunce = measure.Unit{Singular: "Imperial Fluid Ounce", Plural: "Imperial Fluid Ounces", Abbreviation: "imp fl oz", Ratio: ImperialCup.Ratio / 10}
	ImperialFluidDram  = measure.Unit{Singular: "Imperial Fluid Dram", Plural: "Imperial Fluid Drams", Abbreviation: "imp fl dr", Ratio: ImperialFluidOunce.Ratio / 8}
	ImperialMinim      = measure.Unit{Singular: "Imperial Minim", Plural: "Imperial Minims", Abbreviation: "imp min", Ratio: USFluidDram.Ratio / 60}
	ImperialBeerBarrel = measure.Unit{Singular: "Imperial Beer Barrel", Plural: "Imperial Beer Barrels", Abbreviation: "imp beer bbl", Ratio: ImperialGallon.Ratio * 36}
)

// ConvertToUnitFromCubicFeet is a helper to convert a volume to another unit using cubic feet
func (v *Volume) ConvertToUnitFromCubicFeet(unit measure.Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		cubicMetres := LitreToCubicMetre(value)
		cubicFeet := CubicMetresToCubicFeet(cubicMetres)
		return unit.OneToRatio(cubicFeet)
	})
}

// CubicFoot converts a volume to its equivalent value in Cubic Feet
func (v *Volume) CubicFoot() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicFoot)
}

// CubicInch converts a volume to its equivalent value in Cubic Inches
func (v *Volume) CubicInch() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicInch)
}

// CubicYard converts a volume to its equivalent value in Cubic Yards
func (v *Volume) CubicYard() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicYard)
}

// CubicMile converts a volume to its equivalent value in Cubic Miles
func (v *Volume) CubicMile() (error, float64) {
	return v.ConvertToUnitFromCubicFeet(CubicMile)
}

// ConvertToCubicInches is a helper to convert a volume to another unit usinig cubic inches
func (v *Volume) ConvertToUnitFromImperialGallons(unit measure.Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		imperialGallons := LitresToImperialGallons(value)
		return unit.OneToRatio(imperialGallons)
	})
}

// ImperialGallon converts a volume to its equivalent value in Imperial Gallons
func (v *Volume) ImperialGallon() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialGallon)
}

// ImperialQuart converts a volume to its equivalent value in Imperial Quarts
func (v *Volume) ImperialQuart() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialQuart)
}

// ImperialPint converts a volume to its equivalent value in Imperial Pints
func (v *Volume) ImperialPint() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialPint)
}

// ImperialGill converts a volume to its equivalent value in Imperial Gills
func (v *Volume) ImperialGill() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialGill)
}

// ImperialJack converts a volume to its equivalent value in Imperial Jacks
func (v *Volume) ImperialJack() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialJack)
}

// ImperialCup converts a volume to its equivalent value in Imperial Cups
func (v *Volume) ImperialCup() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialCup)
}

// ImperialFluidOunce converts a volume to its equivalent value in Imperial Fluid Ounces
func (v *Volume) ImperialFluidOunce() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialFluidOunce)
}

// ImperialFluidDram converts a volume to its equivalent value in Imperial Fluid Drams
func (v *Volume) ImperialFluidDram() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialFluidDram)
}

// ImperialMinim converts a volume to its equivalent value in Imperial Minims
func (v *Volume) ImperialMinim() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialMinim)
}

// ImperialBeerBarrel converts a volume to its equivalent value in Imperial Beer Barrels
func (v *Volume) ImperialBeerBarrel() (error, float64) {
	return v.ConvertToUnitFromImperialGallons(ImperialBeerBarrel)
}
