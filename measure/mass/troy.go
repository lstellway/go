package mass

import "github.com/lstellway/go/measure"

var (
	TroyOunce   = measure.Unit{Singular: "Troy Ounce", Plural: "Troy Ounces", Abbreviation: "oz t", Ratio: Grain.Ratio * 480}
	TroyPound   = measure.Unit{Singular: "Troy Pound", Plural: "Troy Pounds", Abbreviation: "oz t", Ratio: TroyOunce.Ratio * 12}
	Pennyweight = measure.Unit{Singular: "Pennyweight", Plural: "Pennyweights", Abbreviation: "dwt", Ratio: TroyOunce.Ratio / 20}
)
