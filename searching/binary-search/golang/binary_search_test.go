package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	is := []int{2, 3, 5, 7, 9}
	toSearch := 7
	expected := 3

	res := BinarySearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}

func TestBinarySearchWithValueAtBeginningOfList(t *testing.T) {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	toSearch := 1
	expected := 0

	res := BinarySearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}

func TestBinarySearchWithValueAtEndOfList(t *testing.T) {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	toSearch := 9
	expected := 8

	res := BinarySearch(is, toSearch)

	if res != expected {
		t.Fatalf("Expected to find %v at %v but got %v", toSearch, expected, res)
	}
}
