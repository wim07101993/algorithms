package delta

import "testing"

func TestEncode(t *testing.T) {
	input := []int{3, 6, 9, 4, 8, 1, 5, 9, 4, 8, 2}
	expected := []int{3, 3, 3, -5, 4, -7, 4, 4, -5, 4, -6}

	encoded := Encode(input)

	for i, v := range expected {
		if v != encoded[i] {
			t.Fatalf("expected %v to encode to %v, but got %v (difference at %v)", input, expected, encoded, i)
			return
		}
	}
}

func TestDecode(t *testing.T) {
	input := []int{3, 3, 3, -5, 4, -7, 4, 4, -5, 4, -6}
	expected := []int{3, 6, 9, 4, 8, 1, 5, 9, 4, 8, 2}

	encoded := Decode(input)

	for i, v := range expected {
		if v != encoded[i] {
			t.Fatalf("expected %v to encode to %v, but got %v (difference at %v)", input, expected, encoded, i)
			return
		}
	}
}
