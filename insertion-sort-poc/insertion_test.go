package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortInts(t *testing.T) {
	data := []int{9, 5, 1, 4, 3}
	InsertionSort(data, func(a, b int) bool { return a < b })
	assert.Equal(t, []int{1, 3, 4, 5, 9}, data)
}

func TestInsertionSortStrings(t *testing.T) {
	data := []string{"delta", "alpha", "charlie", "bravo"}
	InsertionSort(data, func(a, b string) bool { return a < b })
	assert.Equal(t, []string{"alpha", "bravo", "charlie", "delta"}, data)
}

func TestAlreadySorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	InsertionSort(data, func(a, b int) bool { return a < b })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, data)
}

func TestReverseSorted(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	InsertionSort(data, func(a, b int) bool { return a < b })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, data)
}

func TestEmptyAndSingle(t *testing.T) {
	var empty []int
	InsertionSort(empty, func(a, b int) bool { return a < b })
	assert.Empty(t, empty)

	single := []int{42}
	InsertionSort(single, func(a, b int) bool { return a < b })
	assert.Equal(t, []int{42}, single)
}
