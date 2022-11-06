package main

import (
	"container/list"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func countPaths(root *tree.Node, target int) int {
	if root == nil {
		return 0
	}

	count := 0
	nodes := list.New()

	var helper func(node *tree.Node, nodes *list.List, sum int)
	helper = func(node *tree.Node, nodes *list.List, sum int) {
		for {
			if nodes.Len() == 0 || sum+node.Value <= target {
				break
			}
			f := nodes.Front()
			sum -= f.Value.(*tree.Node).Value
			nodes.Remove(f)
		}

		nodes.PushBack(node)
		sum += node.Value

		if sum == target {
			count++
		}

		var temp *list.List

		if node.Left != nil {
			temp = list.New()
			temp.PushFrontList(nodes)
			helper(node.Left, temp, sum)
		}

		if node.Right != nil {
			temp = list.New()
			temp.PushFrontList(nodes)
			helper(node.Right, temp, sum)
		}
	}

	helper(root, nodes, 0)

	return count
}
