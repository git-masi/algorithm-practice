package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

func calcAverages(bt *tree.Node) []float64 {
	result := []float64{}

	if bt == nil {
		return result
	}

	q := []*tree.Node{bt}

	for {
		size := len(q)

		if size == 0 {
			break
		}

		sum := float64(0)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum += float64(node.Value)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		result = append(result, sum/float64(size))
	}

	return result
}
