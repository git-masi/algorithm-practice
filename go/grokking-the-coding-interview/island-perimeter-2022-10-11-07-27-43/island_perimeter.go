package main

type Matrix [][]int

func islandPerimeter(m Matrix) int {
	if len(m) == 0 {
		return 0
	}

	var trav func(row, col int) int
	p := 0
	end := false
	trav = func(row, col int) int {
		s := 0

		// Mark current cell as visited
		m[row][col] = 2

		// Calculate number of sides to add to perimeter
		// top
		if row == 0 || m[row-1][col] == 0 {
			s += 1
		}

		// right
		if col+1 == len(m[0]) || m[row][col+1] == 0 {
			s += 1
		}

		// bottom
		if row+1 == len(m) || m[row+1][col] == 0 {
			s += 1
		}

		// left
		if col == 0 || m[row][col-1] == 0 {
			s += 1
		}

		// Traverse adjacent land
		// top
		if row != 0 && m[row-1][col] == 1 {
			s += trav(row-1, col)
		}

		// right
		if col+1 < len(m[0]) && m[row][col+1] == 1 {
			s += trav(row, col+1)
		}

		// bottom
		if row+1 < len(m) && m[row+1][col] == 1 {
			s += trav(row+1, col)
		}

		// left
		if col != 0 && m[row][col-1] == 1 {
			s += trav(row, col-1)
		}

		return s
	}

	for row, sub := range m {
		for col := range sub {
			if m[row][col] == 1 {
				p += trav(row, col)
				end = true
			}

			if end {
				break
			}
		}

		if end {
			break
		}
	}

	return p
}
