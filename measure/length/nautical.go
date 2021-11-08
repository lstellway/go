package length

import "github.com/lstellway/go/measure"

var (
	// Nautical units
	NauticalMile = measure.Unit{Singular: "Nautical Mile", Plural: "Nautical Miles", Abbreviation: "NM", Ratio: Foot.Ratio * 6076}
	Cable        = measure.Unit{Singular: "Cable", Plural: "Cables", Abbreviation: "Nc", Ratio: NauticalMile.Ratio / 10}
	Fathom       = measure.Unit{Singular: "Fathom", Plural: "Fathoms", Abbreviation: "fm", Ratio: Foot.Ratio * 6}
	Shackle      = measure.Unit{Singular: "Shackle", Plural: "Shackles", Abbreviation: "Ns", Ratio: Fathom.Ratio * 15}
	League       = measure.Unit{Singular: "League", Plural: "Leagues", Abbreviation: "Lg", Ratio: NauticalMile.Ratio * 3}
)

// NauticalMile converts a length to its equivalent in Nautical Miles
func (l *Length) NauticalMile() (error, float64) {
	return l.ConvertToFeet(NauticalMile)
}

// Cable converts a length to its equivalent in Cables
func (l *Length) Cable() (error, float64) {
	return l.ConvertToFeet(Cable)
}

// Fathom converts a length to its equivalent in Fathoms
func (l *Length) Fathom() (error, float64) {
	return l.ConvertToFeet(Fathom)
}

// Shackle converts a length to its equivalent in Shackles
func (l *Length) Shackle() (error, float64) {
	return l.ConvertToFeet(Shackle)
}

// League converts a length to its equivalent in Leagues
func (l *Length) League() (error, float64) {
	return l.ConvertToFeet(League)
}
