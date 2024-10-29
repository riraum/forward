package geometry

// Precision for floating point comparisons
const epsilon = 1e-9

func float64Eq(a, b float64) bool {
	return a-b < epsilon && b-a < epsilon
}
