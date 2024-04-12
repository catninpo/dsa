package twocrystalballs

import (
	"math"
)

// Two Crystal Ball Problem:
// Given two crystal balls that will break if dropped from a high enough distance,
// determine the exact spot in which they will break in the most optimised fashion.
func Solve(n []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(n)))))

	i := jump
	for ; i < len(n); i += jump {
		if n[i] {
			break
		}
	}

	// TODO: Need to walk up not down the array.
	for j := i - 1; j >= i-jump; j-- {
		if n[j] {
			return j
		}
	}

	return -1
}
