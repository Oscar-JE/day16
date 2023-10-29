package intopr

import "math"

func ArrayEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	equal := true
	for i := range a {
		equal = equal && a[i] == b[i]
	}
	return equal
}

func PowerOfTwo(exponet int) int {
	return int(math.Pow(2.0, float64(exponet)))
}
