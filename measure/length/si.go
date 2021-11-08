package length

import (
	"fmt"

	"github.com/lstellway/go/measure"
)

var (
	// International System of Units (meters)
	Metre      = measure.Unit{Singular: "Metre", Plural: "Metres", Abbreviation: "m", Ratio: 1}
	Yottametre = measure.Yotta.Unit("%smetre", "%sm")
	Zettametre = measure.Zetta.Unit("%smetre", "%sm")
	Exametre   = measure.Exa.Unit("%smetre", "%sm")
	Petametre  = measure.Peta.Unit("%smetre", "%sm")
	Terametre  = measure.Tera.Unit("%smetre", "%sm")
	Gigametre  = measure.Giga.Unit("%smetre", "%sm")
	Megametre  = measure.Mega.Unit("%smetre", "%sm")
	Kilometre  = measure.Kilo.Unit("%smetre", "%sm")
	Hectometre = measure.Hecto.Unit("%smetre", "%sm")
	Decametre  = measure.Deca.Unit("%smetre", "%sm")
	Decimetre  = measure.Deci.Unit("%smetre", "%sm")
	Centimetre = measure.Centi.Unit("%smetre", "%sm")
	Millimetre = measure.Milli.Unit("%smetre", "%sm")
	Micrometre = measure.Micro.Unit("%smetre", "%sm")
	Nanometre  = measure.Nano.Unit("%smetre", "%sm")
	Picometre  = measure.Pico.Unit("%smetre", "%sm")
	Femtometre = measure.Femto.Unit("%smetre", "%sm")
	Attometre  = measure.Atto.Unit("%smetre", "%sm")
	Zeptometre = measure.Zepto.Unit("%smetre", "%sm")
	Yoctometre = measure.Yocto.Unit("%smetre", "%sm")
)

// Metres converts a length to its equivalent in metres
func (l *Length) Metre() (error, float64) {
	switch l.Measurement.Unit {
	case Metre:
		// Value is already in metres
	case Yottametre, Zettametre, Exametre, Petametre, Terametre, Gigametre, Megametre, Kilometre, Hectometre, Decametre, Decimetre, Centimetre, Millimetre, Micrometre, Nanometre, Picometre, Femtometre, Attometre, Zeptometre, Yoctometre:
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

// Yottametre converts a length to its equivalent in Yottametres
func (l *Length) Yottametre() (error, float64) {
	return l.ConvertFromMetre(Yottametre.OneToRatio)
}

// Zettametre converts a length to its equivalent in Zettametres
func (l *Length) Zettametre() (error, float64) {
	return l.ConvertFromMetre(Zettametre.OneToRatio)
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

// Zeptometre converts a length to its equivalent in Zeptometres
func (l *Length) Zeptometre() (error, float64) {
	return l.ConvertFromMetre(Zeptometre.OneToRatio)
}

// Yoctometre converts a length to its equivalent in Yoctometres
func (l *Length) Yoctometre() (error, float64) {
	return l.ConvertFromMetre(Yoctometre.OneToRatio)
}
