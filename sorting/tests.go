package sorting

import "testing"

type InPlaceSortFunc func(list []int)

func TestInPlaceSortingAlgorithm(t *testing.T, sort InPlaceSortFunc) {
	list := []int{2, 5, 6, 3, 8, 4, 7, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sort(list)

	for i := range list {
		if list[i] != expected[i] {
			t.Fatalf("expected %v at position %v, got %v\r\nexpected: %v\r\nrecieved: %v", expected[i], i, list[i], expected, list)
		}
	}
}
