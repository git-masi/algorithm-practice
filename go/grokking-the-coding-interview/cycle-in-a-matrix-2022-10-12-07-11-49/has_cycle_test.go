package main

import "testing"

func TestHasCycle(t *testing.T) {
	t.Run("It should return 'false' for an empty matrix", func(t *testing.T) {
		got := hasCycle(Matrix{})
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return 'true' for a matrix with all land cells", func(t *testing.T) {
		got := hasCycle(Matrix{
			{1, 1},
			{1, 1},
		})
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return 'false' for a matrix with some land but no cycle", func(t *testing.T) {
		got := hasCycle(Matrix{
			{0, 1, 1},
			{0, 1, 0},
			{1, 1, 0},
		})
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return 'true' for an island with a lake", func(t *testing.T) {
		got := hasCycle(Matrix{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		})
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return 'true' if there are multiple islands and only one has cycle", func(t *testing.T) {
		got := hasCycle(Matrix{
			{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 0},
			{0, 1, 0, 1, 0, 0},
			{0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1},
			{0, 0, 0, 1, 1, 1},
		})
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
