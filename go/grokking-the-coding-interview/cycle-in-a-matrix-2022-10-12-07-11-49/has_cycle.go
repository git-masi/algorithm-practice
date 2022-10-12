package main

type Matrix [][]int

func hasCycle(m Matrix) bool {
	if len(m) < 2 {
		return false
	}

	var trav func(row, col, prow, pcol int) bool
	result := false

	trav = func(row, col, prow, pcol int) bool {
		m[row][col] = 2

		var top int
		var right int
		var bottom int
		var left int
		v := false
		isprev := func(r, c int) bool {
			return r == prow && c == pcol
		}

		if row != 0 {
			top = m[row-1][col]
		} else {
			top = -1
		}

		if col+1 < len(m[0]) {
			right = m[row][col+1]
		} else {
			right = -1
		}

		if row+1 < len(m) {
			bottom = m[row+1][col]
		} else {
			bottom = -1
		}

		if col != 0 {
			left = m[row][col-1]
		} else {
			left = -1
		}

		if top == 2 && !isprev(row-1, col) {
			return true
		}

		if top == 1 {
			v = v || trav(row-1, col, row, col)
		}

		if right == 2 && !isprev(row, col+1) {
			return true
		}

		if right == 1 {
			v = v || trav(row, col+1, row, col)
		}

		if bottom == 2 && !isprev(row+1, col) {
			return true
		}

		if bottom == 1 {
			v = v || trav(row+1, col, row, col)
		}

		if left == 2 && !isprev(row, col-1) {
			return true
		}

		if left == 1 {
			v = v || trav(row, col-1, row, col)
		}

		return v
	}

	for row, sub := range m {
		for col, el := range sub {
			if el == 1 {
				result = result || trav(row, col, 0, 0)
			}
		}
	}

	return result
}
