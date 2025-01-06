package random

import (
	"math/rand"
)

func Choose(values []int) int {
	// Select a random value from slice of values
	return values[rand.Intn(len(values))]
}
