package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func hasSeq(root *tree.Node, seq []int) bool {
	if root == nil || seq == nil || len(seq) == 0 {
		return false
	}

	var helper func(node *tree.Node, count int) bool
	helper = func(node *tree.Node, count int) bool {
		if node == nil {
			return false
		}

		if count >= len(seq) || node.Value != seq[count] {
			return false
		}

		if node.Left == nil && node.Right == nil {
			return count == len(seq)-1
		}

		return helper(node.Left, count+1) || helper(node.Right, count+1)
	}

	return helper(root, 0)
}
