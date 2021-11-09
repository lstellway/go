package mass

import "github.com/lstellway/go/measure"

var (
	ApothecaryScruple = measure.Unit{Singular: "Apothecary Scruple", Plural: "Apothecary Scruples", Abbreviation: "℈", Ratio: Grain.Ratio * 20}
	ApothecaryDrachm  = measure.Unit{Singular: "Apothecary Drachm", Plural: "Apothecary Drachms", Abbreviation: "ʒ", Ratio: ApothecaryScruple.Ratio * 3}
	ApothecaryOunce   = measure.Unit{Singular: "Apothecary Ounce", Plural: "Apothecary Ounces", Abbreviation: "℥", Ratio: Grain.Ratio * 480}
	ApothecaryPound   = measure.Unit{Singular: "Apothecary Pound", Plural: "Apothecary Pounds", Abbreviation: "lb", Ratio: Grain.Ratio * 480}
)

// ApothecaryScruple converts a mass to its equivalent in Apothecary Scruples
func (m *Mass) ApothecaryScruple() (error, float64) {
	return m.ConvertToUnitFromOunces(ApothecaryScruple)
}

// ApothecaryDrachm converts a mass to its equivalent in Apothecary Drachms
func (m *Mass) ApothecaryDrachm() (error, float64) {
	return m.ConvertToUnitFromOunces(ApothecaryDrachm)
}

// ApothecaryOunce converts a mass to its equivalent in Apothecary Ounces
func (m *Mass) ApothecaryOunce() (error, float64) {
	return m.ConvertToUnitFromOunces(ApothecaryOunce)
}

// ApothecaryPound converts a mass to its equivalent in Apothecary Pounds
func (m *Mass) ApothecaryPound() (error, float64) {
	return m.ConvertToUnitFromOunces(ApothecaryPound)
}
