package binary

func Search(haystack []int, needle int) int {
	low := 0
	high := len(haystack)

	for low < high {
		mid := low + (high-low)/2
		value := haystack[mid]

		switch {
		case needle == value:
			return mid
		case needle > value:
			low = mid + 1 // Drop the mid point
		case needle < value:
			high = mid
		}
	}

	return -1
}

func SearchRecursive(haystack []int, needle int) int {
	return recurse(haystack, needle, 0, len(haystack))
}

func recurse(haystack []int, needle, low, high int) int {
	if low >= high {
		return -1
	}

	mid := low + (high-low)/2
	value := haystack[mid]

	switch {
	case needle == value:
		return mid
	case needle > value:
		low = mid + 1
		return recurse(haystack, needle, low, high)
	case needle < value:
		high = mid
		return recurse(haystack, needle, low, high)
	}

	return -1
}
