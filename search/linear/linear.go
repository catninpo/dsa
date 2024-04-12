package linear

// Search operates in O(n) time, time performance will grow in relation to the
// size of haystack linearly at worse case (where needle is in the last position
// of the array).
func Search(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
