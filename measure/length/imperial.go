package length

import (
	"github.com/lstellway/go/measure"
)

// Imperial
var (
	InchesPerFoot float64 = 12.0
	FeetPerYard float64 = 3.0
	FeetPerMile float64 = 5280.0

	Inch = measure.Unit{
		Singular: "Inch",
		Plural: "Inches",
		Abbreviation: "in",
		Ratio: 1,
	}

	Foot = measure.Unit{
		Singular: "Foot",
		Plural: "Feet",
		Abbreviation: "ft",
		Ratio: InchesPerFoot,
	}

	Yard = measure.Unit{
		Singular: "Yard",
		Plural: "Yards",
		Abbreviation: "yd",
		Ratio: InchesPerFoot * FeetPerYard,
	}

	Mile = measure.Unit{
		Singular: "Mile",
		Plural: "Miles",
		Abbreviation: "mi",
		Ratio: InchesPerFoot * FeetPerMile,
	}
)

// ConvertToFoot is a helper to convert a length value from Metres to its equivalent in Feet
func (l *Length) ConvertToFeet(unit measure.Unit) (error, float64) {
	return l.ConvertFromMetre(func(value float64) float64 {
		feet := MetresToFeet(value)
		return unit.OneToRatio(feet)
	})
}

// Inch converts a length to its equivalent in Inches
func (l *Length) Inch() (error, float64) {
	if l.Measurement.Unit == Inch {
		return nil, l.Measurement.Value
	}

	return l.ConvertToFeet(Inch)
}

// Foot converts a length to its equivalent in Feet
func (l *Length) Foot() (error, float64) {
	if l.Measurement.Unit == Foot {
		return nil, l.Measurement.Value
	}

	return l.ConvertToFeet(Foot)
}

// Yard converts a length to its equivalent in Yards
func (l *Length) Yard() (error, float64) {
	if l.Measurement.Unit == Yard {
		return nil, l.Measurement.Value
	}

	return l.ConvertToFeet(Yard)
}

// Mile converts a length to its equivalent in Miles
func (l *Length) Mile() (error, float64) {
	if l.Measurement.Unit == Mile {
		return nil, l.Measurement.Value
	}

	return l.ConvertToFeet(Mile)
}
