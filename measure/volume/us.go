package volume

import "github.com/lstellway/go/measure"

var (
	// US Liquid
	USGallon       = measure.Unit{Singular: "US Gallon", Plural: "US Gallons", Abbreviation: "US gal", Ratio: 1}
	USQuart        = measure.Unit{Singular: "US Quart", Plural: "US Quarts", Abbreviation: "US qt", Ratio: USGallon.Ratio / 4}
	USPint         = measure.Unit{Singular: "US Pint", Plural: "US Pints", Abbreviation: "US pt", Ratio: USQuart.Ratio / 2}
	USGill         = measure.Unit{Singular: "US Gill", Plural: "US Gills", Abbreviation: "US gi", Ratio: USPint.Ratio / 4}
	USJack         = measure.Unit{Singular: "US Jack", Plural: "US Jacks", Abbreviation: "US jack", Ratio: USGill.Ratio / 2}
	USCup          = measure.Unit{Singular: "US Cup", Plural: "US Cups", Abbreviation: "US C", Ratio: USPint.Ratio / 2}
	USFluidOunce   = measure.Unit{Singular: "US Fluid Ounce", Plural: "US Fluid Ounces", Abbreviation: "US fl oz", Ratio: USCup.Ratio / 8}
	USFluidDram    = measure.Unit{Singular: "US Fluid Dram", Plural: "US Fluid Drams", Abbreviation: "US fl dr", Ratio: USFluidOunce.Ratio / 8}
	USMinim        = measure.Unit{Singular: "US Minim", Plural: "US Minims", Abbreviation: "US min", Ratio: USFluidDram.Ratio / 60}
	USTeaspoon     = measure.Unit{Singular: "US Teaspoon", Plural: "US Teaspoons", Abbreviation: "US tsp", Ratio: USFluidOunce.Ratio / 6}
	USDessertspoon = measure.Unit{Singular: "US Dessertspoon", Plural: "US Dessertspoons", Abbreviation: "US dstspn", Ratio: USTeaspoon.Ratio * 2}
	USTablespoon   = measure.Unit{Singular: "US Tablespoon", Plural: "US Tablespoons", Abbreviation: "US tbsp", Ratio: USTeaspoon.Ratio * 3}
	USOilBarrel    = measure.Unit{Singular: "US Oil Barrel", Plural: "US Oil Barrels", Abbreviation: "US oil bbl", Ratio: USGallon.Ratio * 42}
	USBeerBarrel   = measure.Unit{Singular: "US Beer Barrel", Plural: "US Beer Barrels", Abbreviation: "US beer bbl", Ratio: USGallon.Ratio * 31}
	USKeg          = measure.Unit{Singular: "US Keg", Plural: "US Kegs", Abbreviation: "US keg", Ratio: USBeerBarrel.Ratio / 2}

	// US Dry
	USDryGallon = measure.Unit{Singular: "US Dry Gallon", Plural: "US Dry Gallons", Abbreviation: "usdrygal", Ratio: 1}
	USPeck      = measure.Unit{Singular: "US Peck", Plural: "US Pecks", Abbreviation: "bsh", Ratio: USDryGallon.Ratio * 2}
	USBushel    = measure.Unit{Singular: "US Bushel", Plural: "US Bushels", Abbreviation: "bsh", Ratio: USDryGallon.Ratio * 8}
	USDryQuart  = measure.Unit{Singular: "US Dry Quart", Plural: "US Dry Quarts", Abbreviation: "usdryqt", Ratio: USDryGallon.Ratio / 4}
	USDryPint   = measure.Unit{Singular: "US Dry Pint", Plural: "US Dry Pints", Abbreviation: "usdrypt", Ratio: USDryQuart.Ratio / 2}
	USDryBarrel = measure.Unit{Singular: "US Dry Barrel", Plural: "US Dry Barrels", Abbreviation: "usdrybbl", Ratio: USDryGallon.Ratio * 26.25}
)

// ConvertToCubicInches is a helper to convert a volume to another unit usinig cubic inches
func (v *Volume) ConvertToUnitFromUSGallons(unit measure.Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		cubicMetres := LitreToCubicMetre(value)
		cubicFeet := CubicMetresToCubicFeet(cubicMetres)
		cubicInches := cubicFeet * CubicInch.RelativeRatio(CubicFoot)
		usGallons := CubicInchesToUSGallons(cubicInches)
		return unit.OneToRatio(usGallons)
	})
}

// USGallon converts a volume to its equivalent value in US Gallons
func (v *Volume) USGallon() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USGallon)
}

// USQuart converts a volume to its equivalent value in US Quarts
func (v *Volume) USQuart() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USQuart)
}

// USPint converts a volume to its equivalent value in US Pints
func (v *Volume) USPint() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USPint)
}

// USGill converts a volume to its equivalent value in US Gills
func (v *Volume) USGill() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USGill)
}

// USJack converts a volume to its equivalent value in US Jacks
func (v *Volume) USJack() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USJack)
}

// USCup converts a volume to its equivalent value in US Cups
func (v *Volume) USCup() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USCup)
}

// USFluidOunce converts a volume to its equivalent value in US FluidOunces
func (v *Volume) USFluidOunce() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USFluidOunce)
}

// USFluidDram converts a volume to its equivalent value in US FluidDrams
func (v *Volume) USFluidDram() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USFluidDram)
}

// USMinim converts a volume to its equivalent value in US Minims
func (v *Volume) USMinim() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USMinim)
}

// USTeaspoon converts a volume to its equivalent value in US Teaspoons
func (v *Volume) USTeaspoon() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USTeaspoon)
}

// USDessertspoon converts a volume to its equivalent value in US Dessertspoons
func (v *Volume) USDessertspoon() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USDessertspoon)
}

// USTablespoon converts a volume to its equivalent value in US Tablespoons
func (v *Volume) USTablespoon() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USTablespoon)
}

// USOilBarrel converts a volume to its equivalent value in US OilBarrels
func (v *Volume) USOilBarrel() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USOilBarrel)
}

// USBeerBarrel converts a volume to its equivalent value in US BeerBarrels
func (v *Volume) USBeerBarrel() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USBeerBarrel)
}

// USKeg converts a volume to its equivalent value in US Kegs
func (v *Volume) USKeg() (error, float64) {
	return v.ConvertToUnitFromUSGallons(USKeg)
}

// ConvertToCubicInches is a helper to convert a volume to another unit usinig cubic inches
func (v *Volume) ConvertToUnitFromUSDryGallons(unit measure.Unit) (error, float64) {
	return v.ConvertFromLitre(func(value float64) float64 {
		dryGallons := LitresToUSDryGallons(value)
		return unit.OneToRatio(dryGallons)
	})
}

// USDryGallon converts a volume to its equivalent value in US Dry Gallons
func (v *Volume) USDryGallon() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USDryGallon)
}

// USPeck converts a volume to its equivalent value in US Pecks
func (v *Volume) USPeck() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USPeck)
}

// USBushel converts a volume to its equivalent value in US Bushels
func (v *Volume) USBushel() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USBushel)
}

// USDryQuart converts a volume to its equivalent value in US Dry Quarts
func (v *Volume) USDryQuart() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USDryQuart)
}

// USDryPint converts a volume to its equivalent value in US Dry Pints
func (v *Volume) USDryPint() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USDryPint)
}

// USDryBarrel converts a volume to its equivalent value in US Dry Barrels
func (v *Volume) USDryBarrel() (error, float64) {
	return v.ConvertToUnitFromUSDryGallons(USDryBarrel)
}
