package huffman

import (
	"bytes"
	"testing"
)

func TestNewTree(t *testing.T) {
	valueLeafs := CalculateValueLeafs("This is an example of the huffman tree.")
	leaf_x := &Leaf{freq: 1, rune: 'x'}
	leaf_r := &Leaf{freq: 1, rune: 'r'}
	leaf_u := &Leaf{freq: 1, rune: 'u'}
	leaf_o := &Leaf{freq: 1, rune: 'o'}
	leaf_l := &Leaf{freq: 1, rune: 'l'}
	leaf_p := &Leaf{freq: 1, rune: 'p'}
	leaf_T := &Leaf{freq: 1, rune: 'T'}
	leaf_Dot := &Leaf{freq: 1, rune: '.'}
	leaf_t := &Leaf{freq: 2, rune: 't'}
	leaf_n := &Leaf{freq: 2, rune: 'n'}
	leaf_m := &Leaf{freq: 2, rune: 'm'}
	leaf_s := &Leaf{freq: 2, rune: 's'}
	leaf_i := &Leaf{freq: 2, rune: 'i'}
	leaf_a := &Leaf{freq: 3, rune: 'a'}
	leaf_f := &Leaf{freq: 3, rune: 'f'}
	leaf_h := &Leaf{freq: 3, rune: 'h'}
	leaf_e := &Leaf{freq: 5, rune: 'e'}
	leaf_Space := &Leaf{freq: 7, rune: ' '}

	leaf_xr := &Leaf{freq: 2, left: leaf_x, right: leaf_r}
	leaf_x.parent = leaf_xr
	leaf_r.parent = leaf_xr
	leaf_uo := &Leaf{freq: 2, left: leaf_u, right: leaf_o}
	leaf_u.parent = leaf_uo
	leaf_o.parent = leaf_uo
	leaf_lp := &Leaf{freq: 2, left: leaf_l, right: leaf_p}
	leaf_l.parent = leaf_lp
	leaf_p.parent = leaf_lp
	leaf_TDot := &Leaf{freq: 2, left: leaf_T, right: leaf_Dot}
	leaf_T.parent = leaf_TDot
	leaf_Dot.parent = leaf_TDot

	leaf_TDotlp := &Leaf{freq: 4, left: leaf_TDot, right: leaf_lp}
	leaf_TDot.parent = leaf_TDotlp
	leaf_lp.parent = leaf_TDotlp
	leaf_uoxr := &Leaf{freq: 4, left: leaf_uo, right: leaf_xr}
	leaf_uo.parent = leaf_uoxr
	leaf_xr.parent = leaf_uoxr
	leaf_tn := &Leaf{freq: 4, left: leaf_t, right: leaf_n}
	leaf_t.parent = leaf_tn
	leaf_n.parent = leaf_tn
	leaf_ms := &Leaf{freq: 4, left: leaf_m, right: leaf_s}
	leaf_m.parent = leaf_ms
	leaf_s.parent = leaf_ms

	leaf_ia := &Leaf{freq: 5, left: leaf_i, right: leaf_a}
	leaf_i.parent = leaf_ia
	leaf_i.parent = leaf_ia

	leaf_fh := &Leaf{freq: 6, left: leaf_f, right: leaf_h}
	leaf_f.parent = leaf_fh
	leaf_h.parent = leaf_fh

	leaf_TDotlpuoxr := &Leaf{freq: 8, left: leaf_TDotlp, right: leaf_uoxr}
	leaf_TDotlp.parent = leaf_TDotlpuoxr
	leaf_uoxr.parent = leaf_TDotlpuoxr
	leaf_tnms := &Leaf{freq: 8, left: leaf_tn, right: leaf_ms}
	leaf_tn.parent = leaf_tnms
	leaf_ms.parent = leaf_tnms

	leaf_iae := &Leaf{freq: 10, left: leaf_ia, right: leaf_e}
	leaf_ia.parent = leaf_iae
	leaf_e.parent = leaf_iae

	leaf_fhSpace := &Leaf{freq: 13, left: leaf_fh, right: leaf_Space}
	leaf_fh.parent = leaf_fhSpace
	leaf_Space.parent = leaf_fhSpace

	leaf_TDotlpuoxrtnms := &Leaf{freq: 16, left: leaf_TDotlpuoxr, right: leaf_tnms}
	leaf_TDotlpuoxr.parent = leaf_TDotlpuoxrtnms
	leaf_tnms.parent = leaf_TDotlpuoxrtnms

	leaf_iaefhSpace := &Leaf{freq: 23, left: leaf_iae, right: leaf_fhSpace}
	leaf_iae.parent = leaf_iaefhSpace
	leaf_fhSpace.parent = leaf_iaefhSpace

	root := &Leaf{freq: 39, left: leaf_TDotlpuoxrtnms, right: leaf_iaefhSpace}
	leaf_TDotlpuoxrtnms.parent = root
	leaf_iaefhSpace.parent = root

	expected := []*Leaf{
		leaf_x, leaf_r, leaf_u, leaf_o, leaf_l, leaf_p, leaf_T, leaf_Dot, // leafs with 1
		leaf_TDot, leaf_lp, leaf_uo, leaf_xr, // nodes with 2
		leaf_t, leaf_n, leaf_m, leaf_s, leaf_i, // leafs with 2
		leaf_a, leaf_f, leaf_h, // leafs with 3
		leaf_TDotlp, leaf_uoxr, leaf_tn, leaf_ms, // nodes with 4
		leaf_ia,                    // nodes with 5
		leaf_e,                     // leafs with 5
		leaf_fh,                    // nodes with 6
		leaf_Space,                 // leafs with 7
		leaf_TDotlpuoxr, leaf_tnms, // nodes with 8
		leaf_iae,            // nodes with 10
		leaf_fhSpace,        // nodes with 13
		leaf_TDotlpuoxrtnms, // nodes with 16
		leaf_iaefhSpace,     // nodes with 23,
		root,
	}

	tree := NewTree(valueLeafs)

	leafs := tree.getLeafs()

	for i, leaf := range leafs {
		if leaf.rune != expected[i].rune {
			t.Fatalf("leaf at %v has rune '%s' which should be '%s'", i, string(leafs[i].rune), string(expected[i].rune))
		}
		if leaf.freq != expected[i].freq {
			t.Fatalf("leaf at %v has frequency %v which should be %v", i, leaf.freq, expected[i].freq)
		}
	}
}

func TestCalculateValueLeafs(t *testing.T) {
	const input = "This is an example of the huffman tree."
	expected := []Leaf{
		{freq: 1, rune: 'x'},
		{freq: 1, rune: 'r'},
		{freq: 1, rune: 'u'},
		{freq: 1, rune: 'o'},
		{freq: 1, rune: 'l'},
		{freq: 1, rune: 'p'},
		{freq: 1, rune: 'T'},
		{freq: 1, rune: '.'},
		{freq: 2, rune: 't'},
		{freq: 2, rune: 'n'},
		{freq: 2, rune: 'm'},
		{freq: 2, rune: 's'},
		{freq: 2, rune: 'i'},
		{freq: 3, rune: 'a'},
		{freq: 3, rune: 'f'},
		{freq: 3, rune: 'h'},
		{freq: 5, rune: 'e'},
		{freq: 7, rune: ' '},
	}

	leafs := CalculateValueLeafs(input)

	for i, leaf := range leafs {
		if leaf.rune != expected[i].rune {
			t.Fatalf("leaf value %s at index %v does not match (expected %s)", string(leaf.rune), i, string(expected[i].rune))
		}
		if leaf.freq != expected[i].freq {
			t.Fatalf("leaf frequency for %v at index %v does not match (expected %v)", leaf.freq, i, expected[i].freq)
		}
	}
}

func TestTree_Encode(t *testing.T) {
	const input = "This is an example of the huffman tree."
	const expectedLength = 152
	const expected = "11111011001111010000011110100000010100100001011110001010001101111001010000110000011000101101100100000110111010011001110000010100100010111110001001011000"
	tree := NewTree(CalculateValueLeafs(input))
	writer := &BinaryStringWriter{}

	err := tree.Encode(input, writer)
	if err != nil {
		t.Fatalf("error while encoding: %v", err)
	}

	encoded := writer.buffer.String()
	if len(encoded) != expectedLength {
		t.Fatalf("the encoded value is %v runes long while %v were expected", len(encoded), expectedLength)
	}
	if encoded != expected {
		t.Fatalf("the encoded value is not what was expected\r\nreceived: %s\r\nexpected: %s", encoded, expected)
	}
}

func TestTree_Decode(t *testing.T) {
	const input = "This is an example of the huffman tree."
	tree := NewTree(CalculateValueLeafs(input))
	buff := bytes.Buffer{}
	writer := NewBitWriter(&buff)
	reader := NewBitReader(&buff)

	err := tree.Encode(input, writer)
	if err != nil {
		t.Fatalf("error while encoding: %v", err)
	}

	decoded, err := tree.Decode(reader)
	if err != nil {
		t.Fatalf("error while decoding: %v", err)
	}

	if decoded != input {
		t.Fatalf("the decoded value is not correct\r\nreceived: %s\r\nexpected: %s", decoded, input)
	}
}
