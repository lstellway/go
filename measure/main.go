package measure

import (
	"fmt"
	"math"
)

// TODO: Units of measurement (Speed, Acceleration, Power, Time, Light Intensity)
// @see https://www.theedkins.co.uk/jo/units/other.htm
// Metric System
// @see https://www.theedkins.co.uk/jo/units/metric.htm
// SI (time, length, mass, electric current, thermodynamic temperature, amount of substance, luminous intensity)
// https://en.wikipedia.org/wiki/International_System_of_Units

type SIUnit struct {
	Name         string
	Abbreviation string
	Power        int
}

type Unit struct {
	Singular     string
	Plural       string
	Abbreviation string
	Ratio        float64
}

type Measurement struct {
	Unit  Unit
	Value float64
}

var (
	Yotta = SIUnit{"Yotta", "Y", 24}
	Zetta = SIUnit{"Zetta", "Z", 21}
	Exa   = SIUnit{"Exa", "E", 18}
	Peta  = SIUnit{"Peta", "P", 15}
	Tera  = SIUnit{"Tera", "T", 12}
	Giga  = SIUnit{"Giga", "G", 9}
	Mega  = SIUnit{"Mega", "M", 6}
	Kilo  = SIUnit{"Kilo", "k", 3}
	Hecto = SIUnit{"Hecto", "h", 2}
	Deca  = SIUnit{"Deca", "da", 1}
	Deci  = SIUnit{"Deci", "d", -1}
	Centi = SIUnit{"Centi", "c", -2}
	Milli = SIUnit{"Milli", "m", -3}
	Micro = SIUnit{"Micro", "μ", -6}
	Nano  = SIUnit{"Nano", "n", -9}
	Pico  = SIUnit{"Pico", "p", -12}
	Femto = SIUnit{"Femto", "f", -15}
	Atto  = SIUnit{"Atto", "a", -18}
	Zepto = SIUnit{"Zepto", "z", -21}
	Yocto = SIUnit{"Yocto", "y", -24}
)

// The Unit method builds a Unit object from the SIUnit object.
// A name format string and abbreviation format string are passed to
// help build the names and abbreviations using values from the SIUnit object.
func (u *SIUnit) Unit(nameFormat string, abbreviationFormat string) Unit {
	return Unit{
		Singular:     fmt.Sprintf(nameFormat, u.Name),
		Plural:       fmt.Sprintf(nameFormat+"s", u.Name),
		Abbreviation: fmt.Sprintf(abbreviationFormat, u.Abbreviation),
		Ratio:        math.Pow10(u.Power),
	}
}

// RatioToOne converts a value to its base "1" equivalent using the Unit Ratio
func (u *Unit) RatioToOne(value float64) float64 {
	return value * u.Ratio
}

// OneToRatio converts a base "1" value to its equivalent using the Unit Ratio
func (u *Unit) OneToRatio(value float64) float64 {
	return value / u.Ratio
}

// RelativeRatio gets the ratio factor of a relative / comparable unit
func (u *Unit) RelativeRatio(consequent Unit) float64 {
	return u.Ratio / consequent.Ratio
}
