package main

import "testing"

func TestBiggestIsland(t *testing.T) {
	t.Run("It should return '0' for an empty matrix", func(t *testing.T) {
		got := biggestIsland(Matrix{})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '1' for: [[1]]", func(t *testing.T) {
		got := biggestIsland(Matrix{{1}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '0' for: [[0]]", func(t *testing.T) {
		got := biggestIsland(Matrix{{0}})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run(
		`
			It should return '2' for:
			[[0,1,0]
			 [0,0,0]
			 [1,1,0]]
		`,
		func(t *testing.T) {
			m := Matrix{
				{0, 1, 0},
				{0, 0, 0},
				{1, 1, 0},
			}
			got := biggestIsland(m)
			want := 2

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})

	t.Run(
		`
			It should return '6' for:
			[[0, 1, 1, 1]
			 [1, 1, 1, 0]
			 [0, 0, 0, 0]
		 	 [1, 1, 1, 1]]
		`,
		func(t *testing.T) {
			m := Matrix{
				{0, 1, 1, 1},
				{1, 1, 1, 0},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
			}

			got := biggestIsland(m)
			want := 6

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
}
