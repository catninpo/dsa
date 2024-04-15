package mergesort

func Sort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	mid := len(n) / 2

	left := Sort(n[:mid])
	right := Sort(n[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}
