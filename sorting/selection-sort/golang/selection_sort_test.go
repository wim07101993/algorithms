package selection_sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{2, 5, 6, 3, 8, 4, 7, 9, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	output := Sort(input, func(a, b int) int {
		if a < b {
			return -1
		}
		if b < a {
			return 1
		}
		return 0
	})

	if len(output) != len(expected) {
		t.Fatal("expected a list with as many items out as in")
	}
	for i := range output {
		if output[i] != expected[i] {
			t.Fatalf("expected %v at position %v, got %v", output[i], i, expected[i])
		}
	}
}
