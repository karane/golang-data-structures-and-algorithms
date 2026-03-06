package selectionsort

func SelectionSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if less(data[j], data[minIdx]) {
				minIdx = j
			}
		}

		if minIdx != i {
			data[i], data[minIdx] = data[minIdx], data[i]
		}
	}
}
