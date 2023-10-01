package run_length

import "testing"

func TestRunLengthEncode(t *testing.T) {
	input := "aabbbccccdddddeeeeeeefffffggg"
	expected := "2a,3b,4c,5d,7e,5f,3g"

	encoded := RunLengthEncode(input)

	if encoded != expected {
		t.Fatalf("Expected %s to encode to %s but got %s", input, expected, encoded)
	}
}

func TestRunLengthDecode(t *testing.T) {
	input := "2a,3b,4c,5d,7e,5f,3g"
	expected := "aabbbccccdddddeeeeeeefffffggg"

	encoded, err := RunLengthDecode(input)

	if err != nil {
		t.Fatalf("Go error while decoding: %v", err)
	}

	if encoded != expected {
		t.Fatalf("Expected %s to decode to %s but got %s", input, expected, encoded)
	}
}
