package main

import "fmt"

type Matrix [][]int

func numDistinctIslands(m Matrix) int {
	if len(m) == 0 {
		return 0
	}

	var trav func(row, col int, p string) string
	paths := map[string]struct{}{}

	trav = func(row, col int, p string) string {
		m[row][col] = 0

		// top
		if row != 0 && m[row-1][col] == 1 {
			p += trav(row-1, col, "t")
		}

		// right
		if col+1 < len(m[0]) && m[row][col+1] == 1 {
			p += trav(row, col+1, "r")
		}

		// bottom
		if row+1 < len(m) && m[row+1][col] == 1 {
			p += trav(row+1, col, "b")
		}

		//left
		if col != 0 && m[row][col-1] == 1 {
			p += trav(row, col-1, "l")
		}

		return p
	}

	for row, sub := range m {
		for col := range sub {
			if m[row][col] == 1 {
				p := trav(row, col, "o")

				fmt.Println(p)

				if _, ok := paths[p]; !ok {
					paths[p] = struct{}{}
				}
			}
		}
	}

	return len(paths)
}
