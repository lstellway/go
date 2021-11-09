package mass

import "github.com/lstellway/go/measure"

var (
	// Avoirdupois weights
	Ounce         = measure.Unit{Singular: "Ounce", Plural: "Ounces", Abbreviation: "oz", Ratio: 1}
	Dram          = measure.Unit{Singular: "Dram", Plural: "Drams", Abbreviation: "dr", Ratio: Ounce.Ratio / 16}
	Pound         = measure.Unit{Singular: "Pound", Plural: "Pounds", Abbreviation: "lb", Ratio: Ounce.Ratio * 16}
	Grain         = measure.Unit{Singular: "Grain", Plural: "Grains", Abbreviation: "gr", Ratio: Pound.Ratio / 7000}
	Stone         = measure.Unit{Singular: "Stone", Plural: "Stones", Abbreviation: "st", Ratio: Pound.Ratio * 14}
	Quarter       = measure.Unit{Singular: "Quarter", Plural: "Quarters", Abbreviation: "qtr", Ratio: Stone.Ratio * 2}
	Hundredweight = measure.Unit{Singular: "Hundredweight", Plural: "Hundredweights", Abbreviation: "cwt", Ratio: Quarter.Ratio * 4}

	ImperialTonne = measure.Unit{Singular: "Imperial Tonne", Plural: "Imperial Tonnes", Abbreviation: "imp t", Ratio: Hundredweight.Ratio * 20}
	USTon         = measure.Unit{Singular: "US Ton", Plural: "US Tons", Abbreviation: "US t", Ratio: Pound.Ratio * 2000}
)

// ConvertToUnitFromOunces is a helper to convert a mass to another unit using ounces
func (m *Mass) ConvertToUnitFromOunces(unit measure.Unit) (error, float64) {
	return m.ConvertFromGram(func(value float64) float64 {
		ounces := GramToOunce(value)
		return unit.OneToRatio(ounces)
	})
}

// Ounce converts a mass to its equivalent in Ounces
func (m *Mass) Ounce() (error, float64) {
	return m.ConvertToUnitFromOunces(Ounce)
}

// Dram converts a mass to its equivalent in Drams
func (m *Mass) Dram() (error, float64) {
	return m.ConvertToUnitFromOunces(Dram)
}

// Pound converts a mass to its equivalent in Pounds
func (m *Mass) Pound() (error, float64) {
	return m.ConvertToUnitFromOunces(Pound)
}

// Grain converts a mass to its equivalent in Grains
func (m *Mass) Grain() (error, float64) {
	return m.ConvertToUnitFromOunces(Grain)
}

// Stone converts a mass to its equivalent in Stones
func (m *Mass) Stone() (error, float64) {
	return m.ConvertToUnitFromOunces(Stone)
}

// Quarter converts a mass to its equivalent in Quarters
func (m *Mass) Quarter() (error, float64) {
	return m.ConvertToUnitFromOunces(Quarter)
}

// Hundredweight converts a mass to its equivalent in Hundredweights
func (m *Mass) Hundredweight() (error, float64) {
	return m.ConvertToUnitFromOunces(Hundredweight)
}

// ImperialTonne converts a mass to its equivalent in Imperial Tonnes
func (m *Mass) ImperialTonne() (error, float64) {
	return m.ConvertToUnitFromOunces(ImperialTonne)
}

// USTon converts a mass to its equivalent in US Tons
func (m *Mass) USTon() (error, float64) {
	return m.ConvertToUnitFromOunces(USTon)
}
