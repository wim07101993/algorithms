package huffman

import (
	"bytes"
	"testing"
)

func TestNewTree(t *testing.T) {
	valueLeafs := CalculateValueLeafs("This is an example of the huffman tree.")
	leaf_x := &Leaf{freq: 1, value: "x"}
	leaf_r := &Leaf{freq: 1, value: "r"}
	leaf_u := &Leaf{freq: 1, value: "u"}
	leaf_o := &Leaf{freq: 1, value: "o"}
	leaf_l := &Leaf{freq: 1, value: "l"}
	leaf_p := &Leaf{freq: 1, value: "p"}
	leaf_T := &Leaf{freq: 1, value: "T"}
	leaf_Dot := &Leaf{freq: 1, value: "."}
	leaf_t := &Leaf{freq: 2, value: "t"}
	leaf_n := &Leaf{freq: 2, value: "n"}
	leaf_m := &Leaf{freq: 2, value: "m"}
	leaf_s := &Leaf{freq: 2, value: "s"}
	leaf_i := &Leaf{freq: 2, value: "i"}
	leaf_a := &Leaf{freq: 3, value: "a"}
	leaf_f := &Leaf{freq: 3, value: "f"}
	leaf_h := &Leaf{freq: 3, value: "h"}
	leaf_e := &Leaf{freq: 5, value: "e"}
	leaf_Space := &Leaf{freq: 7, value: " "}

	leaf_xr := &Leaf{freq: 2, value: "xr", left: leaf_x, right: leaf_r}
	leaf_x.parent = leaf_xr
	leaf_r.parent = leaf_xr
	leaf_uo := &Leaf{freq: 2, value: "uo", left: leaf_u, right: leaf_o}
	leaf_u.parent = leaf_uo
	leaf_o.parent = leaf_uo
	leaf_lp := &Leaf{freq: 2, value: "lp", left: leaf_l, right: leaf_p}
	leaf_l.parent = leaf_lp
	leaf_p.parent = leaf_lp
	leaf_TDot := &Leaf{freq: 2, value: "T.", left: leaf_T, right: leaf_Dot}
	leaf_T.parent = leaf_TDot
	leaf_Dot.parent = leaf_TDot

	leaf_ms := &Leaf{freq: 4, value: "ms", left: leaf_m, right: leaf_s}
	leaf_m.parent = leaf_ms
	leaf_s.parent = leaf_ms
	leaf_tn := &Leaf{freq: 4, value: "tn", left: leaf_t, right: leaf_n}
	leaf_t.parent = leaf_tn
	leaf_n.parent = leaf_tn
	leaf_uoxr := &Leaf{freq: 4, value: "uoxr", left: leaf_uo, right: leaf_xr}
	leaf_uo.parent = leaf_uoxr
	leaf_xr.parent = leaf_uoxr
	leaf_TDotlp := &Leaf{freq: 4, value: "T.lp", left: leaf_TDot, right: leaf_lp}
	leaf_TDot.parent = leaf_TDotlp
	leaf_lp.parent = leaf_TDotlp

	leaf_ia := &Leaf{freq: 5, value: "ia", left: leaf_i, right: leaf_a}
	leaf_i.parent = leaf_ia
	leaf_a.parent = leaf_ia

	leaf_fh := &Leaf{freq: 6, value: "fh", left: leaf_f, right: leaf_h}
	leaf_f.parent = leaf_fh
	leaf_h.parent = leaf_fh

	leaf_uoxrTDotlp := &Leaf{freq: 8, value: "uoxrT.lp", left: leaf_uoxr, right: leaf_TDotlp}
	leaf_TDotlp.parent = leaf_uoxrTDotlp
	leaf_uoxr.parent = leaf_uoxrTDotlp
	leaf_mstn := &Leaf{freq: 8, value: "mstn", left: leaf_ms, right: leaf_tn}
	leaf_tn.parent = leaf_mstn
	leaf_ms.parent = leaf_mstn

	leaf_iae := &Leaf{freq: 10, value: "iae", left: leaf_ia, right: leaf_e}
	leaf_ia.parent = leaf_iae
	leaf_e.parent = leaf_iae

	leaf_fhSpace := &Leaf{freq: 13, value: "fh ", left: leaf_fh, right: leaf_Space}
	leaf_fh.parent = leaf_fhSpace
	leaf_Space.parent = leaf_fhSpace

	leaf_uoxrTDotlpmstn := &Leaf{freq: 16, value: "uoxrT.lpmstn", left: leaf_uoxrTDotlp, right: leaf_mstn}
	leaf_uoxrTDotlp.parent = leaf_uoxrTDotlpmstn
	leaf_mstn.parent = leaf_uoxrTDotlpmstn

	leaf_iaefhSpace := &Leaf{freq: 23, value: "iaefh ", left: leaf_iae, right: leaf_fhSpace}
	leaf_iae.parent = leaf_iaefhSpace
	leaf_fhSpace.parent = leaf_iaefhSpace

	root := &Leaf{freq: 39, value: "uoxrT.lpmstniaefh ", left: leaf_uoxrTDotlpmstn, right: leaf_iaefhSpace}
	leaf_uoxrTDotlpmstn.parent = root
	leaf_iaefhSpace.parent = root

	expected := []*Leaf{
		leaf_x, leaf_r, leaf_u, leaf_o, leaf_l, leaf_p, leaf_T, leaf_Dot, // leafs with 1
		leaf_TDot, leaf_lp, leaf_uo, leaf_xr, // nodes with 2
		leaf_t, leaf_n, leaf_m, leaf_s, leaf_i, // leafs with 2
		leaf_a, leaf_f, leaf_h, // leafs with 3
		leaf_ms, leaf_tn, leaf_uoxr, leaf_TDotlp, // nodes with 4
		leaf_ia,                    // nodes with 5
		leaf_e,                     // leafs with 5
		leaf_fh,                    // nodes with 6
		leaf_Space,                 // leafs with 7
		leaf_uoxrTDotlp, leaf_mstn, // nodes with 8
		leaf_iae,            // nodes with 10
		leaf_fhSpace,        // nodes with 13
		leaf_uoxrTDotlpmstn, // nodes with 16
		leaf_iaefhSpace,     // nodes with 23,
		root,
	}

	tree := NewTree(valueLeafs)

	for i, leaf := range tree.leafs {
		if leaf.value != expected[i].value {
			t.Fatalf("leaf at %v has value '%s' which should be '%s'", i, tree.leafs[i].value, expected[i].value)
		}
		if leaf.freq != expected[i].freq {
			t.Fatalf("leaf %s has frequency %v which should be %v", leaf.value, leaf.freq, expected[i].freq)
		}
		if leaf.parent == nil && expected[i].parent != nil {
			t.Fatalf("leaf %s has no parent while it should", leaf.value)
		}
		if leaf.parent != nil && expected[i].parent == nil {
			t.Fatalf("leaf %s has no parent while it should", leaf.value)
		}
		if leaf.left == nil && expected[i].left != nil {
			t.Fatalf("leaf %s has no left child while %s whas expected", leaf.value, expected[i].left)
		} else if leaf.left != nil && expected[i].left == nil {
			t.Fatalf("leaf %s has a left child while none whas expected", leaf.value)
		} else if leaf.left != nil && leaf.left.value != expected[i].left.value {
			t.Fatalf("leaf %s has left child %s while %s whas expected", leaf.value, leaf.left, expected[i].left)
		}
		if leaf.right == nil && expected[i].right != nil {
			t.Fatalf("leaf %s has no right child while %s whas expected", leaf.value, expected[i].right)
		} else if leaf.right != nil && expected[i].right == nil {
			t.Fatalf("leaf %s has a right child while none whas expected", leaf.value)
		} else if leaf.right != nil && leaf.right.value != expected[i].right.value {
			t.Fatalf("leaf %s has right child %s while %s whas expected", leaf.value, leaf.right, expected[i].right)
		}
	}
}

func TestCalculateValueLeafs(t *testing.T) {
	const input = "This is an example of the huffman tree."
	expected := []Leaf{
		{freq: 1, value: "x"},
		{freq: 1, value: "r"},
		{freq: 1, value: "u"},
		{freq: 1, value: "o"},
		{freq: 1, value: "l"},
		{freq: 1, value: "p"},
		{freq: 1, value: "T"},
		{freq: 1, value: "."},
		{freq: 2, value: "t"},
		{freq: 2, value: "n"},
		{freq: 2, value: "m"},
		{freq: 2, value: "s"},
		{freq: 2, value: "i"},
		{freq: 3, value: "a"},
		{freq: 3, value: "f"},
		{freq: 3, value: "h"},
		{freq: 5, value: "e"},
		{freq: 7, value: " "},
	}

	leafs := CalculateValueLeafs(input)

	for i, leaf := range leafs {
		if leaf.value != expected[i].value {
			t.Fatalf("leaf value %s at index %v does not match (expected %s)", string(leaf.value), i, string(expected[i].value))
		}
		if leaf.freq != expected[i].freq {
			t.Fatalf("leaf frequency for %v at index %v does not match (expected %v)", leaf.freq, i, expected[i].freq)
		}
	}
}

func TestTree_WriteBinaryValueOfLeaf(t *testing.T) {
	const input = "This is an example of the huffman tree."
	tree := NewTree(CalculateValueLeafs(input))
	writer := &BinaryStringWriter{}
	expect := func(r rune, expected string) {
		leaf, err := tree.GetLeafForRune(r)
		if err != nil {
			t.Fatal(err)
		}

		err = tree.WriteBinaryValueOfLeaf(leaf, writer)
		if err != nil {
			t.Fatal(err)
		}

		s := writer.buffer.String()
		if s != expected {
			t.Fatalf("expected %s to encode to %s but got %s", string(r), expected, s)
		}
	}

	expect('T', "11000")
	expect('h', "0010")
	expect('i', "0111")
	expect('s', "1000")
	expect(' ', "000")
	expect('a', "0110")
	expect('n', "1010")
	expect('e', "010")
	expect('x', "11111")
	expect('m', "1001")
	expect('p', "11000")
	expect('l', "11011")
	expect('t', "1011")
	expect('f', "0011")
	expect('n', "1010")
	expect('r', "11110")
	expect('.', "11000")
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
