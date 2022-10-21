package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func connect(bt *tree.Node) {
	if bt == nil {
		return
	}

	q := []*tree.Node{bt}

	for {
		size := len(q)

		if size == 0 {
			break
		}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			node.Left = nil
			node.Right = nil

			if i+1 < size {
				node.Right = q[0]
			}
		}
	}
}
