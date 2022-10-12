package main

import "testing"

func TestNumDistinctIslands(t *testing.T) {
	t.Run("It should return '0' for an empty matrix", func(t *testing.T) {
		got := numDistinctIslands(Matrix{})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '1' for a single land", func(t *testing.T) {
		got := numDistinctIslands(Matrix{{1}})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '0' for a single water", func(t *testing.T) {
		got := numDistinctIslands(Matrix{{0}})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '1' for for two indistinct islands", func(t *testing.T) {
		got := numDistinctIslands(Matrix{
			{1, 0},
			{0, 1},
		})
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '2' for for two distinct islands", func(t *testing.T) {
		got := numDistinctIslands(Matrix{
			{1, 0, 0},
			{0, 1, 0},
			{0, 1, 0},
		})
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '2' for for two distinct islands among three total islands", func(t *testing.T) {
		got := numDistinctIslands(Matrix{
			{1, 1, 0, 1, 1},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 1, 1, 0},
		})
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '3' for for three distinct islands", func(t *testing.T) {
		got := numDistinctIslands(Matrix{
			{0, 1, 1, 0, 1},
			{1, 1, 1, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0},
		})
		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
