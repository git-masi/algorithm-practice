package main

import "testing"

func TestNumIslands(t *testing.T) {
	t.Run("It should return `0` if the input slice is empty", func(t *testing.T) {
		got := numIslands([][]int{})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return `1` for [[1]]", func(t *testing.T) {
		got := numIslands([][]int{{1}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return `1` for [[1, 0], [0, 0]]", func(t *testing.T) {
		got := numIslands([][]int{{1, 0}, {0, 0}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return `1` for [[1, 1], [0, 0]]", func(t *testing.T) {
		got := numIslands([][]int{{1, 1}, {0, 0}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return `1` for [[1, 1], [1, 0]]", func(t *testing.T) {
		got := numIslands([][]int{{1, 1}, {1, 0}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return `1` [[1, 1], [1, 1]]", func(t *testing.T) {
		got := numIslands([][]int{{1, 1}, {1, 1}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run(
		`It should return "1" for:
			[[0,1,0],
			 [1,1,1],
			 [0,1,0]]
		`,
		func(t *testing.T) {
			got := numIslands([][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			})
			want := 1

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		},
	)

	t.Run(
		`It should return "2" for:
			[[0,1,0,0],
			 [1,1,1,0],
			 [0,1,0,0],
			 [0,0,0,1]]
		`,
		func(t *testing.T) {
			got := numIslands([][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 1},
			})
			want := 2

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		},
	)
}
