package selectionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var intAsc = func(a, b int) bool { return a < b }
var strAsc = func(a, b string) bool { return a < b }

func TestSelectionSortInts(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{7}, []int{7}},
		{[]int{}, []int{}},
	}

	for _, c := range cases {
		SelectionSort(c.input, intAsc)
		assert.Equal(t, c.expected, c.input)
	}
}

func TestSelectionSortStrings(t *testing.T) {
	data := []string{"pear", "apple", "orange"}
	SelectionSort(data, strAsc)
	assert.Equal(t, []string{"apple", "orange", "pear"}, data)
}
