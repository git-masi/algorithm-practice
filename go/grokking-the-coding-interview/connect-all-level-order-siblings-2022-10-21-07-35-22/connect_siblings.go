package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func connectSiblings(bt *tree.Node) {
	if bt == nil {
		return
	}

	q := []*tree.Node{bt}

	for {
		if len(q) == 0 {
			break
		}

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

		if len(q) > 0 {
			node.Right = q[0]
		}
	}
}
