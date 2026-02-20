package mergesort

func MergeSort[T any](data []T, less func(a, b T) bool) {
	if len(data) < 2 {
		return
	}

	temp := make([]T, len(data))
	mergeSort(data, temp, less)
}

func mergeSort[T any](data, temp []T, less func(a, b T) bool) {
	if len(data) <= 1 {
		return
	}

	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	mergeSort(left, temp[:mid], less)
	mergeSort(right, temp[mid:], less)

	merge(left, right, temp, less)

	copy(data, temp[:len(data)])
}

func merge[T any](left, right, temp []T, less func(a, b T) bool) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) {
			temp[k] = left[i]
			i++
		} else {
			temp[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		temp[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		temp[k] = right[j]
		j++
		k++
	}
}
