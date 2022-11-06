package main

import (
	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func diameter(root *tree.Node) int {
	d := 0

	var helper func(node *tree.Node) int
	helper = func(node *tree.Node) int {
		if node == nil {
			return 0
		}

		left := helper(node.Left)
		right := helper(node.Right)

		d = max(left+right+1, d)

		return max(left, right) + 1
	}

	helper(root)

	return d
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
