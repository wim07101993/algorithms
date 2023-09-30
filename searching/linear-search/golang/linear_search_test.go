package linear_search

import "testing"

func TestLinearSearch(t *testing.T) {
	is := []int{3, 7, 1, 8, 9}
	toSearch := 7
	expected := 1

	res := LinearSearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}

func TestLinearSearchWithValueAtBeginningOfList(t *testing.T) {
	is := []int{7, 9, 3, 8, 5, 6, 2, 4, 0}
	toSearch := 7
	expected := 0

	res := LinearSearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}

func TestLinearSearchWithValueAtEndOfList(t *testing.T) {
	is := []int{1, 7, 3, 2, 9, 6, 4, 5, 8}
	toSearch := 8
	expected := 8

	res := LinearSearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}
