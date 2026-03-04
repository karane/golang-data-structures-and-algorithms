package bubblesort

func BubbleSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if less(data[j+1], data[j]) {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true
			}
		}

		// Optimization: stop if already sorted
		if !swapped {
			return
		}
	}
}
