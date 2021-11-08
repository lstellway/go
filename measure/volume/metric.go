package volume

import "github.com/lstellway/go/measure"

var (
	// Metric
	MetricCup          = measure.Unit{Singular: "Metric Cup", Plural: "Metric Cups", Abbreviation: "C", Ratio: Millilitre.Ratio * 250}
	MetricTeaspoon     = measure.Unit{Singular: "Metric Teaspoon", Plural: "Metric Teaspoons", Abbreviation: "tsp", Ratio: Millilitre.Ratio * 5}
	MetricDessertspoon = measure.Unit{Singular: "Metric Dessertspoon", Plural: "Metric Dessertspoons", Abbreviation: "dstspn", Ratio: Millilitre.Ratio * 10}
	MetricTablespoon   = measure.Unit{Singular: "Metric Tablespoon", Plural: "Metric Tablespoons", Abbreviation: "tbsp", Ratio: Millilitre.Ratio * 15}
	// Theoretical (non-existent)
	MetricPint   = measure.Unit{Singular: "Metric Pint", Plural: "Metric Pints", Abbreviation: "metric pt", Ratio: MetricCup.Ratio * 2}
	MetricQuart  = measure.Unit{Singular: "Metric Quart", Plural: "Metric Quarts", Abbreviation: "metric qt", Ratio: MetricPint.Ratio * 2}
	MetricGallon = measure.Unit{Singular: "Metric Gallon", Plural: "Metric Gallons", Abbreviation: "metric gal", Ratio: MetricQuart.Ratio * 4}
)

// MetricCup converts a volume to its equivalent in Metric Cups
func (v *Volume) MetricCup() (error, float64) {
	return v.ConvertFromLitre(MetricCup.OneToRatio)
}

// MetricTeaspoon converts a volume to its equivalent in Metric Teaspoons
func (v *Volume) MetricTeaspoon() (error, float64) {
	return v.ConvertFromLitre(MetricTeaspoon.OneToRatio)
}

// MetricDessertspoon converts a volume to its equivalent in Metric Dessertspoons
func (v *Volume) MetricDessertspoon() (error, float64) {
	return v.ConvertFromLitre(MetricDessertspoon.OneToRatio)
}

// MetricTablespoon converts a volume to its equivalent in Metric Tablespoons
func (v *Volume) MetricTablespoon() (error, float64) {
	return v.ConvertFromLitre(MetricTablespoon.OneToRatio)
}

// MetricPint converts a volume to its equivalent in a non-existent unit of Metric Pints
func (v *Volume) MetricPint() (error, float64) {
	return v.ConvertFromLitre(MetricPint.OneToRatio)
}

// MetricQuart converts a volume to its equivalent in a non-existent unit of Metric Quarts
func (v *Volume) MetricQuart() (error, float64) {
	return v.ConvertFromLitre(MetricQuart.OneToRatio)
}

// MetricGallon converts a volume to its equivalent in a non-existent unit of Metric Gallons
func (v *Volume) MetricGallon() (error, float64) {
	return v.ConvertFromLitre(MetricGallon.OneToRatio)
}
