package recipe

import (
	"fmt"
)

var (
	Celcius   = Unit{Singular: "Degree Celcius", Plural: "Degrees Celcius", Abbreviation: "°C"}
	Farenheit = Unit{Singular: "Degree Farenheit", Plural: "Degrees Farenheit", Abbreviation: "°F"}
	Kelvin    = Unit{Singular: "Degree Kelvin", Plural: "Degrees Kelvin", Abbreviation: "K"}
	Rankine   = Unit{Singular: "Degree Rankine", Plural: "Degrees Rankine", Abbreviation: "°Ra"}
	Reaumur   = Unit{Singular: "Degree Réaumur", Plural: "Degrees Réaumur", Abbreviation: "°Re"}
)

type Temperature struct {
	Measurement Measurement
}

// Convert a value in degrees Celcius to degrees Farenheit
func CelciusToFarenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

// Convert a value in degrees Celcius to degrees Kelvin
func CelciusToKelvin(value float64) float64 {
	return value + 273.15
}

// Convert a value in degrees Celcius to degrees Rankine
func CelciusToRankine(value float64) float64 {
	return (value * 9 / 5) + 491.67
}

// Convert a value in degrees Celcius to degrees Reaumur
func CelciusToReaumur(value float64) float64 {
	return value * 4 / 5
}

// Convert a value in degrees Farenheit to degrees Celcius
func FarenheitToCelcius(value float64) float64 {
	return (value - 32) * 5 / 9
}

// Convert a value in degrees Kelvin to degrees Celcius
func KelvinToCelcius(value float64) float64 {
	return value - 273.15
}

// Convert a value in degrees Rankine to degrees Celcius
func RankineToCelcius(value float64) float64 {
	return (value - 491.67) * 5 / 9
}

// Convert a value in degrees Reaumur to degrees Celcius
func ReaumurToCelcius(value float64) float64 {
	return value * 5 / 4
}

// Celcius converts a temperature to its equivalent in degress Celcius
func (t *Temperature) Celcius() (error, float64) {
	switch t.Measurement.Unit {
	case Celcius:
		// Measurement is already in Celcius
	case Farenheit:
		t.Measurement.Value = FarenheitToCelcius(t.Measurement.Value)
	case Kelvin:
		t.Measurement.Value = KelvinToCelcius(t.Measurement.Value)
	case Rankine:
		t.Measurement.Value = RankineToCelcius(t.Measurement.Value)
	case Reaumur:
		t.Measurement.Value = ReaumurToCelcius(t.Measurement.Value)
	default:
		err := fmt.Errorf("Unsupported Unit, %v", t.Measurement.Unit)
		return err, 0
	}

	t.Measurement.Unit = Celcius

	return nil, t.Measurement.Value
}

func (t *Temperature) ConvertFromCelcius(handle func(float64) float64) (error, float64) {
	err, celcius := t.Celcius()
	if err != nil {
		return err, 0
	}

	return nil, handle(celcius)
}

// Farenheit converts a temperature to its equivalent in degress Farenheit
func (t *Temperature) Farenheit() (error, float64) {
	return t.ConvertFromCelcius(CelciusToFarenheit)
}

// Kelvin converts a temperature to its equivalent in degress Kelvin
func (t *Temperature) Kelvin() (error, float64) {
	return t.ConvertFromCelcius(CelciusToKelvin)
}

// Rankine converts a temperature to its equivalent in degress Rankine
func (t *Temperature) Rankine() (error, float64) {
	return t.ConvertFromCelcius(CelciusToRankine)
}

// Reaumur converts a temperature to its equivalent in degress Reaumur
func (t *Temperature) Reaumur() (error, float64) {
	return t.ConvertFromCelcius(CelciusToReaumur)
}