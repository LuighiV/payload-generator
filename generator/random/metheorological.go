// Package random provides functions to generate random values for
// metheorological data
package random

import (
	"fmt"
	"math/rand"
)

func generateRandom(basevalue float64, rangevariation float64) error {
	return basevalue + rand.Float64*rangevariation/2 - rangevariation/2
}
