package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func rightView(root *tree.Node) []*tree.Node {
	result := []*tree.Node{}

	if root == nil {
		return result
	}

	level := []*tree.Node{root}

	for {
		size := len(level)

		if size == 0 {
			break
		}

		for i := 0; i < size; i++ {
			node := level[0]
			level = level[1:]

			if i+1 == size {
				result = append(result, node)
			}

			if node.Left != nil {
				level = append(level, node.Left)
			}

			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
	}

	return result
}
