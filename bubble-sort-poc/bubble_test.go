package bubblesort

import (
	"reflect"
	"testing"
)

var (
	intAsc    = func(a, b int) bool { return a < b }
	stringAsc = func(a, b string) bool { return a < b }
)

func TestBubbleSortInts(t *testing.T) {
	data := []int{5, 1, 4, 2, 8}
	expected := []int{1, 2, 4, 5, 8}

	BubbleSort(data, intAsc)

	if !reflect.DeepEqual(data, expected) {
		t.Fatalf("expected %v, got %v", expected, data)
	}
}

func TestBubbleSortStrings(t *testing.T) {
	data := []string{"d", "a", "c", "b"}
	expected := []string{"a", "b", "c", "d"}

	BubbleSort(data, stringAsc)

	if !reflect.DeepEqual(data, expected) {
		t.Fatalf("expected %v, got %v", expected, data)
	}
}

func TestAlreadySorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	BubbleSort(data, intAsc)

	if !reflect.DeepEqual(data, expected) {
		t.Fatalf("expected %v, got %v", expected, data)
	}
}

func TestEmptyAndSingle(t *testing.T) {
	var empty []int
	BubbleSort(empty, intAsc)

	one := []int{42}
	BubbleSort(one, intAsc)

	if len(one) != 1 || one[0] != 42 {
		t.Fatalf("unexpected result for single element")
	}
}
