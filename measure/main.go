/**
The measure package provides a unified API for converting between units of measure.
*/
package measure

import (
	"fmt"
	"math"
)

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
	Yotta = SIUnit{Name: "Yotta", Abbreviation: "Y", Power: 24}
	Zetta = SIUnit{Name: "Zetta", Abbreviation: "Z", Power: 21}
	Exa   = SIUnit{Name: "Exa", Abbreviation: "E", Power: 18}
	Peta  = SIUnit{Name: "Peta", Abbreviation: "P", Power: 15}
	Tera  = SIUnit{Name: "Tera", Abbreviation: "T", Power: 12}
	Giga  = SIUnit{Name: "Giga", Abbreviation: "G", Power: 9}
	Mega  = SIUnit{Name: "Mega", Abbreviation: "M", Power: 6}
	Kilo  = SIUnit{Name: "Kilo", Abbreviation: "k", Power: 3}
	Hecto = SIUnit{Name: "Hecto", Abbreviation: "h", Power: 2}
	Deca  = SIUnit{Name: "Deca", Abbreviation: "da", Power: 1}
	Deci  = SIUnit{Name: "Deci", Abbreviation: "d", Power: -1}
	Centi = SIUnit{Name: "Centi", Abbreviation: "c", Power: -2}
	Milli = SIUnit{Name: "Milli", Abbreviation: "m", Power: -3}
	Micro = SIUnit{Name: "Micro", Abbreviation: "μ", Power: -6}
	Nano  = SIUnit{Name: "Nano", Abbreviation: "n", Power: -9}
	Pico  = SIUnit{Name: "Pico", Abbreviation: "p", Power: -12}
	Femto = SIUnit{Name: "Femto", Abbreviation: "f", Power: -15}
	Atto  = SIUnit{Name: "Atto", Abbreviation: "a", Power: -18}
	Zepto = SIUnit{Name: "Zepto", Abbreviation: "z", Power: -21}
	Yocto = SIUnit{Name: "Yocto", Abbreviation: "y", Power: -24}
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
