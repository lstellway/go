package length

import "github.com/lstellway/go/measure"

var (
	// Land measurements
	Furlong      = measure.Unit{Singular: "Furlong", Plural: "Furlongs", Abbreviation: "fur", Ratio: Mile.Ratio / 8}
	Chain        = measure.Unit{Singular: "Chain", Plural: "Chains", Abbreviation: "ch", Ratio: Furlong.Ratio / 10}
	Link         = measure.Unit{Singular: "Link", Plural: "Links", Abbreviation: "link", Ratio: Chain.Ratio / 100}
	Rod          = measure.Unit{Singular: "Rod", Plural: "Rods", Abbreviation: "rod", Ratio: Chain.Ratio / 4}
	RamsdenChain = measure.Unit{Singular: "Ramsden Chain", Plural: "Ramsden Chains", Abbreviation: "ram ch", Ratio: Foot.Ratio * 100}
)

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
