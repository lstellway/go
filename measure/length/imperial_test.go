package length

import (
	"fmt"
	"testing"
)

func TestConvertToFeet(t *testing.T) {
	fmt.Println("ConvertToFeet")

	// Inches
	inches_to_feet := map[float64]float64 {
		12: 1,
		21: 1.75,
		24: 2,
	}

	ConductLengthTest(t, inches_to_feet, Inch, Foot)

	// Feet
	feet_to_feet := map[float64]float64 {
		12: 12,
		21: 21,
		24.1875: 24.1875,
	}

	ConductLengthTest(t, feet_to_feet, Foot, Foot)

	// Yards
	yard_to_feet := map[float64]float64 {
		1: 3,
		12: 36,
		1.5: 4.5,
	}

	ConductLengthTest(t, yard_to_feet, Yard, Foot)

	// Miles
	miles_to_feet := map[float64]float64 {
		1: 5280,
	}

	ConductLengthTest(t, miles_to_feet, Mile, Foot)
}
