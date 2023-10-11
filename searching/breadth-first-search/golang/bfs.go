package bfs

import "errors"

type Node struct {
	children []*Node
	value    string
}

type Explorer struct {
	current *Node
	history []int
}

func (n *Node) FindPathToValue(value string) (explorer Explorer, err error) {
	paths := make([]Explorer, len(n.children))
	for i, c := range n.children {
		paths[i].current = c
		paths[i].history = []int{i}
	}

	for len(paths) > 0 {
		var newPaths []Explorer
		for _, p := range paths {
			if p.current.value == value {
				return p, nil
			}

			childPaths := make([]Explorer, len(p.current.children))
			for i := 0; i < len(p.current.children); i++ {
				childPaths[i] = Explorer{
					current: p.current.children[i],
					history: append(p.history, i),
				}
			}
			newPaths = append(newPaths, childPaths...)
		}
		paths = newPaths
	}

	return Explorer{}, errors.New("no node found with the given value")
}
