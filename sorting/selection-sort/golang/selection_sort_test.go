package selection_sort

import (
	"sorting"
	"testing"
)

func TestSort(t *testing.T) {
	sorting.TestInPlaceSortingAlgorithm(t, Sort)
}
