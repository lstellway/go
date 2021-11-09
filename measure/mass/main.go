package mass

import "github.com/lstellway/go/measure"

var (
	// Ratios
	OunceToGramRatio      = 28.3495
	GrainToMilligramRatio = 64.79891
)

type Mass struct {
	Measurement measure.Measurement
}

// OunceToGram converts a value provided in ounces to grams
func OunceToGram(value float64) float64 {
	return value * OunceToGramRatio
}

// OunceToGram converts a value provided in grams to ounces
func GramToOunce(value float64) float64 {
	return value / OunceToGramRatio
}

// GrainToMilligram converts a value provided in grains to milligrams
func GrainToMilligram(value float64) float64 {
	return value * GrainToMilligramRatio
}

// MilligramToGrain converts a value provided in milligrams to grains
func MilligramToGrain(value float64) float64 {
	return value / GrainToMilligramRatio
}
