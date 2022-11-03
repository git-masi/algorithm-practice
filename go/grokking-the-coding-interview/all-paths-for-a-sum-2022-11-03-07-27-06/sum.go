package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func sum(root *tree.Node, target int) [][]*tree.Node {
	result := [][]*tree.Node{}

	if root == nil {
		return result
	}

	var helper func(node *tree.Node, path []*tree.Node, sum int)

	helper = func(node *tree.Node, path []*tree.Node, sum int) {
		path = append(path, node)
		sum += node.Value

		if node.Left == nil && node.Right == nil {
			if sum == target {
				result = append(result, path)
			}
			return
		}

		if node.Left != nil {
			helper(node.Left, path, sum)
		}

		if node.Right != nil {
			helper(node.Right, path, sum)
		}
	}

	helper(root, []*tree.Node{}, 0)

	return result
}
