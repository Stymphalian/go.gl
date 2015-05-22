package jgl

const (
	epsilon = 0.000000001
)

// Checks if two floats are equal. Doing a comparision using a small epsilon value
func closeEq(a, b, eps float32) bool {
	if a > b {
		return ((a - b) < eps)
	} else {
		return ((b - a) < eps)
	}
}

// Checks if two floats are equal. Doing a comparision using a small epsilon value
func closeEqf64(a, b, eps float64) bool {
	if a > b {
		return ((a - b) < eps)
	} else {
		return ((b - a) < eps)
	}
}
