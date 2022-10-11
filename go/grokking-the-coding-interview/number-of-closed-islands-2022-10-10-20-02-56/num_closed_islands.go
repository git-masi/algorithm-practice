package main

type Matrix [][]int

func numClosedIslands(m Matrix) int {
	if len(m) < 3 {
		return 0
	}

	var trav func(row, col int) int

	count := 0
	trav = func(row, col int) int {
		// return early if this is one of the outermost cells
		if row == 0 || row+1 == len(m) || col == 0 || col+1 == len(m) {
			return 0
		}

		m[row][col] = 0

		// top
		if m[row-1][col] == 1 && trav(row-1, col) == 0 {
			return 0
		}

		// right
		if m[row][col+1] == 1 && trav(row, col+1) == 0 {
			return 0
		}

		// bottom
		if m[row+1][col] == 1 && trav(row+1, col) == 0 {
			return 0
		}

		// left
		if m[row][col-1] == 1 && trav(row, col-1) == 0 {
			return 0
		}

		return 1
	}

	for i, sub := range m {
		for j := range sub {
			if m[i][j] == 1 {
				count += trav(i, j)
			}
		}
	}

	return count
}
