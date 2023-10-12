package selection_sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{2, 5, 6, 3, 8, 4, 7, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Sort(list, func(a, b int) bool {
		return a < b
	})

	for i := range list {
		if list[i] != expected[i] {
			t.Fatalf("expected %v at position %v, got %v", list[i], i, expected[i])
		}
	}
}
