// Package mymath provides simple math operations
package mymath

// Sum returns a sum of unlimited number of integers
func Sum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}

// go doc mymath
// go doc mymath.Sum
