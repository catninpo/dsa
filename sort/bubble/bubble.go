package bubble

func Sort(n []int) {
	for i := len(n) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
}
