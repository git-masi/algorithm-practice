package main

import binarytree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

type LevelOrderValues [][]int

func travLevelOrder(bt binarytree.Node) LevelOrderValues {
	values := LevelOrderValues{}
	// use linked list for better performance
	q := []*binarytree.Node{&bt}

	for {
		size := len(q)

		if size == 0 {
			break
		}

		v := []int{}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			v = append(v, node.Value)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		values = append(values, v)
	}

	return values
}
