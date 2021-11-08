package length

import "github.com/lstellway/go/measure"

var (
	// Imperial
	Foot = measure.Unit{Singular: "Foot", Plural: "Feet", Abbreviation: "ft", Ratio: 1}
	Inch = measure.Unit{Singular: "Inch", Plural: "Inches", Abbreviation: "in", Ratio: Foot.Ratio / 12}
	Yard = measure.Unit{Singular: "Yard", Plural: "Yards", Abbreviation: "yd", Ratio: Foot.Ratio * 3}
	Mile = measure.Unit{Singular: "Mile", Plural: "Miles", Abbreviation: "mi", Ratio: Foot.Ratio * 5280}
)

// ConvertToFoot is a helper to convert a length value from Metres to its equivalent in Feet
func (l *Length) ConvertToFeet(unit measure.Unit) (error, float64) {
	return l.ConvertFromMetre(func(value float64) float64 {
		return unit.OneToRatio(MetresToFeet(value))
	})
}

// Foot converts a length to its equivalent in Feet
func (l *Length) Foot() (error, float64) {
	return l.ConvertToFeet(Foot)
}

// Inch converts a length to its equivalent in Inches
func (l *Length) Inch() (error, float64) {
	return l.ConvertToFeet(Inch)
}

// Yard converts a length to its equivalent in Yards
func (l *Length) Yard() (error, float64) {
	return l.ConvertToFeet(Yard)
}

// Mile converts a length to its equivalent in Miles
func (l *Length) Mile() (error, float64) {
	return l.ConvertToFeet(Mile)
}
