package mass

import (
	"fmt"

	"github.com/lstellway/go/measure"
)

var (
	Gram      = measure.Unit{Singular: "Gram", Plural: "Grams", Abbreviation: "g", Ratio: 1}
	Yottagram = measure.Yotta.Unit("%sgram", "%sg")
	Zettagram = measure.Zetta.Unit("%sgram", "%sg")
	Exagram   = measure.Exa.Unit("%sgram", "%sg")
	Petagram  = measure.Peta.Unit("%sgram", "%sg")
	Teragram  = measure.Tera.Unit("%sgram", "%sg")
	Gigagram  = measure.Giga.Unit("%sgram", "%sg")
	Megagram  = measure.Mega.Unit("%sgram", "%sg")
	Kilogram  = measure.Kilo.Unit("%sgram", "%sg")
	Hectogram = measure.Hecto.Unit("%sgram", "%sg")
	Decagram  = measure.Deca.Unit("%sgram", "%sg")
	Decigram  = measure.Deci.Unit("%sgram", "%sg")
	Centigram = measure.Centi.Unit("%sgram", "%sg")
	Milligram = measure.Milli.Unit("%sgram", "%sg")
	Microgram = measure.Micro.Unit("%sgram", "%sg")
	Nanogram  = measure.Nano.Unit("%sgram", "%sg")
	Picogram  = measure.Pico.Unit("%sgram", "%sg")
	Femtogram = measure.Femto.Unit("%sgram", "%sg")
	Attogram  = measure.Atto.Unit("%sgram", "%sg")
	Zeptogram = measure.Zepto.Unit("%sgram", "%sg")
	Yoctogram = measure.Yocto.Unit("%sgram", "%sg")

	MetricTon = measure.Unit{Singular: "Metric Tonne", Plural: "Metric Tonnes", Abbreviation: "t", Ratio: Kilogram.Ratio * 1000}
)

func (m *Mass) Gram() (error, float64) {
	switch m.Measurement.Unit {
	// International System of Units (SI)
	case Gram, Yottagram, Zettagram, Exagram, Petagram, Teragram, Gigagram, Megagram, Kilogram, Hectogram, Decagram, Decigram, Centigram, Milligram, Microgram, Nanogram, Picogram, Femtogram, Attogram, Zeptogram, Yoctogram, MetricTon:
		m.Measurement.Value = m.Measurement.Unit.RatioToOne(m.Measurement.Value)
	// Imperial
	case Ounce, Dram, Pound, Grain, Stone, Quarter, Hundredweight, ImperialTonne, USTon:
		ounces := m.Measurement.Unit.RatioToOne(m.Measurement.Value)
		m.Measurement.Value = OunceToGram(ounces)
	// Apothecary
	case ApothecaryScruple, ApothecaryDrachm, ApothecaryOunce, ApothecaryPound:
		ounces := m.Measurement.Unit.RatioToOne(m.Measurement.Value)
		m.Measurement.Value = OunceToGram(ounces)
	// Troy
	case TroyOunce, TroyPound, Pennyweight:
		ounces := m.Measurement.Unit.RatioToOne(m.Measurement.Value)
		m.Measurement.Value = OunceToGram(ounces)
	default:
		err := fmt.Errorf("Unknown mass unit, %v", m.Measurement.Unit)
		return err, 0
	}

	m.Measurement.Unit = Gram

	return nil, m.Measurement.Value
}

// ConvertFromGram is a helper to convert a value from its value in Grams
func (m *Mass) ConvertFromGram(handle func(float64) float64) (error, float64) {
	err, litre := m.Gram()
	if err != nil {
		return err, 0
	}

	return nil, handle(litre)
}

// Yottagram converts a mass to its equivalent value in Yottagrams
func (m *Mass) Yottagram() (error, float64) {
	return m.ConvertFromGram(Yottagram.OneToRatio)
}

// Zettagram converts a mass to its equivalent value in Zettagrams
func (m *Mass) Zettagram() (error, float64) {
	return m.ConvertFromGram(Zettagram.OneToRatio)
}

// Exagram converts a mass to its equivalent value in Exagrams
func (m *Mass) Exagram() (error, float64) {
	return m.ConvertFromGram(Exagram.OneToRatio)
}

// Petagram converts a mass to its equivalent value in Petagrams
func (m *Mass) Petagram() (error, float64) {
	return m.ConvertFromGram(Petagram.OneToRatio)
}

// Teragram converts a mass to its equivalent value in Teragrams
func (m *Mass) Teragram() (error, float64) {
	return m.ConvertFromGram(Teragram.OneToRatio)
}

// Gigagram converts a mass to its equivalent value in Gigagrams
func (m *Mass) Gigagram() (error, float64) {
	return m.ConvertFromGram(Gigagram.OneToRatio)
}

// Megagram converts a mass to its equivalent value in Megagrams
func (m *Mass) Megagram() (error, float64) {
	return m.ConvertFromGram(Megagram.OneToRatio)
}

// Kilogram converts a mass to its equivalent value in Kilograms
func (m *Mass) Kilogram() (error, float64) {
	return m.ConvertFromGram(Kilogram.OneToRatio)
}

// Hectogram converts a mass to its equivalent value in Hectograms
func (m *Mass) Hectogram() (error, float64) {
	return m.ConvertFromGram(Hectogram.OneToRatio)
}

// Decagram converts a mass to its equivalent value in Decagrams
func (m *Mass) Decagram() (error, float64) {
	return m.ConvertFromGram(Decagram.OneToRatio)
}

// Decigram converts a mass to its equivalent value in Decigrams
func (m *Mass) Decigram() (error, float64) {
	return m.ConvertFromGram(Decigram.OneToRatio)
}

// Centigram converts a mass to its equivalent value in Centigrams
func (m *Mass) Centigram() (error, float64) {
	return m.ConvertFromGram(Centigram.OneToRatio)
}

// Milligram converts a mass to its equivalent value in Milligrams
func (m *Mass) Milligram() (error, float64) {
	return m.ConvertFromGram(Milligram.OneToRatio)
}

// Microgram converts a mass to its equivalent value in Micrograms
func (m *Mass) Microgram() (error, float64) {
	return m.ConvertFromGram(Microgram.OneToRatio)
}

// Nanogram converts a mass to its equivalent value in Nanograms
func (m *Mass) Nanogram() (error, float64) {
	return m.ConvertFromGram(Nanogram.OneToRatio)
}

// Picogram converts a mass to its equivalent value in Picograms
func (m *Mass) Picogram() (error, float64) {
	return m.ConvertFromGram(Picogram.OneToRatio)
}

// Femtogram converts a mass to its equivalent value in Femtograms
func (m *Mass) Femtogram() (error, float64) {
	return m.ConvertFromGram(Femtogram.OneToRatio)
}

// Attogram converts a mass to its equivalent value in Attograms
func (m *Mass) Attogram() (error, float64) {
	return m.ConvertFromGram(Attogram.OneToRatio)
}

// Zeptogram converts a mass to its equivalent value in Zeptograms
func (m *Mass) Zeptogram() (error, float64) {
	return m.ConvertFromGram(Zeptogram.OneToRatio)
}

// Yoctogram converts a mass to its equivalent value in Yoctograms
func (m *Mass) Yoctogram() (error, float64) {
	return m.ConvertFromGram(Yoctogram.OneToRatio)
}
