package huffman

import (
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

func (tree *Tree) getLeafs() []*Leaf {
	return tree.leafs
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

		if leaf.rune != 0 {
			builder.WriteRune(leaf.rune)
			leaf = root
		}
	}
}

func (tree *Tree) Encode(s string, w BitWriter) error {
	for _, r := range []rune(s) {
		for _, l := range tree.leafs {
			if l.rune == r {
				if err := tree.writeBinaryValueOfLeaf(l, w); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func (tree *Tree) writeBinaryValueOfLeaf(leaf *Leaf, w BitWriter) error {
	parent := leaf.parent
	for {
		if parent == nil {
			return nil
		}

		var err error
		if leaf == parent.left {
			err = w.WriteZero()
		} else {
			err = w.WriteOne()
		}

		if err != nil {
			return err
		}
		leaf = parent
		parent = leaf.parent
	}
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
	rune   rune
}

func CalculateValueLeafs(s string) []*Leaf {
	var leafs []*Leaf

	for _, r := range []rune(s) {
		alreadyExists := false
		for i := range leafs {
			if r == leafs[i].rune {
				leafs[i].freq++
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			leafs = append(leafs, &Leaf{rune: r, freq: 1})
		}
	}

	sort.Slice(leafs, func(i int, j int) bool {
		return leafs[i].freq < leafs[j].freq
	})

	return leafs
}

func (l Leaf) String() string {
	if l.rune == 0 {
		return fmt.Sprintf("leaf(left %v, right %v, parent %v)", l.left, l.right, l.parent)
	}
	return fmt.Sprintf("leaf(value %s, freq %v, parent %v)", string(l.rune), l.freq, l.parent)
}
