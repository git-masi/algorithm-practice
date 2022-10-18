package main

import (
	binarytree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

type LevelOrderValues [][]int

func travLevelOrder(bt binarytree.Node) LevelOrderValues {
	values := LevelOrderValues{}
	q := []*binarytree.Node{&bt}
	idx := 0

	for {
		if idx == len(q) {
			break
		}

		levelSize := len(q) - idx
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			levelValues = append(levelValues, q[idx].Value)

			if q[idx].Left != nil {
				q = append(q, q[idx].Left)
			}

			if q[idx].Right != nil {
				q = append(q, q[idx].Right)
			}

			idx++
		}

		values = append(values, levelValues)
	}

	return values
}
