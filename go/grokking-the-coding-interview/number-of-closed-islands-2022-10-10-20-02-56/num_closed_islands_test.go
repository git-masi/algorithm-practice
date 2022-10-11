package main

import (
	"testing"
)

func TestNumClosedIslands(t *testing.T) {
	t.Run("It should return '0' if the input matrix is empty", func(t *testing.T) {
		m := Matrix{}
		got := numClosedIslands(m)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '0' if the input matrix is < 3x3", func(t *testing.T) {
		m := Matrix{
			{1, 1},
			{1, 1},
		}
		got := numClosedIslands(m)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '1' if the only island is surrounded by water", func(t *testing.T) {
		m := Matrix{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		got := numClosedIslands(m)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return '2' if there are two islands surrounded by water", func(t *testing.T) {
		m := Matrix{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
		}
		got := numClosedIslands(m)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return the correct value regardless of island size", func(t *testing.T) {
		m := Matrix{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		got := numClosedIslands(m)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
