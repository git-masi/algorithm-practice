package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func minDepth(bt *tree.Node) int {
	if bt == nil {
		return 0
	}

	level := 0
	q := []*tree.Node{bt}

	for {
		level++
		size := len(q)

		if size == 0 {
			break
		}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return level
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return level
}
