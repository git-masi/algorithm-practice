package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func travZigZag(bt *tree.Node) [][]int {
	result := [][]int{}

	if bt == nil {
		return result
	}

	q := []*tree.Node{bt}
	idx := 0
	leftToRight := false

	for {
		size := len(q) - idx

		if size == 0 {
			break
		}

		v := []int{}

		for i := 0; i < size; i++ {
			node := q[idx]

			if leftToRight {
				v = append([]int{node.Value}, v...)
			} else {
				v = append(v, node.Value)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			idx++
		}

		result = append(result, v)
		leftToRight = !leftToRight
	}

	return result
}
