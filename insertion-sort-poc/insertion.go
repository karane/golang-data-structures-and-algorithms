package insertionsort

func InsertionSort[T any](data []T, less func(a, b T) bool) {
	n := len(data)
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && less(key, data[j]) {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}
