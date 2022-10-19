package main

import (
	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func revLevelOrder(bt *tree.Node) list.LinkedList[[]int] {
	result := list.LinkedList[[]int]{}
	q := list.LinkedList[*tree.Node]{}

	q.Push(bt)

	for {
		size := q.Size

		if size == 0 {
			break
		}

		row := []int{}

		for i := 0; i < size; i++ {
			node := q.Shift().Value
			row = append(row, node.Value)

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		result.Unshift(row)
	}

	return result
}
