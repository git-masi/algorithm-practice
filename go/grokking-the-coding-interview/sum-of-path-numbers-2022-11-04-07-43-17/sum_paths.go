package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func sumPaths(root *tree.Node) int {
	if root == nil {
		return 0
	}

	total := 0

	var helper func(node *tree.Node, sum int)
	helper = func(node *tree.Node, sum int) {
		sum = sum*10 + node.Value

		if node.Left == nil && node.Right == nil {
			total += sum
		}

		if node.Left != nil {
			helper(node.Left, sum)
		}

		if node.Right != nil {
			helper(node.Right, sum)
		}
	}

	helper(root, 0)

	return total
}
