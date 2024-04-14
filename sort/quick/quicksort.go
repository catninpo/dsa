package quicksort

func Sort(n []int) {
	qSort(n, 0, len(n)-1)
}

func qSort(n []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(n, low, high)

	qSort(n, low, pivotIndex-1)
	qSort(n, pivotIndex+1, high)
}

func partition(n []int, low, high int) int {
	pivot := n[high]
	index := low - 1

	for i := low; i < high; i++ {
		if n[i] <= pivot {
			index++
			n[i], n[index] = n[index], n[i]
		}
	}

	index++
	n[high], n[index] = n[index], n[high]

	return index
}
