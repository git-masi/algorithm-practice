package main

type Matrix [][]int

func numIslandsRecursive(m Matrix) int {
	if len(m) == 0 {
		return 0
	}

	var trav func(row, col int)
	count := 0
	trav = func(row, col int) {
		// Mark cell as visited
		m[row][col] = 0

		// Visit top
		if row != 0 && m[row-1][col] == 1 {
			trav(row-1, col)
		}

		// Visit right
		if col+1 < len(m[0]) && m[row][col+1] == 1 {
			trav(row, col+1)
		}

		// Visit bottom
		if row+1 < len(m) && m[row+1][col] == 1 {
			trav(row+1, col)
		}

		// Visit left
		if col != 0 && m[row][col-1] == 1 {
			trav(row, col-1)
		}

	}

	for row, sub := range m {
		for col, cell := range sub {
			if cell == 1 {
				count++
				trav(row, col)
			}
		}
	}

	return count
}

func numIslandsIterative(m Matrix) int {
	if len(m) == 0 {
		return 0
	}

	count := 0

	for r, sub := range m {
		for c, cell := range sub {
			if cell == 1 {
				count++
				// List of cells to visit
				cells := [][2]int{{r, c}}

				// Mark cell as visited
				m[r][c] = 0

				for {
					// Base case
					if len(cells) == 0 {
						break
					}

					// Get the current cell
					cell := cells[len(cells)-1]
					row := cell[0]
					col := cell[1]

					// Remove cell from list of cells to visit
					cells = cells[:len(cells)-1]

					// Visit top
					if row != 0 && m[row-1][col] == 1 {
						m[row-1][col] = 0
						cells = append(cells, [2]int{row - 1, col})
					}

					// Visit right
					if col+1 < len(m[0]) && m[row][col+1] == 1 {
						m[row][col+1] = 0
						cells = append(cells, [2]int{row, col + 1})
					}

					// Visit bottom
					if row+1 < len(m) && m[row+1][col] == 1 {
						m[row+1][col] = 0
						cells = append(cells, [2]int{row + 1, col})
					}

					// Visit left
					if col != 0 && m[row][col-1] == 1 {
						m[row][col-1] = 0
						cells = append(cells, [2]int{row, col - 1})
					}
				}
			}
		}
	}

	return count
}
