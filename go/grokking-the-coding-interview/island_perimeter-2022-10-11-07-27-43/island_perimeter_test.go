package main

import "testing"

func TestIslandPerimeter(t *testing.T) {
	t.Run("It should return '0' for an empty matrix", func(t *testing.T) {
		got := islandPerimeter(Matrix{})
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return '4' for a single land", func(t *testing.T) {
		got := islandPerimeter(Matrix{{1}})
		want := 4

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return '0' for a single water", func(t *testing.T) {
		got := islandPerimeter(Matrix{{0}})
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return '8'", func(t *testing.T) {
		got := islandPerimeter(Matrix{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		})
		want := 8

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return '14'", func(t *testing.T) {
		got := islandPerimeter(Matrix{
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
		})
		want := 14

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
