package recipe

import "fmt"

var (
	// Ratios
	MetreToFootRatio = 39.37 / 12

	// International System of Units (meters)
	Metre      = Unit{Singular: "Metre", Plural: "Metres", Abbreviation: "m", Ratio: 1}
	Exametre   = Exa.Unit("%smetre", "%sm")
	Petametre  = Peta.Unit("%smetre", "%sm")
	Terametre  = Tera.Unit("%smetre", "%sm")
	Gigametre  = Giga.Unit("%smetre", "%sm")
	Megametre  = Mega.Unit("%smetre", "%sm")
	Kilometre  = Kilo.Unit("%smetre", "%sm")
	Hectometre = Hecto.Unit("%smetre", "%sm")
	Decametre  = Deca.Unit("%smetre", "%sm")
	Decimetre  = Deci.Unit("%smetre", "%sm")
	Centimetre = Centi.Unit("%smetre", "%sm")
	Millimetre = Milli.Unit("%smetre", "%sm")
	Micrometre = Micro.Unit("%smetre", "%sm")
	Nanometre  = Nano.Unit("%smetre", "%sm")
	Picometre  = Pico.Unit("%smetre", "%sm")
	Femtometre = Femto.Unit("%smetre", "%sm")
	Attometre  = Atto.Unit("%smetre", "%sm")

	// Imperial
	Foot = Unit{Singular: "Foot", Plural: "Feet", Abbreviation: "ft", Ratio: 1}
	Inch = Unit{Singular: "Inch", Plural: "Inches", Abbreviation: "in", Ratio: Foot.Ratio / 12}
	Yard = Unit{Singular: "Yard", Plural: "Yards", Abbreviation: "yd", Ratio: Foot.Ratio * 3}
	Mile = Unit{Singular: "Mile", Plural: "Miles", Abbreviation: "mi", Ratio: Foot.Ratio * 5280}

	// Land measurements
	Furlong      = Unit{Singular: "Furlong", Plural: "Furlongs", Abbreviation: "fur", Ratio: Mile.Ratio / 8}
	Chain        = Unit{Singular: "Chain", Plural: "Chains", Abbreviation: "ch", Ratio: Furlong.Ratio / 10}
	Link         = Unit{Singular: "Link", Plural: "Links", Abbreviation: "link", Ratio: Chain.Ratio / 100}
	Rod          = Unit{Singular: "Rod", Plural: "Rods", Abbreviation: "rod", Ratio: Chain.Ratio / 4}
	RamsdenChain = Unit{Singular: "Ramsden Chain", Plural: "Ramsden Chains", Abbreviation: "ram ch", Ratio: Foot.Ratio * 100}

	// Nautical units
	NauticalMile = Unit{Singular: "Nautical Mile", Plural: "Nautical Miles", Abbreviation: "NM", Ratio: Foot.Ratio * 6076}
	Cable        = Unit{Singular: "Cable", Plural: "Cables", Abbreviation: "Nc", Ratio: NauticalMile.Ratio / 10}
	Fathom       = Unit{Singular: "Fathom", Plural: "Fathoms", Abbreviation: "fm", Ratio: Foot.Ratio * 6}
	Shackle      = Unit{Singular: "Shackle", Plural: "Shackles", Abbreviation: "Ns", Ratio: Fathom.Ratio * 15}

	// Unofficial
	League = Unit{Singular: "League", Plural: "Leagues", Abbreviation: "Lg", Ratio: NauticalMile.Ratio * 3}
)

// FeetToMetres converts a given value in Feet to its equivalent in Metres
func FeetToMetres(value float64) float64 {
	return value / MetreToFootRatio
}

// MetresToFeet converts a given value in Metres to its equivalent Feet
func MetresToFeet(value float64) float64 {
	return value * MetreToFootRatio
}

type Length struct {
	Measurement Measurement
}

// Metres converts a length to its equivalent in metres
func (l *Length) Metre() (error, float64) {
	switch l.Measurement.Unit {
	case Metre:
		// Value is already in metres
	case Exametre, Petametre, Terametre, Gigametre, Megametre, Kilometre, Hectometre, Decametre, Decimetre, Centimetre, Millimetre, Micrometre, Nanometre, Picometre, Femtometre, Attometre:
		l.Measurement.Value = l.Measurement.Unit.OneToRatio(l.Measurement.Value)
	case Foot, Inch, Yard, Mile, Furlong, Chain, Link, Rod, RamsdenChain, NauticalMile, Cable, Fathom, Shackle, League:
		feet := l.Measurement.Unit.RatioToOne(l.Measurement.Value)
		l.Measurement.Value = FeetToMetres(feet)
	default:
		err := fmt.Errorf("Unknown length unit, %v", l.Measurement.Unit)
		return err, 0
	}

	l.Measurement.Unit = Metre

	return nil, l.Measurement.Value
}

// ConvertFromMetre is a helper to convert a length value from its value in Metres
func (l *Length) ConvertFromMetre(handle func(float64) float64) (error, float64) {
	err, metre := l.Metre()
	if err != nil {
		return err, 0
	}

	return nil, handle(metre)
}

// Exametre converts a length to its equivalent in Exametres
func (l *Length) Exametre() (error, float64) {
	return l.ConvertFromMetre(Exametre.OneToRatio)
}

// Petametre converts a length to its equivalent in Petametres
func (l *Length) Petametre() (error, float64) {
	return l.ConvertFromMetre(Petametre.OneToRatio)
}

// Terametre converts a length to its equivalent in Terametres
func (l *Length) Terametre() (error, float64) {
	return l.ConvertFromMetre(Terametre.OneToRatio)
}

// Gigametre converts a length to its equivalent in Gigametres
func (l *Length) Gigametre() (error, float64) {
	return l.ConvertFromMetre(Gigametre.OneToRatio)
}

// Megametre converts a length to its equivalent in Megametres
func (l *Length) Megametre() (error, float64) {
	return l.ConvertFromMetre(Megametre.OneToRatio)
}

// Kilometre converts a length to its equivalent in Kilometres
func (l *Length) Kilometre() (error, float64) {
	return l.ConvertFromMetre(Kilometre.OneToRatio)
}

// Hectometre converts a length to its equivalent in Hectometres
func (l *Length) Hectometre() (error, float64) {
	return l.ConvertFromMetre(Hectometre.OneToRatio)
}

// Decametre converts a length to its equivalent in Decametres
func (l *Length) Decametre() (error, float64) {
	return l.ConvertFromMetre(Decametre.OneToRatio)
}

// Decimetre converts a length to its equivalent in Decimetres
func (l *Length) Decimetre() (error, float64) {
	return l.ConvertFromMetre(Decimetre.OneToRatio)
}

// Centimetre converts a length to its equivalent in Centimetres
func (l *Length) Centimetre() (error, float64) {
	return l.ConvertFromMetre(Centimetre.OneToRatio)
}

// Millimetre converts a length to its equivalent in Millimetres
func (l *Length) Millimetre() (error, float64) {
	return l.ConvertFromMetre(Millimetre.OneToRatio)
}

// Micrometre converts a length to its equivalent in Micrometres
func (l *Length) Micrometre() (error, float64) {
	return l.ConvertFromMetre(Micrometre.OneToRatio)
}

// Nanometre converts a length to its equivalent in Nanometres
func (l *Length) Nanometre() (error, float64) {
	return l.ConvertFromMetre(Nanometre.OneToRatio)
}

// Picometre converts a length to its equivalent in Picometres
func (l *Length) Picometre() (error, float64) {
	return l.ConvertFromMetre(Picometre.OneToRatio)
}

// Femtometre converts a length to its equivalent in Femtometres
func (l *Length) Femtometre() (error, float64) {
	return l.ConvertFromMetre(Femtometre.OneToRatio)
}

// Attometre converts a length to its equivalent in Attometres
func (l *Length) Attometre() (error, float64) {
	return l.ConvertFromMetre(Attometre.OneToRatio)
}

// ConvertToFoot is a helper to convert a length value from Metres to its equivalent in Feet
func (l *Length) ConvertToFeet(unit Unit) (error, float64) {
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

// Furlong converts a length to its equivalent in Furlongs
func (l *Length) Furlong() (error, float64) {
	return l.ConvertToFeet(Furlong)
}

// Chain converts a length to its equivalent in Chains
func (l *Length) Chain() (error, float64) {
	return l.ConvertToFeet(Chain)
}

// Link converts a length to its equivalent in Links
func (l *Length) Link() (error, float64) {
	return l.ConvertToFeet(Link)
}

// Rod converts a length to its equivalent in Rods
func (l *Length) Rod() (error, float64) {
	return l.ConvertToFeet(Rod)
}

// RamsdenChain converts a length to its equivalent in Ramsden Chains
func (l *Length) RamsdenChain() (error, float64) {
	return l.ConvertToFeet(RamsdenChain)
}

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
