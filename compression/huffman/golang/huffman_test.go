package huffman

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	valueLeafs := CalculateValueLeafs("wooooooow this is short")
	leaf_r := &Leaf{freq: 1, value: "r"}
	leaf_w := &Leaf{freq: 2, value: "w"}
	leaf_t := &Leaf{freq: 2, value: "t"}
	leaf_h := &Leaf{freq: 2, value: "h"}
	leaf_i := &Leaf{freq: 2, value: "i"}
	leaf_Space := &Leaf{freq: 3, value: " "}
	leaf_s := &Leaf{freq: 3, value: "s"}
	leaf_o := &Leaf{freq: 8, value: "o"}

	leaf_rw := &Leaf{freq: 3, value: "rw", left: leaf_r, right: leaf_w}
	leaf_r.parent = leaf_rw
	leaf_w.parent = leaf_rw

	leaf_th := &Leaf{freq: 4, value: "th", left: leaf_t, right: leaf_h}
	leaf_t.parent = leaf_th
	leaf_h.parent = leaf_th

	leaf_irw := &Leaf{freq: 5, value: "irw", left: leaf_i, right: leaf_rw}
	leaf_i.parent = leaf_irw
	leaf_rw.parent = leaf_irw

	leaf_Spaces := &Leaf{freq: 6, value: " s", left: leaf_Space, right: leaf_s}
	leaf_Space.parent = leaf_Spaces
	leaf_s.parent = leaf_Spaces

	leaf_thirw := &Leaf{freq: 9, value: "thirw", left: leaf_th, right: leaf_irw}
	leaf_th.parent = leaf_thirw
	leaf_irw.parent = leaf_thirw

	leaf_Spaceso := &Leaf{freq: 14, value: " so", left: leaf_Spaces, right: leaf_o}
	leaf_Spaces.parent = leaf_Spaceso
	leaf_o.parent = leaf_Spaceso

	root := &Leaf{freq: 23, value: "thirw so", left: leaf_thirw, right: leaf_Spaceso}
	leaf_thirw.parent = root
	leaf_Spaceso.parent = root

	expected := []*Leaf{
		leaf_r,                         // 1
		leaf_w, leaf_t, leaf_h, leaf_i, // 2
		leaf_rw, leaf_Space, leaf_s, // 3
		leaf_th,      // 4
		leaf_irw,     // 5
		leaf_Spaces,  // 6
		leaf_o,       // 8
		leaf_thirw,   // 9
		leaf_Spaceso, //14
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
	const input = "wooooooow this is short"
	expected := []Leaf{
		{freq: 1, value: "r"},
		{freq: 2, value: "w"},
		{freq: 2, value: "t"},
		{freq: 2, value: "h"},
		{freq: 2, value: "i"},
		{freq: 3, value: " "},
		{freq: 3, value: "s"},
		{freq: 8, value: "o"},
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
	const input = "wooooooow this is short"
	tree := NewTree(CalculateValueLeafs(input))
	expect := func(r rune, expected string) {
		writer := &BinaryStringWriter{}
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

	expect('r', "0110")
	expect('w', "0111")
	expect('t', "000")
	expect('h', "001")
	expect('i', "010")
	expect(' ', "100")
	expect('s', "101")
	expect('o', "11")
}

func TestTree_Encode(t *testing.T) {
	const input = "wooooooow this is short"
	const expected = "0111111111111111110111100000001010101100010101100101001110110000"
	tree := NewTree(CalculateValueLeafs(input))
	writer := &BinaryStringWriter{}

	err := tree.Encode(input, writer)
	if err != nil {
		t.Fatalf("error while encoding: %v", err)
	}

	encoded := writer.buffer.String()
	if len(encoded) != len(expected) {
		t.Fatalf("the encoded value is %v runes long while %v were expected", len(encoded), len(expected))
	}
	if encoded != expected {
		t.Fatalf("the encoded value is not what was expected\r\nreceived: %s\r\nexpected: %s", encoded, expected)
	}
}

func TestTree_Decode(t *testing.T) {
	const input = "0111111111111111110111100000001010101100010101100101001110110000"
	const expected = "wooooooow this is short"
	tree := NewTree(CalculateValueLeafs(expected))
	reader := NewBinaryStringReader(input)

	decoded, err := tree.Decode(reader)
	if err != nil {
		t.Fatalf("error while decoding: %v", err)
	}

	if decoded != expected {
		t.Fatalf("the decoded value is not correct\r\nreceived: %s\r\nexpected: %s", decoded, input)
	}
}
