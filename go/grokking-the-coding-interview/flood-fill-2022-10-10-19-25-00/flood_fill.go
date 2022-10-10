package main

type Matrix [][]int

func floodFill(m Matrix, v, row, col int) {
	if len(m) == 0 {
		return
	}

	curr := m[row][col]
	m[row][col] = v

	// top
	if row != 0 && m[row-1][col] == curr {
		floodFill(m, v, row-1, col)
	}

	// right
	if col+1 < len(m[0]) && m[row][col+1] == curr {
		floodFill(m, v, row, col+1)
	}

	// bottom
	if row+1 < len(m) && m[row+1][col] == curr {
		floodFill(m, v, row+1, col)
	}

	// left
	if col != 0 && m[row][col-1] == curr {
		floodFill(m, v, row, col-1)
	}
}
