package huffman

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Tree struct {
	leafs []*Leaf
}

func NewTree(endLeafs []*Leaf) *Tree {
	tree := Tree{leafs: endLeafs}
	if len(tree.leafs) < 2 {
		panic("To create a huffman tree, at least two characters are needed.")
	}

	for tree.hasMultipleRoots() {
		parent := &Leaf{}
		for i := 0; i < len(tree.leafs); i++ {
			if tree.leafs[i].parent != nil {
				continue
			}
			parent.freq += tree.leafs[i].freq
			parent.value += tree.leafs[i].value
			if parent.left == nil {
				parent.left = tree.leafs[i]
				tree.leafs[i].parent = parent
			} else {
				parent.right = tree.leafs[i]
				tree.leafs[i].parent = parent
				break
			}
		}

		for i := len(tree.leafs) - 1; i >= 0; i-- {
			if tree.leafs[i].freq < parent.freq {
				tree.leafs = append(tree.leafs[:i+1], tree.leafs[i:]...)
				tree.leafs[i+1] = parent
				break
			}
		}
	}

	return &tree
}

func (tree *Tree) Decode(r BitReader) (string, error) {
	builder := strings.Builder{}
	root := tree.getRoot()
	leaf := root
	for {
		value, err := r.Read()
		if err != nil {
			if err == io.EOF {
				return builder.String(), nil
			}
			return "", err
		}

		if value {
			leaf = leaf.right
		} else {
			leaf = leaf.left
		}

		if leaf.left == nil || leaf.right == nil {
			builder.WriteString(leaf.value)
			leaf = root
		}
	}
}

func (tree *Tree) Encode(s string, w BitWriter) error {
	for _, r := range []rune(s) {
		l, err := tree.GetLeafForRune(r)
		if err != nil {
			return err
		}
		if err := tree.WriteBinaryValueOfLeaf(l, w); err != nil {
			return err
		}
	}
	return nil
}

func (tree *Tree) GetLeafForRune(r rune) (*Leaf, error) {
	value := string(r)
	for _, l := range tree.leafs {
		if l.value == value {
			return l, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("could not find value %s in tree", string(r)))
}

func (tree *Tree) WriteBinaryValueOfLeaf(leaf *Leaf, w BitWriter) error {
	parent := leaf.parent
	var bits []bool
	for {
		if parent == nil {
			break
		}

		if leaf == parent.left {
			bits = append(bits, false)
		} else {
			bits = append(bits, true)
		}

		leaf = parent
		parent = leaf.parent
	}

	for i := len(bits) - 1; i >= 0; i-- {
		var err error
		if bits[i] {
			err = w.WriteOne()
		} else {
			err = w.WriteZero()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (tree *Tree) hasMultipleRoots() bool {
	n := 0
	for _, l := range tree.leafs {
		if l.parent == nil {
			if n > 0 {
				return true
			} else {
				n++
			}
		}
	}
	return false
}

func (tree *Tree) getRoot() *Leaf {
	for _, leaf := range tree.leafs {
		if leaf.parent == nil {
			return leaf
		}
	}
	panic("tree has no root")
}

type Leaf struct {
	left   *Leaf
	right  *Leaf
	parent *Leaf
	freq   int
	value  string
	bin    string
}

func CalculateValueLeafs(s string) []*Leaf {
	var leafs []*Leaf

	for _, r := range []rune(s) {
		alreadyExists := false
		value := string(r)
		for i := range leafs {
			if value == leafs[i].value {
				leafs[i].freq++
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			leafs = append(leafs, &Leaf{value: value, freq: 1})
		}
	}

	sort.Slice(leafs, func(i int, j int) bool {
		return leafs[i].freq < leafs[j].freq
	})

	return leafs
}

func (l Leaf) String() string {
	if l.left == nil && l.right == nil {
		return fmt.Sprintf("leaf(value %s, freq %v, parent %v)", l.value, l.freq, l.parent)
	}
	return fmt.Sprintf("leaf(value %s, freq %v, parent %v, left %v, right %v)", l.value, l.freq, l.parent, l.left, l.right)
}
