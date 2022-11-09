package common

import "math"

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
