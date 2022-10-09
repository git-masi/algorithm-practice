package main

import "fmt"

// The coordinates indicate which sub array and which element (0 indexed)
// as such [2,0] means the 3rd sub array and the 0th element
type Coordinates struct {
	row int
	col int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("%d#%d", c.row, c.col)
}

func numIslands(m [][]int) int {
	if len(m) == 0 {
		return 0
	}

	var traverse func(coord Coordinates)

	seen := map[string]string{}
	count := 0
	traverse = func(c Coordinates) {
		s := c.String()

		// Prevent infinite loops
		if _, ok := seen[s]; ok {
			return
		}

		seen[s] = s

		// top
		if c.row > 0 && m[dec(c.row)][c.col] == 1 {
			traverse(Coordinates{row: dec(c.row), col: c.col})
		}

		// right
		if c.col+1 < len(m[0]) && m[c.row][inc(c.col)] == 1 {
			traverse(Coordinates{row: c.row, col: inc(c.col)})
		}

		// bottom
		if c.row+1 < len(m) && m[inc(c.row)][c.col] == 1 {
			traverse(Coordinates{row: inc(c.row), col: c.col})
		}

		// left
		if c.col > 0 && m[c.row][dec(c.col)] == 1 {
			traverse(Coordinates{row: c.row, col: dec(c.col)})
		}
	}

	for i, arr := range m {
		for j, el := range arr {
			if el == 1 {
				c := Coordinates{row: i, col: j}
				if _, ok := seen[c.String()]; !ok {
					count++
					traverse(c)
				}
			}
		}
	}

	return count
}

func inc(num int) int {
	return num + 1
}

func dec(num int) int {
	return num - 1
}
