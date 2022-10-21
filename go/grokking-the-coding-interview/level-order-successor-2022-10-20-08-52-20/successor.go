package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func successor(bt, target *tree.Node) *tree.Node {
	if bt == target {
		return nil
	}

	found := false
	q := []*tree.Node{bt}

	for {
		if len(q) == 0 {
			break
		}

		node := q[0]
		q = q[1:]

		if found {
			return node
		}

		if node == target {
			found = true
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return nil
}
