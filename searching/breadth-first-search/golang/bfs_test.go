package bfs

import "testing"

func TestNode_FindPathToValue(t *testing.T) {
	const toFind = "here it is"
	expectedHistory := []int{0, 1, 0}
	n := &Node{
		value: "root",
		children: []*Node{
			{
				value: "child1",
				children: []*Node{
					{
						value: "nested1",
						children: []*Node{
							{value: "this is not it"},
							{value: "this isn't it either"},
						},
					},
					{
						value: "nested3",
						children: []*Node{
							{value: toFind},
						},
					},
				},
			},
			{
				value: "child2",
				children: []*Node{
					{value: "a"},
					{value: "b"},
					{value: "c"},
				},
			},
		},
	}

	p, err := n.FindPathToValue(toFind)

	if err != nil {
		t.Fatal(err)
	}

	if p.current.value != toFind {
		t.Fatalf("expected to find node with value %s but got value %s", toFind, p.current.value)
	}
	if len(p.history) != len(expectedHistory) {
		t.Fatalf("expected to find path of lenth %v but got length %v", len(expectedHistory), len(p.history))
	}
	for i := range expectedHistory {
		if expectedHistory[i] != p.history[i] {
			t.Fatalf("expected to go to branch %v at index %v, but went to %v", expectedHistory[i], i, p.history[i])
		}
	}
}
