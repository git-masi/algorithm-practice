package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func sum(root *tree.Node, target int) bool {
	if root == nil {
		return false
	}

	return helper(root, 0, target)
}

func helper(node *tree.Node, sum, target int) bool {
	if node == nil {
		return sum == target
	}

	sum += node.Value

	left := helper(node.Left, sum, target)
	right := helper(node.Right, sum, target)

	return left || right
}
