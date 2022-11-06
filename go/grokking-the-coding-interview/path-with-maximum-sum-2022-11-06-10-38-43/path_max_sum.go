package main

import tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"

type Path []*tree.Node

func pathMaxSum(root *tree.Node) Path {
	path := Path{}
	sum := 0

	if root == nil {
		return path
	}

	var helper func(node *tree.Node) (Path, int)
	helper = func(node *tree.Node) (Path, int) {
		if node == nil {
			return Path{}, 0
		}

		pL, sL := helper(node.Left)
		pR, sR := helper(node.Right)

		if sL+sR+node.Value > sum {
			sum = sL + sR + node.Value
			path = Path{node}
			path = append(path, pL...)
			path = append(path, pR...)
		}

		p := Path{node}

		if sL > sR {
			p = append(p, pL...)
			return p, sL + node.Value
		}

		p = append(p, pR...)
		return p, sR + node.Value
	}

	helper(root)

	return path
}
