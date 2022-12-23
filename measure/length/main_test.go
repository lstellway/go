package length

import (
	"fmt"
	"testing"
	"github.com/lstellway/go/measure"
)

func ConductLengthTest(
	t *testing.T,
	subjects map[float64]float64,
	controlUnit measure.Unit,
	hypothesisUnit measure.Unit,
) {
	for control, hypothesis := range subjects {
		l := CreateLength(controlUnit, control)

		err, value := l.Foot()

		fmt.Printf(
			"%v %s = %v %s\n",
			control,
			controlUnit.Plural,
			value,
			hypothesisUnit.Plural,
		)

		if err != nil || value != hypothesis {
			t.Fatalf(
				`%v %s should be %v %s, got %v`,
				control,
				controlUnit.Plural, 
				hypothesis,
				hypothesisUnit.Plural,
				value,
			)
		}
	}
}
