package main

type Matrix [][]int

func biggestIsland(m Matrix) int {
	if len(m) == 0 {
		return 0
	}

	var trav func(m Matrix, row int, col int) int
	largest := 0

	trav = func(m Matrix, row int, col int) int {
		s := 1

		// note that we have processed the current coordinates
		m[row][col] = 0

		// top
		if row != 0 && m[dec(row)][col] == 1 {
			s += trav(m, dec(row), col)
		}

		// right
		if inc(col) < len(m[0]) && m[row][inc(col)] == 1 {
			s += trav(m, row, inc(col))
		}

		// bottom
		if inc(row) < len(m) && m[inc(row)][col] == 1 {
			s += trav(m, inc(row), col)
		}

		// left
		if col != 0 && m[row][dec(col)] == 1 {
			s += trav(m, row, dec(col))
		}

		return s
	}

	for i, slc := range m {
		for j := range slc {
			if m[i][j] == 1 {
				s := trav(m, i, j)

				if s > largest {
					largest = s
				}
			}
		}
	}

	return largest
}

func inc(num int) int {
	return num + 1
}

func dec(num int) int {
	return num - 1
}
