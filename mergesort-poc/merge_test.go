package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortInts(t *testing.T) {
	data := []int{38, 27, 43, 3, 9, 82, 10}
	expected := []int{3, 9, 10, 27, 38, 43, 82}

	MergeSort(data, func(a, b int) bool { return a < b })

	assert.Equal(t, expected, data)
}

func TestMergeSortStrings(t *testing.T) {
	data := []string{"delta", "alpha", "charlie", "bravo"}
	expected := []string{"alpha", "bravo", "charlie", "delta"}

	MergeSort(data, func(a, b string) bool { return a < b })

	assert.Equal(t, expected, data)
}

func TestAlreadySorted(t *testing.T) {
	data := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	MergeSort(data, func(a, b int) bool { return a < b })

	assert.Equal(t, expected, data)
}

func TestReverseSorted(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	MergeSort(data, func(a, b int) bool { return a < b })

	assert.Equal(t, expected, data)
}

func TestEmptyAndSingle(t *testing.T) {
	var empty []int
	MergeSort(empty, func(a, b int) bool { return a < b })
	assert.Empty(t, empty)

	one := []int{42}
	MergeSort(one, func(a, b int) bool { return a < b })
	assert.Equal(t, []int{42}, one)
}
